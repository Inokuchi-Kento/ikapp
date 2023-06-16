FROM golang:1.20-alpine3.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd cmd && go build -o main .

EXPOSE 8010

CMD ["./cmd/main"]