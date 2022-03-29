// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"testing"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/encoder"
	"gotest.tools/assert"
)

func TestE2SmMetRanfunctionDescription(t *testing.T) {
	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"

	newE2SmMetPdu, err := CreateE2SmMetRanfunctionDescription(rfSn, rfE2SMoid, rfd)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMetPdu != nil)

	per, err := encoder.PerEncodeE2SmMetRanFunctionDescription(newE2SmMetPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RANfunctionDescription PER is \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMetRanFunctionDescription(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RANfunctionDescription PER - decoded\n%v", result)
}
