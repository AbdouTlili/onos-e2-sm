// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMetIndicationMessageFormat1(subscriptionID int64, measData *e2smmet.MeasurementData) (*e2smmet.E2SmMetIndicationMessage, error) {

	e2SmMetPdu := e2smmet.E2SmMetIndicationMessage{
		IndicationMessageFormats: &e2smmet.IndicationMessageFormats{
			E2SmMetIndicationMessage: &e2smmet.IndicationMessageFormats_IndicationMessageFormat1{
				IndicationMessageFormat1: &e2smmet.E2SmMetIndicationMessageFormat1{
					SubscriptId: &e2smmet.SubscriptionId{
						Value: subscriptionID,
					},
					MeasData: measData,
				},
			},
		},
	}

	if err := e2SmMetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMetIndicationMessageFormat1(): error validating E2SmMetPDU %s", err.Error())
	}
	return &e2SmMetPdu, nil
}

func CreateMeasurementRecordItemInteger(integer int64) *e2smmet.MeasurementRecordItem {

	return &e2smmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smmet.MeasurementRecordItem_Integer{
			Integer: integer,
		},
	}
}

func CreateMeasurementRecordItemReal(real float64) *e2smmet.MeasurementRecordItem {

	return &e2smmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smmet.MeasurementRecordItem_Real{
			Real: real,
		},
	}
}

func CreateMeasurementRecordItemNoValue() *e2smmet.MeasurementRecordItem {

	return &e2smmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smmet.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
}

func CreateMeasurementData(mrList []*e2smmet.MeasurementRecord) (*e2smmet.MeasurementData, error) {

	mdi := e2smmet.MeasurementData{
		Value: mrList,
	}

	if err := mdi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementData(): error validating MeasurementData %s", err.Error())
	}
	return &mdi, nil
}

func CreateMeasurementRecord(ueid *e2smmet.Ueid, uetag *e2smmet.Uetag, mrItems []*e2smmet.MeasurementRecordItem) (*e2smmet.MeasurementRecord, error) {

	mdi := e2smmet.MeasurementRecord{
		MeasRecordItem: mrItems,
		UeId:           ueid,
		UeTag:          uetag,
	}

	if err := mdi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementRecord(): error validating MeasurementRecord %s", err.Error())
	}
	return &mdi, nil
}

func CreateMeasurementInfoList(measInfoItems []*e2smmet.MeasurementInfoItem) (*e2smmet.MeasurementInfoList, error) {
	mil := e2smmet.MeasurementInfoList{
		Value: measInfoItems,
	}

	if err := mil.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementInfoList(): error validating MeasurementInfoList %s", err.Error())
	}
	return &mil, nil
}

func CreateMeasurementInfoItem(measurementType string) (*e2smmet.MeasurementInfoItem, error) {
	miItem := e2smmet.MeasurementInfoItem{
		MeasType: measurementType,
	}

	if err := miItem.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementInfoItem(): error validating MeasurementInfoItem %s", err.Error())
	}
	return &miItem, nil
}

// some usefull code in the Test file
