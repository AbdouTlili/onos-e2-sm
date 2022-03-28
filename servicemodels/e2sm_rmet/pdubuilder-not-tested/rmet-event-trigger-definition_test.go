package pdubuilder

import (
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRmetEventTriggerDefinition(t *testing.T) {
	var rtPeriod int64 = 15

	newE2SmRmetPdu, err := CreateE2SmRmetEventTriggerDefinition(rtPeriod)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRmetPdu != nil)
}
