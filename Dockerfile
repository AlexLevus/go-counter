FROM golang:1.16-alpine

RUN apk update && apk add  curl

WORKDIR /go/src/app
COPY ./src .

RUN go get -d -v
RUN go build -v
RUN echo $PATH
RUN ls
RUN pwd

CMD [ "./go-counter" ]