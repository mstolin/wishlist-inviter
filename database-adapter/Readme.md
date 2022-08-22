# Database Adapter

This adapter is used to abstract the usage of a PostgresDB. It will perform
CRUD operations

# Development

## Configuration

The following environment variables are:

```
ADDRESS=:8080
JWT_SIGN_KEY=SUPER_SECRET
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin123
DB_NAME=wishlist-inviter
```

## Run

For development purposes, this service can be started using the following 
command from the root directory:

```
$ ADDRESS=:8061 \
  JWT_SIGN_KEY=SUPER_SECRET \
  DB_HOST=localhost \
  DB_PORT=5432 \
  DB_USER=admin \
  DB_PASSWORD=admin123 \
  DB_NAME=wishlist-inviter \
  go run ./database-adapter
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
  localhost/wishlist-inviter/database-adapter
```

Another way is to use `podman-compose`. For that, follow the instructions from
[Set up](../README.md#set-up).

Use this command to only run the Amazon-Adapter. By default no ports are exposed
to the outside.

```
$ sudo podman-compose up database-adapter
```
