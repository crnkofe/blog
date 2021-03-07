import asyncio
import datetime
import json
import requests
import logging
from enum import Enum
from bs4 import BeautifulSoup
from typing import *
from quart import Quart, websocket, request, abort
from hypercorn.config import Config
from hypercorn.asyncio import serve


logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

app = Quart(__name__)


def is_valid_url(url):
    # simplistic check
    return url != None and str(url).startswith("http://") or str(url).startswith("https://")


async def parse_url(request_id: int, url: str) -> Any:
    try:
        headers = {
            "Accept": "text/html",
            "User-Agent": "Crawler 0.1"
        }
        result = requests.get(url, allow_redirects=True, timeout=3, verify=False, headers=headers)
        if result.status_code == 200:
            site_text = result.text
            soup = BeautifulSoup(site_text, 'html.parser')
            # Extracting all the <a> tags into a list.
            links = [tag.get('href') for tag in soup.find_all('a') if is_valid_url(tag.get('href'))]
            await queue.put({
                "request_id": request_id,
                "url": url,
                "links": links
            })
        else:
            logger.info("Failed parsing url content: %s (%d, %s)", url, result.status_code, result)
    except Exception as e:
        logger.error("Failed scrapping url: {}".format(url), e)


@app.route('/urls/<request_id>', methods=['PUT'])
async def process_url(request_id):
    if request_id is None:
        logger.info("request_id is missing")
        abort(400, "missing request_id")

    data = await request.get_json()
    if data is None:
        logger.info("No data in payload: %s", data)
        abort(400, "missing request_id")

    url = data['url']

    # create task asynchronously - will execute after this function if finished
    task = asyncio.create_task(parse_url(request_id, url))
    return ({
        "status": "accepted"
    }, 201)


@app.websocket('/scraped_urls')
async def notify_scraped_urls():
    try:
        while True:
            try:
                data = queue.get_nowait()
                serialized_data = json.dumps(data)
                await websocket.send(serialized_data)
            except asyncio.QueueEmpty:
                pass

            await asyncio.sleep(1)
    except asyncio.CancelledError:
        logger.info("Stopping loop due to client cancellation")
        raise


if __name__ == '__main__':
    global queue
    try:
         # Create a queue that we will use to store our "workload".
        queue = asyncio.Queue()
        config = Config()
        config.bind = ["localhost:8080"]
        asyncio.run(serve(app, config))
    except KeyboardInterrupt:
        logger.info("Stopping")
