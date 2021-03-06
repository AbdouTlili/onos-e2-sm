// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

import (
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeNRCGI(t *testing.T) {

	var plmnIDnrcgi = "ONF"

	nrCGI := &e2sm_rc_pre_v2.Nrcgi{
		PLmnIdentity: &e2sm_rc_pre_v2.PlmnIdentity{
			Value: []byte(plmnIDnrcgi),
		},
		NRcellIdentity: &e2sm_rc_pre_v2.NrcellIdentity{
			Value: &e2sm_rc_pre_v2.BitString{
				Value: []byte{0xd4, 0xbc, 0x90},
				Len:   22, //uint32
			},
		},
	}

	xer, err := xerEncodeNRCGI(nrCGI)
	assert.NilError(t, err)
	t.Logf("NRCGI XER\n%s", string(xer))
}
