import requests
import uvicorn
from fastapi import FastAPI, HTTPException, Depends, Request, status
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse
from fastapi.security import OAuth2PasswordBearer

from amazon_adapter.scrapper.scrapper import Scrapper
from amazon_adapter.authHandler import AuthHandler
from amazon_adapter.jsonHTTPException import JSONHTTPException
from amazon_adapter.settings import Settings

app = FastAPI()
settings = Settings()
auth_handler = AuthHandler(settings)


@app.exception_handler(JSONHTTPException)
async def json_exception_handler(request: Request, exc: JSONHTTPException):
    return JSONResponse(
        status_code=exc.status, 
        content={
            "error": {
                "status": exc.status,
                "error": exc.error,
                "message": exc.message
            }
        }
    )

@app.on_event("startup")
async def on_startup():
    app.requests_session = requests.Session()

@app.on_event("shutdown")
async def on_shutdown():
    app.requests_session.close()

@app.get("/wishlists/{id}", dependencies=[Depends(auth_handler.auth_wrapper)])
async def wishlist_handler(id: str):
    scrapper = Scrapper(app.requests_session)
    try:
        wishlist = scrapper.scrap_wishlist(id)
        return JSONResponse(content=jsonable_encoder(wishlist))
    except HTTPException as exc:
        raise JSONHTTPException("", exc.status_code, exc.detail)
    except Exception as exc:
        raise JSONHTTPException("Internal Server Error", 500, str(exc))


def start():
    """Launched with `poetry run start` at root level"""
    uvicorn.run("amazon_adapter.main:app", reload=False)