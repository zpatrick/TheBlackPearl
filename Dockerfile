FROM alpine
RUN apk add --no-cache ca-certificates
ADD ./static /static
ADD ./views /views
ADD ./tbp /
CMD ["/tbp"]
