# Mango
Mango is the core module which allows us to easily connect applications and services.
The services and applications built using this framework and Docker allow us to have a some control over a lot of moving parts.

Every module can be individually edited, compiled and deployed, you don't even need Go or Dart installed, via Docker.   

Applications are built using;
* Dart for Front-end logic only.
* Go for Back-end logic
* Templates are rendered using Go's default HTML Template language.
* Husk for Storage
* Docker for Networking and scalability

HTTPS is the only option for running these modules, as modern-day requirements don't allow for un-secure connections and dealing with security at deploy time is not fun.

[![CodeFactor](https://www.codefactor.io/repository/github/louisevanderlith/mango/badge)](https://www.codefactor.io/repository/github/louisevanderlith/mango)

## Project Requirements
* Docker
* docker-compose

## Development Requirements
* Go
* Dart
* NPM 
* meta (npm i -g meta)
* gulp

## Project Layout
This repository used to contain all of the APP and API modules, but they have since been moved to their own repositories.
Now it contains shared logic for Controllers, Service Discovery, etc. As well as some information on how to setup and run everything in union.

The 'Router' and 'Gate' applications are key to running any of the other products found within this profile.
If you don't require fancy named URLs, then running 'Router' is all you need to get up and running.

After cloning the desired application, navigate to it and run using docker-compose.
Specific details on running every module can be found in it's README.md

Please see /secure repository for User login information.

## On Paged Data
Mango API's support paging out of the box, as we use Husk which demands all results be requested by page.
Every GET request, where the intention is to get many results, you have to specify the :pagesize Data.
https://ads.localhost/advert/all/:pageSize
https://ads.localhost/advert/all/A6

Where `A6` is `Page A` & `6 items per Page` is requested.
We support 26 pages (A-Z) and any positve amount of results per page. 



# Everything you need to know about the mango project.
## Creating machines to run this project
1. Create 'docker-machines'
    We need to setup atleast 2 docker-machines to run a swarm stack.
    * First create a machine for running the Gate and UI applications. (*-box) naming is used to identify 'Gate' machines.
    ```docker-machine create --swarm --driver "virtualbox" --virtualbox-disk-size "5000" mango-box```
    * Create a machine for running the Web services. (*-api) naming is used to identify 'Web service' machines
    ```docker-machine create --swarm --driver "virtualbox" --virtualbox-disk-size "5000" mango-api```

2. Setup Swarm
3. Load Balancer 
https://auth0.com/blog/load-balancing-nodejs-applications-with-nginx-and-docker/ 
https://superuser.openstack.org/articles/run-load-balanced-service-docker-containers-openstack/

## Deploying the project
1. 
1.5 Configure /etc/hosts