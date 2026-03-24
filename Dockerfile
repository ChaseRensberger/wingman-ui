FROM oven/bun:1.2.22-alpine AS css
WORKDIR /app

COPY package.json bun.lock input.css index.html ./
COPY components ./components

RUN bun install --frozen-lockfile
RUN bun run css:build

FROM golang:1.26-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=css /app/output.css ./output.css

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o wingman-ui .

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app

COPY --from=builder /app/wingman-ui /app/wingman-ui
COPY --from=builder /app/index.html /app/index.html
COPY --from=builder /app/components /app/components
COPY --from=builder /app/output.css /app/output.css

EXPOSE 3000
CMD ["/app/wingman-ui"]
