package encoder

import (
	"encoding/hex"
	"fmt"

	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmMetIndicationHeader(ih *e2smxtdd.E2SmXTddIndicationHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-XTDD-IndicationHeader message is\n%v", ih)
	if err := ih.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-IndicationHeader PDU %s", err.Error())
	}
	per, err := aper.MarshalWithParams(ih, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		//REVIEW
		fmt.Println("error in aper.MarshalWithParams(ih, \"valueExt\", e2smxtdd.E2smMetChoicemap, nil)")
		return nil, err
	}
	log.Debugf("Encoded E2SM-XTDD-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMetIndicationHeader(per []byte) (*e2smxtdd.E2SmXTddIndicationHeader, error) {

	log.Debugf("Obtained E2SM-XTDD-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	result := e2smxtdd.E2SmXTddIndicationHeader{}
	//REVIEW
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smxtdd.E2smXtddChoicemap, nil)
	if err != nil {
		fmt.Println("aper.UnmarshalWithParams(per, &result, \"valueExt\", e2smxtdd.E2smMetChoicemap, nil)")
		return nil, err
	}

	log.Debugf("Decoded E2SM-XTDD-IndicationHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-XTDD-IndicationHeader PDU %s", err.Error())
	}

	return &result, nil
}
