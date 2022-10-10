package encoder

import (
	"encoding/hex"

	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/appengine/log"
)

func PerEncodeE2SmXtddRanFunctionDescription(rfd *e2smxtdd.E2SmXTddRAnfunctionDescription) ([]byte, error) {

	log.Debugf("Obtained E2SM-XTDD-RANfunctionDescription message is\n%v", rfd)
	if err := rfd.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-RANfunctionDescription PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(rfd, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-XTDD-RANfunctionDescription PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmXtddRanFunctionDescription(per []byte) (*e2smxtdd.E2SmXTddRAnfunctionDescription, error) {

	log.Debugf("Obtained E2SM-XTDD-RANfunctionDescription PER bytes are\n%v", hex.Dump(per))

	result := e2smxtdd.E2SmXTddRAnfunctionDescription{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-XTDD-RANfunctionDescription from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-RANfunctionDescription PDU %s", err.Error())
	}

	return &result, nil
}
