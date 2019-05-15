FROM golang:1.12.4

ENV GOPATH /go

ENV PROJECT_NAME smart-grow-api
ENV PROJECT github.com/aanciaes/$PROJECT_NAME
ENV APP_ENV prod

COPY . /go/src/$PROJECT

WORKDIR /go/src/$PROJECT

RUN chmod +x setup.sh

EXPOSE 8000

CMD ./setup.sh \
    && go install /go/src/$PROJECT \
    && $GOPATH/bin/$PROJECT_NAME