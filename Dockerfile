FROM golang:1.15.6 as builder
LABEL maintainer="Nagy Salem <me@muhnagy.com>"
WORKDIR /app
COPY . .
RUN make build

FROM golang:1.15.6
WORKDIR /app
COPY --from=builder /app/pipedrive .
EXPOSE 3000
CMD ./pipedrive