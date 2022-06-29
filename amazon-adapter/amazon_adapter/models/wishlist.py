from pydantic import BaseModel
from typing import List

from amazon_adapter.models.item import Item


class Wishlist(BaseModel):
    id: str
    vendor: str
    name: str
    items: List[Item] = []
