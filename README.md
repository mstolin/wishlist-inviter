# TL;TR

# Design

## Architecture

![architecture](./docs/img/architecture.png)

## Database

![database](./docs/img/erm.png)

# Documentation

- User-Endpoint
    + [API](./docs/user-endpoint.md)
    + [Usage](./user-endpoint/Readme.md)
- User-Service
    + [API](./docs/user-service.md)
    + [Usage](./user-service/Readme.md)
- Mail-Service
    + [API](./docs/mail-service.md)
    + [Usage](./mail-service/Readme.md)
- Scrapper-Facade
    + [API](./docs/scrapper-facade.md)
    + [Usage](./scrapper-facade/Readme.md)
- Database-Adapter
    + [API](./docs/database-adapter.md)
    + [Usage](./database-adapter/Readme.md)
- GMail-Adapter
    + [API](./docs/gmail-adapter.md)
    + [Usage](./gmail-adapter/Readme.md)
- Amazon-Adapter
    + [API](./docs/amazon-adapter.md)
    + [Usage](./amazon-adapter/Readme.md)

# Usage

## Setup

### JWT

For authentication purposes, a JWT token is necessary. You can create a random 
secret at https://www.random.org/bytes/.

Save the secret at `/env/JWT.env`.

```
JWT_SIGN_KEY=THIS-IS-THE-RANDOM-SECRET
```

### Google-Mail

To send mails via the Google-Mail SMTP server an APP password is necessary. This is
explained at https://support.google.com/mail/answer/185833?hl=en. Save you GMail
address and the app password at `/env/gmail-adapter/.env`.

```
GMAIL_MAIL=yourgmail@gmail.com
GMAIL_PASSWORD=THE-APP-PASSWORD
```

## Run

This project uses [Podman](https://podman.io/) to run all services in a container.
A `docker-compose.yml` file exists in the root directory of the project, that can
be used with [podman-compose](https://github.com/containers/podman-compose) to
start the whole multi-service application. After that, the *User-Endpoint* is
available at [http://localhost:8080](http://localhost:8080).

```
$ sudo podman-compose up --build
```

You can also use podman-compose to build and run a specific collection of services.

```
$ sudo podman-compose up --build database-adapter database
```

**Important to mention**: By default only the *User-Endpoint* and *Amazon-Clone* are
accessible through the outside. If you wanto to test/access specific services using 
Podman, you have to publish its ports in the `docker-compose.yml`.

## Testing

To test this project, a [Postman](https://www.postman.com/) collection is provided at
`/Present-Roulette.postman_collection.json`.

# License

[Creative Commons Attribution 4.0 International Public License](./LICENSE)
