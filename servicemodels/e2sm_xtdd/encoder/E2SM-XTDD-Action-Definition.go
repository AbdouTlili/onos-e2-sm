package encoder

import (
	"encoding/hex"

	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmXtddActionDefinition(ad *e2smxtdd.E2SmXTddActionDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-XTDD-ActionDefinition message is\n%v", ad)
	if err := ad.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-ActionDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ad, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-XTDD-ActionDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmXtddActionDefinition(per []byte) (*e2smxtdd.E2SmXTddActionDefinition, error) {

	log.Debugf("Obtained E2SM-XTDD-ActionDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smxtdd.E2SmXTddActionDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-XTDD-ActionDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-ActionDefinition PDU %s", err.Error())
	}

	return &result, nil
}
