FROM alpine
ADD app /bin/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk -Uuv add --no-cache ca-certificates tini tzdata && \
  ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  mkdir /kuu
WORKDIR /kuu
ADD . .
ENTRYPOINT ["/sbin/tini","--", "app"]