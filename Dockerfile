FROM golang:latest AS development
RUN git clone --progress --verbose --depth=1 https://github.com/Bpazy/padavan_exporter /padavan_exporter
WORKDIR /padavan_exporter
RUN go env && CGO_ENABLED=0 go build ./cmd/padavan_exporter

FROM alpine:latest AS production
COPY --from=development /padavan_exporter/padavan_exporter /padavan_exporter/padavan_exporter
WORKDIR /padavan_exporter
ENTRYPOINT ./padavan_exporter
