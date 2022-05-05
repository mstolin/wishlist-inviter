FORMAT: 1A
HOST: http://mail-service/

# Mail-Service

The Mail-Service is responsible to handle all mail-specific requests.

Currently it only supports invitations for specific products.
When a request has been made, the service will genrate mail message
(mail-header + body) containing a text and forwards it to the
*Gmail-Adapter*.

In the future, this service should also be used for any other mail related tasks,
like user registration, *forgot password?*, and so on.

## Invitations Endpoint [/invitations]

This endpoint is used to generate an invitation message, for one or more products.
Then it will be redirected to the *Gmail-Adapter* and send to the recipient.

### Send Invitation [POST]

+ Request (application/json)

        {
            "recipient": "recipient@domain.tld",
            "subject": "You have been invited",
            "user_id": "asfdasdfasdf",
            "items": [1, 3, 4]
        }

+ Response 200 (application/json)

        {
            "recipient": "recipient@domain.tld",
            "subject": "You have been invited",
            "message": "This is a longer text ..."
        }
        
+ Response 400

        {
            "error": "recipient, subject, user_id, and items are required fields"
        }

+ Response 500

        {
            "error": "GENERIC ERROR MESSAGE"
        }
