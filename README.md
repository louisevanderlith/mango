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