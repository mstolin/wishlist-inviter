# Amazon Scrapper

This is a web scraper written in Python, that simply scraps a wishlist from 
Amazon.

# Development

## Configuration

Tu run properly, the service requires the following environment variables:

```
HOST=localhost
PORT=8080
JWT_SIGN_KEY=SUPER_SECRET
AMAZON_URL=http://localhost:8041
```

## Set up

This project uses [Poetry](https://python-poetry.org/) for dependency
management. TO instell dependencies, run the following command:

```
$ poetry install
```

## Run using Poetry

Run this project for development purposes using Poetry:

```
$ HOST=localhost PORT=8042 AMAZON_URL=http://localhost:8041 JWT_SIGN_KEY=SECRET poetry run start
```

## Build and Run using Podman

First you can build the image using the following command:

```
$ sudo podman build -t localhost/wishlist-inviter/amazon-scrapper .
```

Next, it is possible to run the service with `podman run`. In this example, an
environment file was created containing all the variables introduced in
[Configuration](#configuration).

```
$ sudo podman run -d --rm \
    -p 8080:8080 \
    --env-file .env \
    localhost/wishlist-inviter/amazon-scrapper
```

Another way is to use `podman-compose`. For that, follow the instructions from
[Set up](../README.md#set-up).

Use this command to only run the Amazon-Adapter. By default no ports are exposed
to the outside.

```
$ sudo podman-compose up amazon-adapter
```
