package encoder

import (
	"encoding/hex"

	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRmetIndicationHeader(ih *e2smrmet.E2SmRmetIndicationHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RMET-IndicationHeader message is\n%v", ih)
	if err := ih.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-IndicationHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ih, "valueExt", e2smrmet.Choicemape2smRmet, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RMET-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRmetIndicationHeader(per []byte) (*e2smrmet.E2SmRmetIndicationHeader, error) {

	log.Debugf("Obtained E2SM-RMET-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	result := e2smrmet.E2SmRmetIndicationHeader{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smrmet.Choicemape2smRmet, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RMET-IndicationHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RMET-IndicationHeader PDU %s", err.Error())
	}

	return &result, nil
}
