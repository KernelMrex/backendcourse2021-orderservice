FROM debian:10-slim

ADD ./bin/orderservice /app/bin/
WORKDIR /app

EXPOSE 8000

CMD [ "/app/bin/orderservice" ]