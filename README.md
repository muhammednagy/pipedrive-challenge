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
You need to have GO 1.15 or higher installed locally and build-essentials.

You would also need MySQL for a database. you can either run it locally or run it via a docker using.
`docker run --name mysqlDB -e MYSQL_DATABASE=pipedrive MYSQL_USER=pipedrive MYSQL_PASSWORD=pipedrive -p 3306:3306`

To build: ```make build```  
To run: ```make run```  
To test: ```make test```  
To run tests and display coverage percentage: ```make cover```   
To clean the binary file: ```make clean```

or you can just run it via docker-compose. please make sure to update pipedrive token first.
`docker-compose up -d` application should be listening on port  3000 afterwards.
### API
The app support create, read and delete operations.
For more details please take a look at the application swagger running at `/documentation/index.html`

* `GET` `/api/v1/people` - get all people who their gists are being tracked
* `GET` `/api/v1/people/:username` - get a person gists since last visit (can also get all gists. please look at swagger for more details)
* `DELETE` `/api/v1/people/:username` - delete a person from app DB and stop tracking the person's github
* `POST` `/api/v1/people` - add a new person.  send username in as formData in the request body.

## Part II - CI /CD
I used github actions to set up CI/CD process. Github actions gets triggered by every commit pushed to master and
every merge request. It checks for a wide variety of things including running tests, build, checking for lint errors, 
errors checking and security check

it will also automatically build an image of the app and tag it if all the previous checks passed.
the docker images will be deployed to google cloud containers registry. project id and service account are defined in 
repo secrets to ensure security of the key and ease of changing the project id or the service key without updating the code.

also, during the CI/CD process we run code quality check on the application also for possible performance issues and anti-patterns
which brings huge value to the software process delivery along with the other checks mentioned earlier.
## Part III - The cloud
* Ror running the checks on user gists i didn't have to use any ops related solution, but I used golang built-in schedulers.
a function that would get all users and then fetch their gists then make activities in pipedrive every 3 hours. 

* The docker image is following the builder pattern where we build on one container then run from another container.
that allows us not to have to store our code on the customer facing container. 

* Kubernetes makes it extremely easy to scale up or down or even auto scale in addition to being the industry standard 
  which is why I choose it.
* I've used terraform for provisioning and controlling the infrastructure. 

* to provision a new cluster you need to: 
  * Download Terraform from [here](https://www.terraform.io/downloads.html) 
  * Get you service account key from Google with at least project editor access. for more info check 
    [here](https://cloud.google.com/iam/docs/creating-managing-service-account-keys)
  * copy your service key to deploy/credentials and name it serviceaccount.json and update provider.tf with the appropriate values
  * cd to deploy/terraform then execute `terraform apply` and write yes when asked for it!
  
  That's it! you have provisioned your cluster.

* To deploy it to your new cluster using github:
    * You can edit the Github secrets with the new project id, cluster name, GKE zone and service account.
    * re-run the latest job for the main branch from github actions.
* To deploy it to your new cluster without using github:
  * cd deploy
  * update COMMIT_SHA with the commit SHA you want to deploy or use latest for latest image deployment
  *  ./deploy.sh with the new project id, cluster name, GKE zone and service account.

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

* for part 2 and part 3 of the task I was trying to decide how I will implement auto deployment. I realized I have 2 options:

  * First option: to have github actions do everything from testing to building to deploying and to deploy somewhere else the user
  would have to edit the variables manually then push which would trigger github actions to deploy to the new place. it is
    clean (you can easily spot which step in the pipeline failed) and simple. 
    
  * Second option: have the github actions pipeline test, build and push to google cloud container registry (gcr) 
  then run a bash script to do the deployment. thus allowing the pipeline to use the same script to deploy and allowing 
    the user to deploy on his own.
    user can easily create a whole new cluster and deploy to it with 2 commands one for terraform apply, and
    the other to run the script
    
  I decided to use the second option