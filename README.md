# mango
Mango monorepo for all of avosa's applications and services.
Please note that this repo is currently maintained on Windows, and all scripts and settings are setup as such.

## Project Requirements
* PostgreSQL v9.1 and above (We use 9.6)
* GO v1.7.3 and above
* Beego v1.8.3 and above

## Running the Project
* $ npm install
* $ gulp (Watchers have been setup for JS & CSS changes.)
* $ To build all projects, run 'build.ps1' and compiled files will be copied to /mango/bin

## Project Layout
* The API folder contains all micro-services and APIs.
* The APP folder contains all websites and applications.
* The DB folder contains all database models and their logic.
* The UTIL folder contains logic used by most applications.

The 'Router' and 'WWW' applications are key to running any of the other products found within this repo.
If you don't require fancy named URLs, then running 'Router' is all you need to get up and running.

### Router
The router API is used hold all versions of an application or API's URLs
For example, we can use the same router for all of our environments like 'LIVE', 'UAT', and 'DEV'
This means that we can ask the router API for our e-mail API (Comms.API) and depending on the environment of the caller
the API, we will get the correct URL.
The functionality provided by this API also ensures that we can't all anything from 'LIVE' when running on 'DEV',
and may be seen as a way of keeping developers safe.

In order for the router to know about a Database, API or Application, we have to register with the router API on start-up.
We don't need to store URLs anywhere in our applications or JavaScript.
We just need to know what application we want.
This decreases effort when deploying our applications and also keeps the code free from URLs which could cause problems in the future.

You will see 'srv.Register(port)' within the main.go of every application.
This function is used to register an application.

### WWW (/APP/www)
The www application acts as the main entry point for all applications, as we can assign names to every registered application.
So, when debugging we can call 'https://comms.localhost:80/' instead of 'http://localhost:8085' for the 'Comms.API'.
This avoids the need to remember the specific port for every application running.

WWW also acts as the default website for avosa.co.za, this website is used as a fallback when
an application can't be found by the Router API.

We require 1(one) instance of www running for every environment we have.

## API Folder
- ### Artifact
  Images, Audio, and other assets should be stored and retrieved using this API.
- ### Comment
  Any comments made on the system, should all be controlled by the Comment API.
- ### Comms
  Email, SMS and other Messages are all to handled by Comms.
- ### Router
  See the description for Router above.
- ### Secure
  User authentication and session control is all done by the Secure API.
- ### Things
  This API will act as a central lookup database.

## APP Folder
- ### Admin
  We should be able to control and monitor every application and it's data from this application.
- ### Classifieds
  Classifieds will act as the central platform for advertising products and services.
- ### Logbook
  Logbook is an application that will be built to provide added value to the Classifieds application.
- ###Shop
  This application is our central e-commerce platform.
- ### WWW
  See the description for www above.
