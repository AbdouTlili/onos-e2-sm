// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRmetIndicationHeader(t *testing.T) {
	bs := asn1.BitString{
		Value: []byte{0x9b, 0xcd, 0x40},
		Len:   22,
	}
	plmnID := []byte{0x21, 0x22, 0x23}
	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	var fileFormatVersion = "txt"
	var senderName = "ONF"
	var senderType = "someType"
	var vendorName = "onf"

	globalRmetNodeID, err := CreateGlobalRmetnodeIDgNBID(&bs, plmnID)
	globalRmetNodeID.GetGNb().GNbCuUpId = &e2smrmet.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalRmetNodeID.GetGNb().GNbDuId = &e2smrmet.GnbDuId{
		Value: gnbDuID,
	}
	assert.NilError(t, err)

	newE2SmRmetPdu, err := CreateE2SmRmetIndicationHeader(timeStamp)
	assert.NilError(t, err)
	newE2SmRmetPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetSenderType(senderType).SetVendorName(vendorName).SetGlobalRmetnodeID(globalRmetNodeID)
	assert.Assert(t, newE2SmRmetPdu != nil)
}
