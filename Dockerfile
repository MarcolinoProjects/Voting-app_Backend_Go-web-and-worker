FROM golang:1.17-alpine as builder

WORKDIR /src
COPY cmd /src/cmd
COPY pkg /src/pkg
COPY go.mod go.sum /src/
RUN ls -la && go build -o main -v ./cmd/*.go

FROM alpine:latest
LABEL org.opencontainers.image.authors="Victor Marcolino<marcolino.victor@gmail.com>"
EXPOSE 8080
COPY --from=builder /src/main /src/main
WORKDIR /src
COPY entrypoint.sh entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]