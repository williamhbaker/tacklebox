FROM golang:1.15.5

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main ./cmd/web
WORKDIR /dist
RUN cp /build/main .
EXPOSE 4000
CMD ["/dist/main"]