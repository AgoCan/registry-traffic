FROM golang:1.16.7 as builder
ENV GOPROXY=https://goproxy.cn
COPY . /root
RUN cd /root && go mod tidy && go build -ldflags "-X main.Version=0.0.1 -x main.CommitID=none"