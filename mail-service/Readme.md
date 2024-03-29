# Mail-Service

The Mail-Service is a business services that is responsible to handle all mail-
related tasks.

# Development

## Configuration

The following environment variables are required:

```
ADDRESS=:8021
JWT_SIGN_KEY=SUPER_SECRET
GMAIL_ADAPTER=http://localhost:8043
DATABASE_ADAPTER=http://localhost:8061
SENDER_MAIL=sender@domain.tld
```

## Run

For development use the following command from the root directory:

```
$ ADDRESS=:8021 \
  JWT_SIGN_KEY=SUPER_SECRET \
  GMAIL_ADAPTER=http://localhost:8043 \
  DATABASE_ADAPTER=http://localhost:8061 \
  SENDER_MAIL=sender@domain.tld \
  go run ./mail-service
```

## Build and Run using Podman

To build this service use `podman build`:

```
$ podman build -t localhost/wishlist-inviter/mail-service .
```

Next, we can start a container using `podman run`. In this example, an
environment file was created containing all the variables introduced in
[Configuration](#configuration).

```
$ podman run -d --rm \
  -p 8080:8080 \
  --env-file .env \
  localhost/wishlist-inviter/mail-service
```

Another way is to use `podman-compose`. For that, follow the instructions from
[Set up](../README.md#set-up).

Use this command to only run the Mail-Service. By default no ports are exposed
to the outside.

```
$ sudo podman-compose up mail-service
```
