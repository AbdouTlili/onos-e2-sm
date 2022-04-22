package encoder

import (
	"encoding/hex"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmMetActionDefinition(ad *e2smmet.E2SmMetActionDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-MET-ActionDefinition message is\n%v", ad)
	if err := ad.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-ActionDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ad, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MET-ActionDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMetActionDefinition(per []byte) (*e2smmet.E2SmMetActionDefinition, error) {

	log.Debugf("Obtained E2SM-MET-ActionDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smmet.E2SmMetActionDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MET-ActionDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-ActionDefinition PDU %s", err.Error())
	}

	return &result, nil
}
