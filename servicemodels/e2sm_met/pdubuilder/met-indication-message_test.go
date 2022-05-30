// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"testing"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"gotest.tools/assert"
)

func TestE2SmMetIndicationMessageFormat1(t *testing.T) {

	var integer64 int64 = 12345
	var integer32 int32 = 150

	// var rl float64 = 6789
	var cellObjID int32 = 15
	var granularity int64 = 21
	var subscriptionID int64 = 12345

	measRecItemList := make([]*e2smmet.MeasurementRecordItem, 0)

	/* the first ue measRecord */

	measRecItemList = append(measRecItemList, CreateMeasurementRecordItemInteger(integer64))
	measRecItemList = append(measRecItemList, CreateMeasurementRecordItemNoValue())
	// measRecItemList = append(measRecItemList, CreateMeasurementRecordItemReal(rl))

	measRec1, err := CreateMeasurementRecord(&e2smmet.Ueid{
		Value: integer32,
	}, &e2smmet.Uetag{
		Value: "eurecom",
	}, measRecItemList)
	assert.NilError(t, err)

	/* end of the first ue measRecord */

	/* the second ue measRecord */
	measRecItemList2 := make([]*e2smmet.MeasurementRecordItem, 0)

	measRecItemList2 = append(measRecItemList2, CreateMeasurementRecordItemInteger(integer64))
	measRecItemList2 = append(measRecItemList2, CreateMeasurementRecordItemNoValue())
	// measRecItemList2 = append(measRecItemList2, CreateMeasurementRecordItemReal(rl))

	measRec2, err := CreateMeasurementRecord(&e2smmet.Ueid{
		Value: integer32,
	}, &e2smmet.Uetag{
		Value: "eurecom2",
	}, measRecItemList2)
	assert.NilError(t, err)
	/* end of the second  ue measRecord */

	measRecList := make([]*e2smmet.MeasurementRecord, 0)

	measRecList = append(measRecList, measRec1, measRec2)
	measData, err := CreateMeasurementData(measRecList)
	assert.NilError(t, err)

	///measInfoList
	// this is a go slice and not a Type defined in pb.go it is used later to create
	// the real MeasurmentInfoList object
	measInfolist := make([]*e2smmet.MeasurementInfoItem, 0)

	// miii1, err := CreateMeasurementTypeMeasID(30)
	// assert.NilError(t, err)
	// miii2, err := CreateMeasurementTypeMeasID(31)
	// assert.NilError(t, err)
	// miii3, err := CreateMeasurementTypeMeasID(32)
	// assert.NilError(t, err)

	mii1, err := CreateMeasurementInfoItem("miii1")
	assert.NilError(t, err)
	mii2, err := CreateMeasurementInfoItem("miii2")
	assert.NilError(t, err)
	mii3, err := CreateMeasurementInfoItem("miii3")
	assert.NilError(t, err)

	measInfolist = append(measInfolist, mii1, mii2, mii3)

	measInfoList, err := CreateMeasurementInfoList(measInfolist)
	assert.NilError(t, err)

	/////

	// measName, err := CreateMeasurementTypeMeasName(measurementName)
	// assert.NilError(t, err)
	// measInfoItem, err := CreateMeasurementInfoItem(measName)
	// assert.NilError(t, err)
	// measInfoItem.SetLabelInfoList(&labelInfoList)

	// measInfoList := e2smmet.MeasurementInfoList{
	// 	Value: make([]*e2smmet.MeasurementInfoItem, 0),
	// }
	// measInfoList.Value = append(measInfoList.Value, measInfoItem)

	// measRecord := e2smmet.MeasurementRecord{
	// 	Value: make([]*e2smmet.MeasurementRecordItem, 0),
	// }
	// measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemInteger(integer))
	// measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemNoValue())
	// measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemReal(rl))

	// measDataItem, err := CreateMeasurementDataItem(&measRecord)
	// assert.NilError(t, err)
	// measDataItem.SetIncompleteFlag()

	// measData := e2smmet.MeasurementData{
	// 	Value: make([]*e2smmet.MeasurementDataItem, 0),
	// }
	// measData.Value = append(measData.Value, measDataItem)

	newE2SmMetPdu, err := CreateE2SmMetIndicationMessageFormat1(subscriptionID, measData)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMetPdu != nil)
	newE2SmMetPdu.SetGranularityPeriod(granularity).SetCellObjectID(cellObjID).SetMeasInfoList(measInfoList)
	t.Logf("Composed IndicationMessage-Format1 is \n %v \n", newE2SmMetPdu)
}
