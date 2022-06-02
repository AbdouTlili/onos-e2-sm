// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package met

import (
	"encoding/hex"
	// "fmt"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"

	// hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"testing"

	"gotest.tools/assert"
)

var perRefBytes = []byte{0x00, 0x32}

func createMeasurementRecordItemInteger() *e2smmet.MeasurementRecordItem {
	return &e2smmet.MeasurementRecordItem{
		Value: "99",
	}
}

func Test_perEncodingMeasurementRecordItemInteger(t *testing.T) {

	mri := createMeasurementRecordItemInteger()

	per, err := aper.MarshalWithParams(mri, "", e2smmet.E2smMetChoicemap, nil)
	// fmt.Printf("%#v--", per)
	assert.NilError(t, err)
	t.Logf("MeasurementRecordItem (Integer) PER\n%v", hex.Dump(per))

	result := e2smmet.MeasurementRecordItem{}
	err = aper.UnmarshalWithParams(per, &result, "", e2smmet.E2smMetChoicemap, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementRecordItem (Integer) PER - decoded\n%v", &result)
	assert.Equal(t, mri.GetValue(), result.GetValue())
}

// func Test_perMeasurementRecordItemIntegerCompareBytes(t *testing.T) {

// 	mri := createMeasurementRecordItemInteger()

// 	per, err := aper.MarshalWithParams(mri, "", e2smmet.E2smMetChoicemap, nil)
// 	assert.NilError(t, err)
// 	t.Logf("MeasurementRecordItem (Integer) PER\n%v", hex.Dump(per))

// 	// //Comparing with reference bytes
// 	assert.NilError(t, err)
// 	assert.DeepEqual(t, per, perRefBytes)
// }

// func Test_perEncodingMeasurementRecordItemString(t *testing.T) {

// 	mri := createMeasurementRecordItemString()
// 	fmt.Println(mri.GetStr())
// 	per, err := aper.MarshalWithParams(mri, "", e2smmet.E2smMetChoicemap, nil)
// 	fmt.Printf("%#v--", per)
// 	assert.NilError(t, err)
// 	// t.Logf("MeasurementRecordItem (Integer) PER\n%v", hex.Dump(per))

// 	// result := e2smmet.MeasurementRecordItem{}
// 	// err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smmet.E2smMetChoicemap, nil)
// 	// assert.NilError(t, err)
// 	// //assert.Assert(t, &result != nil)
// 	// t.Logf("MeasurementRecordItem (Integer) PER - decoded\n%v", &result)
// 	// assert.Equal(t, mri.GetInteger(), result.GetInteger())
// }
