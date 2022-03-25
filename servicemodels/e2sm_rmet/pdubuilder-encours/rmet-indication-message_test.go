// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"gotest.tools/assert"
	"testing"
)

func TestMeasurementLabelFull(t *testing.T) {
	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qfi int32 = 62
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2smrmet.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	var plo = e2smrmet.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	startEndIndication := e2smrmet.StartEndInd_START_END_IND_START

	labelInfoItem, err := CreateLabelInfoItem(plmnID, sst, sd, &fiveQI, &qfi, &qci, &qciMax, &qciMin, &arpMax, &arpMin,
		&bitrateRange, &layerMuMimo, &sum, &distX, &distY, &distZ, &plo, &startEndIndication)
	assert.NilError(t, err)
	assert.Assert(t, labelInfoItem != nil)
	t.Logf("Composed (full) MeasurementLabel is \n%v", labelInfoItem.GetMeasLabel())
}

func TestMeasurementLabelPartially(t *testing.T) {
	//plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	//var qfi int32 = 62
	var qci int32 = 15
	//var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	//var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	//var sum = e2smrmet.SUM_SUM_TRUE
	//var distX int32 = 123
	//var distY int32 = 456
	//var distZ int32 = 789
	//var plo = e2smrmet.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	//startEndIndication := e2smrmet.StartEndInd_START_END_IND_START

	labelInfoItem, err := CreateLabelInfoItem(nil, sst, sd, &fiveQI, nil, &qci, &qciMax, nil, &arpMax, nil,
		&bitrateRange, &layerMuMimo, nil, nil, nil, nil, nil, nil)
	assert.NilError(t, err)
	assert.Assert(t, labelInfoItem != nil)
	t.Logf("Composed (partially) MeasurementLabel is \n%v", labelInfoItem.GetMeasLabel())
}

func TestE2SmRmetIndicationMessageFormat1(t *testing.T) {
	var integer int64 = 12345
	var rl float64 = 6789
	var cellObjID = "onf"
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qfi int32 = 62
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 10
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var sum = e2smrmet.SUM_SUM_TRUE
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(measName)
	assert.NilError(t, err)

	measInfoList := e2smrmet.MeasurementInfoList{
		Value: make([]*e2smrmet.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2smrmet.MeasurementRecord{
		Value: make([]*e2smrmet.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, CreateMeasurementRecordItemReal(rl))

	measDataItem, err := CreateMeasurementDataItem(&measRecord)
	assert.NilError(t, err)
	measDataItem.SetIncompleteFlag()

	measData := e2smrmet.MeasurementData{
		Value: make([]*e2smrmet.MeasurementDataItem, 0),
	}
	measData.Value = append(measData.Value, measDataItem)

	newE2SmRmetPdu, err := CreateE2SmRmetIndicationMessageFormat1(subscriptionID, &measData)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmRmetPdu != nil)
	newE2SmRmetPdu.SetGranularityPeriod(granularity).SetCellObjectID(cellObjID).SetMeasInfoList(&measInfoList)
	t.Logf("Composed IndicationMessage-Format1 is \n %v \n", newE2SmRmetPdu)
}
