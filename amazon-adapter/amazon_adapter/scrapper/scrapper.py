import os
from fastapi import HTTPException
from requests import Session
from bs4 import BeautifulSoup, Tag
from typing import List, Optional

from amazon_adapter.models.wishlist import Wishlist
from amazon_adapter.models.item import Item

#
# Links looks like: https://www.amazon.de/hz/wishlist/ls/194N1KF03IPTL
#


class Scrapper:

    def __init__(self, requests_session: Session) -> None:
        self._requests_session = requests_session

    def _parse_item(self, item_elem: Tag) -> Item:
        id = item_elem["data-itemid"]
        price = self._get_price(item_elem)
        title_elem = item_elem.find(id=f"itemName_{id}")
        if title_elem is not None and price is not None:
            name = title_elem.text.strip()
            return Item(id=id, name=name, price=price, vendor="amazon")
        else:
            return None

    def _get_price(self, item_elem: Tag) -> Optional[float]:
        price_text = item_elem["data-price"]
        if price_text != "-Infinity": # Items with no price will have this value
            return float(price_text)
        else:
            return None

    def _parse_items(self, wrapper: Tag) -> List[Item]:
        items = wrapper.select("ul#g-items li.g-item-sortable")
        if len(items) > 0:
            parsed_items = [self._parse_item(item) for item in items]
            return list(filter(lambda product: product is not None, parsed_items))
        else:
            return []

    def _parse_wishlist(self, id: str, page: str) -> Optional[Wishlist]:
        soup = BeautifulSoup(page, "lxml")
        wishlist_wrapper = soup.find(id="wishlist-page")
        if wishlist_wrapper is None:
            raise Exception("Wishlist wrapper with id wishlist-page not found")
        else:
            name = wishlist_wrapper.find(id="profile-list-name").text.strip()
            items = self._parse_items(wishlist_wrapper)
            return Wishlist(id=id, name=name, items=items, vendor="amazon")


    def scrap_wishlist(self, id: str) -> Wishlist:
        amazon_url = os.environ.get("AMAZON_URL")
        if amazon_url is None:
            raise Exception("Environment variable AMAZON_URL is not defined")

        url = f"{amazon_url}/hz/wishlist/ls/{id}"
        try:
            page = self._requests_session.get(url, timeout=10000)

            if page.status_code == 404:
                raise HTTPException(status_code=404, detail=f"wishlist at url {url} not found")
            elif page.status_code != 200:
                raise Exception(f"HTTP status {page.status_code} for URL {url}")

            wishlist = self._parse_wishlist(id, page.text)
            if wishlist is None:
                raise Exception(f"Unable to parse wishlist with ID {id} from {url}")
            else:
                return wishlist
        except Exception as exc:
            raise exc from Exception


