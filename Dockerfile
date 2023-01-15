FROM golang:1.14.1

RUN \
  : install requirements \
  && apt-get update \
  && apt-get install -y \
    git \
    make \
  && apt-get clean \
  && useradd golang --create-home

USER golang
WORKDIR /home/golang

CMD ["bash"]
