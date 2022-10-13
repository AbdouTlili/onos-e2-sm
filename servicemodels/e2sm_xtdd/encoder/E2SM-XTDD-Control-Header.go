package encoder

import (
	"encoding/hex"
	"fmt"

	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmXtddControlHeader(ch *e2smxtdd.E2SmXTddControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-XTDD-ControlHeader message is\n%v", ch)
	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-ControlHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ch, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		fmt.Println("error aper.MarshalWithParams(ch, \"valueExt\", e2smxtdd.E2smXtddChoicemap, nil)")
		return nil, err
	}
	log.Debugf("Encoded E2SM-XTDD-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmXtddControlHeader(per []byte) (*e2smxtdd.E2SmXTddControlHeader, error) {

	log.Debugf("Obtained E2SM-XTDD-ControlHeader PER bytes are\n%v", hex.Dump(per))

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
