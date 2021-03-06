// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmctypes

import (
	e2sm_kpm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/v1beta1/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodePlmnIdentity(t *testing.T) {

	var plmnID = "ONF"

	plmnIdentity := &e2sm_kpm_ies.PlmnIdentity{
		Value: []byte(plmnID),
	}

	xer, err := xerEncodePlmnIdentity(plmnIdentity)
	assert.NilError(t, err)
	t.Logf("PLMN-Identity XER\n%s", string(xer))
}
