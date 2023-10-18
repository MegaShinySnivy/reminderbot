# ReminderBot
This is a test bot designed to remind users of actions that need to take place on a certain date.
## Setup instructions
### Requirements
This bot assumes you have a container runtime and cron installed on your system of choice
### Setup
Using cron, have it set up to run the container every day at the time you would like the reminder to go off, and set the env vars BOT_TOKEN, CHANNEL_ID and MESSAGE to your requirements. Then let the bot do its thing.
## Current limits
* This currently only goes off on the 1st. Soon I will have an option to pick the day
* It currently relies on cron to set the time at which it goes off. This is also on the roadmap for addition