FROM alpine:latest

COPY out .

EXPOSE 8080

CMD ["./out"]
