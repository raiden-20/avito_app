FROM debian as BUILD

RUN apt update -y && apt upgrade -y 
RUN apt install lsb-release curl gpg -y
RUN apt install git -y
RUN curl -O https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
RUN tar xvf go1.12.7.linux-amd64.tar.gz
RUN chown -R root:root ./go

ENV PATH="/go/bin:/usr/bin/git:{$PATH}"
ENV GOPATH="/go"

COPY ./src /app/src
COPY  cert.pem /cert.pem
COPY  key.pem /key.pem
WORKDIR /app/src

RUN git config --global http.sslVerify false
RUN go mod tidy
RUN go build -o /main main.go

CMD ["/main"]