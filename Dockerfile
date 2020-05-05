FROM golang:1.13-alpine AS builder

ENV GOPROXY="direct"

WORKDIR /go/src/github.com/mtpereira/hugo-post-preview
COPY . .

RUN adduser -D -g '' go \
        && apk add --no-cache git ca-certificates tzdata \
        && update-ca-certificates
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
    go build -ldflags='-w -s -extldflags "-static"' -o /go/bin/hugo-post-preview ./cmd/hugo-post-preview \
        && chmod 500 /go/bin/hugo-post-preview

FROM scratch

COPY --from=builder /etc/passwd /etc/group /etc/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder --chown=go /go/bin/hugo-post-preview /go/bin/hugo-post-preview

USER go
ENTRYPOINT [ "/go/bin/hugo-post-preview" ]
