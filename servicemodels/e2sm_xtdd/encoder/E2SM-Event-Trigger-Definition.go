package encoder

import (
	"encoding/hex"

	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2SmXtddEventTriggerDefinition(etd *e2smxtdd.E2SmXTddEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-XTDD-EventTriggerDefinition message is\n%v", etd)
	if err := etd.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-EventTriggerDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(etd, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-XTDD-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmXtddEventTriggerDefinition(per []byte) (*e2smxtdd.E2SmXTddEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-XTDD-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smxtdd.E2SmXTddEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-XTDD-EventTriggerDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-EventTriggerDefinition PDU %s", err.Error())
	}

	return &result, nil
}
