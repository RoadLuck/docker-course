FROM golang:alpine
MAINTAINER RoadLuck(Vicente)

RUN mkdir /code
WORKDIR code/
COPY . .
EXPOSE 8080
CMD [ "go", "run", "http-server.go" ]
