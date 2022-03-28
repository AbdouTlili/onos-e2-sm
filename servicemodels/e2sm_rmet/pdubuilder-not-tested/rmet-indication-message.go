// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRmetIndicationMessageFormat1(subscriptionID int64, measData *e2smrmet.MeasurementData) (*e2smrmet.E2SmRmetIndicationMessage, error) {

	e2SmRmetPdu := e2smrmet.E2SmRmetIndicationMessage{
		IndicationMessageFormats: &e2smrmet.IndicationMessageFormats{
			E2SmRmetIndicationMessage: &e2smrmet.IndicationMessageFormats_IndicationMessageFormat1{
				IndicationMessageFormat1: &e2smrmet.E2SmRmetIndicationMessageFormat1{
					SubscriptId: &e2smrmet.SubscriptionId{
						Value: subscriptionID,
					},
					MeasData: measData,
				},
			},
		},
	}

	if err := e2SmRmetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRmetIndicationMessageFormat1(): error validating E2SmRmetPDU %s", err.Error())
	}
	return &e2SmRmetPdu, nil
}

func CreateMeasurementRecordItemInteger(integer int64) *e2smrmet.MeasurementRecordItem {

	return &e2smrmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smrmet.MeasurementRecordItem_Integer{
			Integer: integer,
		},
	}
}

func CreateMeasurementRecordItemReal(real float64) *e2smrmet.MeasurementRecordItem {

	return &e2smrmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smrmet.MeasurementRecordItem_Real{
			Real: real,
		},
	}
}

func CreateMeasurementRecordItemNoValue() *e2smrmet.MeasurementRecordItem {

	return &e2smrmet.MeasurementRecordItem{
		MeasurementRecordItem: &e2smrmet.MeasurementRecordItem_NoValue{
			NoValue: 0,
		},
	}
}

func CreateMeasurementDataItem(mr *e2smrmet.MeasurementRecord) (*e2smrmet.MeasurementDataItem, error) {

	mdi := e2smrmet.MeasurementDataItem{
		MeasRecord: mr,
	}

	if err := mdi.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementDataItem(): error validating MatchingUeidItem %s", err.Error())
	}
	return &mdi, nil
}
