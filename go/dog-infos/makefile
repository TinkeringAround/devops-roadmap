BINARY_NAME=90DaysOfDevOps

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin src/*.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux src/*.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows src/*.go

run:
	./${BINARY_NAME}-darwin

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows