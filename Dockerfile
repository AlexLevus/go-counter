FROM golang:1.16-alpine

RUN apk update && apk add  curl

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /counter

EXPOSE 8080

CMD [ "/counter" ]