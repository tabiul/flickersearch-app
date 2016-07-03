# A Simple Flicker Search App

The webApp allows the user to search for a specific photo available in Flicker(https://www.flickr.com/) via a keyword search

## Technologies

### why Go

 * cross-platform programming language that allows creation of REST API without the need of any additional framework
 * concurrency support is build into the language

### why AngularJS

 * easy framework to build SPA web application. No need to learn new language to use like __REACT__
 * no need of additional tooling for compilation

## Requirements

   * Go 1.6 or later
   * Golint (https://github.com/golang/lint). Make sure `golint` is added to __PATH__. This is optional
   * Python 2.7 (used for building)

## Setup

   * download Go (https://golang.org/dl/)
   * setup `GOROOT` environment variable (refers to the above link for further details for os specific instruction)
   * add `$GOROOT/bin` to __PATH__ environment variables
   * clone this repo
      * git clone https://github.com/tabiul/flickersearch-app.git
   * setup `GOPATH` environment variable (refers to location where the project is cloned, let say you cloned the project in the directory c:\flickersearch-app then this will be the value of the variable)

## Build
     python build.py

## Test
     python build.py --test

## Clean
     python build.py --clean

## Running

### Linux

navigate to folder `bin/linux_amd64`

    ./flickersearch -apiKey <apiKey> -port 8080 -webapp <path to webapp folder that is found in the root folder>

### Windows

navigate to folder `bin\windows_386`

    flickersearch.exe -apiKey <apiKey> -port 8080 -webapp <path to webapp folder that is found in the root folder>

The webapp assumes port ___8080___. Should you decide to change the port then you will need to update the js scripts accordingly

## Features

### WebApp

   * Create an account to perform search. A default account ___admin___ with password ___admin___ is available (it is hardcoded, not a good idea but good for quick testing as the users are stored in memory)
   * Enter the preferred search criteria and click on the __search__ button. The last 10 searches are available as autocomplete
   * A listing of 5 thumbnail photo will be shown. Click on any thumbnail to view a larger photo
   * Click on __<<__ and __>>__ to navigate to previous and next page
   * Click ___logout___ to logout


### Server

Following REST API are available

## image

search for images

    /image?search=<search criteria>&page=<page number>&username=<username>

   * __username__ is optional. provide this if you want to keep track of history of specific user searches
   * __page__ is optional. provide this if you want specific page


## history

retrieve user search history. system keep tracks of only last 10 searches

    /history?username=<username>

if there are any history for the user then it will be returned else HTTP status __No Content__ will be response

## user

create new user

    /user

__POST__ keys

   * username
   * password

username and password is stored in a `map` for simplicity. In addition the password is stored in __base64__ encoding which is also terrible idea

### improvements

   * store username and password in proper db
   * use bcrypt with salt for password storage


## authentication

authenticate user

    /authenticate

__POST__ keys

   * username
   * password


## TODO

   * add server side authentication using token
   * add db for persistence
   * improve the ui (terrible at the moment)

