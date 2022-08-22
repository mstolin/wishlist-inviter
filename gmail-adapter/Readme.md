# GMail Adapter

This is an adapter service to send mails via the Google Mail SMTP server.

# Development

## Configuration

The following environment variables are required:

```
ADDRESS=:8080
JWT_SIGN_KEY=SUPER_SECRET
GMAIL_HOST=smtp.gmail.com
GMAIL_PORT=587
GMAIL_MAIL=yourmail@gmail.com
GMAIL_PASSWORD=APP_PASSWORD
```

You can find an explanation on how to create app passwords at 
https://support.google.com/mail/answer/185833.

### Run

For development purposes, this service can be started using the following 
command from the root directory:

```
$ ADDRESS=:8061 \
  JWT_SIGN_KEY=SUPER_SECRET \
  GMAIL_HOST=smtp.gmail.com \
  GMAIL_PORT=587 \
  GMAIL_MAIL=yourmail@gmail.com \
  GMAIL_PASSWORD=APP_PASSWORD \
  go run ./gmail-adapter
```

## Build and Run using Podman

Use this command to build an image of the service:

```
$ podman build -t localhost/wishlist-inviter/database-adapter .
```

Next, we can start a container using `podman run`. In this example, an
environment file was created containing all the variables introduced in
[Configuration](#configuration).

```
$ podman run -d --rm \
  -p 8080:8080 \
  --env-file .env \
  localhost/wishlist-inviter/gmail-adapter
```

Another way is to use `podman-compose`. For that, follow the instructions from
[Set up](../README.md#set-up).

Use this command to only run the Amazon-Adapter. By default no ports are exposed
to the outside.

```
$ sudo podman-compose up gmail-adapter
```
