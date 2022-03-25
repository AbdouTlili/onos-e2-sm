#!/bin/sh


docker run -it \
    -v `pwd`:/go/src/github.com/AbdouTlili/onos-e2-sm \
    -v `pwd`/../onos-lib-go:/go/src/github.com/onosproject/onos-lib-go \
    -w /go/src/github.com/AbdouTlili/onos-e2-sm \
    --entrypoint /go/src/github.com/AbdouTlili/onos-e2-sm/build/bin/generateChoice-helper-script.sh \
    onosproject/protoc-go:v1.0.2 