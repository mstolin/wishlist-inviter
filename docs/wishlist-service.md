FORMAT: 1A
HOST: http://product-service

# Product-Service

Product-Service is an API, that allows to collect product data from different vendors.
For example, collecting information about a single product or collecting a whole wishlist
with all its products.

## Amazon Collection

This collection provides the endpoints to receive data from Amazon.
Currently only wishlists can be collected.

### Amazon wishlists [/amazon/wishlist/{wishlistId}]

+ Parameters
    + wishlistId (string) - ID of the wishlist, given by Amazon.

#### Collect a wishlist [GET]

+ Response 200 (application/json)

        {
            "id": "194N1KF03IPTL",
            "vendor": "amazon",
            "name": "Kaffee-Zeug",
            "items": [
                {
                    "id": "I3UCMMATCW0ATV",
                    "vendor": "amazon",
                    "name": "Hario 400 ml Olive Wood New Coffee Server, Transparent",
                    "price": 54.84
                },
                {
                    "id": "IP0OBIK4UO9AG",
                    "vendor": "amazon",
                    "name": "Hario V60 Glass Coffee Dripper",
                    "price": 50.8
                }
            ]
        }
        
+ Response 404 (application/json)

        {
            "status_code": 404,
            "message": "Resource not found"
        }
        
+ Response 500 (application/json)

        {
            "status_code": 500,
            "message": "Internal server error"
        }
