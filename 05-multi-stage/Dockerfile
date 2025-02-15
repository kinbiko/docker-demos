FROM golang:1.24.0-alpine AS builder

# Specified and overwritten when building.
ARG VERSION=0.0.1-dev

WORKDIR /app/
COPY . .
RUN apk add -U --no-cache ca-certificates \
  && adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid 10001 appuser \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o demo -trimpath -ldflags="-w -s -X main.version=${VERSION}" .


FROM scratch

ENV PORT=8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /app/
COPY --from=builder /app/demo .

EXPOSE 8080
USER appuser:appuser
CMD [ "./demo" ]
