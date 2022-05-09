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
            
### Delete a specific User [DELETE]

This endpoint deletes the user for the given id.

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
        
## User Items Collection [/users/{userId}/items]

### Add a collection of new Items to a User [POST]

Using this endpoint, a new item can be added to a specific user.

+ Request (application/json)

        [
            {
                "name": "Test Item 3", 
                "price": 4.89, 
                "vendor": "amazon", 
                "vendor_id": "SOME_ID"
            }
        ]

+ Response 200 (application/json)

        [
            {
                "id":4,
                "created_at":"2022-05-09T21:05:08.375312841Z",
                "updated_at":"2022-05-09T21:05:08.375312841Z",
                "name":"Test Item 3",
                "price":4.89,
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

### Get all Items of a User [GET]

This endpoint return all items that are associated to a specific user.

+ Response 200 (application/json)

        {
            
        }

+ Response 404 (application/json)

        {
            "error":{
                "status":404,
                "error":"Not Found",
                "message":"The requested resource is not available."
            }
        }
        
## Specific User Items Collection [/users/{userId}/items/{itemId}]

### Update a specific Item [PUT]

Via this endpoint, an existing item, associated to a specific user, can be altered.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }
        
### Delete a specific Item [DELETE]

This endpoint deletes the item with the given id.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }

### Get a specific Item [GET]

This endpoint returns the user item with the given id.

+ Response 200 (application/json)

        {
            
        }
