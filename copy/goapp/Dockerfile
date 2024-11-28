FROM debian:latest
RUN apt update
RUN apt-get install -y git make musl-dev golang

ENV GOROOT=/usr/lib/go
ENV GOPATH=/go
ENV PATH=/go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /usr/src/work

COPY main.go .
RUN mkdir template
COPY template/* template/.

EXPOSE 8080 
RUN go mod init work

RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq

CMD ["go","run","main.go"]


