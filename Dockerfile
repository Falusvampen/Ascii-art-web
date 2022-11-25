FROM golang:1.19

WORKDIR /app

LABEL maintainer="Falusvampen and Kevazy"

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["/app/main"]