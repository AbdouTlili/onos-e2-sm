package pdubuilder

import (
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRmetRanfunctionDescription(rfSn string, rfE2SMoid string, rfd string) (*e2smrmet.E2SmRmetRanfunctionDescription, error) {

	e2SmRmetPdu := e2smrmet.E2SmRmetRanfunctionDescription{
		RanFunctionName: &e2smrmet.RanfunctionName{
			RanFunctionShortName:   rfSn,
			RanFunctionE2SmOid:     rfE2SMoid,
			RanFunctionDescription: rfd,
		},
	}

	if err := e2SmRmetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRmetRanfunctionDescription(): error validating E2SmRmetRanfunctionDescription %s", err.Error())
	}
	return &e2SmRmetPdu, nil
}

func CreateRicRmetnodeItem(globalRmetnodeID *e2smrmet.GlobalRmetnodeId) *e2smrmet.RicRmetnodeItem {

	res := e2smrmet.RicRmetnodeItem{
		RicRmetnodeType: globalRmetnodeID,
	}

	return &res
}

func CreateCellMeasurementObjectItem(cellObjID string, cellGlobalID *e2smrmet.CellGlobalId) *e2smrmet.CellMeasurementObjectItem {

	return &e2smrmet.CellMeasurementObjectItem{
		CellObjectId: &e2smrmet.CellObjectId{
			Value: cellObjID,
		},
		CellGlobalId: cellGlobalID,
	}
}

func CreateCellGlobalIDNRCGI(plmnID []byte, cellIDBits36 []byte) (*e2smrmet.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if len(cellIDBits36) != 5 {
		return nil, errors.NewInvalid("expecting 5 bits to carry NRcellIdentity, got %d", len(cellIDBits36))
	}
	if cellIDBits36[4]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", cellIDBits36[4])
	}
	bs := asn1.BitString{
		Value: cellIDBits36,
		Len:   36,
	}

	return &e2smrmet.CellGlobalId{
		CellGlobalId: &e2smrmet.CellGlobalId_NrCgi{
			NrCgi: &e2smrmet.Nrcgi{
				PLmnIdentity: &e2smrmet.PlmnIdentity{
					Value: plmnID,
				},
				NRcellIdentity: &e2smrmet.NrcellIdentity{
					Value: &bs,
				},
			},
		},
	}, nil
}

func CreateCellGlobalIDEUTRACGI(plmnID []byte, bs *asn1.BitString) (*e2smrmet.CellGlobalId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if len(bs.GetValue()) != 4 {
		return nil, errors.NewInvalid("expecting 4 bits to carry EutraCellIdentity, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[3]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[3])
	}

	return &e2smrmet.CellGlobalId{
		CellGlobalId: &e2smrmet.CellGlobalId_EUtraCgi{
			EUtraCgi: &e2smrmet.Eutracgi{
				PLmnIdentity: &e2smrmet.PlmnIdentity{
					Value: plmnID,
				},
				EUtracellIdentity: &e2smrmet.EutracellIdentity{
					Value: bs,
				},
			},
		},
	}, nil
}

func CreateRicEventTriggerStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32) *e2smrmet.RicEventTriggerStyleItem {

	return &e2smrmet.RicEventTriggerStyleItem{
		RicEventTriggerStyleType: &e2smrmet.RicStyleType{
			Value: ricStyleType,
		},
		RicEventTriggerStyleName: &e2smrmet.RicStyleName{
			Value: ricStyleName,
		},
		RicEventTriggerFormatType: &e2smrmet.RicFormatType{
			Value: ricFormatType,
		},
	}
}

func CreateRicReportStyleItem(ricStyleType int32, ricStyleName string, ricFormatType int32,
	measInfoActionList *e2smrmet.MeasurementInfoActionList, indHdrFormatType int32,
	indMsgFormatType int32) *e2smrmet.RicReportStyleItem {

	return &e2smrmet.RicReportStyleItem{
		RicReportStyleType: &e2smrmet.RicStyleType{
			Value: ricStyleType,
		},
		RicReportStyleName: &e2smrmet.RicStyleName{
			Value: ricStyleName,
		},
		RicActionFormatType: &e2smrmet.RicFormatType{
			Value: ricFormatType,
		},
		MeasInfoActionList: measInfoActionList,
		RicIndicationHeaderFormatType: &e2smrmet.RicFormatType{
			Value: indHdrFormatType,
		},
		RicIndicationMessageFormatType: &e2smrmet.RicFormatType{
			Value: indMsgFormatType,
		},
	}
}

func CreateMeasurementInfoActionItem(measTypeName string) *e2smrmet.MeasurementInfoActionItem {

	return &e2smrmet.MeasurementInfoActionItem{
		MeasName: &e2smrmet.MeasurementTypeName{
			Value: measTypeName,
		},
	}
}
