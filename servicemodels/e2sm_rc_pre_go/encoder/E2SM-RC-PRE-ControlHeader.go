// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRcPreControlHeader(ch *e2smrcprev2.E2SmRcPreControlHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader message is\n%v", ch)
	if err := ch.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RC-PRE-ControlHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(ch, "choiceExt", e2smrcprev2.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreControlHeader(per []byte) (*e2smrcprev2.E2SmRcPreControlHeader, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlHeader PER bytes are\n%v", hex.Dump(per))

	result := e2smrcprev2.E2SmRcPreControlHeader{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2smrcprev2.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RC-PRE-ControlHeader PDU %s", err.Error())
	}

	return &result, nil
}
