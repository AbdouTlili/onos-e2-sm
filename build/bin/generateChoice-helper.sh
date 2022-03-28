export proto_imports=".:${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate:${GOPATH}/src"
echo $GOPATH
make build_protoc_gen_choice
protoc -I=$proto_imports:${GOPATH}/src/github.com/onosproject/onos-lib-go/api:${GOPATH}/src/github.com/onosproject/onos-e2-sm/servicemodels \
  --proto_path=servicemodels \
  --choice_out=./servicemodels/ \
  e2sm_met/v1/e2sm_met.proto