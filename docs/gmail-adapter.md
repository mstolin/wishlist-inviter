FORMAT: 1A
HOST: http://gmail-adapter/

# GMail-Adapter

The Gmail-Adapter is an adapter-service, of the adapter-service layer, that is responsible
to facade the usage of the Google Mail service.

Overall, it is only used to send mails via Google Mail.

## Mail Endpoint [/mail]

The `/mail` endpoint is used to send a mail.

### Send Mail [GET]

+ Request (application/json)

        {
            "recipient": "recipient@domain.tld",
            "subject": "You have been invited",
            "message": "This is a longer text ..."
        }

+ Response 200 (application/json)

        {
            "recipient": "recipient@domain.tld",
            "subject": "You have been invited",
            "message": "This is a longer text ..."
        }
        
+ Response 400

        {
            "error": "recipient, subject, and message are required fields"
        }

+ Response 500

        {
            "error": "recipient, subject, and message are required fields"
        }
