# Database Adapter

This is a simple adapter service (written in Go), to access a Postgresql database.

## Configuration

Create a `.env` file in the root directory of this project, that contains the following environment variables:

```
SERVICE_ADDRESS=:8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin123
DB_NAME=present-roulette
```

## Development

For development purposes, just run the `run.sh` script.
If needed, set the environment varibles accordingly.

```
$ ./run.sh
# or from root directory
$ DB_HOST=localhost DB_PORT=5432 DB_USER=admin DB_PASSWORD=admin123 DB_NAME=present-roulette go run ./database-adapter
```

## Build and Run

The first step to build the image of this service. This can be either done with Docker or Padman.

```
$ podman build -t localhost/present-roulette/database-adapter .
```

Next, it is possible to run the service using the following command:

```
$ podman run -d -p 8080:8080 --rm --env-file .env localhost/present-roulette/database-adapter
```

It is important to map the exact same OS port to the container port.

## REST Endpoint

```
$ curl -X POST http://localhost:8080/users \
  -H 'Content-Type: application/json' \
  -d '{}'
```

```
$ curl http://localhost:8080/users/c8dc276b-176a-4468-b6b6-4af63f1b98f1
```

```
$ curl -X DELETE http://localhost:8080/users/c8dc276b-176a-4468-b6b6-4af63f1b98f1
```

```
$ curl http://localhost:8080/users/7dd27df6-5af3-4968-92cd-ad28b6e644f6/items
```

```
$ curl -X POST http://localhost:8080/users/7dd27df6-5af3-4968-92cd-ad28b6e644f6/items \
  -H 'Content-Type: application/json' \
  -d '[{"name": "Test Item 1", "price": 4.89, "vendor": "amazon", "vendor_id": "SOME_ID"}, {"name": "Test Item 2", "price": 4.89, "vendor": "amazon", "vendor_id": "SOME_ID"}]'
```

```
$ curl http://localhost:8080/users/7dd27df6-5af3-4968-92cd-ad28b6e644f6/items/14
```

```
$ curl -X DELETE http://localhost:8080/users/7dd27df6-5af3-4968-92cd-ad28b6e644f6/items/14
```
