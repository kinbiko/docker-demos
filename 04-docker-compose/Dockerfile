FROM golang:1.24

COPY . .
RUN go build -o backend .
EXPOSE 8080
CMD ["./backend"]
