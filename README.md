# request2discord

![Build Status](https://github.com/greyhammer/request2discord/actions/workflows/dockerimage.yml/badge.svg)

## Background 
request2discord is a simple go app, that runs a simple webserver on port 8080 and upon receiving any type of request, then dumps that request data to a discord webhook. 

This project was created when moving services to a new server or hosting provider, and wanting to monitor the old one for any requests before shutting down. 


## Usage

Create a Discord Webhook - [Instructions](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks)

Simple docker run: `docker run -d -e DISCORD_WEBHOOK_URL=<webhook url> -e DISCORD_WEBHOOK_BOTNAME=<name for messages to appear from> -p 8080:8080 greyhammer/request2discord`

## Contributing Guide
Contributions Welcome

Please use a feature branch and then open a MR/PR with any changes you wanna make. 

### Todo 
* Add versioning and tags to docker image
* Verify main branch is protected
* Tests ? 
* Better request dump data, different module or something
* Better formatting of discord message
* Add option for saving to db ? 

