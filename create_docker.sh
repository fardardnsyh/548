#!/bin/bash
docker image build -f Dockerfile -t forum-image .

docker images


docker container run -p :8080 --detach --name container forum-image

docker ps -a

docker exec -it container /bin/bash

ls -l