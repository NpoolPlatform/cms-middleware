FROM uhub.service.ucloud.cn/entropypool/alpine:3.16

WORKDIR /app

COPY cms-middleware /app

ENTRYPOINT ["/app/cms-middleware"]

EXPOSE 50800 50801 50802 50803

CMD ["run"]
