# Published Endpoints
GET    /search                   --> /search?engine=google&engine=bing&query=<phrase_to_search>
GET    /status
GET    /profiling


# Build and Run docker container
dockebuild --tag searchengine:1.0 .
docker run --publish 9091:9091 --name searchEngine searchengine:1.0

