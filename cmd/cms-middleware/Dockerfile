FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY cms-middleware /app

ENTRYPOINT ["/app/cms-middleware"]

EXPOSE 50600 50601 50602 50603

CMD ["run"]
