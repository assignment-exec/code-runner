FROM golang

ENV GOMODULE=on
ENV GOFLAGS=-mod=vendor

COPY . /code-runner
WORKDIR /code-runner

RUN apt-get install gcc

RUN go build -o code-runner-server

EXPOSE 8082

CMD ["./code-runner-server"]

