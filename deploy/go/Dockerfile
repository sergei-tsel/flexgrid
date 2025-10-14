FROM golang:1.25

WORKDIR /app
COPY ./web/backend .
RUN go build -o main ./cmd/flexgrid

EXPOSE 8080
ENTRYPOINT ["/app/main"]
