FROM ubuntu:18.04

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

RUN apt-get update

RUN apt-get install -y \
  iputils-ping \
  curl \
  wget \
  telnet \
  htop \
  vim \
  nano
  # set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf

ENV GOLANG_VERSION 1.15

RUN set -eux; \
	apt-get install -y \
		bash \
		gcc \
		musl-dev \
		openssl \
		golang-go \
	;

COPY ./package/go1.15.src.tar.gz ./

RUN export \
# set GOROOT_BOOTSTRAP such that we can actually build Go
		GOROOT_BOOTSTRAP="$(go env GOROOT)" \
# ... and set "cross-building" related vars to the installed system's values so that we create a build targeting the proper arch
# (for example, if our build host is GOARCH=amd64, but our build env/image is GOARCH=386, our build needs GOARCH=386)
		GOOS="$(go env GOOS)" \
		GOARCH="$(go env GOARCH)" \
		GOHOSTOS="$(go env GOHOSTOS)" \
		GOHOSTARCH="$(go env GOHOSTARCH)" \
	; \
# also explicitly set GO386 and GOARM if appropriate
# https://github.com/docker-library/golang/issues/184
	apkArch="$(apk --print-arch)"; \
	case "$apkArch" in \
		armhf) export GOARM='6' ;; \
		x86) export GO386='387' ;; \
	esac; \
	\
	wget -O go.tgz "https://src.fedoraproject.org/lookaside/pkgs/golang/go$GOLANG_VERSION.src.tar.gz/sha512/7d85382bcc6a0625dfa3d07196ab364860846367ed67697a7b1516b0af551a72bc4873882141fc3c7a60d39f2e27b33f6693e8b18b608de76fc9a55b5eac55ea/go$GOLANG_VERSION.src.tar.gz"; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	\
	cd /usr/local/go/src; \
	./make.bash; \
	\
	rm -rf \
# https://github.com/golang/go/blob/0b30cf534a03618162d3015c8705dd2231e34703/src/cmd/dist/buildtool.go#L121-L125
		/usr/local/go/pkg/bootstrap \
# https://golang.org/cl/82095
# https://github.com/golang/build/blob/e3fe1605c30f6a3fd136b561569933312ede8782/cmd/release/releaselet.go#L56
		/usr/local/go/pkg/obj \
	; \
	export PATH="/usr/local/go/bin:$PATH"; \
	go version
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

#RUN clang+llvm-7.0.1-x86_64-linux-gnu-ubuntu-18.04\

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-extldflags -static" -o app main.go

RUN ls

RUN mv /go/release/app /app

RUN pwd

CMD ["/app"]
