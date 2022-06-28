##
## Build
##
FROM golang:1.17.7-buster AS build

WORKDIR /app

COPY . .
RUN go mod init && go mod tidy

#COPY *.go ./

RUN go build -o /docker-gs-ping

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-gs-ping /docker-gs-ping

#EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]