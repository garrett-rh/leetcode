FROM golang:alpine AS base

WORKDIR /code

COPY . .

FROM base as test
RUN ./scripts/test.sh

FROM base as lint
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /bin v1.49.0
RUN ./scripts/lint.sh

from base as cover
RUN ./scripts/cover.sh
