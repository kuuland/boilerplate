FROM alpine
ADD app /bin/
RUN apk -Uuv add --no-cache ca-certificates tini tzdata && \
  ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  mkdir /boilerplate
WORKDIR /boilerplate
COPY kuu.json docs ./
ENTRYPOINT ["/sbin/tini","--", "app"]