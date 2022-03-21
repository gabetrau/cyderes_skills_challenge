# cyderes_skills_challenge
Creating GCP API utilizing API Gateway, Cloud Functions, Cloud Run, Terraform, and Cloud Build

Looking for:
	* programming style/best practices
	* Git repo setup
	* testing methodology
		- CI integration
		- 100% test coverage
	* Strengths/Weaknesses
	* Use Terraform to create infrastructure
	* Creativity

URL: https://library-63f35y33qa-uc.a.run.app/
* / -> GET request that returns welcome text
* /books -> GET request that returns sorted list by title of all the books in the library
* /books/<author> -> GET request that returns sorted list of books by a specified author
* /books -> POST request with JSON attached to add book to the library
	
ex. curl -X POST https://library-63f35y33qa-uc.a.run.app/books
-H 'Content-Type: application/json'
-d '{"id": "13", "title": "Harry Potter and the Order of the Pheonix", "author": "J.K Rowling", "genre": "fiction"}'
	

Decided on GCP for cloud platform 
* chose this because of CYDERES close relationship with Google Chronicle.
* I also enjoy the talks google cloud tech gives about its services. 
* Serverless possibilities with GCP: Cloud Run, Cloud Engine, Cloud Functions
	- Cloud Run = bringing serverless to containers using Google's Kuebernetes Engine 
	- Cloud Engine = run applications on cloud
	- Cloud Functions = cloud services <-> cloud functions -> invokes other services
* API Gateway = I tried to use this with Cloud Run and Cloud Functions but wasn't able to do it

Resource I used for inspiration: https://github.com/GoogleCloudPlatform/serverless-expeditions/tree/main/terraform-serverless
  
I decided to use Cloud Run. The source code is in the goservice folder. There is also a folder for a Cloud Function.
The function works, but I did not get a chance to integrate it with the API.

  
CI/CD done with cloudbuild.yaml
 * this file contains the logic for building an image out of the source code contained in the goservice directory

  
Terraform files => main.tf, project.tf, storage.tf, variables.tf, goservice.tf, gofunc.tf, outputs.tf
  - I used these files to create my infrasture and should only require running the following commands

  
# Steps to Create Infrastructure
  
* install gloud and connect with your project ID
  
* clone this repository
	
```
gcloud builds submit
```
```
terraform init
terraform apply
```
  - provide project id
  - yes
  

	
