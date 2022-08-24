import jwt
from fastapi import Security, HTTPException
from fastapi.responses import JSONResponse
from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer
from datetime import datetime, timedelta

from amazon_adapter.jsonHTTPException import JSONHTTPException

class AuthHandler:

    _security = HTTPBearer()
    _secret = ""

    def set_secret(self, secret: str):
        self._secret = secret

    def decode_token(self, token: str):
        try:
            payload = jwt.decode(token, self._secret, algorithms=['HS256'])
            return payload
        except jwt.ExpiredSignatureError:
            raise JSONHTTPException("Unauthorized", 401, "token is expired")
            # raise HTTPException(status_code=401, detail="token is expired")
            #return JSONResponse(
            #    status_code=401,
            #    content={
            #        "error": {
            #            "status": "Unauthorized",
            #            "message": "token is expired"
            #        }
            #    }
            #)
        except jwt.InvalidTokenError:
            raise JSONHTTPException("Unauthorized", 401, "token is expired")
            # raise HTTPException(status_code=401, detail="invalid token")
            #return JSONResponse(
            #    status_code=401,
            #    content={
            #        "error": {
            #            "status": "Unauthorized",
            #            "message": "invalid token"
            #        }
            #    }
            #)

    def auth_wrapper(self, auth: HTTPAuthorizationCredentials = Security(_security)):
        return self.decode_token(auth.credentials)
