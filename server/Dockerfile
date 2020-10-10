FROM golang:1.15.0-buster

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR cmd/server
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
