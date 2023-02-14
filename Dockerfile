####################### Build stage #######################
FROM golang:1.19-alpine3.17 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main /app/main.go

####################### Run stage #######################
FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/main .

COPY . .

EXPOSE 8080

CMD [ "/app/main" ]