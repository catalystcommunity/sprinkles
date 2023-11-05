FROM golang:1.21 as builder

WORKDIR /workspace
COPY protos ./protos
COPY server/go.mod server/go.mod
COPY server/go.sum server/go.sum
# download dependencies
WORKDIR server
RUN go mod download
# copy source code
COPY server .
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o main main.go

# Use distroless as minimal base image
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/server/main .
USER 65532:65532

ENTRYPOINT ["/main"]
