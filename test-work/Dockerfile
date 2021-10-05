FROM golang:1.17-bullseye
WORKDIR /app
RUN apt-get update && apt-get install -y postgresql-client
COPY . .
RUN cd /app/cmd && go build
ENTRYPOINT ["/app/cmd/cmd"]
