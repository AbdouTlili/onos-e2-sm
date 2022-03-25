package pdubuilder

import (
	"encoding/hex"
	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/encoder"
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRmetRanfunctionDescription(t *testing.T) {
	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"
	var rfi int32 = 21

	plmnID := []byte{0x21, 0x22, 0x23}
	bs := asn1.BitString{
		Value: []byte{0xd4, 0xbc, 0x08},
		Len:   22,
	}
	cellIDbits := []byte{0x12, 0xF0, 0xDE, 0xBC, 0xF0}
	cellGlobalID, err := CreateCellGlobalIDNRCGI(plmnID, cellIDbits) // 36 bit
	assert.NilError(t, err)
	var cellObjID = "ONF"
	cellMeasObjItem := CreateCellMeasurementObjectItem(cellObjID, cellGlobalID)

	var gnbCuUpID int64 = 12345
	var gnbDuID int64 = 6789
	globalRmetnodeID, err := CreateGlobalRmetnodeIDgNBID(&bs, plmnID)
	assert.NilError(t, err)
	globalRmetnodeID.GetGNb().GNbCuUpId = &e2smrmet.GnbCuUpId{
		Value: gnbCuUpID,
	}
	globalRmetnodeID.GetGNb().GNbDuId = &e2smrmet.GnbDuId{
		Value: gnbDuID,
	}

	cmol := make([]*e2smrmet.CellMeasurementObjectItem, 0)
	cmol = append(cmol, cellMeasObjItem)

	rmetNodeItem := CreateRicRmetnodeItem(globalRmetnodeID).SetCellMeasurementObjectList(cmol)

	rknl := make([]*e2smrmet.RicRmetnodeItem, 0)
	rknl = append(rknl, rmetNodeItem)

	var ricStyleType int32 = 11
	var ricStyleName = "onf"
	var ricFormatType int32 = 15
	retsi := CreateRicEventTriggerStyleItem(ricStyleType, ricStyleName, ricFormatType)

	retsl := make([]*e2smrmet.RicEventTriggerStyleItem, 0)
	retsl = append(retsl, retsi)

	measInfoActionList := e2smrmet.MeasurementInfoActionList{
		Value: make([]*e2smrmet.MeasurementInfoActionItem, 0),
	}

	var measTypeName = "OpenNetworking"
	var measTypeID int32 = 24
	measInfoActionItem := CreateMeasurementInfoActionItem(measTypeName)
	measInfoActionItem.MeasId = &e2smrmet.MeasurementTypeId{
		Value: measTypeID,
	}
	measInfoActionList.Value = append(measInfoActionList.Value, measInfoActionItem)

	var indMsgFormat int32 = 24
	var indHdrFormat int32 = 47
	rrsi := CreateRicReportStyleItem(ricStyleType, ricStyleName, ricFormatType, &measInfoActionList, indHdrFormat, indMsgFormat)

	rrsl := make([]*e2smrmet.RicReportStyleItem, 0)
	rrsl = append(rrsl, rrsi)

	newE2SmRmetPdu, err := CreateE2SmRmetRanfunctionDescription(rfSn, rfE2SMoid, rfd)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRmetPdu != nil)
	newE2SmRmetPdu.SetRanFunctionInstance(rfi).SetRicEventTriggerStyleList(retsl).SetRicRmetNodeList(rknl).SetRicReportStyleList(rrsl)

	per, err := encoder.PerEncodeE2SmRmetRanFunctionDescription(newE2SmRmetPdu)
	assert.NilError(t, err)
	t.Logf("E2SM-RANfunctionDescription PER is \n%v", hex.Dump(per))

	result, err := encoder.PerDecodeE2SmRmetRanFunctionDescription(per)
	assert.NilError(t, err)
	t.Logf("E2SM-RANfunctionDescription PER - decoded\n%v", result)
}
