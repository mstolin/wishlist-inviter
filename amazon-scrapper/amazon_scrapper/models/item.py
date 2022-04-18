from pydantic import BaseModel


class Item(BaseModel):
    id: str
    vendor: str
    name: str
    price: float
