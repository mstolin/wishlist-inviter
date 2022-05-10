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

### Create a new User [POST]

Use this endpoint to create a new User dataset.
To create a new user, just send an empty json request.

+ Request (application/json)

        {
        }

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
    
+ Response 404 (application/json)

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "The requested resource is not available."
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

### Get a specific User [GET]

This endpoint will response the user data of the requested user id.

+ Response 200 (application/json)

        {
            "id":"c8dc276b-176a-4468-b6b6-4af63f1b98f1",
            "created_at":"2022-05-09T14:37:17.916636Z",
            "updated_at":"2022-05-09T14:37:17.916636Z",
            "items":[]
        }

+ Response 404 (application/json)

        {
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"user with id df345dfsg345 not found"
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

This endpoint deletes the user for the given id.
After a successful requests, it responds with user
instance, equal to a request made to *Get a specific User*.

+ Response 200 (application/json)

         {
            "id":"4b33f6f3-fcb9-4137-8a4e-c7e2c9eadf3f",
            "created_at":"2022-05-10T14:47:41.836503+02:00",
            "updated_at":"2022-05-10T14:51:34.85178+02:00",
            "items":[
                {
                    "id":1,
                    "created_at":"2022-05-10T14:51:34.852161+02:00",
                    "updated_at":"2022-05-10T14:51:34.852161+02:00",
                    "name":"Test Item 1",
                    "price":4.89,
                    "vendor":"amazon",
                    "vendor_id":"SOME_ID"
                },
                {
                    "id":2,
                    "created_at":"2022-05-10T14:51:34.852161+02:00",
                    "updated_at":"2022-05-10T14:51:34.852161+02:00",
                    "name":"Test Item 2",
                    "price":4.89,
                    "vendor":"amazon",
                    "vendor_id":"SOME_ID"
                }
            ]
        }

+ Response 404 (application/json)

        {
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"user with id df345dfsg345 not found"
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

### Add a collection of new Items to a User [POST]

Using this endpoint, a collection of items can be added to a specific user.
The collection can consist of one or more items.
After a successful request, the server respons with the overall item collection
of the user, equal to *Get all Items of a User*.

+ Request (application/json)

        [
            {
                "name": "Test Item 2", 
                "price": 7.99, 
                "vendor": "amazon", 
                "vendor_id": "SOME_ID"
            },
            {
                "name": "Test Item 3", 
                "price": 12.60, 
                "vendor": "amazon", 
                "vendor_id": "SOME_ID"
            },
        ]

+ Response 200 (application/json)

        [
            {
                "id":1,
                "created_at":"2022-05-09T20:06:50.141576+02:00",
                "updated_at":"2022-05-09T20:06:50.141576+02:00",
                "name":"Test Item 1",
                "price":54.78,
                "vendor":"amazon",
                "vendor_id":"SOME_ID"
            },
            {
                "id":2,
                "created_at":"2022-05-09T23:04:34.929649+02:00",
                "updated_at":"2022-05-09T23:04:34.929649+02:00",
                "name":"Test Item 2",
                "price":7.99,
                "vendor":"amazon",
                "vendor_id":"SOME_ID"
            },
            {
                "id":3,
                "created_at":"2022-05-09T23:04:45.939046+02:00",
                "updated_at":"2022-05-09T23:04:45.939046+02:00",
                "name":"Test Item 3",
                "price":12.60,
                "vendor":"amazon",
                "vendor_id":"SOME_ID"
            }
        ]

+ Response 404 (application/json)

        {
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"The requested resource is not available."
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

This endpoint return all items that are associated to a specific user.

+ Response 200 (application/json)

        [
            {
                "id":1,
                "created_at":"2022-05-09T20:06:50.141576+02:00",
                "updated_at":"2022-05-09T20:06:50.141576+02:00",
                "name":"Test Item 1",
                "price":54.78,
                "vendor":"amazon",
                "vendor_id":"SOME_ID"
            },
            {
                "id":2,
                "created_at":"2022-05-09T23:04:34.929649+02:00",
                "updated_at":"2022-05-09T23:04:34.929649+02:00",
                "name":"Test Item 2",
                "price":7.99,
                "vendor":"amazon",
                "vendor_id":"SOME_ID"
            },
            {
                "id":3,
                "created_at":"2022-05-09T23:04:45.939046+02:00",
                "updated_at":"2022-05-09T23:04:45.939046+02:00",
                "name":"Test Item 3",
                "price":12.60,
                "vendor":"amazon",
                "vendor_id":"SOME_ID"
            }
        ]

+ Response 404 (application/json)

        {
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"The requested resource is not available."
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
        
### Delete a specific Item [DELETE]

This endpoint deletes the item with the given id.
After a successful delete operation, it will response with
the deleted Item object.

+ Response 200 (application/json)

        {
            "id":1,
            "created_at":"2022-05-09T20:06:50.141576+02:00",
            "updated_at":"2022-05-09T20:06:50.141576+02:00",
            "name":"Test Item 1",
            "price":54.78,
            "vendor":"amazon",
            "vendor_id":"SOME_ID"
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

This endpoint returns the user item with the given id.

+ Response 200 (application/json)

        {
            "id":1,
            "created_at":"2022-05-09T20:06:50.141576+02:00",
            "updated_at":"2022-05-09T20:06:50.141576+02:00",
            "name":"Test Item 1",
            "price":54.78,
            "vendor":"amazon",
            "vendor_id":"SOME_ID"
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
