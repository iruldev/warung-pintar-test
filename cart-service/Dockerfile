############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o cart-service

############################
# STEP 2 build a small image
############################
FROM alpine
WORKDIR /app
COPY --from=builder /app/cart-service /app

CMD ["./cart-service"]