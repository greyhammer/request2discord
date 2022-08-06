# request2discord

## Background 
request2discord is a simple go app, that runs a simple webserver in port 8080 and upon receiving any type of request, then dumps that request data to a discord webhook. 
This project was created when moving services to a new server or hosting provider, and wanting to monitor the old one for any requests before shutting down. 


## Usage
Simple docker run: `docker run -d -p 8080:8080 kjjanak/request2discord`

## Contributing Guide
Contributions Welcome
Please use a feature branch and then open a MR/PR with any changes you wanna make. 

### Todo 
* Make discord url and bot name docker env vars
* Tests ? 
* Better request dump data, different module or something
* Better formatting of discord message


