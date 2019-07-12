# Build the manager binary
FROM golang:1.10.3 as builder

# Copy in the go src
WORKDIR /go/src/github.com/cloudnativedaysjp/showks-concourseci-pipeline-operator
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

RUN wget https://github.com/concourse/concourse/releases/download/v5.3.0/fly-5.3.0-linux-amd64.tgz
RUN tar xvzf fly-5.3.0-linux-amd64.tgz

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager github.com/cloudnativedaysjp/showks-concourseci-pipeline-operator/cmd/manager

# Copy the controller-manager into a thin image
FROM ubuntu:latest
WORKDIR /
COPY --from=builder /go/src/github.com/cloudnativedaysjp/showks-concourseci-pipeline-operator/manager .
COPY --from=builder /go/src/github.com/cloudnativedaysjp/showks-concourseci-pipeline-operator/fly .
ENTRYPOINT ["/manager"]
