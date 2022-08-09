FORMAT: 1A
HOST: http://user-service/

# User-Service

The user-Service is responsible to handle _User_ related tasks.
This includes, creating and deleting a user using the _Database-Adapter_,
as well as handling all requests concerning _User_ _Items_.

## Users collection [/users]

Through this endpoint, a new user instance can be created by
sending a `POST` request. It will return the new user in a 
JSON representation if the request was successful.
f the request is invalid, a 400 error is send. Otherwise,
if any error appears on the server, a 500 error is send.

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


## Specific user collection [/users/{userId}]

Use this endpoint to either receive information for a specific
user, through a `GET` request, or delete an existing user via
`DELETE`.

If the user is not found, a 404 error is send.
If any other error occurs on the server side, a 500 error is
send.

### Get a specific user [GET]

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


### Delete a specific user [DELETE]

-   Response 200 (application/json)

            {
                "id": "6eaade67-e087-4327-8bd5-92934baf58ed",
                "created_at": "2022-08-07T16:18:39.827566Z",
                "updated_at": "2022-08-07T16:18:39.827566Z",
                "items": []
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


## User items collection [/users/{userId}/items]

This collect provides to receive wished items for a specific user (`GET`), 
or to add new items to the wishlist (`POST`).

If the user does not exists, a 404 error is thrown.
If the request for the `POST` request is invalid, a 400 error is thrown.
For any other errors on the server side, a 500 error is thrown.

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

## Specific user item collection [/users/{userId}/items/{itemId}]

Via this endpoint the client can either receive information about a
specific item from a specific user (`GET`), update the specific item (`PUT`), 
or delete the item `DELETE`.

The fields that can be updated are `name`, `price`, `vendor_id`,
and `has_been_bought`. After an item has been updated, the 
system will update its meta value `updated_at`.

if the user or the item do not exist, a 404 error is thrown.
If the request for the `PUT` request is invalid, a 400
error is thrown.
Otherwise, if an error occurs on the server side, it will response 
with a 500 error.

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


### Delete item [DELETE]

-   Response 200 (application/json)

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

### Get a specific Item [GET]

-   Response 200 (application/json)

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
