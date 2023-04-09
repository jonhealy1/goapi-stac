# Building the binary of the App
FROM golang:1.17 AS build

# `boilerplate` should be replaced with your project name
WORKDIR /go/src/go-stac-api

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

# `boilerplate` should be replaced here as well
COPY --from=build /go/src/go-stac-api/app .
COPY --from=build /go/src/go-stac-api/.env .

# Exposes port 6002 because our program listens on that port
EXPOSE 6002

CMD ["./app"]