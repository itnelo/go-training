ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}

ARG APP_DEPLOYMENT_PATH

WORKDIR ${APP_DEPLOYMENT_PATH}

COPY . .

RUN go get -d -v
RUN go mod verify
RUN go mod vendor -v
RUN go install -mod vendor -v

CMD ["go-training"]
