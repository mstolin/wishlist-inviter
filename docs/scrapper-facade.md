FORMAT: 1A
HOST: http://scrapper-facade

# Scrapper-Facade

The Scrapper-Facade service is used to gather product information from
different vendors. It forwards the request to the wanted vendor scrapper 
and returns its response. At this point only Amazon wishlists can be 
gathered.

## Amazon collection [/amazon/wishlists/{wishlistId}]

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