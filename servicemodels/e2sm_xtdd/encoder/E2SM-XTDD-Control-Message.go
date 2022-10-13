package encoder

import (
	"encoding/hex"
	"fmt"

	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmXtddControlMessage(cm *e2smxtdd.E2SmXTddControlMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-XTDD-IndicationMessage message is\n%v", cm)
	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-ControlMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(cm, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		fmt.Println("error aper.MarshalWithParams(cm, \"valueExt\", e2smxtdd.E2smXtddChoicemap, nil)")
		return nil, err
	}
	log.Debugf("Encoded E2SM-XTDD-ControlMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmXtddControlMessage(per []byte) (*e2smxtdd.E2SmXTddControlMessage, error) {

	log.Debugf("Obtained E2SM-XTDD-ControlMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smxtdd.E2SmXTddControlMessage{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-XTDD-ControlMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-ControlMessage PDU %s", err.Error())
	}

	return &result, nil
}
