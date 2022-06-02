// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package met

import (
	"encoding/hex"
	// "fmt"

	// "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"

	// hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"testing"

	"gotest.tools/assert"
)

var perRefBytesMeasRec = []byte{0x00, 0x01, 0x00, 0x41, 0x00, 0x01, 0x00, 0x39, 0x01, 0x31, 0x30}

func createMeasurementRecord() (*e2smmet.MeasurementRecord, error) {
	var a int64 = 1
	res := &e2smmet.MeasurementRecord{
		MeasRecordItemList: &e2smmet.MeasurementRecordItemList{},
		UeId: &e2smmet.Ueid{
			Value: a,
		},
		UeTag: &e2smmet.Uetag{
			Value: "A",
		},
	}

	item1 := pdubuilder.CreateMeasurementRecordItemInteger(9)
	item2 := pdubuilder.CreateMeasurementRecordItemInteger(10)

	res.MeasRecordItemList.Value = append(res.MeasRecordItemList.Value, item1, item2)

	if err := res.Validate(); err != nil {
		// fmt.Printf("---------no validate")
		return nil, err
	}
	return res, nil
}

func Test_perEncodingMeasurementRecord(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smmet.E2smMetChoicemap
	per, err := aper.MarshalWithParams(mr, "valueExt", e2smmet.E2smMetChoicemap, nil)

	// fmt.Printf("bytes len %d \n --Perbytes  : %#v", len(per), per)
	assert.NilError(t, err)

	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

	result := e2smmet.MeasurementRecord{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	// //assert.Assert(t, &result != nil)
	t.Logf("MeasurementRecord PER - decoded\n%v", &result)
	assert.Equal(t, 2, len(mr.MeasRecordItemList.Value))
	assert.Equal(t, mr.GetUeId().GetValue(), result.GetUeId().GetValue())
	assert.Equal(t, mr.GetUeTag().GetValue(), result.GetUeTag().GetValue())
	assert.Equal(t, mr.MeasRecordItemList.GetValue()[0].GetValue(), result.MeasRecordItemList.GetValue()[0].GetValue())
}

func Test_perMeasurementRecordCompareBytes(t *testing.T) {

	mr, err := createMeasurementRecord()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smmet.E2smMetChoicemap
	per, err := aper.MarshalWithParams(mr, "valueExt", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

	// Comparing with reference bytes
	// perRefBytes, err := hexlib.DumpToByte(refPerMeasRecordNoReal)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytesMeasRec)
}
