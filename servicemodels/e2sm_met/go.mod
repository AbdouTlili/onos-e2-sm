module github.com/AbdouTlili/onos-e2-sm/servicemodels/e2smmet

go 1.16

// replace github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met => ./

require (
	github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met v0.0.0-20221019140119-0047e49196e4
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/onosproject/onos-api/go v0.9.43
	github.com/onosproject/onos-lib-go v0.8.17
	google.golang.org/protobuf v1.27.1
	gotest.tools v2.2.0+incompatible
)
