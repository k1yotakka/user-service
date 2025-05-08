FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service .

RUN apk add --no-cache netcat-openbsd
RUN chmod +x wait-for.sh

EXPOSE 8001

ENTRYPOINT ["./wait-for.sh", "postgres", "5432", "./user-service"]
