// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package met

import (
	"encoding/hex"
	"fmt"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	// hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var perRefBytes = []byte{0x00, 0x32}

/*  we dont text the no value use case  */
//var refPerMRIreal = "0000000  20 03 80 01 0b                                    | ....|"
var refPerMRInovalue = "00000000  40                                                |@|"

func createMeasurementRecordItemInteger() *e2smmet.MeasurementRecordItem {
	return &e2smmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smmet.MeasurementRecordItem_Integer{
			Integer: 50,
		},
	}
}

//func createMeasurementRecordItemReal() *e2smmet.MeasurementRecordItem {
//	return &e2smmet.MeasurementRecordItem{
//		MeasurementRecordItem: &e2smmet.MeasurementRecordItem_Real{
//			Real: 22,
//		},
//	}
//}

func createMeasurementRecordItemNoValue() *e2smmet.MeasurementRecordItem {
	return &e2smmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smmet.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
}

func Test_perEncodingMeasurementRecordItemInteger(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mri, "valueExt", e2smmet.E2smMetChoicemap, nil)
	fmt.Printf("%#v--", per)
	assert.NilError(t, err)
	t.Logf("MeasurementRecordItem (Integer) PER\n%v", hex.Dump(per))

	result := e2smmet.MeasurementRecordItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementRecordItem (Integer) PER - decoded\n%v", &result)
	assert.Equal(t, mri.GetInteger(), result.GetInteger())
}

func Test_perMeasurementRecordItemIntegerCompareBytes(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mri, "valueExt", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementRecordItem (Integer) PER\n%v", hex.Dump(per))

	// //Comparing with reference bytes
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}

//func Test_perEncodingMeasurementRecordItemReal(t *testing.T) {
//
//	mri := createMeasurementRecordItemReal()
//
//	aper.ChoiceMap = e2smmet.Choicemape2smKpm
//	per, err := aper.MarshalWithParams(mri, "valueExt")
//	assert.NilError(t, err)
//	t.Logf("MeasurementRecordItem (Real) PER\n%v", hex.Dump(per))
//
//	result := e2smmet.MeasurementRecordItem{}
//	err = aper.UnmarshalWithParams(per, &result, "valueExt")
//	assert.NilError(t, err)
//	assert.Assert(t, &result != nil)
//	t.Logf("MeasurementRecordItem (Real) PER - decoded\n%v", &result)
//	assert.Equal(t, mri.GetReal(), result.GetReal())
//}
//
//func Test_perMeasurementRecordItemRealCompareBytes(t *testing.T) {
//
//	mri := createMeasurementRecordItemReal()
//
//	aper.ChoiceMap = e2smmet.Choicemape2smKpm
//	per, err := aper.MarshalWithParams(mri, "valueExt")
//	assert.NilError(t, err)
//	t.Logf("MeasurementRecordItem (Real) PER\n%v", hex.Dump(per))
//
//	//Comparing with reference bytes
//	perRefBytes, err := hexlib.DumpToByte(refPerMRIreal)
//	assert.NilError(t, err)
//	assert.DeepEqual(t, per, perRefBytes)
//}

// func Test_perEncodingMeasurementRecordItemNull(t *testing.T) {

// 	mri := createMeasurementRecordItemNoValue()

// 	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
// 	per, err := aper.MarshalWithParams(mri, "valueExt", e2smmet.Choicemape2smKpm, nil)
// 	assert.NilError(t, err)
// 	t.Logf("MeasurementRecordItem (No value) PER\n%v", hex.Dump(per))

// 	result := e2smmet.MeasurementRecordItem{}
// 	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.Choicemape2smKpm, nil)
// 	assert.NilError(t, err)
// 	//assert.Assert(t, &result != nil)
// 	t.Logf("MeasurementRecordItem (No value) PER - decoded\n%v", &result)
// 	assert.Equal(t, mri.GetNoValue(), result.GetNoValue())
// }

// func Test_perMeasurementRecordItemNullCompareBytes(t *testing.T) {

// 	mri := createMeasurementRecordItemNoValue()

// 	//aper.ChoiceMap = e2smmet.Choicemape2smKpm
// 	per, err := aper.MarshalWithParams(mri, "valueExt", e2smmet.Choicemape2smKpm, nil)
// 	assert.NilError(t, err)
// 	t.Logf("MeasurementRecordItem (No value) PER\n%v", hex.Dump(per))

// 	//Comparing with reference bytes
// 	perRefBytes, err := hexlib.DumpToByte(refPerMRInovalue)
// 	assert.NilError(t, err)
// 	assert.DeepEqual(t, per, perRefBytes)
// }
