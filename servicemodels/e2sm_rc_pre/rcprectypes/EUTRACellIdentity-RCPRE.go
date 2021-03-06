// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "EUTRACellIdentity-RCPRE.h"
import "C"
import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"unsafe"
)

func xerEncodeEUTRACellIdentity(eUTRACellIdentity *e2sm_rc_pre_v2.EutracellIdentity) ([]byte, error) {
	eUTRACellIdentityCP, err := newEUTRACellIdentity(eUTRACellIdentity)
	if err != nil {
		return nil, err
	}

	bytes, err := encodeXer(&C.asn_DEF_EUTRACellIdentity_RCPRE, unsafe.Pointer(eUTRACellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeEUTRACellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeEUTRACellIdentity(eUTRACellIdentity *e2sm_rc_pre_v2.EutracellIdentity) ([]byte, error) {
	eUTRACellIdentityCP, err := newEUTRACellIdentity(eUTRACellIdentity)
	if err != nil {
		return nil, err
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_EUTRACellIdentity_RCPRE, unsafe.Pointer(eUTRACellIdentityCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeEUTRACellIdentity() %s", err.Error())
	}
	return bytes, nil
}

func newEUTRACellIdentity(eUTRACellIdentity *e2sm_rc_pre_v2.EutracellIdentity) (*C.EUTRACellIdentity_RCPRE_t, error) {
	eUTRACellIdentityC, err := newBitString(eUTRACellIdentity.Value)
	if err != nil {
		return nil, err
	}

	return eUTRACellIdentityC, nil
}

func decodeEUTRACellIdentity(eUTRACellIdentityC *C.EUTRACellIdentity_RCPRE_t) (*e2sm_rc_pre_v2.EutracellIdentity, error) {
	eUTRACellIdentity := new(e2sm_rc_pre_v2.EutracellIdentity)

	eUTRACellIdentityBs, err := decodeBitString(eUTRACellIdentityC)
	if err != nil {
		return nil, fmt.Errorf("decodeEUTRACellIdentity() %s", err.Error())
	}
	eUTRACellIdentity.Value = eUTRACellIdentityBs

	return eUTRACellIdentity, nil
}
