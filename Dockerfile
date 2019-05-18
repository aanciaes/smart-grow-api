FROM golang:1.12.4

ENV GOPATH /go

ENV PROJECT_NAME smart-grow-api
ENV PROJECT github.com/aanciaes/$PROJECT_NAME
ENV APP_ENV prod

COPY . /go/src/$PROJECT

WORKDIR /go/src/$PROJECT

RUN chmod +x setup.sh
RUN chmod +x wait-for-it.sh

CMD ./wait-for-it.sh -t 0 mysql-db:3306 -- echo "mysql-db is up" \
    && ./setup.sh \
    && go install /go/src/$PROJECT \
    && $GOPATH/bin/$PROJECT_NAME