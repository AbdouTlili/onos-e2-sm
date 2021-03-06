// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/encoder"
	e2smrcprev2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go/v2/e2sm-rc-pre-v2-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreIndicationMsg(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)
	dlArfcn, err := CreateEArfcn(253)
	assert.NilError(t, err)
	var pci int32 = 11
	cellSize := e2smrcprev2.CellSize_CELL_SIZE_MACRO

	cellID := asn1.BitString{
		Value: []byte{0xac, 0xd4, 0xbc, 0x90},
		Len:   28, //uint32
	}
	cgi, err := CreateCellGlobalIDEUTRACGI(plmnIDBytes, &cellID)
	assert.NilError(t, err)

	neighbor, err := CreateNrt(cgi, dlArfcn, cellSize, &e2smrcprev2.Pci{
		Value: pci,
	})
	assert.NilError(t, err)

	neighbors := make([]*e2smrcprev2.Nrt, 0)
	neighbors = append(neighbors, neighbor)
	neighbors = append(neighbors, neighbor)

	earfcn, err := CreateEArfcn(253)
	assert.NilError(t, err)

	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes, earfcn, cellSize, pci, neighbors)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err := encoder.PerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Indication Message: \n%v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreIndicationMessage(per)
	assert.NilError(t, err)
	t.Logf("PER decoded RC-PRE-IndicationMessage is \n%v", resultPer)
	assert.Equal(t, newE2SmRcPrePdu.String(), resultPer.String())
}

func TestE2SmRcPreIndicationMsgNoNeighbors(t *testing.T) {
	var plmnID = "12f410"
	plmnIDBytes, err := hex.DecodeString(plmnID)
	assert.NilError(t, err)
	var pci int32 = 11
	cellSize := e2smrcprev2.CellSize_CELL_SIZE_MACRO

	earfcn, err := CreateEArfcn(253)
	assert.NilError(t, err)

	newE2SmRcPrePdu, err := CreateE2SmRcPreIndicationMsgFormat1(plmnIDBytes, earfcn, cellSize, pci, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)

	per, err := encoder.PerEncodeE2SmRcPreIndicationMessage(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("PER Encoded Indication Message: \n%v", hex.Dump(per))

	resultPer, err := encoder.PerDecodeE2SmRcPreIndicationMessage(per)
	assert.NilError(t, err)
	t.Logf("PER decoded RC-PRE-IndicationMessage is \n%v", resultPer)
	assert.Equal(t, newE2SmRcPrePdu.String(), resultPer.String())
}
