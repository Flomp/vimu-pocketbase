FROM golang:1.19-alpine
RUN apk add build-base

WORKDIR /

COPY go.mod ./
COPY go.sum ./
COPY migrations ./migrations

RUN go mod download

COPY *.go ./

RUN go build -o /vimu-pocketbase
RUN /vimu-pocketbase migrate


EXPOSE 8090

ENTRYPOINT ["/vimu-pocketbase", "serve", "--http=0.0.0.0:8090", "--dir=/pb_data"]