# Scrapper-Facade

The Scrapper-Facade abstracts the usage of all vendor-specific adapter services
like the Amazon-Adapter.

# Development

## Configuration

The following environment variables are required:

```
SERVICE_ADDRESS=:8022
JWT_SIGN_KEY=SUPER_SECRET
AMAZON_SCRAPPER=http://localhost:8042
```

## Run

For development use the following command from the root directory:

```
$ SERVICE_ADDRESS=:8022 \
  JWT_SIGN_KEY=SUPER_SECRET \
  AMAZON_SCRAPPER=http://localhost:8042 \
  JWT_SIGN_KEY=SUPER_SECRET \
  go run ./scrapper-facade
```

## Build and Run using Podman

To build this service use `podman build`:

```
$ podman build -t localhost/wishlist-inviter/scrapper-facade .
```

Next, we can start a container using `podman run`. In this example, an
environment file was created containing all the variables introduced in
[Configuration](#configuration).

```
$ podman run -d --rm \
  -p 8080:8080 \
  --env-file .env \
  localhost/wishlist-inviter/scrapper-facade
```

Another way is to use `podman-compose`. For that, follow the instructions from
[Set up](../README.md#set-up).

Use this command to only run the Scrapper-Facade. By default no ports are
exposed to the outside.

```
$ sudo podman-compose up scrapper-facade
```
