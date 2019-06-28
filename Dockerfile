FROM golang:alpine AS builder
RUN apk add git
WORKDIR $GOPATH/src/github.com/ruspatrick/stan-svc
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o news .
FROM scratch
WORKDIR /app/bin
COPY --from=builder /go/src/github.com/ruspatrick/stan-svc/news .
CMD [ "/app/bin/news" ]