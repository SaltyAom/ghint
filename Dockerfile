FROM golang:alpine AS builder

WORKDIR /usr/src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# * ---------- * ---------- * ---------- * ---------- * #
FROM busybox

WORKDIR /usr/app
COPY --from=builder /usr/src/app .

EXPOSE 3000
CMD ["sh", "-c", "./app"]
