FROM golang:1.15.6
LABEL maintainer="Nagy Salem <me@muhnagy.com>"
WORKDIR /app
COPY . .
RUN make build
ENV PIPEDRIVE_TOKEN "pipedrive token here"
CMD ["./pipedrive"]