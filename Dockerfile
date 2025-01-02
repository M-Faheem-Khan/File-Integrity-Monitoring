# Build Stage
FROM golang:bookworm@sha256:2e838582004fab0931693a3a84743ceccfbfeeafa8187e87291a1afea457ff7a AS builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go mod tidy
RUN go build -o fim cmd/main.go

# Deploy Stage
FROM debian:bookworm@sha256:b877a1a3fdf02469440f1768cf69c9771338a875b7add5e80c45b756c92ac20a
WORKDIR /app
COPY --from=builder /build/fim /app/
WORKDIR /app

CMD [ "sh" ]
