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
)

func PerEncodeE2SmMetIndicationMessage(im *e2smmet.E2SmMetIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-MET-IndicationMessage message is\n%v", im)
	if err := im.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-IndicationMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(im, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		fmt.Println("error aper.MarshalWithParams(im, \"valueExt\", e2smmet.E2smMetChoicemap, nil)")
		return nil, err
	}
	log.Debugf("Encoded E2SM-MET-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMetIndicationMessage(per []byte) (*e2smmet.E2SmMetIndicationMessage, error) {

	log.Debugf("Obtained E2SM-MET-IndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smmet.E2SmMetIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MET-IndicationMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-IndicationMessage PDU %s", err.Error())
	}

	return &result, nil
}
