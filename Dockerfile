FROM golang:1.14.4 AS builder
LABEL maintainer="Ming Cheng"

# Using 163 mirror for Debian Strech
#RUN sed -i 's/deb.debian.org/mirrors.163.com/g' /etc/apt/sources.list
#RUN apt-get update

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PACKAGE github.com/mingcheng/genpasswd.go
ENV BUILD_DIR ${GOPATH}/src/${PACKAGE}
ENV GOPROXY https://mirrors.cloud.tencent.com/go,https://goproxy.cn,https://goproxy.io,direct

# Print go version
RUN echo "GOROOT is ${GOROOT}, GOPATH is ${GOPATH}" && \
	echo ${GOROOT}/bin/go version

# Build
COPY . ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
RUN make clean && \
  env GO111MODULE=on GO15VENDOREXPERIMENT=1 make build && \
  ${BUILD_DIR}/genpasswd -version && \
  mv ${BUILD_DIR}/genpasswd /usr/bin/genpasswd

# Stage2
FROM alpine:3.11.6

# @from https://mirrors.ustc.edu.cn/help/alpine.html
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

COPY --from=builder /usr/bin/genpasswd /bin/genpasswd
ENTRYPOINT ["/bin/genpasswd"]
