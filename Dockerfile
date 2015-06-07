FROM google/golang-runtime

MAINTAINER dockerlover@zju.edu.cn

ENV APPDIR $GOPATH/src/github.com/Go-Docker-Hackathon/team-ZJU-FT

RUN mkdir -p $APPDIR

ADD [".",$APPDIR]

WORKDIR $APPDIR

RUN go build

EXPOSE 8082

CMD ["./team-ZJU-FT"]
