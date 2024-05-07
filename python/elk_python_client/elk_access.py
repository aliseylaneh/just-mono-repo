from datetime import datetime
from elasticsearch import Elasticsearch


if __name__ == "__main__":
    client = Elasticsearch(
        "https://localhost:9200",
        basic_auth=("elastic", "ajhD8*c6qwWj4KjX=ybM"),
        verify_certs=False,
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
