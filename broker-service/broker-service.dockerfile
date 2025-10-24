# # Base go image
# FROM golang:1.25-alpine as builder
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# #change: .internal/api
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/broker-server
# RUN chmod +x /app/brokerApp

# # Build a tiny docker image
# FROM alpine:latest
# RUN mkdir /app
# COPY --from=builder /app/brokerApp /app
# CMD [ "/app/brokerApp" ]


# With Makefile
FROM alpine:latest
RUN mkdir /app
COPY brokerApp /app
CMD [ "/app/brokerApp" ]