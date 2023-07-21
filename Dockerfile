FROM       golang:1.20.3 as builder
WORKDIR    /src/
COPY       . ./
RUN        CGO_ENABLED=0 go build -o bin/jppp .

FROM       scratch
WORKDIR    /go/
COPY       --from=builder /src/bin/jppp ./
EXPOSE     8000
ENTRYPOINT ["./jppp"]
