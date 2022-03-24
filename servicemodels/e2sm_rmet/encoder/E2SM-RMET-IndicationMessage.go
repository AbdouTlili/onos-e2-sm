package encoder

import (
	"encoding/hex"

	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRmetIndicationMessage(im *e2smrmet.E2SmRmetIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-RMET-IndicationMessage message is\n%v", im)
	if err := im.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-IndicationMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(im, "valueExt", e2smrmet.Choicemape2smRmet, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RMET-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRmetIndicationMessage(per []byte) (*e2smrmet.E2SmRmetIndicationMessage, error) {

	log.Debugf("Obtained E2SM-RMET-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smrmet.E2SmRmetIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smrmet.Choicemape2smRmet, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RMET-IndicationMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-IndicationMessage PDU %s", err.Error())
	}

	return &result, nil
}
