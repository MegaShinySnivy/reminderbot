# ReminderBot
This is a simple bot designed to remind users of actions that need to take place on a certain date.
## Setup instructions
### Requirements
This bot assumes you have a container runtime and cron installed on your system of choice
### Setup
Using your CRI, set the env vars BOT_TOKEN, CHANNEL_ID and MESSAGE to your requirements. Then let the bot run via cron set to the time and day you want a reminder. 
## Why?
I'm very bad at remembering to do specific things, so having a bot ping me on those dates would be useful
## Why cron tho?
Because why have the program handle that when I have a perfectly good program to handle that for me? Plus this was designed to run as a stateless kubernetes pod, so why bother with the extra overhead?
## Can you really call this a bot?
¯\\\_(ツ)_/¯