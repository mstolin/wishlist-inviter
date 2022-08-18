import jwt
from fastapi import Security, HTTPException
from fastapi.security import HTTPAuthorizationCredentials, HTTPBearer
from datetime import datetime, timedelta

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
            raise HTTPException(status_code=401, detail="token is expired")
        except jwt.InvalidTokenError:
            raise HTTPException(status_code=401, detail="invalid token")

    def auth_wrapper(self, auth: HTTPAuthorizationCredentials = Security(_security)):
        return self.decode_token(auth.credentials)
