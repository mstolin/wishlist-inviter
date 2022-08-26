# User-Endpoint

The User-Endpoint serves as the gateway for all client requests. It communicates
with the underlying business services.

# Development

## Configuration

The following environment variables are required:

```
SERVICE_ADDRESS=:8080
JWT_SIGN_KEY=SUPER_SECRET
USER_SERVICE=http://localhost:8071
MAIL_SERVICE=http://localhost:8021
SCRAPPER_FACADE=http://localhost:8022
```

## Run

For development use the following command from the root directory:

```
$ ADDRESS=:8080 \
  USER_SERVICE=http://localhost:8071 \
  MAIL_SERVICE=http://localhost:8021 \
  SCRAPPER_FACADE=http://localhost:8022 \
  JWT_SIGN_KEY=SUPER_SECRET \
  go run ./user-endpoint
```

## Build and Run using Podman

To build this service use `podman build`:

```
$ podman build -t localhost/wishlist-inviter/user-endpoint .
```

Next, we can start a container using `podman run`. In this example, an
environment file was created containing all the variables introduced in
[Configuration](#configuration).

```
$ podman run -d --rm \
  -p 8080:8080 \
  --env-file .env \
  localhost/wishlist-inviter/user-endpoint
```

Another way is to use `podman-compose`. For that, follow the instructions from
[Set up](../README.md#set-up).

Use this command to only run the User-Endpoint. By default no ports are exposed
to the outside.

```
$ sudo podman-compose up user-endpoint
```
