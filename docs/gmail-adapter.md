FORMAT: 1A
HOST: http://gmail-adapter/

# GMail-Adapter

The Gmail-Adapter is an adapter-service, of the adapter-service layer, that is responsible
to facade the usage of the Google Mail service.

Overall, it is only used to send mails via Google Mail.

```bash
$ curl -X POST http://gmail-adapter/mail \
  -H 'Content-Type: application/json' \
  -d '{"recipient":"test@name.org","subject":"Test Subject","body":"This is a sample text"}'
```

## Mail Endpoint [/mail]

The `/mail` endpoint is used to send a mail.

### Send Mail [GET]

+ Request (application/json)

        {
            "recipient": "recipient@domain.tld",
            "subject": "Email subject",
            "body": "Email body"
        }

+ Response 200 (application/json)

        {
            "recipient": "recipient@domain.tld",
            "subject": "Email subject",
            "body": "Email body"
        }

+ Response 400

        {
            "error": {
                "status": 400,
                "error": "Bad Request",
                "message": "GENERIC ERROR MESSAGE"
            } 
        }
    
+ Response 404

        {
            "error": {
                "status": 404,
                "error": "Not Found",
                "message": "The requested resource is not available."
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
