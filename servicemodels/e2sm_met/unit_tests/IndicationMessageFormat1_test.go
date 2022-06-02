package met

import (
	// "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"

	"encoding/hex"
	"fmt"
	"testing"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"gotest.tools/assert"
	// hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
)

var aperBytesIndMes1 = []byte{0x00, 0x09, 0x00, 0x01, 0x00, 0x01, 0x00,
	0x41, 0x00, 0x01, 0x00, 0x39, 0x01, 0x31, 0x30, 0x00, 0x01, 0x00,
	0x41, 0x00, 0x01, 0x00, 0x39, 0x01, 0x31, 0x30}

func createIndicationMessage1() (*e2smmet.E2SmMetIndicationMessageFormat1, error) {
	mr1, err := createMeasurementRecord()

	if err != nil {
		return &e2smmet.E2SmMetIndicationMessageFormat1{}, err
	}

	mr2, err := createMeasurementRecord()
	if err != nil {
		return &e2smmet.E2SmMetIndicationMessageFormat1{}, err
	}

	mdata := e2smmet.MeasurementData{
		Value: make([]*e2smmet.MeasurementRecord, 0),
	}

	mdata.Value = append(mdata.Value, mr1, mr2)

	ind1 := e2smmet.E2SmMetIndicationMessageFormat1{
		SubscriptId: &e2smmet.SubscriptionId{Value: 10},
		MeasData:    &mdata,
	}

	return &ind1, nil
}

func Test_perEncodingIndHeader1(t *testing.T) {
	imes1, err := createIndicationMessage1()
	assert.NilError(t, err)

	per, err := aper.MarshalWithParams(imes1, "valueExt", e2smmet.E2smMetChoicemap, nil)

	fmt.Printf("bytes len %d \n --Perbytes for indMessage1 : %#v", len(per), per)
	assert.NilError(t, err)

	t.Logf("indication message1 PER\n%v", hex.Dump(per))

	result := e2smmet.E2SmMetIndicationMessageFormat1{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	// //assert.Assert(t, &result != nil)
	t.Logf("indication message1 PER - decoded\n%v", &result)
	assert.Equal(t, 2, len(result.MeasData.Value))
	assert.Equal(t, imes1.MeasData.Value[0].GetUeId().GetValue(), result.MeasData.Value[0].GetUeId().GetValue())
	assert.Equal(t, imes1.MeasData.Value[0].GetUeTag().GetValue(), result.MeasData.Value[0].GetUeTag().GetValue())
	assert.Equal(t, imes1.MeasData.Value[0].MeasRecordItemList.GetValue()[0].GetValue(), result.MeasData.Value[0].MeasRecordItemList.GetValue()[0].GetValue())
}

func Test_perIndicationMessage1CompareBytes(t *testing.T) {

	imes1, err := createIndicationMessage1()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smmet.E2smMetChoicemap
	per, err := aper.MarshalWithParams(imes1, "valueExt", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

	// Comparing with reference bytes
	// perRefBytes, err := hexlib.DumpToByte(refPerMeasRecordNoReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, aperBytesIndMes1)
}
