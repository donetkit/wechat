FROM alpine

ENV TZ Asia/Shanghai
# tar czvf src.tar.gz config docs main ts
COPY src.tar.gz /usr/local/src/src.tar.gz
WORKDIR /app
RUN set -xe \
    && tar -xvf /usr/local/src/src.tar.gz -C /usr/local/src \
    && rm -rf /usr/local/src/src.tar.gz \
    && chmod +x /usr/local/src/main \
    && mv /usr/local/src/* /app \
    && rm -rf /usr/local/src/*

ENTRYPOINT ["./main"]
