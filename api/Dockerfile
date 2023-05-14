FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 3211

RUN chmod +x main

CMD ["./main"]