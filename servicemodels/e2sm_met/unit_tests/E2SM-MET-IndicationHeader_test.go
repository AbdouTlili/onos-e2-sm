package met

import (
	"encoding/hex"
	// "fmt"

	// "fmt"
	"testing"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/encoder"
	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"gotest.tools/assert"
)

var refPerE2SmMetIndicationHeader = []byte{0x04, 0x62, 0x94, 0xb9, 0x22, 0x00, 0x01, 0x02, 0x6d, 0x63, 0x73, 0x02, 0x70, 0x68, 0x72, 0x00, 0x09}

func createE2SmMetIndicationHeader() (*e2smmet.E2SmMetIndicationHeader, error) {

	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	// var fileFormatVersion = "txt"
	// var senderName = "ONF"

	globalMetNodeID, err := pdubuilder.CreateGlobalMetnodeID(10)

	mil, _ := createMeasurementInfoList()

	newE2SmMetPdu, _ := pdubuilder.CreateE2SmMetIndicationHeader(timeStamp, mil)

	if err != nil {
		return nil, err
	}
	// newE2SmMetPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetGlobalMETnodeID(globalMetNodeID)

	newE2SmMetPdu.SetGlobalMETnodeID(globalMetNodeID)

	return newE2SmMetPdu, nil
}

func Test_perEncodingE2SmMetIndicationHeader(t *testing.T) {

	ih, err := createE2SmMetIndicationHeader()
	// fmt.Println(ih)
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmMetIndicationHeader(ih)
	assert.NilError(t, err)
	t.Logf("E2SM-MET-IndicationHeader PER\n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMetIndicationHeader(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SM-MET-IndicationHeader PER - decoded\n%v", result)

	// fmt.Println(result.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMeasInfoList().Value[1].Value)

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

func Test_perE2SmMetIndicationHeaderCompareBytes(t *testing.T) {

	ih, err := createE2SmMetIndicationHeader()
	assert.NilError(t, err)

	per, err := encoder.PerEncodeE2SmMetIndicationHeader(ih)
	assert.NilError(t, err)
	// fmt.Printf("bydfgggggtes len %d \n --Perbytes  : %#v", len(per), per)
	t.Logf("E2SM-MET-IndicationHeader PER\n%v", hex.Dump(per))

	_, err = encoder.PerDecodeE2SmMetIndicationHeader(refPerE2SmMetIndicationHeader)
	assert.NilError(t, err)

	// fmt.Println(indH.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().MeasInfoList.Value[0].Value)

	//Comparing with reference bytes
	// perRefBytes, err := hexlib.DumpToByte(refPerE2SmMetIndicationHeader)
	// assert.NilError(t, err)
	// assert.DeepEqual(t, per, perRefBytes)
}
