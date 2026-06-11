FROM golang:1.22-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /religare .

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends wireless-tools ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /religare /app/religare
COPY runConfigurations/ /app/runConfigurations/
COPY models/sampleTextContent /app/models/sampleTextContent
COPY docker/entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh /app/religare

ENTRYPOINT ["/app/entrypoint.sh"]
CMD []
