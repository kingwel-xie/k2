FROM alpine

COPY ./k2 /
EXPOSE 8000

CMD ["/k2","server","-c", "/config/settings.yml"]
