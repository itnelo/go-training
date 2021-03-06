ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}

ARG DEPLOYMENT_PATH
ARG TIMEZONE

WORKDIR ${DEPLOYMENT_PATH}

COPY . .

# timezone
RUN ln -snf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime && echo ${TIMEZONE} > /etc/timezone && \
    date

RUN go get -d -v
RUN go mod verify
RUN go mod vendor -v
RUN go build -mod vendor -o /usr/local/bin/app -v

USER www-data

ENTRYPOINT ["app", "-port", "9696"]
