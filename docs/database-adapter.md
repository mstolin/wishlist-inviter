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

+ Request (application/json)

        {
        }

+ Response 200 (application/json)

        {
        }
        
## Specific User Collection [/users/{userId}]

### Get a specific User [GET]

This endpoint will response the user data of the requested user id.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }
            
### Delete a specific User [DELETE]

This endpoint deletes the user for the given id.

+ Request (application/json)

        {
        }

+ Response 200 (application/json)

        {

        }
        
## User Items Collection [/users/{userId}/items]

### Add a new Item to a User [POST]

Using this endpoint, a new item can be added to a specific user.

+ Request (application/json)

        {
            
        }

+ Response 200 (application/json)

        {
            
        }

### Get all Items of a User [GET]

This endpoint return all items that are associated to a specific user.

+ Response 200 (application/json)

        {
            
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
