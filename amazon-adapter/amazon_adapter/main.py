import os
import requests
import uvicorn
from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse

from amazon_adapter.scrapper.scrapper import Scrapper

app = FastAPI()

@app.on_event("startup")
async def on_startup():
    app.requests_session = requests.Session()

@app.on_event("shutdown")
async def on_shutdown():
    app.requests_session.close()

@app.get("/wishlists/{id}")
async def wishlist_handler(id: str):
    scrapper = Scrapper(app.requests_session)
    try:
        wishlist = scrapper.scrap_wishlist(id)
        return JSONResponse(content=jsonable_encoder(wishlist))
    except HTTPException as exc:
        return JSONResponse(
            status_code=exc.status_code,
            content={
                "error": {
                    "status": exc.status_code,
                    "message": exc.detail
                }
            }
        )
    except Exception as exc:
        return JSONResponse(
            status_code=500,
            content={
                "error": {
                    "status": exc.status_code,
                    "error": "Internal Server Error",
                    "message": str(exc)
                }
            }
        )


def start():
    """Launched with `poetry run start` at root level"""
    load_dotenv()
    host = os.environ.get("HOST")
    port = int(os.environ.get("PORT"))
    uvicorn.run("amazon_adapter.main:app", host=host, port=port, reload=False)