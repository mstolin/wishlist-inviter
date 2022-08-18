# TL;TR

The goal of this project is to present the learning of the course Service Design
and Engineering. The requirements are to develop a project that consists of 
multiple microservices and use techniques and technologies presented during the
course. The idea of the project is an application where the user can create
wishlists, and invites other persons to buy an item from that wishlist via 
email. The content of a wishlist is based on a vendor wishlist (e.g. 
Amazon.com). To make a wishlist available, the application extracts the content 
from a given wishlist URL and saves it in a structured format to the database.

# Design

This chapter describes the architecture and database design.

## Architecture

![architecture](./docs/img/architecture.png)

The figure illustrates the architecture of the application. It consists of four 
different layers, each serving a different purpose. Each layer (except the Data 
Layer) includes services that are designed to solve a specific task and be 
self-describing and self-continued. The architecture is designed as a 
choreography. No service takes responsibility for the whole process and services 
are aware of collaborating.

## Database

![database](./docs/img/erm.png)

Figure 2 illustrates the UML diagram of the database model. There exist two 
models, a user, and an item. A single user owns multiple items. To simplify the 
database model, an item can only be owned by a single user. A user only exists 
of an ID that identifies a user and is unique for each one. The User-Service 
creates new user instances and randomly creates a UUIDv4 for a new user. An item 
is identified by its ID. Furthermore, the item model also saves the vendor ID 
which refers to the ID that was given by the vendor. Important to mention is 
also the has been bought property, the reason why an item can only be owned by a
single user. It represents if an item has been bought by a buyer or not.

# Documentation

-   User-Endpoint
    -   [API](./docs/user-endpoint.md)
    -   [Usage](./user-endpoint/Readme.md)
-   User-Service
    -   [API](./docs/user-service.md)
    -   [Usage](./user-service/Readme.md)
-   Mail-Service
    -   [API](./docs/mail-service.md)
    -   [Usage](./mail-service/Readme.md)
-   Scrapper-Facade
    -   [API](./docs/scrapper-facade.md)
    -   [Usage](./scrapper-facade/Readme.md)
-   Database-Adapter
    -   [API](./docs/database-adapter.md)
    -   [Usage](./database-adapter/Readme.md)
-   GMail-Adapter
    -   [API](./docs/gmail-adapter.md)
    -   [Usage](./gmail-adapter/Readme.md)
-   Amazon-Adapter
    -   [API](./docs/amazon-adapter.md)
    -   [Usage](./amazon-adapter/Readme.md)

# Usage

## Set up

### JWT

For authentication purposes, a JWT token is necessary. You can create a random
secret at https://www.random.org/bytes/.

Save the secret at `/env/JWT.env`.

```
JWT_SIGN_KEY=THIS-IS-THE-RANDOM-SECRET
```

### Google-Mail

To send mails via the Google-Mail SMTP server an app password is necessary. This 
is explained at https://support.google.com/mail/answer/185833?hl=en. Save your 
GMail address and the app password at `/env/gmail-adapter/.env`.

```
GMAIL_MAIL=yourgmail@gmail.com
GMAIL_PASSWORD=THE-APP-PASSWORD
```

## Run

This project uses [Podman](https://podman.io/) to run all services in a 
container. A `docker-compose.yml` file exists in the root directory of the 
project, that can be used with 
[podman-compose](https://github.com/containers/podman-compose) to start the 
whole multi-service application. After that, the _User-Endpoint_ is available at 
[http://localhost:8080](http://localhost:8080).

```
$ sudo podman-compose up --build
```

You can also use podman-compose to build and run a specific collection of 
services.

```
$ sudo podman-compose up --build database-adapter database
```

**Important to mention**: By default only the _User-Endpoint_ and _Amazon-Clone_
are accessible through the outside. If you wanto to test/access specific 
services using Podman, you have to publish its ports in the 
`docker-compose.yml`.

## Testing

To test this project, a [Postman](https://www.postman.com/) collection is 
provided at `/Present-Roulette.postman_collection.json`.

# License

[Creative Commons Attribution 4.0 International Public License](./LICENSE)
