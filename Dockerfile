FROM golang:1.15.6 as builder
LABEL maintainer="Nagy Salem <me@muhnagy.com>"
WORKDIR /app
COPY . .
RUN make build

FROM golang:1.15.6
WORKDIR /app
COPY --from=builder /app/pipedrive .
COPY --from=builder /app/wait-for-mysql.sh /usr/local/bin/wait-for-mysql.sh
EXPOSE 3000
CMD ./pipedrive