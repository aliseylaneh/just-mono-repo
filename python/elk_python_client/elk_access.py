from datetime import datetime
from elasticsearch import Elasticsearch


if __name__ == "__main__":
    client = Elasticsearch(
        "http://localhost:9200",
    )
    doc = {
        "author": "author_name",
        "text": "Interesting content...",
        "timestamp": datetime.now(),
    }
    resp = client.index(index="test-index", id=3, document=doc)
    print(resp["result"])

    resp = client.search(index="test-index")
    for hit in resp["hits"]["hits"]:
        print(hit["_source"])
