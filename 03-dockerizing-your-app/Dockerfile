FROM golang:1.24.0-alpine

COPY . .
ENV PORT=8080
RUN go build -o demo .
#EXPOSE 8080
CMD ["./demo"]

