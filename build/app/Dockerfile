FROM golang:1.15-alpine AS builder
RUN apk update && mkdir /app
WORKDIR /app
#ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 go build -o  dist/app -a -ldflags "-w -s" ./cmd/app

FROM alpine:3.10

ARG REPO_URL
ARG BRANCH
ARG COMMIT_REF
LABEL repo-url=${REPO_URL}
LABEL branch=${BRANCH}
LABEL commit-ref=${COMMIT_REF}

RUN apk update \
  && apk upgrade \
  && apk add --no-cache \
  ca-certificates tzdata bash \
  && update-ca-certificates 2>/dev/null || true \
  && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && mkdir /app
WORKDIR /app
COPY --from=builder /app/dist/app /app/app

EXPOSE 80

CMD ["/app/app"]