FROM golang:1.22.3-alpine3.19

WORKDIR /app

COPY . .

RUN go build -o server .

CMD ["./server"]