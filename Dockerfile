# ---- Build Go Backend ----
FROM golang:1.23.2 AS backend-builder

WORKDIR /app
COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/ ./server/

WORKDIR /app/server/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/jotjournal main.go

WORKDIR /app/server/db/migrate
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/migrate main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=backend-builder /app/jotjournal .
COPY --from=backend-builder /app/migrate .
COPY --from=backend-builder /app/server/db/migrate/migrations ./migrations

EXPOSE 8080

CMD ["./jotjournal"]
