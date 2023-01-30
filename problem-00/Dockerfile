FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o p0 && \
    chmod +x p0

EXPOSE 8080

CMD ["/app/p0"]
