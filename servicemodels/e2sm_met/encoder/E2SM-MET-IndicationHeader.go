// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"fmt"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger()

func PerEncodeE2SmMetIndicationHeader(ih *e2smmet.E2SmMetIndicationHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-MET-IndicationHeader message is\n%v", ih)
	if err := ih.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-IndicationHeader PDU %s", err.Error())
	}
	per, err := aper.MarshalWithParams(ih, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		//REVIEW
		fmt.Println("error in aper.MarshalWithParams(ih, \"valueExt\", e2smmet.E2smMetChoicemap, nil)")
		return nil, err
	}
	log.Debugf("Encoded E2SM-MET-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMetIndicationHeader(per []byte) (*e2smmet.E2SmMetIndicationHeader, error) {

	log.Debugf("Obtained E2SM-MET-IndicationHeader PER bytes are\n%v", hex.Dump(per))

	result := e2smmet.E2SmMetIndicationHeader{}
	//REVIEW
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		fmt.Println("aper.UnmarshalWithParams(per, &result, \"valueExt\", e2smmet.E2smMetChoicemap, nil)")
		return nil, err
	}

	log.Debugf("Decoded E2SM-MET-IndicationHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-IndicationHeader PDU %s", err.Error())
	}

	return &result, nil
}
