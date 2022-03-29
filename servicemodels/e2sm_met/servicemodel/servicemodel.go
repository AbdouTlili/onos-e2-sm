// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"
	"fmt"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/encoder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"google.golang.org/protobuf/proto"
)

type MetServiceModel string

const smName = "e2sm_met"
const smVersion = "v1_go"
const moduleName = "e2smmet.so.2.0"
const smOIDMet = "1.3.6.1.4.1.53148.1.2.2.2"

func (sm MetServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDMet,
	}
	return smData
}

func (sm MetServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMetIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm MetServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmet.E2SmMetIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMetIndicationHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm MetServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMetIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm MetServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmet.E2SmMetIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMetIndicationMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm MetServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMetRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetRanFunctionDescription to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm MetServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmet.E2SmMetRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetRanfunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMetRanFunctionDescription(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm MetServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMetEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetEventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm MetServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmet.E2SmMetEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMetEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm MetServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmMetActionDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetActionDefinitio to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetActionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm MetServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smmet.E2SmMetActionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetActionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmMetActionDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm MetServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}

func (sm MetServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}

func (sm MetServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}

func (sm MetServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}

func (sm MetServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}

func (sm MetServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}

func (sm MetServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2smmet.E2SmMetRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	e2Cells := request.E2Cells
	serviceModel := serviceModels[smOIDMet]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	ranFunction := &topoapi.METRanFunction{}

	for _, metNode := range ranFunctionDescription.RicMetNodeList {
		for _, cell := range metNode.CellMeasurementObjectList {
			cellObject := &topoapi.E2Cell{
				CellObjectID: cell.GetCellObjectId().GetValue(),
				CellGlobalID: &topoapi.CellGlobalID{},
			}
			switch cellGlobalID := cell.CellGlobalId.GetCellGlobalId().(type) {
			case *e2smmet.CellGlobalId_NrCgi:
				cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.NrCgi.NRcellIdentity.Value.Value, int(cellGlobalID.NrCgi.NRcellIdentity.Value.Len)))
				cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_NRCGI
			case *e2smmet.CellGlobalId_EUtraCgi:
				cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Value, int(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Len)))
				cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_ECGI
			}

			*e2Cells = append(*e2Cells, cellObject)
		}
	}

	for _, reportStyle := range ranFunctionDescription.GetRicReportStyleList() {
		metReportStyle := &topoapi.METReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		var measurements []*topoapi.METMeasurement
		for _, meanInfoItem := range reportStyle.GetMeasInfoActionList().GetValue() {
			measurements = append(measurements, &topoapi.METMeasurement{
				ID:   meanInfoItem.GetMeasId().String(),
				Name: meanInfoItem.GetMeasName().GetValue(),
			})
		}

		metReportStyle.Measurements = measurements
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, kpmReportStyle)
	}
	ranFunctionAny, err := prototypes.MarshalAny(ranFunction)
	if err != nil {
		return err
	}

	serviceModel.RanFunctions = append(serviceModel.RanFunctions, ranFunctionAny)
	return nil
}

func bitStringToUint64(bitString []byte, bitCount int) uint64 {
	var result uint64
	for i, b := range bitString {
		result += uint64(b) << ((len(bitString) - i - 1) * 8)
	}
	if bitCount%8 != 0 {
		return result >> (8 - bitCount%8)
	}
	return result
}
