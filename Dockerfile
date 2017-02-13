FROM golang:1.8
RUN apt-get update && apt-get install -y zip
ADD . /go/src/github.com/gopheracademy/learn
EXPOSE 3000
ENV SESSION_SECRET 99bottlesOfb33rOnTheWa119tyNineB0ttlesOfB33r!
ENV GO_ENV production
WORKDIR /go/src/github.com/gopheracademy/learn
RUN go get -t -u -v github.com/gobuffalo/buffalo/...
CMD ["./run.sh"]
