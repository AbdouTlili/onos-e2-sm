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

func createRicReportStyleItem() *e2sm_rc_pre_v2.RicReportStyleList {

	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56

	ricReportStyleItem := &e2sm_rc_pre_v2.RicReportStyleList{
		RicReportStyleType: &e2sm_rc_pre_v2.RicStyleType{
			Value: ricReportStyleType, //int32
		},
		RicReportStyleName: &e2sm_rc_pre_v2.RicStyleName{
			Value: ricReportStyleName, //string
		},
		RicIndicationHeaderFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: ricIndicationHeaderFormatType, //int32
		},
		RicIndicationMessageFormatType: &e2sm_rc_pre_v2.RicFormatType{
			Value: ricIndicationMessageFormatType, //int32
		},
	}

	return ricReportStyleItem
}

func Test_xerEncodeRicReportStyleList(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	xer, err := xerEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	//assert.Equal(t, 308, len(xer))
	t.Logf("RIC-ReportStyle-List XER\n%s", string(xer))
}

func Test_xerDecodeRicReportStyleList(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	xer, err := xerEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	//assert.Equal(t, 308, len(xer))
	t.Logf("RIC-ReportStyle-List XER\n%s", string(xer))

	result, err := xerDecodeRicReportStyleItem(xer)
	assert.NilError(t, err)
	t.Logf("RIC-ReportStyle-List XER - decoded\n%v", result)
	assert.Equal(t, ricReportStyleItem.RicReportStyleType.Value, result.RicReportStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicReportStyleName.Value, result.RicReportStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationHeaderFormatType.Value, result.RicIndicationHeaderFormatType.Value, "Encoded and decoded RicIndicationHeaderFormatType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationMessageFormatType.Value, result.RicIndicationMessageFormatType.Value, "Encoded and decoded RicIndicationMessageFormatType values are not the same")
}

func Test_perEncodeRicReportStyleList(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	per, err := perEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 17, len(per))
	t.Logf("RIC-ReportStyle-List PER\n%v", hex.Dump(per))
}

func Test_perDecodeRicReportStyleList(t *testing.T) {

	ricReportStyleItem := createRicReportStyleItem()

	per, err := perEncodeRicReportStyleItem(ricReportStyleItem)
	assert.NilError(t, err)
	assert.Equal(t, 17, len(per))
	t.Logf("RIC-ReportStyle-List PER\n%v", hex.Dump(per))

	result, err := perDecodeRicReportStyleItem(per)
	assert.NilError(t, err)
	t.Logf("RIC-ReportStyle-List PER - decoded\n%v", result)
	assert.Equal(t, ricReportStyleItem.RicReportStyleType.Value, result.RicReportStyleType.Value, "Encoded and decoded RicStyleType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicReportStyleName.Value, result.RicReportStyleName.Value, "Encoded and decoded RicStyleName values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationHeaderFormatType.Value, result.RicIndicationHeaderFormatType.Value, "Encoded and decoded RicIndicationHeaderFormatType values are not the same")
	assert.Equal(t, ricReportStyleItem.RicIndicationMessageFormatType.Value, result.RicIndicationMessageFormatType.Value, "Encoded and decoded RicIndicationMessageFormatType values are not the same")
}
