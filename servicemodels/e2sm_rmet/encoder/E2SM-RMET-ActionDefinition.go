package encoder

import (
	"encoding/hex"

	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2SmRmetActionDefinition(ad *e2smrmet.E2SmRmetActionDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-RMET-ActionDefinition message is\n%v", ad)
	if err := ad.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-ActionDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ad, "valueExt", e2smrmet.E2smRmetChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RMET-ActionDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRmetActionDefinition(per []byte) (*e2smrmet.E2SmRmetActionDefinition, error) {

	log.Debugf("Obtained E2SM-RMET-ActionDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smrmet.E2SmRmetActionDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smrmet.E2smRmetChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RMET-ActionDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-ActionDefinition PDU %s", err.Error())
	}

	return &result, nil
}
