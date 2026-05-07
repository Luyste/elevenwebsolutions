# ---- tailwind stage ----
FROM node:22-alpine AS tailwind

WORKDIR /app
RUN npm install -g @tailwindcss/cli
COPY web ./web
RUN npx @tailwindcss/cli -i web/static/css/input.css -o web/static/css/style.css --minify

# ---- build stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=tailwind /app/web/static/css/style.css ./web/static/css/style.css
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# ---- run stage ----
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/web ./web

EXPOSE 8080
CMD ["./server"]
