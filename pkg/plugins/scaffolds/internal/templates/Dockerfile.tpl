FROM golang:1.18-buster

ARG GOPROXY

WORKDIR /workspace
COPY . .
ENV GOPROXY=${GOPROXY:-https://proxy.golang.org}

RUN make csi
RUN chmod u+x /workspace/bin/csi

ENTRYPOINT ["/workspace/bin/csi"]
