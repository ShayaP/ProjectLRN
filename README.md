# Introduction: 
LRN is a tutoring app that connects those who want to learn further about academic subjects with those who want to teach others in academic subjects. 
The goal of the app is to connect these users with each other so they can go offline and start LRN’ing. 
This app only suggests other tutors and tutees, it is up to the users to connect to each other and see if they are a good fit for each other

# Login Credentials: 
There are several accounts that were created for testing purposes:
Username:  lrnc99913@gmail.com 
Password: LRNsystem123

Username: test1234btefla@gmail.com 
Password: f65hdsr3$

Username: test988test@gmail.com
Password: butterskin4%

Username: test988atutor@gmail.com 
Password: butterskin4%

Username: test988bobbytutor@gmail.com
Password: butterskin4%

# Requirements: 
This is a web application and the user must have:
Internet Connection
A web browser such as Chrome or Firefox

To run the live application, visit
https://www.lrn-cse110.me

# Dev Requirements
This app requires the latest versions of:
Golang (V. 1.11)
Sqlite3
Buffalo
Goth
Google metadata

# Installation Instruction: 
In order to install the requirements follow the steps below:
1. Install golang and SQLite3:
    1.1. https://golang.org/doc/install
    1.2. http://www.sqlitetutorial.net/download-install-sqlite/ 
2. Install Buffalo w/ SQLite3
    2.1. Run “go get -u -v -tags sqlite github.com/gobuffalo/buffalo/buffalo”
    2.2. Run “go get -u -v -tags sqlite github.com/gobuffalo/buffalo-pop”
3. Install dependencies
    3.1. Run “go get github.com/markbates/goth”
    3.2. Run “go get cloud.google.com/go/compute/metadata”


In order to run the app follow the commands below:
Run “cd ~/go/src/github.com/cileonard/”
Run “buffalo new lrn-test --db-type sqlite3”
If ‘Error: exit status 243’ run ‘sudo npm install -g yarn’
Git clone https://github.com/CILeonard/lrn.git
Mv lrn-test/node_modules lrn/
Mv lrn-test/database.yml lrn/
Run “cd lrn/”
Run “buffalo dev”


# Known Bugs:
If you get a "401" no matching session found, clear your cookies for this website. 


# Contacts:

Technical Support: 
Conner Leonard: Software Architect    cileonard08@gmail.com     
Other: 
Shaya Parsa:  Software Development Lead     sparsa@ucsd.edu

# FAQs:

I get 401 error saying the session could not be found when logging in:
clear cookies for localhost and for 127.0.0.1
then access the site via 127.0.0.1:3000

When I do buffalo dev the app does not start running:
Run “buffalo-pop pop drop”
Run “buffalo-pop pop create -a”
Run “buffalo-pop pop migrate up”



