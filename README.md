# pipedirve-challenge
DevOps challenge from pipedrive


## Part I - Github API & Pipedrive API

### Requirements
Variables could also be passed as cli parameters if you would like run with -h to see a list of possible parameters

* `PIPEDRIVE_TOKEN` must be set as environment variable or as a cli parameter
* `GITHUB_TOKEN` is optional but adding it reduces possibility of rate limits errors
* `PORT` is optional default is 3000
* `DATABASE_NAME` is optional default is database.sqlite3 it can also be used a path if you want to save your database
  somewhere else than the root of the app
### Build
You need to have go 1.15 or higher installed locally and build-essentials.

To build: ```make build```  
To run: ```make run```  
To test: ```make test```  
To run tests and display coverage percentage: ```make cover```   
To clean the binary file: ```make clean```

### API
The app support create, read and delete operations.
For more details please take a look at the application swagger running at `/documentation/index.html`

* `GET` `/api/v1/people` - get all people who their gists are being tracked
* `GET` `/api/v1/person/:username` - get a person gists since last visit (can also get all gists. please look at swagger for more details)
* `DELETE` `/api/v1/person/:username` - delete a person from app DB and stop tracking the person's github
* `POST` `/api/v1/person` - add a new person.  send username in as formData in the request body.

### Notes

* Upon trying to understand  the difference between pipedrive terms deal and activity it seemed like activity makes
  sense as a gist
  
* I decided to use sqlite3 since the app usage is fairly simple and having a big postgresql or mysql server might be
  too much at least in the beginning of the API. switching to postgres or mysql should require around 3 lines of code
  change thanks to using Gorm (ORM)

* a github user seems like it could be a person, and I don't find any harm in storing the username and email (we get it
  from Github), so I decided to make a DB person correspond to Pipedrive Person and having gists as notes that would
  have the gist raw url (or gist pull url if it's marked as truncated due to having so big files)
  
* Checking for persons new gists will occur every 3 hours. a ticker will tick every 3 hours causing a check on all persons
  new gists by calling `exporter.ExportGists(dbConnection, configuration)`.
  
* Adding a person who is already added will result in a BadRequest response.