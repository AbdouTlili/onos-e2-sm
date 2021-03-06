// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package testsmctypes

import (
	"encoding/hex"
	test_sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/test_sm_aper_go_lib/v1/test-sm-ies"
	"gotest.tools/assert"
	"testing"
)

func createConstrainedChoice2Msg() (*test_sm_ies.ConstrainedChoice2, error) {

	constrainedChoice2 := test_sm_ies.ConstrainedChoice2{
		//ConstrainedChoice2: &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2A{
		//	ConstrainedChoice2A: 15,
		//},
		ConstrainedChoice2: &test_sm_ies.ConstrainedChoice2_ConstrainedChoice2B{
			ConstrainedChoice2B: 156,
		},
	}

	return &constrainedChoice2, nil
}

func Test_xerEncodingConstrainedChoice2(t *testing.T) {

	constrainedChoice2, err := createConstrainedChoice2Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice2 PDU")

	xer, err := xerEncodeConstrainedChoice2(constrainedChoice2)
	assert.NilError(t, err)
	t.Logf("ConstrainedChoice2 XER\n%s", string(xer))

	result, err := xerDecodeConstrainedChoice2(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice2 XER - decoded\n%v", result)
	//assert.Equal(t, constrainedChoice2.GetConstrainedChoice2A(), result.GetConstrainedChoice2A())
	assert.Equal(t, constrainedChoice2.GetConstrainedChoice2B(), result.GetConstrainedChoice2B())
}

func Test_perEncodingConstrainedChoice2(t *testing.T) {

	constrainedChoice2, err := createConstrainedChoice2Msg()
	assert.NilError(t, err, "Error creating ConstrainedChoice2 PDU")

	per, err := perEncodeConstrainedChoice2(constrainedChoice2)
	assert.NilError(t, err)
	t.Logf("ConstrainedChoice2 PER\n%v", hex.Dump(per))

	result, err := perDecodeConstrainedChoice2(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("ConstrainedChoice2 PER - decoded\n%v", result)
	//assert.Equal(t, constrainedChoice2.GetConstrainedChoice2A(), result.GetConstrainedChoice2A())
	assert.Equal(t, constrainedChoice2.GetConstrainedChoice2B(), result.GetConstrainedChoice2B())
}
