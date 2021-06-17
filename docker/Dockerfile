FROM golang:1.16-buster

ENV SRC_DIR=/go/src/github.com/ioannisGiak89/account-api-client

COPY provisioning/entrypoint.sh /usr/bin

WORKDIR $SRC_DIR

RUN chmod +x /usr/bin/entrypoint.sh

CMD ["entrypoint.sh"]
