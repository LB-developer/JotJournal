# ---- Build Go Backend ----
FROM public.ecr.aws/docker/library/golang:1.23.2 AS backend-builder

WORKDIR /app
COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/ ./server/

WORKDIR /app/server/cmd
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o /app/jotjournal main.go

FROM public.ecr.aws/debian/debian:stable
RUN apt-get update && apt-get install -y

WORKDIR /root/
COPY --from=backend-builder /app/jotjournal .

EXPOSE 8080

CMD ["./jotjournal"]
