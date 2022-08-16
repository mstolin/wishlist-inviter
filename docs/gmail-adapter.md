FORMAT: 1A
HOST: http://gmail-adapter/

# GMail-Adapter

The GMail-Adapter is a service that sends mails via Google-Mail through SMTP.

## Mail Endpoint [/mail]

This endpoint will send an email through Google-Mails SMTP server.
It requires a simple JSON request, given the recipient of the mail, the subject,
and the mail body.
Usually this information should be generated via the [Mail-Servive](./mail-service.md).

If the given request is invalid, the endpoint responds a 400 error. Unauthorized clients
will receive a 401 error. If any other error will occur, a 500 error is send.

### Send Mail [GET]

+ Request (application/json)

    + Headers

        Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjA3NDk4MTksImlhdCI6MTY2MDY2MzQxOX0.fTF35iyBNsflkDlv2vdIQNjH6X0GexD7Q5MaEpg_T8o

    + Body

        {
            "recipient": "recipient@domain.tld",
            "subject": "Email subject",
            "body": "Email body"
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
