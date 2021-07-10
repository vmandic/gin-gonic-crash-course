FROM golang:1.16.5

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o server

CMD ["./server"]