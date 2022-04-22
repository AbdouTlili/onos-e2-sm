// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"testing"

	"gotest.tools/assert"
)

func TestE2SmMetEventTriggerDefinition(t *testing.T) {
	var rtPeriod int64 = 15

	newE2SmMetPdu, err := CreateE2SmMetEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMetPdu != nil)
}
