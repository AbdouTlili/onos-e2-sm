// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMetActionDefinitionFormat1(ricStyleType int32,
	actionDefinition *e2smmet.E2SmMetActionDefinitionFormat1) (*e2smmet.E2SmMetActionDefinition, error) {

	e2SmMetPdu := e2smmet.E2SmMetActionDefinition{
		RicStyleType: &e2smmet.RicStyleType{
			Value: ricStyleType,
		},
		ActionDefinitionFormats: &e2smmet.ActionDefinitionFormats{
			E2SmMetActionDefinition: &e2smmet.ActionDefinitionFormats_ActionDefinitionFormat1{
				ActionDefinitionFormat1: actionDefinition,
			},
		},
	}

	if err := e2SmMetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMetActionDefinitionFormat1(): error validating E2SmMetActionDefinition %s", err.Error())
	}
	return &e2SmMetPdu, nil
}

func CreateActionDefinitionFormat1(cellObjID int32, measInfoList *e2smmet.MeasurementInfoList,
	granularity int64, subID int64) (*e2smmet.E2SmMetActionDefinitionFormat1, error) {

	actionDefinitionFormat1 := e2smmet.E2SmMetActionDefinitionFormat1{
		CellObjId: &e2smmet.CellObjectId{
			Value: cellObjID,
		},
		MeasInfoList: measInfoList,
		GranulPeriod: &e2smmet.GranularityPeriod{
			Value: granularity,
		},
		SubscriptId: &e2smmet.SubscriptionId{
			Value: subID,
		},
	}

	if err := actionDefinitionFormat1.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateActionDefinitionFormat1(): error validating E2SmMetActionDefinitionFormat1 %s", err.Error())
	}
	return &actionDefinitionFormat1, nil
}

func CreateMeasurementInfoItem(measType *e2smmet.MeasurementType) (*e2smmet.MeasurementInfoItem, error) {

	item := e2smmet.MeasurementInfoItem{
		MeasType: measType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementInfoItem(): error validating MeasurementInfoItem %s", err.Error())
	}
	return &item, nil
}

func CreateMeasurementTypeMeasID(measTypeID int32) (*e2smmet.MeasurementType, error) {
	measType := e2smmet.MeasurementType{
		MeasurementType: &e2smmet.MeasurementType_MeasId{
			MeasId: &e2smmet.MeasurementTypeId{
				Value: measTypeID,
			},
		},
	}

	if err := measType.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementTypeMeasID(): error validating MeasurementType %s", err.Error())
	}
	return &measType, nil
}

func CreateMeasurementTypeMeasName(measName string) (*e2smmet.MeasurementType, error) {
	measType := e2smmet.MeasurementType{
		MeasurementType: &e2smmet.MeasurementType_MeasName{
			MeasName: &e2smmet.MeasurementTypeName{
				Value: measName,
			},
		},
	}

	if err := measType.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementTypeMeasName(): error validating MeasurementType %s", err.Error())
	}
	return &measType, nil
}
