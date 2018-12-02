# WARNING This Dockerfile is used as a generic template for all applicatons
# DO NOT Use directly.
FROM golang:1.11 as builder

ARG WRKDIR

WORKDIR ${WRKDIR}
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM alpine:latest

ARG WRKDIR

#COPY WHAT YOU NEED HERE ONLY
COPY --from=builder ${WRKDIR}/app .
COPY --from=builder ${WRKDIR}/conf .

CMD ["./app"]
