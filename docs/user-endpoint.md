FORMAT: 1A
HOST: http://user-endpoint/

# User-Endpoint

User-endpoint is the process centric API of this project. It provides an access point to the end-user
and is the only service that the user is supposed to communicate with directly.

All requests made to this service are forwarded to the _more specialized_ service.
For example all requests made to `/mail/` are forwarded to he _mail-service_.

## User Collection [/users/]

This collection is responsible to handle all user related requests.

Using a `POST` request a new user instance is being created.
A user consist only of a ID which is a _UUIDv4_. This is
completely random generated. An empty JSON request is 
enough to create a new user instance.

### Create a new user [POST]

-   Request (application/json)

            {}

-   Response 200 (application/json)

            {
                "id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7",
                "created_at": "2022-08-07T13:54:49.964166093Z",
                "updated_at": "2022-08-07T13:54:49.964166093Z",
                "items": null
            }

-   Response 400 (application/json)

            {
                "error": {
                    "status": 400,
                    "error": "Bad Request",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 500 (application/json)

            {
                "error": {
                    "status": 500,
                    "error": "Internal Server Error",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

## Specific user collection [/users/{userId}/]

This endpoint returns all infos about a user.
Its ID, some meta information, and all wished items.

### Get specific user [GET]

-   Response 200 (application/json)

            {
                "id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7",
                "created_at": "2022-08-07T13:54:49.964166Z",
                "updated_at": "2022-08-07T13:54:49.964166Z",
                "items": []
            }

-   Response 404 (application/json)

            {
                "error": {
                    "status": 404,
                    "error": "Not Found",
                    "message": "user with id I_DO_NOT_EXIST not found"
                }
            }

-   Response 500 (application/json)

            {
                "error": {
                    "status": 500,
                    "error": "Internal Server Error",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

## User items collection [/users/{userId}/items/]

Use this endpoint to request all items of a specific user or to
add a list of one or multiple items to the user.

### Get items [GET]

-   Response 200 (application/json)

            [
                {
                    "id": 1,
                    "created_at": "2022-08-07T14:04:46.329234Z",
                    "updated_at": "2022-08-07T14:04:46.329234Z",
                    "name": "Hario 400 ml Olive Wood New Coffee Server, Transparent",
                    "price": 54.84,
                    "vendor": "amazon",
                    "vendor_id": "I3UCMMATCW0ATV",
                    "has_been_baught": false
                },
                {
                    "id": 2,
                    "created_at": "2022-08-07T14:04:46.329234Z",
                    "updated_at": "2022-08-07T14:04:46.329234Z",
                    "name": "Hario V60 Glass Coffee Dripper",
                    "price": 50.8,
                    "vendor": "amazon",
                    "vendor_id": "IP0OBIK4UO9AG",
                    "has_been_baught": false
                }
            ]

-   Response 404 (application/json)

            {
                "error": {
                    "status": 404,
                    "error": "Not Found",
                    "message": "The requested resource is not available."
                }
            }

-   Response 500 (application/json)

            {
                "error": {
                    "status": 500,
                    "error": "Internal Server Error",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

### Add items [POST]

-   Request (application/json)

            [
                {
                    "name": "Hario 400 ml Olive Wood New Coffee Server, Transparent",
                    "price": 54.84,
                    "vendor_id": "I3UCMMATCW0ATV",
                    "vendor": "amazon"
                },
                {
                    "name": "Hario V60 Glass Coffee Dripper",
                    "price": 50.8,
                    "vendor": "amazon",
                    "vendor_id": "IP0OBIK4UO9AG"
                }
            ]

-   Response 200 (application/json)

            [
                {
                    "id": 1,
                    "created_at": "2022-08-07T14:04:46.329234Z",
                    "updated_at": "2022-08-07T14:04:46.329234Z",
                    "name": "Hario 400 ml Olive Wood New Coffee Server, Transparent",
                    "price": 54.84,
                    "vendor": "amazon",
                    "vendor_id": "I3UCMMATCW0ATV",
                    "has_been_baught": false
                },
                {
                    "id": 2,
                    "created_at": "2022-08-07T14:04:46.329234Z",
                    "updated_at": "2022-08-07T14:04:46.329234Z",
                    "name": "Hario V60 Glass Coffee Dripper",
                    "price": 50.8,
                    "vendor": "amazon",
                    "vendor_id": "IP0OBIK4UO9AG",
                    "has_been_baught": false
                }
            ]

-   Response 400 (application/json)

            {
                "error": {
                    "status": 400,
                    "error": "Bad Request",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 404 (application/json)

            {
                "error": {
                    "status": 404,
                    "error": "Not Found",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 500 (application/json)

            {
                "error": {
                    "status": 500,
                    "error": "Internal Server Error",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

## Single user item collection [/users/{userId}/items/{itemId}]

This endpoint is used to receive information, update, or delete a
specific item of a user.

A `PUT` request can be used to update an existing item.
The fields that can be updated are `name`, `price`, `vendor_id`,
and `has_been_bought`. After an item has been updated, the 
system will update its meta value `updated_at`.

### Update item [PUT]

-   Request (application/json)

            {
                "name": "Hario V60 Glass Coffee Dripper",
                "price": 50.8,
                "vendor_id": "IP0OBIK4UO9AG",
                "has_been_baught": true
            }

-   Response 200 (application/json)

            {
                "id": 2,
                "created_at": "2022-08-07T14:04:46.329234Z",
                "updated_at": "2022-08-07T14:14:39.86491697Z",
                "name": "Hario V60 Glass Coffee Dripper",
                "price": 50.8,
                "vendor": "amazon",
                "vendor_id": "IP0OBIK4UO9AG",
                "has_been_baught": true
            }

-   Response 400 (application/json)

            {
                "error": {
                    "status": 400,
                    "error": "Bad Request",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 404 (application/json)

            {
                "error": {
                    "status": 404,
                    "error": "Not Found",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 500 (application/json)

            {
                "error": {
                    "status": 500,
                    "error": "Internal Server Error",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

# Wishlist Collection [/items/]

Through this collection, products from different vendors
can be scapped. Currently only amazon wishlists are available.

## Amazon wishlist [/items/amazon/wishlists/{wishlistId}]

This endpoint is used to get a structured JSON representation
of a specific Amazon wishlist.

A 404 error is sent if the wanted wishlist does not exists.
If any error on the server side occurs a 500 response is sent.

### Get an Amazon Wishlist [GET]

-   Response 200 (application/json)

            {
                "id": "3IGO1OHMRSUUM",
                "vendor": "amazon",
                "name": "Games",
                "items": [
                    {
                        "id": "I9RYNJAL9GAYA",
                        "name": "Hades [Nintendo Switch]",
                        "price": 64.99,
                        "vendor": "amazon"
                    },
                    {
                        "id": "I7DF1D943PWDV",
                        "name": "The Witcher 3: Wild Hunt - Complete Edition - [Nintendo Switch]",
                        "price": 52.23,
                        "vendor": "amazon"
                    },
                    {
                        "id": "ITTLSQBBXRYE3",
                        "name": "The Legend of Zelda: Breath of the Wild [Nintendo Switch]",
                        "price": 54.99,
                        "vendor": "amazon"
                    },
                    {
                        "id": "I2NU97P58RCZFA",
                        "name": "The Legend of Zelda: Link's Awakening [Nintendo Switch]",
                        "price": 45.99,
                        "vendor": "amazon"
                    },
                    {
                        "id": "I37NVYF3F7GSV1",
                        "name": "Nintendo Luigi's Mansion 3 - [Nintendo Switch]",
                        "price": 47.94,
                        "vendor": "amazon"
                    }
                ]
            }

-   Response 404 (application/json)

            {
                "error": {
                    "status": 404,
                    "error": "Not Found",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 500 (application/json)

            {
                "error": {
                    "status": 500,
                    "error": "Internal Server Error",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

# Mail Collection [/mail/]

Through this collection, it is possible to send an e-mail via Google-Mail.

## Invitations [/mail/invitations]

Use this endpoint to send an invitation e-mail to a specific recipient.
It is important to add the wanted items, the user want the recipient to buy.

### Send an invitation [POST]

-   Request (application/json)

            {
                "recipient": "mail@domain.tld",
                "user_id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7",
                "items": [1, 2]
            }

-   Response 200 (application/json)

            {
                "recipient": "marcelpascal.stolin@studenti.unitn.it",
                "body": "To: <marcelpascal.stolin@studenti.unitn.it>\nFrom: <marcelstolin@gmail.com>\nSubject: Someone has invitited you to his wishlist\nHi,\nyou have been invited to buy the following items:\n\n  - Hario 400 ml Olive Wood New Coffee Server, Transparent, 54.84€ (https://www.amazon.com/dp/I3UCMMATCW0ATV/)\n  - Hario V60 Glass Coffee Dripper, 50.80€ (https://www.amazon.com/dp/IP0OBIK4UO9AG/)\n\nCheers!"
            }

-   Response 400 (application/json)

            {
                "error": {
                    "status": 400,
                    "error": "Bad Request",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 404 (application/json)

            {
                "error": {
                    "status": 404,
                    "error": "Not Found",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }

-   Response 500 (application/json)

            {
                "error": {
                    "status": 500,
                    "error": "Internal Server Error",
                    "message": "GENERIC ERROR MESSAGE"
                }
            }
