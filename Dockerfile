FROM golang:1.17-bullseye
WORKDIR /usr/src/app

# local copy of sqlite3 for people exec-ing into the container
RUN apt-get update && apt-get install -y sqlite3

# refresh deps
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build --tags "json1" -v -o /usr/local/bin/app .
ENTRYPOINT ["app"]
