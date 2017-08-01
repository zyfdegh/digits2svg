FROM container4armhf/armhf-alpine:latest
MAINTAINER zyfdegh <zyfdegg@gmail.com>

USER root
WORKDIR /root

COPY bin/digits2svg /root/digits2svg

CMD ["./digits2svg"]
