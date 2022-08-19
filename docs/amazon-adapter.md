FORMAT: 1A
HOST: http://amazon-adapter

# Amazon-Adapter

This service provides functionality to gather information about products from
Amazon. Currently only Amazon wishlists are supported.

## Amazon wishlist [/wishlists/{wishlistId}]

This endpoint is used to get a structured JSON representation
of a specific Amazon wishlist.

Unauthorized clients will receive a 401 error.
A 404 error is sent if the wanted wishlist does not exists.
If any error on the server side occurs a 500 response is sent.

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