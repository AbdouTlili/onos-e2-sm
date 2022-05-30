// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"testing"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"gotest.tools/assert"
)

func createMeasurementInfoList() (*e2smmet.MeasurementInfoList, error) {

	res := &e2smmet.MeasurementInfoList{
		Value: make([]*e2smmet.MeasurementInfoItem, 0),
	}
	mii, err := createMeasurementInfoItem("mcs")
	if err != nil {
		return nil, err
	}
	res.Value = append(res.Value, mii)

	mii2, err := createMeasurementInfoItem("phr")
	if err != nil {
		return nil, err
	}
	res.Value = append(res.Value, mii2)

	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func createMeasurementInfoItem(name string) (*e2smmet.MeasurementInfoItem, error) {

	res := &e2smmet.MeasurementInfoItem{
		Value: name,
	}
	if err := res.Validate(); err != nil {
		return nil, err
	}

	return res, nil
}

func TestE2SmMetIndicationHeader(t *testing.T) {

	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}

	var fileFormatVersion = "txt"
	var senderName = "ONF"

	globalMetNodeID, err := CreateGlobalMetnodeID(15)
	assert.NilError(t, err)

	mil, err := createMeasurementInfoList()
	assert.NilError(t, err)

	newE2SmMetPdu, err := CreateE2SmMetIndicationHeader(timeStamp, mil)
	assert.NilError(t, err)
	newE2SmMetPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetGlobalMETnodeID(globalMetNodeID)
	assert.Assert(t, newE2SmMetPdu != nil)
}
