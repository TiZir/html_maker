FROM golang:1.22
WORKDIR /app

RUN apt-get update && apt-get -y install postgresql-client && apt-get -y install netcat-openbsd

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
COPY ./entrypoint.sh /entrypoint.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]