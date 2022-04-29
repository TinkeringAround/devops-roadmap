tag = secure-ubuntu:1

build:
	docker build . -t $(tag) --progress=plain

run:
	docker run -it --rm --name secure-ubuntu $(tag) /bin/bash