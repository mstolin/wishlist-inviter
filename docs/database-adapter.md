FORMAT: 1A
HOST: http://database-adapter/

# Database-Adapter

The Database-Adapter is an adapter service located in the Adapter Service Layer.
It is responsible to directly connect to the underlying database, and therefore,
to execute all database related operations.

The underlying Database is a Postgresql database. The goal of this adapter is to
provide an interface where the underlying databse can easily be changed, without
effecting the other services of this project.

## Users Collection [/users]

Via this endpoint, a new user can be created by sending
an empty request through `POST`.
If the request was successful, it will respond the newly created
user instance.

If the request is invalid, a 400 error is thrown. Otherwise, for
any other error on the server side, a 500 error is thrown.

### Create a new User [POST]

Use this endpoint to create a new User dataset.
To create a new user, just send an empty json request.

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

        {}

+ Response 200 (application/json)

        {
            "id":"c8dc276b-176a-4468-b6b6-4af63f1b98f1",
            "created_at":"2022-05-09T14:37:17.91663653Z",
            "updated_at":"2022-05-09T14:37:17.91663653Z",
            "items":null
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

+ Response 500 (application/json)

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }
        
## Specific User Collection [/users/{userId}]

Through this endpoint, information about a specific user
can be requested (`GET`), or a specific user can be deleted
(`DELETE`).

The response for both methods is equal. Even if the user has
been deleted, the response will be the data of the user.

If the user does not exist, a 404 error is thrown.
For any other error that happens on the server side, a
500 error is send.

