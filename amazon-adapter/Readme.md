# Amazon Scrapper

This is a web scrapper, written in Python, that simply scraps a wishlist from Amazon.

## Configuration

Create a `.env` file in the root directory of this project, that contains the following environment variables:

```
AMAZON_HOST=http://localhost:8080
```

## Development

For development purposes, run the following using [Poetry](https://python-poetry.org/).

```
$ poetry install
$ HOST=localhost PORT=8080 poetry run start
```

## Build and Run

The first step to build the image of this service. This can be either done with Docker or Padman.

```
$ podman build -t localhost/present-roulette/amazon-scrapper .
```

Next, it is possible to run the service using the following command:

```
$ podman run -d -p 8080:8080 --rm --env-file .env localhost/present-roulette/amazon-scrapper
```

It is important to map the exact same OS port to the container port.

## REST Endpoints

-   Endpoint: `/wishlist/WISHLIST_ID` \
    Method: GET

## Send Requests

```
$ curl http://localhost:8080/wishlist/194N1KF03IPTL
```
