# User Service

The User-Service is responsible to handle all user-related business tasks. It is
able to talk to the Database-Adapter.

# Development

## Configuration

The following environment variables are required:

```
SERVICE_ADDRESS=:8080
JWT_SIGN_KEY=SUPER_SECRET
DATABASE_ADAPTER=http://localhost:8061
```

## Run

For development use the following command from the root directory:

```
$ ADDRESS=:8071 \
  JWT_SIGN_KEY=SUPER_SECRET \
  DATABASE_ADAPTER=http://localhost:8061 \
  go run ./user-service
```

## Build and Run using Podman

To build this service use `podman build`:

```
$ podman build -t localhost/wishlist-inviter/user-service .
```

Next, we can start a container using `podman run`. In this example, an
environment file was created containing all the variables introduced in
[Configuration](#configuration).

```
$ podman run -d --rm \
  -p 8080:8080 \
  --env-file .env \
  localhost/wishlist-inviter/user-service
```

Another way is to use `podman-compose`. For that, follow the instructions from
[Set up](../README.md#set-up).

Use this command to only run the User-Service. By default no ports are exposed
to the outside.

```
$ sudo podman-compose up user-service
```
