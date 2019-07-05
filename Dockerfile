FROM golang:1.12 AS builder

WORKDIR '/app'

# Copy over the dependencies declaration and install them
COPY go.mod .
RUN go get -d -v ./...

# Copy over the application code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go


FROM alpine
WORKDIR '/app'
COPY --from=builder /app/main .

EXPOSE 8080
EXPOSE 50051

CMD ["/app/main"]