FROM golang:1.17-bullseye
RUN apt-get update && apt-get install -y postgresql-client
WORKDIR /app
COPY . .
RUN cd /app/cmd && go build
EXPOSE 8080
ENTRYPOINT ["/app/cmd/cmd"]
