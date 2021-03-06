// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmKpmEventTriggerDefinition(t *testing.T) {
	var rtPeriod int32 = 12 // range is from 0 to 19

	newE2SmKpmPdu, err := CreateE2SmKpmEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
