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

func PerEncodeE2SmRcPreControlMessage(cm *e2smrcprev2.E2SmRcPreControlMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlMessage message is\n%v", cm)
	if err := cm.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RC-PRE-ControlMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(cm, "choiceExt", e2smrcprev2.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RC-PRE-ControlMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcPreControlMessage(per []byte) (*e2smrcprev2.E2SmRcPreControlMessage, error) {

	log.Debugf("Obtained E2SM-RC-PRE-ControlMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smrcprev2.E2SmRcPreControlMessage{}
	err := aper.UnmarshalWithParams(per, &result, "choiceExt", e2smrcprev2.RcPreChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RC-PRE-ControlMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RC-PRE-ControlMessage PDU %s", err.Error())
	}

	return &result, nil
}
