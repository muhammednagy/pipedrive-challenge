# pipedirve-challenge
DevOps challenge from pipedrive


## Part I - Github API & Pipedrive API

### Requirements
Variables could also be passed as cli parameters if you would like run with -h to see a list of possible parameters

* `PIPEDRIVE_TOKEN` must be set as environment variable or as a cli parameter
* `GITHUB_TOKEN` is optional but adding it reduces possibility of rate limits errors
* `DATABASE_NAME` name of mysql database
* `TEST_DATABASE_NAME` name of mysql database for testing
* `DATABASE_USERNAME` mysql username
* `DATABASE_PASSWORD` mysql password 
* `DATABASE_HOST` mysql host
* `DATABASE_PORT` mysql port
### Build
You need to have go 1.15 or higher installed locally and build-essentials.

You would also need MySQL for a database. you can either run it locally or run it via a docker using.
`docker run --name mysqlDB -e MYSQL_DATABASE=pipedrive MYSQL_USER=pipedrive MYSQL_PASSWORD=pipedrive -p 3306:3306`

To build: ```make build```  
To run: ```make run```  
To test: ```make test```  
To run tests and display coverage percentage: ```make cover```   
To clean the binary file: ```make clean```

### API
The app support create, read and delete operations.
For more details please take a look at the application swagger running at `/documentation/index.html`

* `GET` `/api/v1/people` - get all people who their gists are being tracked
* `GET` `/api/v1/people/:username` - get a person gists since last visit (can also get all gists. please look at swagger for more details)
* `DELETE` `/api/v1/people/:username` - delete a person from app DB and stop tracking the person's github
* `POST` `/api/v1/people` - add a new person.  send username in as formData in the request body.

### Notes

* Upon trying to understand  the difference between pipedrive terms deal and activity it seemed like activity makes
  sense as a gist
  
* I started with using sqlite3 but then realized that scaling with sqlite3 is going to be very hard due to it being
a file and being really hard to be used by more than 1 replica sets at the same time. so I decided to convert to MySQL instead.
  however, changing database to for example PostgreSQL is fairly simple and easy due to using an ORM.
  Changing to sqlite3 would take a bit more changes than a couple of code lines due to the fact that sqlite is based on a file.
  Current ORM does all action by default in a transaction which would mean that if two replicas tried to update
  at the same time exactly then only 1 would update and the other won't. 

* a github user seems like it could be a person, and I don't find any harm in storing the username and email (we get it
  from Github), so I decided to make a DB person correspond to Pipedrive Person and having gists as notes that would
  have the gist raw url (or gist pull url if it's marked as truncated due to having so big files)
  
* Checking for persons new gists will occur every 3 hours. a ticker will tick every 3 hours causing a check on all persons
  new gists by calling `exporter.ExportGists(dbConnection, configuration)`.
  
* Adding a person who is already added will result in a BadRequest response.