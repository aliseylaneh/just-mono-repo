import asyncio
from collections import Counter
from http import HTTPStatus
from pathlib import Path
import httpx
import tqdm
from captures_flags import BASE_URL
from flag_common import main, DownloadStatus, save_flag

DEFAULT_CONCUR_REQ = 5
MAX_CONCUR_REQ = 1000


async def get_flag(client: httpx.AsyncClient, base_url: str, cc: str) -> bytes:
    url = f"{base_url}/{cc}/{cc}.gif".lower()
    response = await client.get(url, timeout=3.1, follow_redirects=True)
    response.raise_for_status()
    return response.read()


async def download_one(
    client: httpx.AsyncClient,
    cc: str,
    base_url: str,
    semaphore: asyncio.Semaphore,
    verbose: bool,
):
    try:
        async with semaphore:
            image = await get_flag(client=client, base_url=base_url, cc=cc)
    except httpx.HTTPStatusError as e:
        response = e.response
        if response.status_code == HTTPStatus.NOT_FOUND:
            status = DownloadStatus.NOT_FOUND
            msg = f"res: {response.url}"
    else:
        await asyncio.to_thread(save_flag, image, f"{cc}.gif")
        status = DownloadStatus.OK
        msg = "OK"
    if verbose and msg:
        print(cc, msg)
    return status
