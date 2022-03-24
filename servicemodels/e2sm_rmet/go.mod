module github.com/AbdouTlili/onos-e2-sm/servicemodels/e2smrmet

go 1.16

replace github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet => ./

require (
	github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet v0.0.0-00010101000000-000000000000
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/gogo/protobuf v1.3.2
	github.com/onosproject/onos-api/go v0.9.8
	github.com/onosproject/onos-lib-go v0.8.9
	google.golang.org/appengine v1.6.7
	google.golang.org/protobuf v1.27.1
)
