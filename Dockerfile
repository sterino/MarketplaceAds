FROM golang:1.22.4-alpine AS builder

WORKDIR /build

COPY . .

RUN go build -o app ./cmd


FROM alpine

WORKDIR /app

#COPY --from=builder /build/.env ./.env
COPY --from=builder /build/migrations ./migrations
COPY --from=builder /build/app ./app

RUN chmod +x ./app

USER root

ENTRYPOINT ["./app"]