# Mail Adapter

This is a simple adapter service (written in Go), that is responsible to send a POST request to the given GMail-Service

## Configuration

Create a `.env` file in the root directory of this project, that contains the following environment variables:

```
ADDRESS=:8021 DATABASE_ADAPTER=http://localhost:8061 GMAIL_ADAPTER=http://localhost:8043 SENDER_MAIL=marcelstolin@gmail.com go run ./mail-service
```

## Build and Run

The first step to build the image of this service. This can be either done with Docker or Padman.

```
$ podman build -t localhost/wishlist-inviter/mail-adapter .
```

Next, it is possible to run the service using the following command:

```
$ podman run -d -p 8080:8080 --rm --env-file .env localhost/wishlist-inviter/mail-adapter
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
