// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/rcprectypes"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcPreRanfunctionDescriptionMsg(t *testing.T) {
	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 3
	var ricEventStyleType int32 = 13
	var ricEventStyleName = "ONFevent"
	var ricEventFormatType int32 = 42
	var ricReportStyleType int32 = 12
	var ricReportStyleName = "ONFreport"
	var ricIndicationHeaderFormatType int32 = 21
	var ricIndicationMessageFormatType int32 = 56
	retsl := make([]*e2sm_rc_pre_v2.RicEventTriggerStyleList, 0)
	retsi := CreateRicEventTriggerStyleItem(ricEventStyleType, ricEventStyleName, ricEventFormatType)
	retsl = append(retsl, retsi)

	rrsl := make([]*e2sm_rc_pre_v2.RicReportStyleList, 0)
	rrsi := CreateRicReportStyleItem(ricReportStyleType, ricReportStyleName, ricIndicationHeaderFormatType,
		ricIndicationMessageFormatType)
	rrsl = append(rrsl, rrsi)
	newE2SmRcPrePdu, err := CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription, retsl, rrsl)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	if newE2SmRcPrePdu != nil {
		newE2SmRcPrePdu.RanFunctionName.RanFunctionInstance = ranFunctionInstance
	}

	per, err := rcprectypes.PerEncodeE2SmRcPreRanfunctionDescription(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("Encoded RanFunctionDescription in PER is \n%v", hex.Dump(per))

	result1, err := rcprectypes.PerDecodeE2SmRcPreRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("Decoded RanFunctionDescription from PER is \n%v", result1)
}

func TestE2SmRcPreRanfunctionDescriptionMsgExcludeOptionalFields(t *testing.T) {
	var ranFunctionShortName = "ONF"
	var ranFunctionE2SmOid = "Oid"
	var ranFunctionDescription = "OpenNetworking"
	var ranFunctionInstance int32 = 3
	newE2SmRcPrePdu, err := CreateE2SmRcPreRanfunctionDescriptionMsg(ranFunctionShortName, ranFunctionE2SmOid,
		ranFunctionDescription, nil, nil)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRcPrePdu != nil)
	if newE2SmRcPrePdu != nil {
		newE2SmRcPrePdu.RanFunctionName.RanFunctionInstance = ranFunctionInstance
	}

	xer, err := rcprectypes.XerEncodeE2SmRcPreRanfunctionDescription(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("Encoded RanFunctionDescription in XER is \n%s", string(xer))

	result, err := rcprectypes.XerDecodeE2SmRcPreRanfunctionDescription(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("Decoded RanFunctionDescription from XER is \n%v", result)

	per, err := rcprectypes.PerEncodeE2SmRcPreRanfunctionDescription(newE2SmRcPrePdu)
	assert.NilError(t, err)
	t.Logf("Encoded RanFunctionDescription in PER is \n%v", hex.Dump(per))

	result1, err := rcprectypes.PerDecodeE2SmRcPreRanfunctionDescription(per)
	assert.NilError(t, err)
	assert.Assert(t, result1 != nil)
	t.Logf("Decoded RanFunctionDescription from PER is \n%v", result1)
}
