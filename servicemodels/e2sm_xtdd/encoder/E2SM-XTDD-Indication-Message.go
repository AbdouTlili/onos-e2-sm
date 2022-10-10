package encoder

import (
	"encoding/hex"
	"fmt"

	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmMetIndicationMessage(im *e2smxtdd.E2SmXTddIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-XTDD-IndicationMessage message is\n%v", im)
	if err := im.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-IndicationMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(im, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		fmt.Println("error aper.MarshalWithParams(im, \"valueExt\", e2smxtdd.E2smMetChoicemap, nil)")
		return nil, err
	}
	log.Debugf("Encoded E2SM-XTDD-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMetIndicationMessage(per []byte) (*e2smxtdd.E2SmXTddIndicationMessage, error) {

	log.Debugf("Obtained E2SM-XTDD-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smxtdd.E2SmXTddIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-XTDD-IndicationMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-IndicationMessage PDU %s", err.Error())
	}

	return &result, nil
}
