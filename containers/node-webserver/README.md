# Nodejs Webserver

## Create a simple Node.js web server

- Create index.js to expose a web server on port 3030

````javascript
// content of index.js
const http = require('http');
const port = 3030;
const requestHandler = (request, response) => {
    console.log(request.url);
    response.end('Hello Node.js Server!');
};
const server = http.createServer(requestHandler);
server.listen(port, (err) => {
    if (err) {
        return console.log('something bad happened', err);
    }
    console.log(`server is listening on ${port}`);
});
````

- Install package.json using `npm install`
- Start using npm start `curl localhost:3030` or open in local browser -Terminate local service after testing

## Write a Dockerfile

- Use a Node.js base image
- Copy package.json over to image and `npm install` at build time
- Run `npm start` at container start time
- `EXPOSE` port 3030 from the container

## Instantiate image and connect to your container

- Verify that you can connect to the web service from within the container
- Use `docker container ls` to show exposed ports for running containers
- Verify that you can connect to the web service from outside the container e.g. on host machine
- Capture all your commands in scripts (good IAC practice)

## Run your container as a non-root user, e.g. user 1000

- Modify your container to run all processes as user 1000
- Set appropriate file permissions to enable access to your non-root user