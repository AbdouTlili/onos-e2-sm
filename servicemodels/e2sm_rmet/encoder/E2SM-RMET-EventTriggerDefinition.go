package encoder

import (
	"encoding/hex"

	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRmetEventTriggerDefinition(etd *e2smrmet.E2SmRmetEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-RMET-EventTriggerDefinition message is\n%v", etd)
	if err := etd.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-EventTriggerDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(etd, "valueExt", e2smrmet.E2smRmetChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RMET-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRmetEventTriggerDefinition(per []byte) (*e2smrmet.E2SmRmetEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-RMET-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smrmet.E2SmRmetEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smrmet.E2smRmetChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RMET-EventTriggerDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-EventTriggerDefinition PDU %s", err.Error())
	}

	return &result, nil
}
