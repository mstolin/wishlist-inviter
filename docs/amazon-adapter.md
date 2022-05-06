FORMAT: 1A
HOST: http://amazon-adapter

# Amazon-Adapter

This API is supposed to be used, to scrap data from Amazon.
Currently, the only data that can be scraped is a wishlist.

## Wishlists Collection [/wishlists/{wishlistId}]

### Get a specific Wishlist [GET]

This endpoint will return a wishlist with all its items for the given id.
The id, is the same id given by Amazon.

+ Response 200 (application/json)

        {
        
        }
        
+ Response 404 (application/json)

        {
            "error": "Wishlist not found"
        }
