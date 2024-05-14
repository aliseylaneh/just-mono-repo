from datetime import datetime
from elasticsearch import AsyncElasticsearch
import asyncio
from typing import Any


async def insert(
    client: AsyncElasticsearch, index_name: str, identifier: int, body: dict[str, Any]
):
    resp = await client.index(
        index=index_name, id=identifier, document=body, pretty=True
    )
    return resp


async def find_all(
    client: AsyncElasticsearch,
    index_name: str,
    pagination_size: int,
    pagination_index: int,
) -> list[dict[str, Any]]:
    resp = await client.search(
        index=index_name, size=pagination_size, from_=pagination_index, pretty=True
    )
    result = []
    for item in resp.body["hits"]["hits"]:
        result.append(item["_source"])
    print(result)
    return result


async def main():
    client = AsyncElasticsearch(
        "http://localhost:9200",
    )
    doc = {
        "author": "author_name",
        "text": "Interesting content...",
        "timestamp": datetime.now(),
    }
    async with client:
        await insert(client=client, identifier=1, index_name="books", body=doc)
        await find_all(
            client=client, index_name="books", pagination_index=0, pagination_size=2
        )


if __name__ == "__main__":
    asyncio.run(main())
