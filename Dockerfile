FROM golang:1.16-buster

ENV SRC_DIR=/go/src/github.com/ioannisGiak89/account-api-client

WORKDIR $SRC_DIR

COPY . $SRC_DIR

RUN cd $SRC_DIR \
    && cp ./docker/provision/entrypoint.sh /usr/bin \
    && chmod +x /usr/bin/entrypoint.sh

CMD ["entrypoint.sh"]
