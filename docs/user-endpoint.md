FORMAT: 1A
HOST: http://user-endpoint/

# User-Endpoint

User-endpoint is the process centric API of this project. It provides an access point to the end-user
and is the only service that the user is supposed to communicate with directly.

All requests made to this service are forwarded to the _more specialized_ service.
For example all requests made to `/mail/` are forwarded to he _mail-service_.

## Authentication collection [/auth]

This endpoint is used to authenticate. The request has to be a
JSON object containing a valid user ID. This service uses JWT for
authentication. Therefore, the response is a 24h valid JWT token.

If the user is not found, the API responds with a 404 error. If
the request is invalid in general, a 400 is sent. Otherwise, a 
500 error.

### Authenticate [POST]

+ Request (application/json)

        {
            "user_id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7"
        }

+ Response 200 (application/json)

        {
            "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o"
        }

+ Response 400 (application/json)

        {
            "error": {
                "status": 400,
                "error": "Bad Request",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

## Users collection [/users]

Through this endpoint, a new user instance can be created by
sending a `POST` request. It will return the new user in a 
JSON representation if the request was successful. The 
registration/creation of a user, does not require a JWT token
for authentication.

If the request is invalid, a 400 error is send. Otherwise, 
if any error appears on the server, a 500 error is send.

### Create a new user [POST]

+ Request (application/json)

            {}

+ Response 200 (application/json)

        {
            "id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7",
            "created_at": "2022-08-07T13:54:49.964166093Z",
            "updated_at": "2022-08-07T13:54:49.964166093Z",
            "items": null
        }

+ Response 400 (application/json)

        {
            "error": {
                "status": 400,
                "error": "Bad Request",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }


## Specific user collection [/users/{userId}]

Use this endpoint to either receive information for a specific
user, through a `GET` request, or delete an existing user via
`DELETE`.

If the user is not found, a 404 error is send. Is the client
unauthorized, a 401 error is repsonded. If any other error occurs 
on the server side, a 500 error is send.

### Get a specific user [GET]

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7",
            "created_at": "2022-08-07T13:54:49.964166Z",
            "updated_at": "2022-08-07T13:54:49.964166Z",
            "items": []
        }

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }


### Delete a specific user [DELETE]

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id": "6eaade67-e087-4327-8bd5-92934baf58ed",
            "created_at": "2022-08-07T16:18:39.827566Z",
            "updated_at": "2022-08-07T16:18:39.827566Z",
            "items": []
        }

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }


## User items collection [/users/{userId}/items]

This collect provides to receive wished items for a specific user (`GET`), 
or to add new items to the wishlist (`POST`).

If the user does not exists, a 404 error is thrown. If the request for the `POST` 
request is invalid, a 400 error is thrown. Unauthenticated clients will receive 
a 401 error.For any other errors on the server side, a 500 error is thrown.

### Get items [GET]

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

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

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

### Add items [POST]

+ Request Add new item (application/json)

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

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

+ Response 200 (application/json)

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

+ Response 400 (application/json)

        {
            "error": {
                "status": 400,
                "error": "Bad Request",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

## Specific user item collection [/users/{userId}/items/{itemId}]

Via this endpoint the client can either receive information about a
specific item from a specific user (`GET`), update the specific item (`PUT`), 
or delete the item `DELETE`.

The fields that can be updated are `name`, `price`, `vendor_id`,
and `has_been_bought`. After an item has been updated, the 
system will update its meta value `updated_at`.

if the user or the item do not exist, a 404 error is thrown.
If the request for the `PUT` request is invalid, a 400
error is thrown. If the client is not authenticated, it will
receive a 401 error. Otherwise, if an error occurs on the server 
side, it will response with a 500 error.

### Update item [PUT]

+ Request Update item (application/json)

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

            {
                "name": "Hario V60 Glass Coffee Dripper",
                "price": 50.8,
                "vendor_id": "IP0OBIK4UO9AG",
                "has_been_baught": true
            }

+ Response 200 (application/json)

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

+ Response 400 (application/json)

        {
            "error": {
                "status": 400,
                "error": "Bad Request",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }


### Delete item [DELETE]

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id": 2,
            "created_at": "2022-08-07T14:04:46.329234Z",
            "updated_at": "2022-08-07T14:16:05.154862Z",
            "name": "Hario V60 Glass Coffee Dripper",
            "price": 50.8,
            "vendor": "amazon",
            "vendor_id": "IP0OBIK4UO9AG",
            "has_been_baught": true
        }

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

### Get a specific Item [GET]

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id": 2,
            "created_at": "2022-08-07T14:04:46.329234Z",
            "updated_at": "2022-08-07T14:16:05.154862Z",
            "name": "Hario V60 Glass Coffee Dripper",
            "price": 50.8,
            "vendor": "amazon",
            "vendor_id": "IP0OBIK4UO9AG",
            "has_been_baught": true
        }

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

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

The amazon collection can be used to gather information about
amazon products. Currently, only Amazon wishlists are supported.
The endpoint will return an Amazon wishlist in JSON format.

A 404 error is sent if the wanted wishlist does not exists. For
aunauthorized access, a 401 error is responded. If any error on 
the server side occurs a 500 response is sent.

### Get an Amazon Wishlist [GET]

+ Request (application/json)

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

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

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500 (application/json)

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

### Send Invitation [POST]

+ Request (application/json)

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

            {
                "recipient": "recipient@domain.tld",
                "subject": "You have been invited",
                "user_id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7",
                "items": [1, 3, 4]
            }

+ Response 200 (application/json)

        {
            "message": "mail has been sent successfully"
        }
        
+ Response 400

        {
            "error": {
                "status": 400,
                "error": "Bad Request",
                "message": "GENERIC ERROR MESSAGE"
            } 
        }

+ Response 401 (application/json)

        {
            "error": {
                "status": 401,
                "error": "Unauthorized",
                "message": "GENERIC ERROR MESSAGE"
            }
        }

+ Response 500

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }
