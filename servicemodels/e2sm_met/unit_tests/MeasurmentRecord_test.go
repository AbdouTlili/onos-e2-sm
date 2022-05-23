// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package met

import (
	// "encoding/hex"
	"fmt"
	// "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	// hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

//var refPerMeasRecord = "00000000  03 00 15 20 03 80 01 0b  40                       |... ....@|"
var perRefBytesMeasRec = []byte{0x40,0x0b,0x00,0xa0,0x80,0x00,0x00,0x00,0x0a}

// func createMeasurementRecord() (*e2smmet.MeasurementRecord, error) {
// 	res := &e2smmet.MeasurementRecord{
// 		MeasRecordItem: make([]*e2smmet.MeasurementRecordItem, 0),
// 		UeId: &e2smmet.Ueid{
// 			Value: 11,
// 		},
// 		UeTag: &e2smmet.Uetag{
// 			Value: "AAAA",
// 		},
// 	}

// 	item1 := &e2smmet.MeasurementRecordItem{
// 		MeasurementRecordItem: &e2smmet.MeasurementRecordItem_Integer{
// 			Integer: 10,
// 		},
// 	}
// 	res.MeasRecordItem = append(res.MeasRecordItem, item1)

// 	//ToDo - bring back once handling of REAL types is implemented
// 	//item2 := &e2smmet.MeasurementRecordItem{
// 	//	MeasurementRecordItem: &e2smmet.MeasurementRecordItem_Real{
// 	//		Real: 22,
// 	//	},
// 	//}
// 	//res.Value = append(res.Value, item2)

// 	if err := res.Validate(); err != nil {
// 		fmt.Printf("---------no validate")
// 		return nil, err
// 	}
// 	return res, nil
// }

func Test_perEncodingMeasurementRecord(t *testing.T) {

	result := e2smmet.MeasurementRecord{}
	err := aper.UnmarshalWithParams(perRefBytesMeasRec, &result, "", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	fmt.Printf("--bytes : %#v",  result)


}

// 	mr, err := createMeasurementRecord()
// 	assert.NilError(t, err)

// 	//aper.ChoiceMap = e2smmet.E2smMetChoicemap
// 	per, err := aper.MarshalWithParams(mr, "", e2smmet.E2smMetChoicemap, nil)
// 	fmt.Printf("bytes len %d \n --bytes : %#v", len(per), per)
// 	assert.NilError(t, err)
// 	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

// 	result := e2smmet.MeasurementRecord{}
// 	err = aper.UnmarshalWithParams(perRefBytesMeasRec, &result, "", e2smmet.E2smMetChoicemap, nil)
// 	assert.NilError(t, err)
// 	// //assert.Assert(t, &result != nil)
// 	t.Logf("MeasurementRecord PER - decoded\n%v", &result)
// 	assert.Equal(t, 1, len(mr.MeasRecordItem))
// 	assert.Equal(t, mr.GetUeId().GetValue(), result.GetUeId().GetValue())
// 	assert.Equal(t, mr.GetUeTag().GetValue(), result.GetUeTag().GetValue())

// 	// //assert.Equal(t, mr.GetValue()[1].GetReal(), result.GetValue()[1].GetReal())
// 	assert.Equal(t, mr.GetMeasRecordItem()[1].GetInteger(), result.GetMeasRecordItem()[1].GetInteger())
// }

// func Test_perMeasurementRecordCompareBytes(t *testing.T) {

// 	mr, err := createMeasurementRecord()
// 	assert.NilError(t, err)

// 	//aper.ChoiceMap = e2smmet.E2smMetChoicemap
// 	per, err := aper.MarshalWithParams(mr, "", e2smmet.E2smMetChoicemap, nil)
// 	assert.NilError(t, err)
// 	t.Logf("MeasurementRecord PER\n%v", hex.Dump(per))

// 	//Comparing with reference bytes
// 	// perRefBytes, err := hexlib.DumpToByte(refPerMeasRecordNoReal)
// 	// assert.NilError(t, err)
// 	// assert.DeepEqual(t, per, perRefBytesMeasRec)
// }
