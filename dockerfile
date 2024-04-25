FROM alpine

RUN apk add ca-certificates

RUN mkdir config

COPY config/env.json config

COPY timely /

EXPOSE 8081
EXPOSE 8082

ENTRYPOINT ["/timely"]