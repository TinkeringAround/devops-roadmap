name = node-webserver
tag = $(name):1

build:
	docker build . -t $(tag) --progress=plain

dev:
	docker run -it --rm --name $(name) -p 3030:3030 $(tag) /bin/sh

run:
	docker run -d --rm --name $(name) -p 3030:3030 $(tag)

destroy:
	docker rm -f $(name)

cleanup:
	docker rmi $(tag)