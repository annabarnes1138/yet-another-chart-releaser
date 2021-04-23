FROM alpine:3.13

RUN apk --no-cache add ca-certificates git

COPY yacr /usr/local/bin/yacr

# Ensure that the binary is available on path and is executable
RUN yacr --help

ENTRYPOINT /usr/local/bin/yacr