### Get a specific User [GET]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id":"4b33f6f3-fcb9-4137-8a4e-c7e2c9eadf3f",
            "created_at":"2022-05-10T14:47:41.836503+02:00",
            "updated_at":"2022-05-10T14:51:34.85178+02:00",
            "items":[
                {
                    "id":1,
                    "created_at":"2022-05-09T20:06:50.141576+02:00",
                    "updated_at":"2022-05-09T20:06:50.141576+02:00",
                    "name": "Hades [Nintendo Switch]",
                    "price": 64.99,
                    "vendor": "amazon"
                    "vendor_id":"I9RYNJAL9GAYA",
                    "has_been_baught": false
                },
                {
                    "id":2,
                    "created_at":"2022-05-09T23:04:34.929649+02:00",
                    "updated_at":"2022-05-09T23:04:34.929649+02:00",
                    "name": "The Witcher 3: Wild Hunt - Complete Edition - [Nintendo Switch]",
                    "price": 52.23,
                    "vendor": "amazon"
                    "vendor_id":"I7DF1D943PWDV",
                    "has_been_baught": false
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
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"GENERIC ERROR MESSAGE"
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
            
### Delete a specific User [DELETE]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id":"4b33f6f3-fcb9-4137-8a4e-c7e2c9eadf3f",
            "created_at":"2022-05-10T14:47:41.836503+02:00",
            "updated_at":"2022-05-10T14:51:34.85178+02:00",
            "items":[
                {
                    "id":1,
                    "created_at":"2022-05-09T20:06:50.141576+02:00",
                    "updated_at":"2022-05-09T20:06:50.141576+02:00",
                    "name": "Hades [Nintendo Switch]",
                    "price": 64.99,
                    "vendor": "amazon"
                    "vendor_id":"I9RYNJAL9GAYA",
                    "has_been_baught": false
                },
                {
                    "id":2,
                    "created_at":"2022-05-09T23:04:34.929649+02:00",
                    "updated_at":"2022-05-09T23:04:34.929649+02:00",
                    "name": "The Witcher 3: Wild Hunt - Complete Edition - [Nintendo Switch]",
                    "price": 52.23,
                    "vendor": "amazon"
                    "vendor_id":"I7DF1D943PWDV",
                    "has_been_baught": false
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
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"GENERIC ERROR MESSAGE
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
        
## User Items Collection [/users/{userId}/items]

Using this endpoint, a list of all whished items of specific user
can be requested (`GET`) or multiple items can be added to the
wishlist of the user (`POST`).
After a successful request, the server respons with the overall item collection
of the user.

If the user does not exists, a 404 error is thrown.
If the request for the `POST` endpoint is invalid, a 400 error is given.
Otherwise, for any other error, the server responds with a 500 error.

### Add a collection of new Items to a User [POST]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

        [
            {
                "name": "Hades [Nintendo Switch]",
                "price": 64.99,
                "vendor": "amazon"
                "vendor_id":"I9RYNJAL9GAYA",
                "has_been_baught": false
            }
        ]

+ Response 200 (application/json)

        [
            {
                "id":1,
                "created_at":"2022-05-09T20:06:50.141576+02:00",
                "updated_at":"2022-05-09T20:06:50.141576+02:00",
                "name": "Hades [Nintendo Switch]",
                "price": 64.99,
                "vendor": "amazon"
                "vendor_id":"I9RYNJAL9GAYA",
                "has_been_baught": false
            },
            {
                "id":2,
                "created_at":"2022-05-09T23:04:34.929649+02:00",
                "updated_at":"2022-05-09T23:04:34.929649+02:00",
                "name": "The Witcher 3: Wild Hunt - Complete Edition - [Nintendo Switch]",
                "price": 52.23,
                "vendor": "amazon"
                "vendor_id":"I7DF1D943PWDV",
                "has_been_baught": false
            },
            {
                "id":3,
                "created_at":"2022-05-09T23:04:45.939046+02:00",
                "updated_at":"2022-05-09T23:04:45.939046+02:00",
                "name": "The Legend of Zelda: Breath of the Wild [Nintendo Switch]",
                "price": 54.99,
                "vendor": "amazon"
                "vendor_id":"ITTLSQBBXRYE3",
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
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"GENERIC ERROR MESSAGE"
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

### Get all Items of a User [GET]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        [
            {
                "id":1,
                "created_at":"2022-05-09T20:06:50.141576+02:00",
                "updated_at":"2022-05-09T20:06:50.141576+02:00",
                "name": "Hades [Nintendo Switch]",
                "price": 64.99,
                "vendor": "amazon"
                "vendor_id":"I9RYNJAL9GAYA",
                "has_been_baught": false
            },
            {
                "id":2,
                "created_at":"2022-05-09T23:04:34.929649+02:00",
                "updated_at":"2022-05-09T23:04:34.929649+02:00",
                "name": "The Witcher 3: Wild Hunt - Complete Edition - [Nintendo Switch]",
                "price": 52.23,
                "vendor": "amazon"
                "vendor_id":"I7DF1D943PWDV",
                "has_been_baught": false
            },
            {
                "id":3,
                "created_at":"2022-05-09T23:04:45.939046+02:00",
                "updated_at":"2022-05-09T23:04:45.939046+02:00",
                "name": "The Legend of Zelda: Breath of the Wild [Nintendo Switch]",
                "price": 54.99,
                "vendor": "amazon"
                "vendor_id":"ITTLSQBBXRYE3",
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
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"GENERIC ERROR MESSAGE"
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
        
## Specific User Items Collection [/users/{userId}/items/{itemId}]

This endpoint allows to request information about a specific item
from a users wishlist (`GET`), delete an item `DELETE`, or update
an item (`PUT`).

If the user or the item does not exists, a 404 error is thrown.
If the data of the `PUT` request is invalid, a 400 error is thrown.
Otherwise, the endpoint will respond with a 500 error for server-side
errors.
        
### Delete a specific Item [DELETE]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id":1,
            "created_at":"2022-05-09T20:06:50.141576+02:00",
            "updated_at":"2022-05-09T20:06:50.141576+02:00",
            "name": "Nintendo Luigi's Mansion 3 - [Nintendo Switch]",
            "price": 47.94,
            "vendor": "amazon",
            "vendor_id":"I37NVYF3F7GSV1",
            "has_been_baught": false
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
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"GENERIC ERROR MESSAGE"
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

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

+ Response 200 (application/json)

        {
            "id":1,
            "created_at":"2022-05-09T20:06:50.141576+02:00",
            "updated_at":"2022-05-09T20:06:50.141576+02:00",
            "name": "Nintendo Luigi's Mansion 3 - [Nintendo Switch]",
            "price": 47.94,
            "vendor": "amazon",
            "vendor_id":"I37NVYF3F7GSV1",
            "has_been_baught": false
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
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"GENERIC ERROR MESSAGE"
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

### Update a specific Item [PUT]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

        {
            "name": "Hades [Nintendo Switch]",
            "price": 64.99,
            "vendor": "amazon"
            "vendor_id":"I9RYNJAL9GAYA",
            "has_been_baught": true
        }

+ Response 200 (application/json)

        {
            "id":1,
            "created_at":"2022-05-09T20:06:50.141576+02:00",
            "updated_at":"2022-05-09T20:08:32.234543+02:00",
            "name": "Hades [Nintendo Switch]",
            "price": 64.99,
            "vendor": "amazon"
            "vendor_id":"I9RYNJAL9GAYA",
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
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"GENERIC ERROR MESSAGE"
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
