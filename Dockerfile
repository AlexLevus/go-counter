FROM golang:1.18

COPY . /go/src/app/

WORKDIR /go/src/app/

RUN go build -o ./cmd/main ./cmd/main.go

ENTRYPOINT [ "./cmd/main" ]
