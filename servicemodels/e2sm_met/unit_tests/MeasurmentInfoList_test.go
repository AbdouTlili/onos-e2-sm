package met

import (
	"encoding/hex"

	// "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"

	// hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"testing"

	"gotest.tools/assert"
)

var perRefBytesMIL = []byte{0x00, 0x01, 0x02, 0x6d, 0x63, 0x73, 0x02, 0x70, 0x68, 0x72}

func createMeasurementInfoList() (*e2smmet.MeasurementInfoList, error) {

	res := &e2smmet.MeasurementInfoList{
		Value: make([]*e2smmet.MeasurementInfoItem, 0),
	}
	mii, err := createMeasurementInfoItem("mcs")
	if err != nil {
		return nil, err
	}
	res.Value = append(res.Value, mii)

	mii2, err := createMeasurementInfoItem("phr")
	if err != nil {
		return nil, err
	}
	res.Value = append(res.Value, mii2)

	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func Test_perEncodeMeasurementInfoList(t *testing.T) {

	mii, err := createMeasurementInfoList()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "", e2smmet.E2smMetChoicemap, nil)

	// fmt.Printf("bydfgggggtes len %d \n --Perbytes  : %#v", len(per), per)

	assert.NilError(t, err)
	t.Logf("MeasurementInfo-List PER\n%v", hex.Dump(per))

	result := e2smmet.MeasurementInfoItem{}
	err = aper.UnmarshalWithParams(per, &result, "", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfo-List PER - decoded\n%v", &result)
	// assert.DeepEqual(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
}

func Test_perMeasurementInfoListCompareBytes(t *testing.T) {

	mii, err := createMeasurementInfoList()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-List PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	assert.DeepEqual(t, per, perRefBytesMIL)
}
