# GMail Service

This is a simple service (written in Go) that is responsible to send an e-mail via Google Mails SMTP server. For communication, this service implements a REST interface.

## Configuration

Create a `.env` file in the root directory of this project, that contains the following environment variables:

```
SERVICE_ADDRESS=:8080

GMAIL_HOST=smtp.gmail.com
GMAIL_PORT=587
GMAIL_MAIL=yourmail@gmail.com
GMAIL_PASSWORD=APP_PASSWORD
```

You can find an explanation on how to create app passwords at https://support.google.com/mail/answer/185833.

## Build and Run

The first step to build the image of this service. This can be either done with Docker or Padman.

```
$ podman build -t localhost/present-roulette/gmail-service .
```

Next, it is possible to run the service using the following command:

```
$ podman run -d -p 8080:8080 --rm --env-file .env localhost/present-roulette/gmail-service
```

It is important to map the exact same OS port to the container port.

## REST Endpoints

-   Endpoint: `/mail/send` \
    Method: POST \
    Content-Type: `application/json` \
    Fields: recipient, subject, message

## Send Requests

```
$ curl -X POST http://localhost:8080/mail/send \
  -H 'Content-Type: application/json' \
  -d '{"recipient":"r@domain.tld","subject":"Test Subject","message":"This is a sample text"}'
```
