FROM node
ENV DEBIAN_FRONTEND noninteractive
ENV TERM xterm
# prerequisites
RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get install -y curl vim unzip zip
# gcc for cgo
RUN apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.8rc3
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 0ff3faba02ac83920a65b453785771e75f128fbf9ba4ad1d5e72c044103f9c7a

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz
# cleanup
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/*
ENV PATH $PATH:/root/bin:/usr/local/go/bin
ENV GOPATH /root
RUN mkdir -p $GOPATH/bin
RUN go version
RUN go get github.com/gobuffalo/buffalo/...
ENV SESSION_SECRET 99bottlesOfb33rOnTheWa119tyNineB0ttlesOfB33r!
ENV GO_ENV production
RUN go get -t -u -v github.com/gobuffalo/buffalo/...
RUN go get -t -u -v github.com/markbates/grift/...
ADD . /root/src/github.com/gopheracademy/learn
WORKDIR /root/src/github.com/gopheracademy/learn
RUN npm i
RUN buffalo build -o bin/learn
EXPOSE 3000
CMD ["/root/src/github.com/gopheracademy/learn/bin/learn"]
