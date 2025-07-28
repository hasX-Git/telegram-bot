# Bank API

## Prerequisites
This bot works in pair with the myBank repository, specifically with its database, so it won't work on its own unless you create a new database for this bot.

You can get myBank repository with

    git clone https://github.com/hasX-Git/myBank.git

It is important that myBank and telegramBot are in one directory.

## How to use
### 1. Pull the repository
Pull the repository on your local machine with

    git clone https://github.com/hasX-Git/telegram-bot.git

### 2. Working with .example files
There are 3 example files: .env_example, Dockerfile_example, and docker-compose_example.yml.

These are templates for .env, Dockerfile and docker-compose.yml respectively.

Wherever there is <ins>#change</ins> comment, you can change those to any preferred name.

Wherever there is <ins>#AP</ins>, which stands for App Port, it must be the same for all ports with #AP comment. For example, if in .env file the port for #AP is 5432, it must be 5432 for #AP in Dockerfile and docker-compose. Same with <ins>#DP</ins>, Database Port

### 3. .env file

To get API key for AI, follow the link:

    https://aistudio.google.com/apikey?_gl=1*1xg6gsd*_ga*MjY2ODYxMTg5LjE3NTI1NzcyNzA.*_ga_P1DBVKWT6V*czE3NTM3MDg0OTkkbzckZzEkdDE3NTM3MDg1MzgkajIxJGwwJGg3MDcwMDg0ODk.

and generate a key. You need to insert it into **GEMINI_API_KEY** field.

To get token for telegram bot, use the telegram BotFather bot:

    @BotFather

and create bot there. It will generate token for your bot, which you need to insert into **TOKEN** field

### 4. Running program

First run the myBank container with

    docker compose up --build -d

Then run tgApp

    docker compose up --build -d
