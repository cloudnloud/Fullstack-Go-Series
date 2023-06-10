FROM golang:1.19-alpine

WORKDIR /app/backend

COPY go.mod go.sum ./

RUN go mod download

COPY ./ .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
