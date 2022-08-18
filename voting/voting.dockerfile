FROM alpine:latest

RUN mkdir /app

COPY votingApp /app

CMD [ "/app/votingApp" ]
