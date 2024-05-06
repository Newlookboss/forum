FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
ENV CGO_ENABLED=1
RUN apk add build-base
RUN apk add --no-cache gcc musl-dev
COPY . . 
RUN go build -o main ./cmd/web/

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main ./
COPY /ui /app/ui/
COPY /forum.db /app/
COPY /internal/models/statements.sql /app/internal/models/

CMD ["./main"]