FROM golang AS build

WORKDIR /app

ADD . .

RUN go mod download

RUN go build -a -o promotion

FROM alpine

COPY --from=build /app/promotion /app/promotion

ENTRYPOINT [ "/app/promotion" ]


