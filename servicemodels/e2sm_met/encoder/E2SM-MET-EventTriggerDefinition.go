// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmMetEventTriggerDefinition(etd *e2smmet.E2SmMetEventTriggerDefinition) ([]byte, error) {

	log.Debugf("Obtained E2SM-MET-EventTriggerDefinition message is\n%v", etd)
	if err := etd.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-EventTriggerDefinition PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(etd, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-MET-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmMetEventTriggerDefinition(per []byte) (*e2smmet.E2SmMetEventTriggerDefinition, error) {

	log.Debugf("Obtained E2SM-MET-EventTriggerDefinition PER bytes are\n%v", hex.Dump(per))

	result := e2smmet.E2SmMetEventTriggerDefinition{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-MET-EventTriggerDefinition from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-MET-EventTriggerDefinition PDU %s", err.Error())
	}

	return &result, nil
}
