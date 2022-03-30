// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package met

import (
	"testing"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/encoder"
	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"gotest.tools/assert"
)

var refPerE2SmMetIndicationHeader = "00000000  1f 21 22 23 24 18 74 78  74 00 00 03 4f 4e 46 40  |.!\"#$.txt...ONF@|\n" +
	"00000010  73 6f 6d 65 54 79 70 65  06 6f 6e 66 0c 37 34 37  |someType.onf.747|\n" +
	"00000020  00 d4 bc 08 80 30 39 20  1a 85                    |.....09 ..|"

func createE2SmMetIndicationHeader() (*e2smmet.E2SmMetIndicationHeader, error) {

	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	var fileFormatVersion = "txt"
	var senderName = "ONF"

	globalMetNodeID, err := pdubuilder.CreateGlobalMetnodeID(15)

	newE2SmMetPdu, err := pdubuilder.CreateE2SmMetIndicationHeader(timeStamp)

	if err != nil {
		return nil, err
	}
	newE2SmMetPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetGlobalMETnodeID(globalMetNodeID)

	return newE2SmMetPdu, nil
}

func Test_perEncodingE2SmMetIndicationHeader(t *testing.T) {

	ih, err := createE2SmMetIndicationHeader()
	// fmt.Println(ih)
	assert.NilError(t, err)

	_, err = encoder.PerEncodeE2SmMetIndicationHeader(ih)
	assert.NilError(t, err)
	// t.Logf("E2SM-MET-IndicationHeader PER\n%v", hex.Dump(per))

	// result, err := encoder.PerDecodeE2SmMetIndicationHeader(per)
	// assert.NilError(t, err)
	// assert.Assert(t, result != nil)
	// t.Logf("E2SM-MET-IndicationHeader PER - decoded\n%v", result)
	// assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetFileFormatversion(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetFileFormatversion())
	// assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGNbDuId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGNbDuId().GetValue())
	// assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGNbCuUpId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGNbCuUpId().GetValue())
	// assert.DeepEqual(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGlobalGNbId().GetPlmnId().GetValue())
	// assert.DeepEqual(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetValue())
	// assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetGNb().GetGlobalGNbId().GetGnbId().GetGnbId().GetLen())
	// assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderName(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderName())
	// assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderType(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetSenderType())
	// assert.DeepEqual(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetColletStartTime().GetValue(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetColletStartTime().GetValue())
	// assert.Equal(t, ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetVendorName(), result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetVendorName())
}

// func Test_perE2SmMetIndicationHeaderCompareBytes(t *testing.T) {

// 	ih, err := createE2SmMetIndicationHeader()
// 	assert.NilError(t, err)

// 	per, err := encoder.PerEncodeE2SmMetIndicationHeader(ih)
// 	assert.NilError(t, err)
// 	t.Logf("E2SM-MET-IndicationHeader PER\n%v", hex.Dump(per))

// 	//Comparing with reference bytes
// 	perRefBytes, err := hexlib.DumpToByte(refPerE2SmMetIndicationHeader)
// 	assert.NilError(t, err)
// 	assert.DeepEqual(t, per, perRefBytes)
// }
