FROM golang:alpine

WORKDIR /code

COPY . .

ENTRYPOINT [ "go", "test" ]
CMD [ "./..." ]

