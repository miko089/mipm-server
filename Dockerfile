FROM golang:1.22-alpine3.19 as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o mipm_backend .

FROM alpine:3.19 AS runner
WORKDIR /runner
COPY --from=builder /build/mipm_backend ./
ENTRYPOINT ["./mipm_backend"]
