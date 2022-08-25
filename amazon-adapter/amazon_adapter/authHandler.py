from jose import jwt
from jose.exceptions import ExpiredSignatureError, JWTError
from fastapi import Security, HTTPException
from fastapi.responses import JSONResponse
from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer
from datetime import datetime, timedelta

from amazon_adapter.jsonHTTPException import JSONHTTPException

class AuthHandler:

    _security = HTTPBearer()

    def __init__(self, settings):
        self._settings = settings

    def decode_token(self, token: str):
        try:
            payload = jwt.decode(token, self._settings.jwt_sign_key, algorithms=['HS256'])
            return payload
        except ExpiredSignatureError:
            raise JSONHTTPException("Unauthorized", 401, "token is expired")
        except JWTError:
            raise JSONHTTPException("Unauthorized", 401, "token is invalid")

    def auth_wrapper(self, auth: HTTPAuthorizationCredentials = Security(_security)):
        return self.decode_token(auth.credentials)
