// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package rcprectypes

import (
	"encoding/hex"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func createRanparameterValueValueInteger() *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueInt{
			ValueInt: 44,
		},
	}
}

func createRanparameterValueValueEnum() *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueEnum{
			ValueEnum: 98,
		},
	}
}

func createRanparameterValueValueBool() *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueBool{
			ValueBool: true,
		},
	}
}

func createRanparameterValueValueOctS() *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueOctS{
			ValueOctS: "onf",
		},
	}
}

func createRanparameterValueValuePrtS() *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValuePrtS{
			ValuePrtS: "onf",
		},
	}
}

func createRanparameterValueValueBitS() *e2sm_rc_pre_v2.RanparameterValue {
	return &e2sm_rc_pre_v2.RanparameterValue{
		RanparameterValue: &e2sm_rc_pre_v2.RanparameterValue_ValueBitS{
			ValueBitS: &e2sm_rc_pre_v2.BitString{
				Value: []byte{0x4d, 0xcb, 0x90},
				Len:   22,
			},
		},
	}
}

func Test_xerEncodeRanParameterValue(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	xer, err := xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	t.Logf("RANparameterValue (Integer) XER\n%s", string(xer))

	rpv = createRanparameterValueValueEnum()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 73, len(xer))
	t.Logf("RANparameterValue (Enum) XER\n%s", string(xer))

	rpv = createRanparameterValueValueBool()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 78, len(xer))
	t.Logf("RANparameterValue (Bool) XER\n%s", string(xer))

	rpv = createRanparameterValueValueOctS()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 79, len(xer))
	t.Logf("RANparameterValue (OctetString) XER\n%s", string(xer))

	rpv = createRanparameterValueValuePrtS()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 74, len(xer))
	t.Logf("RANparameterValue (PrintableString) XER\n%s", string(xer))

	rpv = createRanparameterValueValueBitS()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 107, len(xer))
	t.Logf("RANparameterValue (BitString) XER\n%s", string(xer))
}

func Test_xerDecodeRanParameterValue(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	xer, err := xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 71, len(xer))
	t.Logf("RANparameterValue (Integer) XER\n%s", string(xer))

	result, err := xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Integer) XER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueInt(), result.GetValueInt())

	rpv = createRanparameterValueValueEnum()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 73, len(xer))
	t.Logf("RANparameterValue (Enum) XER\n%s", string(xer))

	result, err = xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Enum) XER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueEnum(), result.GetValueEnum())

	rpv = createRanparameterValueValueBool()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 78, len(xer))
	t.Logf("RANparameterValue (Bool) XER\n%s", string(xer))

	result, err = xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Bool) XER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueBool(), result.GetValueBool())

	rpv = createRanparameterValueValueOctS()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 79, len(xer))
	t.Logf("RANparameterValue (OctetString) XER\n%s", string(xer))

	result, err = xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (OctetString) XER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueOctS(), result.GetValueOctS())

	rpv = createRanparameterValueValuePrtS()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 74, len(xer))
	t.Logf("RANparameterValue (PrintableString) XER\n%s", string(xer))

	result, err = xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (PrintableString) XER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValuePrtS(), result.GetValuePrtS())

	rpv = createRanparameterValueValueBitS()

	xer, err = xerEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	//assert.Equal(t, 107, len(xer))
	t.Logf("RANparameterValue (BitString) XER\n%s", string(xer))

	result, err = xerDecodeRanparameterValue(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (BitString) XER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueBitS().GetLen(), result.GetValueBitS().GetLen())
	assert.DeepEqual(t, rpv.GetValueBitS().GetValue(), result.GetValueBitS().GetValue())
}

func Test_perEncodeRanParameterValue(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	per, err := perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RANparameterValue (Integer) PER\n%v", hex.Dump(per))

	rpv = createRanparameterValueValueEnum()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Enum) PER\n%v", hex.Dump(per))

	rpv = createRanparameterValueValueBool()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("RANparameterValue (Bool) PER\n%v", hex.Dump(per))

	rpv = createRanparameterValueValueOctS()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("RANparameterValue (OctetString) PER\n%v", hex.Dump(per))

	rpv = createRanparameterValueValuePrtS()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("RANparameterValue (PrintableString) PER\n%v", hex.Dump(per))

	rpv = createRanparameterValueValueBitS()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("RANparameterValue (BitString) PER\n%v", hex.Dump(per))
}

func Test_perDecodeRanParameterValue(t *testing.T) {

	rpv := createRanparameterValueValueInteger()

	per, err := perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 2, len(per))
	t.Logf("RANparameterValue (Integer) PER\n%v", hex.Dump(per))

	result, err := perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Integer) PER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueInt(), result.GetValueInt())

	rpv = createRanparameterValueValueEnum()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 3, len(per))
	t.Logf("RANparameterValue (Enum) PER\n%v", hex.Dump(per))

	result, err = perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Enum) PER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueEnum(), result.GetValueEnum())

	rpv = createRanparameterValueValueBool()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(per))
	t.Logf("RANparameterValue (Bool) PER\n%v", hex.Dump(per))

	result, err = perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (Bool) PER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueBool(), result.GetValueBool())

	rpv = createRanparameterValueValueOctS()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("RANparameterValue (OctetString) PER\n%v", hex.Dump(per))

	result, err = perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (OctetString) PER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueOctS(), result.GetValueOctS())

	rpv = createRanparameterValueValuePrtS()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("RANparameterValue (PrintableString) PER\n%v", hex.Dump(per))

	result, err = perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (PrintableString) PER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValuePrtS(), result.GetValuePrtS())

	rpv = createRanparameterValueValueBitS()

	per, err = perEncodeRanparameterValue(rpv)
	assert.NilError(t, err)
	assert.Equal(t, 5, len(per))
	t.Logf("RANparameterValue (BitString) PER\n%v", hex.Dump(per))

	result, err = perDecodeRanparameterValue(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("RANparameterValue (BitString) PER - decoded\n%v", result)
	assert.Equal(t, rpv.GetValueBitS().GetLen(), result.GetValueBitS().GetLen())
	assert.DeepEqual(t, rpv.GetValueBitS().GetValue(), result.GetValueBitS().GetValue())
}
