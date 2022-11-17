FROM golang:latest
WORKDIR /dockerapp
LABEL maintainer="Falusvampen and Kevazy"
COPY . .
RUN go mod download
EXPOSE 8080
RUN go build -o server
CMD [ "./server" ]