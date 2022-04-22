// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"testing"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"gotest.tools/assert"
)

func TestE2SmKpmActionDefinitionFormat1(t *testing.T) {
	var ricStyleType int32 = 12
	var cellObjID int32 = 32
	var granularity int64 = 21
	var subscriptionID int64 = 12345
	var measurementName = "trial"

	measName, err := CreateMeasurementTypeMeasName(measurementName)
	assert.NilError(t, err)
	measInfoItem, err := CreateMeasurementInfoItem(measName)
	assert.NilError(t, err)

	measInfoList := e2smmet.MeasurementInfoList{
		Value: make([]*e2smmet.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	actionDefinition, err := CreateActionDefinitionFormat1(cellObjID, &measInfoList, granularity, subscriptionID)
	assert.NilError(t, err)

	newE2SmKpmPdu, err := CreateE2SmMetActionDefinitionFormat1(ricStyleType, actionDefinition)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmKpmPdu != nil)
}
