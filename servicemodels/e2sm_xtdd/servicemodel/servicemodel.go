package servicemodel

import (
	"encoding/hex"
	"fmt"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/encoder"
	e2smxtdd "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_xtdd/v1/e2sm-xtdd-go"
	types "github.com/onosproject/onos-api/go/onos/e2t/e2sm"
	"google.golang.org/protobuf/proto"
)

type XtddServiceModel string

const smName = "e2sm-xtdd"
const smVersion = "v1"
const moduleName = "e2smxtdd.so.2.0"
const smOIDMet = "1.3.6.1.4.1.53148.1.2.2.97"

func (sm XtddServiceModel) ServiceModelData() types.ServiceModelData {
	smData := types.ServiceModelData{
		Name:       smName,
		Version:    smVersion,
		ModuleName: moduleName,
		OID:        smOIDMet,
	}
	return smData
}

func (sm XtddServiceModel) IndicationHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmXtddIndicationHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetIndicationHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetIndicationHeader %s", err)
	}

	return protoBytes, nil
}

func (sm XtddServiceModel) IndicationHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smxtdd.E2SmXTddIndicationHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetIndicationHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmXtddIndicationHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetIndicationHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm XtddServiceModel) IndicationMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmXtddIndicationMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetIndicationMessage to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetIndicationMessage %s", err)
	}

	return protoBytes, nil
}

func (sm XtddServiceModel) IndicationMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smxtdd.E2SmXTddIndicationMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		fmt.Println("error Unmarshalling ")
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetIndicationMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmXtddIndicationMessage(protoObj)
	if err != nil {
		fmt.Println("error encoder.PerEncodeE2SmMetIndicationMessage(protoObj) ")
		return nil, errors.NewInvalid("error encoding E2SmMetIndicationMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm XtddServiceModel) RanFuncDescriptionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmXtddRanFunctionDescription(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetRanFunctionDescription to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetRanFunctionDescription %s", err)
	}

	return protoBytes, nil
}

func (sm XtddServiceModel) RanFuncDescriptionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smxtdd.E2SmXTddRAnfunctionDescription)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetRanfunctionDescription %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmXtddRanFunctionDescription(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetRanfunctionDescription to PER %s", err)
	}

	return perBytes, nil
}

func (sm XtddServiceModel) EventTriggerDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmXtddEventTriggerDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetEventTriggerDefinition to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetEventTriggerDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm XtddServiceModel) EventTriggerDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smxtdd.E2SmXTddEventTriggerDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetEventTriggerDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmXtddEventTriggerDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetEventTriggerDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm XtddServiceModel) ActionDefinitionASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmXtddActionDefinition(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmMetActionDefinitio to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmMetActionDefinition %s", err)
	}

	return protoBytes, nil
}

func (sm XtddServiceModel) ActionDefinitionProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smxtdd.E2SmXTddActionDefinition)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmMetActionDefinition %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmXtddActionDefinition(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmMetActionDefinition to PER %s", err)
	}

	return perBytes, nil
}

func (sm XtddServiceModel) ControlHeaderASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmXtddControlHeader(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmXtddControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmXtddControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm XtddServiceModel) ControlHeaderProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smxtdd.E2SmXTddControlHeader)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmXTddControlHeader %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmXtddControlHeader(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmXTddControlHeader to PER %s", err)
	}

	return perBytes, nil
}

func (sm XtddServiceModel) ControlMessageASN1toProto(asn1Bytes []byte) ([]byte, error) {
	perBytes, err := encoder.PerDecodeE2SmXtddControlMessage(asn1Bytes)
	if err != nil {
		return nil, errors.NewInvalid("error decoding E2SmXtddControlHeader to PER %s\n%v", err, hex.Dump(asn1Bytes))
	}

	protoBytes, err := proto.Marshal(perBytes)
	if err != nil {
		return nil, errors.NewInvalid("error marshalling asn1Bytes to E2SmXtddControlHeader %s", err)
	}

	return protoBytes, nil
}

func (sm XtddServiceModel) ControlMessageProtoToASN1(protoBytes []byte) ([]byte, error) {
	protoObj := new(e2smxtdd.E2SmXTddControlMessage)
	if err := proto.Unmarshal(protoBytes, protoObj); err != nil {
		return nil, errors.NewInvalid("error unmarshalling protoBytes to E2SmXTddControlMessage %s", err)
	}

	perBytes, err := encoder.PerEncodeE2SmXtddControlMessage(protoObj)
	if err != nil {
		return nil, errors.NewInvalid("error encoding E2SmXTddControlMessage to PER %s", err)
	}

	return perBytes, nil
}

func (sm XtddServiceModel) ControlOutcomeASN1toProto(asn1Bytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}

func (sm XtddServiceModel) ControlOutcomeProtoToASN1(protoBytes []byte) ([]byte, error) {
	return nil, errors.NewInvalid("not implemented on MET")
}
