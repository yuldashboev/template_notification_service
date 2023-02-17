# workspace (GOPATH) configured at /go
FROM golang:1.18.7 as builder

#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/template/template_notification_service
WORKDIR $GOPATH/src/gitlab.udevs.io/template/template_notification_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/template_notification_service /



FROM alpine
COPY --from=builder template_notification_service .
ENTRYPOINT ["/template_notification_service"]



