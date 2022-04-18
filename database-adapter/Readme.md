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

## REST Endpoints

### `/wishlist`

Creates a new wishlist

- Method: `POST`
- Content-Type: `application/json`

```
$ curl -X POST http://localhost:8080/wishlist \
  -H 'Content-Type: application/json' \
  -d '{"recipient":"r@domain.tld","subject":"Test Subject","message":"This is a sample text"}'
```

### `/wishlist/WISHLIST_ID`

Returns the wishlist for the requested ID.

- Method: `GET`

```
$ curl http://localhost:8080/wishlist/1
```

### `/wishlist/WISHLIST_ID`

Updates the requested wishlist with the given data.

- Method: `PUT`
- Content-Type: `application/json`

```
$ curl -X PUT http://localhost:8080/wishlist/1 \
  -H 'Content-Type: application/json' \
  -d '{"recipient":"r@domain.tld","subject":"Test Subject","message":"This is a sample text"}'
```

### `/wishlist/WISHLIST_ID`

Deletes the requested wishlist.

- Method: `DELETE`

```
$ curl -X DELETE http://localhost:8080/wishlist/1
```

### `/wishlist/WISHLIST_ID/item`

Adds a new item to the given wishlist

- Method: `POST`
- Content-Type: `application/json`

```
$ curl -X POST http://localhost:8080/wishlist/1/item \
  -H 'Content-Type: application/json' \
  -d '{"recipient":"r@domain.tld","subject":"Test Subject","message":"This is a sample text"}'
```

### `/wishlist/WISHLIST_ID/item/ITEM_ID`

Returns the requested item from the wishlist.

- Method: `GET`

```
# Return item with ID 3 from wishlist with ITEM 1
$ curl http://localhost:8080/wishlist/1/item/3
```

### `/wishlist/WISHLIST_ID/item/ITEM_ID`

Updates the requested item with the given data.

- Method: `PUT`
- Content-Type: `application/json`

```
$ curl -X PUT http://localhost:8080/wishlist/1/item/3 \
  -H 'Content-Type: application/json' \
  -d '{"recipient":"r@domain.tld","subject":"Test Subject","message":"This is a sample text"}'
```

### `/wishlist/WISHLIST_ID/item/ITEM_ID`

Deletes the requested item from the wishlist.

- Method: `DELETE`

```
$ curl -X DELETE http://localhost:8080/wishlist/1/item/3
```

### `/user`

Creates a new user.

```
$ curl -X POST http://localhost:8080/user \
  -H 'Content-Type: application/json' \
  -d '{}'
```

### `/user/USER_ID`

Returns the data of the requested user.

```
$ curl http://localhost:8080/user/2b7ec5fd-2623-483a-8b15-f376ff2297c6
```

### `/user/USER_ID`

Deletes the requested user.

- Method: `DELETE`

```
$ curl -X DELETE http://localhost:8080/user/2b7ec5fd-2623-483a-8b15-f376ff2297c6
```
