package pdubuilder

import (
	"encoding/hex"
	"testing"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/encoder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"gotest.tools/assert"
)

func TestE2SmMetRanfunctionDescription(t *testing.T) {
	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"

	var ricStyleType int32 = 11
	var ricStyleName = "onf"
	var ricFormatType int32 = 15
	retsi := CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2smmet.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2smmet.MeasurementInfoActionList{
		Value: make([]*e2smmet.MeasurementInfoActionItem, 0),
	}

	var measTypeName = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2smmet.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 24
	var indHdrFormat int32 = 47
	rrsi := CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2smmet.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmMetPdu, err := CreateE2SmMetRanfunctionDescription(rfSn, rfE2SMoid, rfd)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMetPdu != nil)
	newE2SmMetPdu.SetRicEventTriggerStyleList(retsl).SetRicReportStyleList(rrsl)

	per, err := encoder.PerEncodeE2SmMetRanFunctionDescription(newE2SmMetPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RANfunctionDescription PER is \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmMetRanFunctionDescription(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RANfunctionDescription PER - decoded\n%v", result)
}
