// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2ctypes

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-v2"
	"gotest.tools/assert"
	"testing"
)

func Test_xerEncodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x30},
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0xc0},
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x38},
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID)
	ngeNbID.GetNgENb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	//assert.Equal(t, 478, len(xer))
	//assert.Equal(t, 448, len(xer)) // without GNbDuID
	t.Logf("GlobalKpmnodeNgEnbID XER\n%s", string(xer))
}

func Test_xerDecodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x30},
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0xc0},
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x38},
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID)
	ngeNbID.GetNgENb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	xer, err := xerEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	//assert.Equal(t, 478, len(xer))
	//assert.Equal(t, 448, len(xer)) // without GNbDuID
	t.Logf("GlobalKpmnodeNgEnbID XER\n%s", string(xer))

	result, err := xerDecodeGlobalKpmnodeNgEnbID(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeNgEnbID XER - decoded\n%s", result)
}

func Test_perEncodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x30},
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0xc0},
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x38},
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID)
	ngeNbID.GetNgENb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	//assert.Equal(t, 15, len(per))
	//assert.Equal(t, 14, len(per)) // without GNbDuID
	t.Logf("GlobalKpmnodeNgEnbID PER\n%v", hex.Dump(per))
}

func Test_perDecodeGlobalKpmnodeNgEnbID(t *testing.T) {

	bs := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x30},
		Len:   20,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	shortMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0xc0},
		Len:   18,
	}
	longMacroEnbID := e2sm_kpm_v2.BitString{
		Value: []byte{0xd4, 0xbc, 0x38},
		Len:   21,
	}
	var gnbDuID int64 = 42

	ngeNbID, err := pdubuilder.CreateGlobalKpmnodeIDngENbID(&bs, plmnID, &shortMacroEnbID, &longMacroEnbID)
	ngeNbID.GetNgENb().GNbDuId = &e2sm_kpm_v2.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	per, err := perEncodeGlobalKpmnodeNgEnbID(ngeNbID.GetNgENb())
	assert.NilError(t, err)
	//assert.Equal(t, 15, len(per))
	//assert.Equal(t, 14, len(per)) // without GNbDuID
	t.Logf("GlobalKpmnodeNgEnbID PER\n%v", hex.Dump(per))

	result, err := perDecodeGlobalKpmnodeNgEnbID(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("GlobalKpmnodeNgEnbID PER - decode\n%v", result)
}
