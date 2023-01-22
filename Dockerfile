FROM golang:1.19-alpine
RUN apk add build-base

WORKDIR /

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN mkdir /pb_migrations
COPY migrations.js ./pb_migrations

RUN go build -o /vimu-pocketbase

EXPOSE 8080

CMD [ "/vimu-pocketbase serve" ]
