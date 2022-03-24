package encoder

import (
	"encoding/hex"

	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/appengine/log"
)

func PerEncodeE2SmKpmRanFunctionDescription(rfd *e2smrmet.E2SmKpmRanfunctionDescription) ([]byte, error) {

	log.Debugf("Obtained E2SM-RMET-RANfunctionDescription message is\n%v", rfd)
	if err := rfd.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-RANfunctionDescription PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(rfd, "valueExt", e2smrmet.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RMET-RANfunctionDescription PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmKpmRanFunctionDescription(per []byte) (*e2smrmet.E2SmKpmRanfunctionDescription, error) {

	log.Debugf("Obtained E2SM-RMET-RANfunctionDescription PER bytes are\n%v", hex.Dump(per))

	result := e2smrmet.E2SmKpmRanfunctionDescription{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smrmet.Choicemape2smKpm, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RMET-RANfunctionDescription from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-RANfunctionDescription PDU %s", err.Error())
	}

	return &result, nil
}
