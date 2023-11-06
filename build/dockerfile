FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o scraper cmd/main.go
ENTRYPOINT ./scraper

FROM alpine
WORKDIR /app
COPY --from=builder /app/scraper .
ENTRYPOINT ["/app/scraper"]

