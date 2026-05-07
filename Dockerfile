# ---- build stage ----
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache libstdc++ libgcc
ADD https://github.com/tailwindlabs/tailwindcss/releases/download/v4.1.7/tailwindcss-linux-x64-musl /usr/local/bin/tailwindcss
RUN chmod +x /usr/local/bin/tailwindcss

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN tailwindcss -i web/static/css/input.css -o web/static/css/style.css --minify
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# ---- run stage ----
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/web ./web

EXPOSE 3000
CMD ["./server"]
