# BUILD STAGE
FROM golang:latest as build

# copy
WORKDIR /go/src/github.com/mchmarny/kuser/
COPY . /src/

# dependancies
WORKDIR /src/
ENV GO111MODULE=on
RUN go mod tidy

# build
WORKDIR /src/
RUN CGO_ENABLED=0 go build -v -o /kuser



# RUN STAGE
FROM alpine as release
RUN apk add --no-cache ca-certificates

# app executable
COPY --from=build /kuser /app/

# run
WORKDIR /app/
ENTRYPOINT ["./kuser"]