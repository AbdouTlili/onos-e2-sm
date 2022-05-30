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

var perRefBytesMII = []byte{0x03, 0x74, 0x65, 0x73, 0x74}

func createMeasurementInfoItem(name string) (*e2smmet.MeasurementInfoItem, error) {

	res := &e2smmet.MeasurementInfoItem{
		Value: name,
	}
	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func Test_perEncodeMeasurementInfoItem(t *testing.T) {

	mii, err := createMeasurementInfoItem("mcs")
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "", e2smmet.E2smMetChoicemap, nil)

	// fmt.Printf("bydfgggggtes len %d \n --Perbytes  : %#v", len(per), per)

	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	result := e2smmet.MeasurementInfoItem{}
	err = aper.UnmarshalWithParams(per, &result, "", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementInfo-Item PER - decoded\n%v", &result)
	// assert.DeepEqual(t, mii.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue(), result.GetLabelInfoList().GetValue()[0].GetMeasLabel().GetPlmnId().GetValue())
}

func Test_perMeasurementInfoItemCompareBytes(t *testing.T) {

	mii, err := createMeasurementInfoItem("test")
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mii, "", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementInfo-Item PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	assert.DeepEqual(t, per, perRefBytesMII)
}
