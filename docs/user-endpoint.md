FORMAT: 1A
HOST: http://user-endpoint/

# User-Endpoint

User-endpoint is the process centric API of this project. It provides an access point to the end-user
and is the only service that the user is supposed to communicate with directly.

All requests made to this service are forwarded to the *more specialized* service.
For example all requests made to `/mail/` are forwarded to he *mail-service*.

# User Collection [/users/]

This collection is responsible to handle all user related requests.

### Create a new User [POST]

This endopoint is used to create a new user instance.
A simple empty JSON request is enough to create a new user.
The response is a *UUIDv4* that uniquly identifies the new user.

+ Request (application/json)

        {
        }

+ Response 200 (application/json)

        {
            "id": "sdfsdfgdfg"
        }

+ Response 500 (application/json)

## Specific User [/users/{userId}/]

This API is used to get informations about a specific user.

### Get User [GET]

+ Response 200 (application/json)

        {
            "id": "",
            "items": [
                {
                    "name": "",
                    "price": 58.0,
                    "vendor": "amazon"
                    "vendor_id": ""
                }
            ]
        }

+ Response 404 (application/json)

        {
            "error": "User not found"
        }

+ Response 500 (application/json)

        {
            "error": "GENERIC ERROR MESSAGE"
        }

## User Items [/users/{userId}/items/]

Use this collection to get all items associated to a specific user.

### Get Items [GET]

+ Response 200 (application/json)

        [
            {
                "name": "sdf"
            }
        ]

+ Response 404 (application/json)

        {
            "error": "User not found"
        }

+ Response 500 (application/json)

        {
            "error": "GENERIC ERROR MESSAGE"
        }

### Add Items [PUT]

+ Request (application/json)

        {
            "name": "",
            "price": 58.0,
            "vendor": "amazon"
            "vendor_id": "sdafasdf"
        }

+ Response 200 (application/json)

        {
            "name": "",
            "price": 58.0,
            "vendor": "amazon"
            "vendor_id": "sdafasdf"
        }

+ Response 400 (application/json)

        {
            "error": "Unable to parse request"
        }

+ Response 500 (application/json)

        {
            "error": "GENERIC ERROR MESSAGE"
        }


# Items Collection [/items/]

Through this collection, products can be received.

## Amazon [/items/amazon/wishlists/{wishlistId}]

This collection is only used for Amazon related products.

### Get an Amazon Wishlist [GET]

+ Response 200 (application/json)

        {
            "name": "asdf",
            "vendor_id": "asasf",
            "items": [
                {
                    "name": "",
                    "price": 58.0,
                    "vendor": "amazon"
                    "vendor_id": "sdafasdf"
                }
            ]
        }

+ Response 404 (application/json)

        {
            "error": "Wishlist not found"
        }

+ Response 500 (application/json)

        {
            "error": "GENERIC ERROR MESSAGE"
        }


# Mail Collection [/mail/]

Through this collection, it is possible to send mail via Google-Mail.
It directly forwards the requests to the *Mail-Service*.

## Invitations [/mail/invitations]

This collection is used to send invitation messages.

### Send an Invitation [POST]

+ Request (application/json)

        {
            "recipient": "mail@domain.tld"
            "user_id": "asdfasdfasdf",
            "items": [1, 4, 6]
        }

+ Response 200 (application/json)

        {
            
        }

+ Response 400 (application/json)

        {
            "error": "Error handling request"
        }

+ Response 404 (application/json)

        {
            "error": "User not found"
        }

+ Response 500 (application/json)

        {
            "error": "GENERIC ERROR MESSAGE"
        }
