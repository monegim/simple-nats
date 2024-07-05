From golang:1.22-alpine AS Builder
WORKDIR /go/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o simple-nats

FROM alpine:3.18
COPY --from=builder  --chown=65534:65534 /go/src/app/simple-nats .
USER 65534
EXPOSE 4222
CMD [ "./simple-nats" ]
