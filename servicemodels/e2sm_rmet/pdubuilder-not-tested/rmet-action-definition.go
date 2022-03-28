package pdubuilder

import (
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRmetActionDefinitionFormat1(ricStyleType int32,
	actionDefinition *e2smrmet.E2SmRmetActionDefinitionFormat1) (*e2smrmet.E2SmRmetActionDefinition, error) {

	e2SmRmetPdu := e2smrmet.E2SmRmetActionDefinition{
		RicStyleType: &e2smrmet.RicStyleType{
			Value: ricStyleType,
		},
		ActionDefinitionFormats: &e2smrmet.ActionDefinitionFormats{
			E2SmRmetActionDefinition: &e2smrmet.ActionDefinitionFormats_ActionDefinitionFormat1{
				ActionDefinitionFormat1: actionDefinition,
			},
		},
	}

	if err := e2SmRmetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRmetActionDefinitionFormat1(): error validating E2SmRmetActionDefinition %s", err.Error())
	}
	return &e2SmRmetPdu, nil
}

func CreateActionDefinitionFormat1(cellObjID string, measInfoList *e2smrmet.MeasurementInfoList,
	granularity int64, subID int64) (*e2smrmet.E2SmRmetActionDefinitionFormat1, error) {

	actionDefinitionFormat1 := e2smrmet.E2SmRmetActionDefinitionFormat1{
		CellObjId: &e2smrmet.CellObjectId{
			Value: cellObjID,
		},
		MeasInfoList: measInfoList,
		GranulPeriod: &e2smrmet.GranularityPeriod{
			Value: granularity,
		},
		SubscriptId: &e2smrmet.SubscriptionId{
			Value: subID,
		},
	}

	if err := actionDefinitionFormat1.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateActionDefinitionFormat1(): error validating E2SmRmetActionDefinitionFormat1 %s", err.Error())
	}
	return &actionDefinitionFormat1, nil
}

func CreateMeasurementInfoItem(measType *e2smrmet.MeasurementType) (*e2smrmet.MeasurementInfoItem, error) {

	item := e2smrmet.MeasurementInfoItem{
		MeasType: measType,
	}

	if err := item.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementInfoItem(): error validating MeasurementInfoItem %s", err.Error())
	}
	return &item, nil
}

func CreateMeasurementTypeMeasID(measTypeID int32) (*e2smrmet.MeasurementType, error) {
	measType := e2smrmet.MeasurementType{
		MeasurementType: &e2smrmet.MeasurementType_MeasId{
			MeasId: &e2smrmet.MeasurementTypeId{
				Value: measTypeID,
			},
		},
	}

	if err := measType.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementTypeMeasID(): error validating MeasurementType %s", err.Error())
	}
	return &measType, nil
}

func CreateMeasurementTypeMeasName(measName string) (*e2smrmet.MeasurementType, error) {
	measType := e2smrmet.MeasurementType{
		MeasurementType: &e2smrmet.MeasurementType_MeasName{
			MeasName: &e2smrmet.MeasurementTypeName{
				Value: measName,
			},
		},
	}

	if err := measType.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateMeasurementTypeMeasName(): error validating MeasurementType %s", err.Error())
	}
	return &measType, nil
}

func CreateSnssai(sst []byte) (*e2smrmet.Snssai, error) {

	if len(sst) != 1 {
		return nil, errors.NewInvalid("SST should be of length 1 byte, got %d", len(sst))
	}
	return &e2smrmet.Snssai{
		SSt: sst,
	}, nil
}
