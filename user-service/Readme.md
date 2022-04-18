# User Service

This is a simple adapter service (written in Go), that is responsible to send a POST request to the given GMail-Service

## Configuration

Create a `.env` file in the root directory of this project, that contains the following environment variables:

```
SERVICE_ADDRESS=:8080

DATABASE_URL=http://database-adapter:8080
```

## Build and Run

The first step to build the image of this service. This can be either done with Docker or Padman.

```
$ podman build -t localhost/present-roulette/mail-adapter .
```

Next, it is possible to run the service using the following command:

```
$ podman run -d -p 8080:8080 --rm --env-file .env localhost/present-roulette/mail-adapter
```

It is important to map the exact same OS port to the container port.

## REST Endpoints

-   Endpoint: `/mail/send/invitation` \
    Method: POST \
    Content-Type: `application/json` \
    Fields: recipient, subject, message

## Send Requests

```
$ curl -X POST http://localhost:8080/mail/send/invitation \
  -H 'Content-Type: application/json' \
  -d '{"recipient":"r@domain.tld","subject":"Test Subject","message":"This is a sample text"}'
```
