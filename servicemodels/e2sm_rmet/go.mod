module github.com/AbdouTlili/onos-e2-sm/servicemodels/e2smrmet

go 1.16

replace github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet => ./

require (
	github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet v0.0.0-20220325143040-651b4b4f5000
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/gogo/protobuf v1.3.2
	github.com/onosproject/onos-api/go v0.9.8
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go v0.8.7
	github.com/onosproject/onos-lib-go v0.8.12
	google.golang.org/protobuf v1.27.1
	gotest.tools v2.2.0+incompatible
)
