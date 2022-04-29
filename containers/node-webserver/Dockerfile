FROM node:16.15.0-alpine
ARG workdir=/usr/src/app
RUN addgroup runner && adduser -D -H webserver -G runner
USER webserver
COPY ./ $workdir
WORKDIR $workdir
EXPOSE 3030
CMD ["node", "index.js"]
