FORMAT: 1A
HOST: http://mail-service/

# Mail-Service

The Mail-Service is responsible to handle all mail-specific requests.

Currently it only supports invitations for specific products.
When a request has been made, the service will generate a mail
(mail-header + body) containing a text and forwards it to the
[GMail-Adapter](./gmail-adapter.md).

In the future, this service should also be used for any other mail related tasks,
like user registration, *forgot password?*, and so on.

## Invitations Endpoint [/invitations]

This endpoint is used to generate an invitation message, for one or more products.
Then it will be redirected to the *Gmail-Adapter* and send to the recipient who is
supposed to buy the products.

If the request is invalid, a 400 error is send. For an unauthorized client, a 401
error is responded. Otherwise, a 500 error is send.

### Send Invitation [POST]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

        {
            "recipient": "recipient@domain.tld",
            "subject": "You have been invited",
            "user_id": "8a8c3b24-8997-43fc-b4b2-86482b3f70e7",
            "items": [1, 3, 4]
        }

+ Response 200 (application/json)

        {
            "message": "mail has been sent successfully"
        }
        
+ Response 400

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

+ Response 500

        {
            "error": {
                "status": 500,
                "error": "Internal Server Error",
                "message": "GENERIC ERROR MESSAGE"
            }
        }
