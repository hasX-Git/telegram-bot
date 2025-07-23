# Bank API

## Prerequisite
This bot works in pair with the myBank repository, specifically with its database, so it won't work on its own unless you create a new database for this bot.

You can get myBank repository with

    git clone https://github.com/hasX-Git/myBank.git

It is important that myBank and telegramBot are in one directory.

## How to use
### 1. Pull the repository
Pull the repository on your local machine with

    git clone https://github.com/hasX-Git/telegram-bot.git

### 2. Working with .example files
There are 3 example files: .env.example, Dockerfile.Example, and docker-compose_example.yml.

These are templates for .env, Dockerfile and docker-compose.yml respectively.

Wherever there is <ins>#change</ins> comment, you can change those to any preferred name.

Wherever there is <ins>#AP</ins>, which stands for App Port, it must be the same for all ports with #AP comment. For example, if in .env file the port for #AP is 5432, it must be 5432 for #AP in Dockerfile and docker-compose. Same with <ins>#DP</ins>, Database Port

### 3. Running program
Run the following line

    docker compose up --build