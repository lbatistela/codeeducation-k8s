FROM golang:1.14.2 as builder
COPY src /go/src
WORKDIR /go/src/codeeducation
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o greeting

FROM scratch
COPY --from=builder /go/src/codeeducation/greeting /
EXPOSE 8000
CMD ["/greeting"]