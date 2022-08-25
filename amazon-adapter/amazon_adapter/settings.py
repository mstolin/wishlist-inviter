from pydantic import BaseSettings

class Settings(BaseSettings):
    host: str = "localhost"
    port: int = 8080
    jwt_sign_key: str
