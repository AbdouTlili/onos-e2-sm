// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"testing"

	"gotest.tools/assert"
)

func TestE2SmMetIndicationHeader(t *testing.T) {

	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}

	var fileFormatVersion = "txt"
	var senderName = "ONF"

	globalMetNodeID, err := CreateGlobalMetnodeID(15)
	assert.NilError(t, err)

	newE2SmMetPdu, err := CreateE2SmMetIndicationHeader(timeStamp)
	assert.NilError(t, err)
	newE2SmMetPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetGlobalMETnodeID(globalMetNodeID)
	assert.Assert(t, newE2SmMetPdu != nil)
}
