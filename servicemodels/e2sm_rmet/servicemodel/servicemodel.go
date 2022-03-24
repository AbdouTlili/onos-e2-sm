package servicemodel

import (
	"encoding/hex"
	"fmt"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/encoder"
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	prototypes "github.com/gogo/protobuf/types"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"google.golang.org/protobuf/proto"
)

type RmetServiceModel string

const smName = "e2sm_rmet"
const smVersion = "v2_go"
const moduleName = "e2smrmet.so.2.0"
const smOIDRmet = "1.3.6.1.4.1.53148.1.2.2.98"

func (sm RmetServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDRmet,
	}
	return smData
}

func (sm RmetServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRmetIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRmetIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRmetIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm RmetServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrmet.E2SmRmetIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRmetIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRmetIndicationHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRmetIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm RmetServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRmetIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRmetIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRmetIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm RmetServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrmet.E2SmRmetIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRmetIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRmetIndicationMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRmetIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm RmetServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRmetRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRmetRanFunctionDescription to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRmetRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm RmetServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrmet.E2SmRmetRanfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRmetRanfunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRmetRanFunctionDescription(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRmetRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm RmetServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRmetEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRmetEventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRmetEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm RmetServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrmet.E2SmRmetEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRmetEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRmetEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRmetEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm RmetServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmRmetActionDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmRmetActionDefinitio to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmRmetActionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm RmetServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smrmet.E2SmRmetActionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmRmetActionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmRmetActionDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmRmetActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm RmetServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on KPM")
}

func (sm RmetServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on KPM")
}

func (sm RmetServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on KPM")
}

func (sm RmetServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on KPM")
}

func (sm RmetServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on KPM")
}

func (sm RmetServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on KPM")
}

func (sm RmetServiceModel) OnSetup(request *types.OnSetupRequest) error {
	protoBytes, err := sm.RanFuncDescriptionASN1toProto(request.RANFunctionDescription)
	if err != nil {
		return err
	}
	ranFunctionDescription := &e2smrmet.E2SmRmetRanfunctionDescription{}
	err = proto.Unmarshal(protoBytes, ranFunctionDescription)
	if err != nil {
		return err
	}
	serviceModels := request.ServiceModels
	e2Cells := request.E2Cells
	serviceModel := serviceModels[smOIDRmet]
	serviceModel.Name = ranFunctionDescription.RanFunctionName.RanFunctionShortName
	ranFunction := &topoapi.KPMRanFunction{}

	for _, rmetNode := range ranFunctionDescription.RicRmetNodeList {
		for _, cell := range rmetNode.CellMeasurementObjectList {
			cellObject := &topoapi.E2Cell{
				CellObjectID: cell.GetCellObjectId().GetValue(),
				CellGlobalID: &topoapi.CellGlobalID{},
			}
			switch cellGlobalID := cell.CellGlobalId.GetCellGlobalId().(type) {
			case *e2smrmet.CellGlobalId_NrCgi:
				cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.NrCgi.NRcellIdentity.Value.Value, int(cellGlobalID.NrCgi.NRcellIdentity.Value.Len)))
				cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_NRCGI
			case *e2smrmet.CellGlobalId_EUtraCgi:
				cellObject.CellGlobalID.Value = fmt.Sprintf("%x", bitStringToUint64(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Value, int(cellGlobalID.EUtraCgi.EUtracellIdentity.Value.Len)))
				cellObject.CellGlobalID.Type = topoapi.CellGlobalIDType_ECGI
			}

			*e2Cells = append(*e2Cells, cellObject)
		}
	}

	for _, reportStyle := range ranFunctionDescription.GetRicReportStyleList() {
		rmetReportStyle := &topoapi.KPMReportStyle{
			Name: reportStyle.RicReportStyleName.Value,
			Type: reportStyle.RicReportStyleType.Value,
		}
		var measurements []*topoapi.KPMMeasurement
		for _, meanInfoItem := range reportStyle.GetMeasInfoActionList().GetValue() {
			measurements = append(measurements, &topoapi.KPMMeasurement{
				ID:   meanInfoItem.GetMeasId().String(),
				Name: meanInfoItem.GetMeasName().GetValue(),
			})
		}

		rmetReportStyle.Measurements = measurements
		ranFunction.ReportStyles = append(ranFunction.ReportStyles, rmetReportStyle)
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
