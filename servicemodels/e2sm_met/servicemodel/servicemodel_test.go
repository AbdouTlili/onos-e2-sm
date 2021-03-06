// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package servicemodel

import (
	"encoding/hex"

	"fmt"
	"testing"

	"github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/pdubuilder"
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert"
)

var metTestSm MetServiceModel

func TestServicemodel_IndicationHeaderProtoToASN1(t *testing.T) {

	timeStamp := []byte{0x21, 0x22, 0x23, 0x24}
	// var fileFormatVersion = "txt"
	// var senderName = "ONF"

	globalMetNodeID, err := pdubuilder.CreateGlobalMetnodeID(5)
	assert.NilError(t, err)

	mil, err := createMeasurementInfoList()
	assert.NilError(t, err)

	newE2SmMetPdu, err := pdubuilder.CreateE2SmMetIndicationHeader(timeStamp, mil)
	assert.NilError(t, err)
	// newE2SmMetPdu.SetFileFormatVersion(fileFormatVersion).SetSenderName(senderName).SetGlobalMETnodeID(globalMetNodeID)

	newE2SmMetPdu.SetGlobalMETnodeID(globalMetNodeID)
	assert.NilError(t, err)

	err = newE2SmMetPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	protoBytes, err := proto.Marshal(newE2SmMetPdu)
	assert.NilError(t, err, "unexpected error marshalling E2SmMetIndicationHeader to bytes")
	// fmt.Printf("%#v", protoBytes)

	// assert.Equal(t, 69, len(protoBytes))

	asn1Bytes, err := metTestSm.IndicationHeaderProtoToASN1(protoBytes)
	fmt.Printf("%#v", asn1Bytes)
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-MET-IndicationHeader (gNB) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

func TestServicemodel_IndicationHeaderASN1toProto(t *testing.T) {
	// This value is taken from Shad and passed as a byte array directly to the function
	// It's the encoding of what's in the file ../test/E2SM-MET-Indication-Header-gNB.xml
	indicationHeaderAsn1Bytes := []byte{0x04, 0x62, 0x72, 0x4b, 0x13, 0x00, 0x09}

	protoBytes, err := metTestSm.IndicationHeaderASN1toProto(indicationHeaderAsn1Bytes)
	assert.NilError(t, err, "unexpected error converting asn1Bytes to protoBytes")
	assert.Assert(t, protoBytes != nil)
	testIH := &e2smmet.E2SmMetIndicationHeader{}
	err = proto.Unmarshal(protoBytes, testIH)
	assert.NilError(t, err)
	t.Logf("Decoded message is \n%v", testIH)
	// assert.DeepEqual(t, []byte{0x21, 0x22, 0x23, 0x24}, testIH.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetColletStartTime().GetValue())
	assert.Equal(t, int64(10), testIH.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().GetMetNodeId().GetValue())
}

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

// func TestServicemodel_IndicationMessageProtoToASN1(t *testing.T) {
// 	var integer int64 = 10
// 	// var rl float64 = 6789
// 	// var cellObjID int32 = 11
// 	// var granularity int64 = 12
// 	var subscriptionID int64 = 10

// 	measRecItemList := make([]*e2smmet.MeasurementRecordItem, 0)

// 	//creating measRecordItems

// 	measRecItemList = append(measRecItemList, pdubuilder.CreateMeasurementRecordItemInteger(integer))

// 	measRecItemList = append(measRecItemList, pdubuilder.CreateMeasurementRecordItemInteger(integer))

// 	measRecItemList = append(measRecItemList, pdubuilder.CreateMeasurementRecordItemInteger(integer))

// 	measRecItemList = append(measRecItemList, pdubuilder.CreateMeasurementRecordItemInteger(integer))

// 	measRecItemList = append(measRecItemList, pdubuilder.CreateMeasurementRecordItemInteger(integer))

// 	/* the first ue measRecord */

// 	measRec1, err := pdubuilder.CreateMeasurementRecord(&e2smmet.Ueid{
// 		Value: 10,
// 	}, &e2smmet.Uetag{
// 		Value: "ABC",
// 	}, measRecItemList)
// 	assert.NilError(t, err)

// 	/* end of the first ue measRecord */

// 	/* the 2nd ue measRecord */

// 	measRec2, err := pdubuilder.CreateMeasurementRecord(&e2smmet.Ueid{
// 		Value: 10,
// 	}, &e2smmet.Uetag{
// 		Value: "ABC",
// 	}, measRecItemList)
// 	assert.NilError(t, err)

// 	/* end of the 2nd ue measRecord */

// 	/* the 3rd ue measRecord */

// 	measRec3, err := pdubuilder.CreateMeasurementRecord(&e2smmet.Ueid{
// 		Value: 10,
// 	}, &e2smmet.Uetag{
// 		Value: "ABC",
// 	}, measRecItemList)
// 	assert.NilError(t, err)

// 	/* end of the 3rd ue measRecord */

// 	/* the 4th ue measRecord */

// 	measRec4, err := pdubuilder.CreateMeasurementRecord(&e2smmet.Ueid{
// 		Value: 10,
// 	}, &e2smmet.Uetag{
// 		Value: "ABC",
// 	}, measRecItemList)
// 	assert.NilError(t, err)

// 	/* end of the 4th ue measRecord */

// 	/* end of the second  ue measRecord */
// 	measRecList := make([]*e2smmet.MeasurementRecord, 0)
// 	measRecList = append(measRecList, measRec1, measRec2, measRec3, measRec4)
// 	measData, err := pdubuilder.CreateMeasurementData(measRecList)
// 	assert.NilError(t, err)

// 	///measInfoList
// 	// this is a go slice and not a Type defined in pb.go it is used later to create
// 	// the real MeasurmentInfoList object
// 	// measInfolist := make([]*e2smmet.MeasurementInfoItem, 0)

// 	// miii1, err := pdubuilder.CreateMeasurementTypeMeasID(30)
// 	// assert.NilError(t, err)
// 	// miii2, err := pdubuilder.CreateMeasurementTypeMeasID(31)
// 	// assert.NilError(t, err)
// 	// miii3, err := pdubuilder.CreateMeasurementTypeMeasID(32)
// 	// assert.NilError(t, err)

// 	// mii1, err := pdubuilder.CreateMeasurementInfoItem(miii1)
// 	// assert.NilError(t, err)
// 	// mii2, err := pdubuilder.CreateMeasurementInfoItem(miii2)
// 	// assert.NilError(t, err)
// 	// mii3, err := pdubuilder.CreateMeasurementInfoItem(miii3)
// 	// assert.NilError(t, err)

// 	// measInfolist = append(measInfolist, mii1, mii2, mii3)

// 	// measInfoList, err := pdubuilder.CreateMeasurementInfoList(measInfolist)
// 	// assert.NilError(t, err)

// 	/////

// 	newE2SmMetPdu, err := pdubuilder.CreateE2SmMetIndicationMessageFormat1(subscriptionID, measData)
// 	assert.NilError(t, err)
// 	assert.Assert(t, newE2SmMetPdu != nil)
// 	// newE2SmMetPdu.SetGranularityPeriod(granularity).SetCellObjectID(cellObjID) //.SetMeasInfoList(measInfoList)

// 	err = newE2SmMetPdu.Validate()
// 	assert.NilError(t, err, "error validating E2SmPDU")

// 	assert.NilError(t, err)
// 	protoBytes, err := proto.Marshal(newE2SmMetPdu)
// 	assert.NilError(t, err, "unexpected error marshalling E2SmMetIndicationMessage (Format1) to bytes")

// 	asn1Bytes, err := metTestSm.IndicationMessageProtoToASN1(protoBytes)
// 	fmt.Printf("%#v - %v", asn1Bytes, len(asn1Bytes))
// 	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
// 	assert.Assert(t, asn1Bytes != nil)
// 	t.Logf("E2SM-MET-IndicationMessage (Format1)) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
// }

func TestServicemodel_IndicationMessageASN1toProto(t *testing.T) {

	indicationMessageAsn1 := []byte{0x00, 0x00, 0x09, 0x05, 0x00, 0x09, 0x01, 0x00, 0x41, 0x42, 0x43, 0x05,
		0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x09, 0x01,
		0x00, 0x41, 0x42, 0x43, 0x05, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00,
		0x01, 0x0a, 0x00, 0x09, 0x01, 0x00, 0x41, 0x42, 0x43, 0x05, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01,
		0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x09, 0x01, 0x00, 0x41, 0x42, 0x43, 0x05, 0x00, 0x01, 0x0a,
		0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x09, 0x01, 0x00, 0x41, 0x42,
		0x43, 0x05, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a, 0x00, 0x01, 0x0a,
	}

	protoBytes, err := metTestSm.IndicationMessageASN1toProto(indicationMessageAsn1)
	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
	assert.Assert(t, protoBytes != nil)
	// assert.Equal(t, 124, len(protoBytes))
	testIM := &e2smmet.E2SmMetIndicationMessage{}
	err = proto.Unmarshal(protoBytes, testIM)
	assert.NilError(t, err)
	//t.Logf("Decoded message is \n%v", testIM)
	// assert.Equal(t, int64(13), testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetSubscriptId().GetValue())
	// assert.Equal(t, int32(11), testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetCellObjId().GetValue())
	// mr := testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasData().GetValue()[1]
	// assert.Equal(t, mr.GetUeId().GetValue(), measRec2.GetUeId().GetValue())
	// assert.Equal(t, mr.GetUeTag().GetValue(), measRec2.GetUeTag().GetValue())
	// mil := testIM.GetIndicationMessageFormats().GetIndicationMessageFormat1().GetMeasInfoList().GetValue()
	// assert.Equal(t, mil[0].GetMeasType(), measInfoList.GetValue()[0].GetMeasType())

}

func TestServicemodel_RanFuncDescriptionProtoToASN1(t *testing.T) {
	var rfSn = "onf"
	var rfE2SMoid = "oid123"
	var rfd = "someDescription"

	newE2SmMetPdu, err := pdubuilder.CreateE2SmMetRanfunctionDescription(rfSn, rfE2SMoid, rfd)
	assert.NilError(t, err)
	assert.Assert(t, newE2SmMetPdu != nil)

	err = newE2SmMetPdu.Validate()
	assert.NilError(t, err, "error validating E2SmPDU")

	assert.NilError(t, err)
	protoBytes, err := proto.Marshal(newE2SmMetPdu)
	// fmt.Printf("\n--- %d", len(protoBytes))
	assert.NilError(t, err, "unexpected error marshalling E2SmMetRanfunctionDescription to bytes")

	asn1Bytes, err := metTestSm.RanFuncDescriptionProtoToASN1(protoBytes)
	//DONE the length is non the same ?
	// the length of a message of the same nature is the same but the lenghth of proto and asn messsages is diffrent
	// fmt.Printf("%#v --- %d", asn1Bytes, len(asn1Bytes))
	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
	assert.Assert(t, asn1Bytes != nil)
	t.Logf("E2SM-MET-RANfunctionDescription asn1Bytes are \n%v", hex.Dump(asn1Bytes))
}

// func TestServicemodel_RanFuncDescriptionASN1toProto(t *testing.T) {
// 	// This message is taken as an output from the function above
// 	ranFuncDescriptionAsn1 := []byte{0x0, 0x80, 0x6f, 0x6e, 0x66, 0x0, 0x0, 0x5, 0x6f,
// 		0x69, 0x64, 0x31, 0x32, 0x33, 0x7, 0x0, 0x73, 0x6f, 0x6d, 0x65, 0x44, 0x65,
// 		0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e}

// 	protoBytes, err := metTestSm.RanFuncDescriptionASN1toProto(ranFuncDescriptionAsn1)
// 	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
// 	assert.Assert(t, protoBytes != nil)
// 	assert.Equal(t, 32, len(protoBytes))
// 	testRFD := &e2smmet.E2SmMetRanfunctionDescription{}
// 	err = proto.Unmarshal(protoBytes, testRFD)
// 	t.Logf("Decoded message is \n%v", testRFD)
// 	assert.NilError(t, err)
// 	assert.Equal(t, "oid123", testRFD.GetRanFunctionName().GetRanFunctionE2SmOid())
// }

// // func TestServicemodel_EventTriggerDefinitionProtoToASN1(t *testing.T) {
// // 	var rtPeriod int64 = 12
// // 	e2SmMetEventTriggerDefinition, err := pdubuilder.CreateE2SmMetEventTriggerDefinition(rtPeriod)
// // 	assert.NilError(t, err, "error creating E2SmPDU")
// // 	assert.Assert(t, e2SmMetEventTriggerDefinition != nil, "Created E2SmPDU is nil")

// // 	err = e2SmMetEventTriggerDefinition.Validate()
// // 	assert.NilError(t, err, "error validating E2SmPDU")

// // 	assert.NilError(t, err)
// // 	protoBytes, err := proto.Marshal(e2SmMetEventTriggerDefinition)
// // 	assert.NilError(t, err, "unexpected error marshalling E2SmMetEventTriggerDefinition to bytes")

// // 	asn1Bytes, err := metTestSm.EventTriggerDefinitionProtoToASN1(protoBytes)
// // 	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
// // 	assert.Assert(t, asn1Bytes != nil)
// // 	t.Logf("E2SM-MET-EventTriggerDefinition asn1Bytes are \n%v", hex.Dump(asn1Bytes))
// // }

// // func TestServicemodel_EventTriggerDefinitionASN1toProto(t *testing.T) {
// // 	// This value is taken from Shad and passed as a byte array directly to the function
// // 	// It's the encoding of what's in the file ../test/E2SM-MET-EventTriggerDefinition.xml
// // 	eventTriggerDefinitionAsn1 := []byte{0x00, 0x0b}

// // 	protoBytes, err := metTestSm.EventTriggerDefinitionASN1toProto(eventTriggerDefinitionAsn1)
// // 	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
// // 	assert.Assert(t, protoBytes != nil)
// // 	assert.Equal(t, 6, len(protoBytes))
// // 	testETD := &e2smmet.E2SmMetEventTriggerDefinition{}
// // 	err = proto.Unmarshal(protoBytes, testETD)
// // 	t.Logf("Decoded message is \n%v", testETD)
// // 	assert.NilError(t, err)
// // 	t.Logf("Reporting period is \n%v", testETD.GetEventDefinitionFormats().GetEventDefinitionFormat1().GetReportingPeriod())
// // 	assert.Equal(t, int64(12), testETD.GetEventDefinitionFormats().GetEventDefinitionFormat1().GetReportingPeriod())
// // }

// // func TestServicemodel_ActionDefinitionProtoToASN1(t *testing.T) {
// // 	var ricStyleType int32 = 12
// // 	var cellObjID = "onf"
// // 	var granularity int64 = 21
// // 	var subscriptionID int64 = 12345
// // 	var measurementName = "trial"

// // 	var valEnum int64 = 201
// // 	tce := e2smmet.TestCondExpression_TEST_COND_EXPRESSION_LESSTHAN
// // 	tci, err := pdubuilder.CreateTestCondInfo(pdubuilder.CreateTestCondTypeRSRP(), tce, pdubuilder.CreateTestCondValueEnum(valEnum))
// // 	assert.NilError(t, err)

// // 	mci, err := pdubuilder.CreateMatchingCondItemTestCondInfo(tci)
// // 	assert.NilError(t, err)

// // 	mcl := &e2smmet.MatchingCondList{
// // 		Value: make([]*e2smmet.MatchingCondItem, 0),
// // 	}
// // 	mcl.Value = append(mcl.Value, mci)

// // 	measName, err := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
// // 	assert.NilError(t, err)
// // 	measCondItem, err := pdubuilder.CreateMeasurementCondItem(measName, mcl)
// // 	assert.NilError(t, err)

// // 	measCondList := e2smmet.MeasurementCondList{
// // 		Value: make([]*e2smmet.MeasurementCondItem, 0),
// // 	}
// // 	measCondList.Value = append(measCondList.Value, measCondItem)

// // 	actionDefinitionFormat3, err := pdubuilder.CreateActionDefinitionFormat3(cellObjID, &measCondList, granularity, subscriptionID)
// // 	assert.NilError(t, err)

// // 	newE2SmMetPdu, err := pdubuilder.CreateE2SmMetActionDefinitionFormat3(ricStyleType, actionDefinitionFormat3)
// // 	assert.NilError(t, err, "error creating E2SmPDU")
// // 	assert.Assert(t, newE2SmMetPdu != nil, "Created E2SmPDU is nil")

// // 	err = newE2SmMetPdu.Validate()
// // 	assert.NilError(t, err, "error validating E2SmPDU")

// // 	assert.NilError(t, err)
// // 	protoBytes, err := proto.Marshal(newE2SmMetPdu)
// // 	assert.NilError(t, err, "unexpected error marshalling E2SmMetActionDefinition to bytes")

// // 	asn1Bytes, err := metTestSm.ActionDefinitionProtoToASN1(protoBytes)
// // 	assert.NilError(t, err, "unexpected error converting protoBytes to asnBytes")
// // 	assert.Assert(t, asn1Bytes != nil)
// // 	t.Logf("E2SM-MET-ActionDefinition (Format3) asn1Bytes are \n%v", hex.Dump(asn1Bytes))
// // }

// // func TestServicemodel_ActionDefinitionASN1toProto(t *testing.T) {
// // 	actionDefinitionAsn1 := []byte{0x00, 0x0c, 0x40, 0x00, 0x03, 0x6f, 0x6e, 0x66, 0x00, 0x00, 0x00, 0x40, 0x74, 0x72, 0x69, 0x61,
// // 		0x6c, 0x00, 0x00, 0x48, 0x21, 0x02, 0x00, 0xc9, 0x00, 0x14, 0x40, 0x30, 0x38}

// // 	protoBytes, err := metTestSm.ActionDefinitionASN1toProto(actionDefinitionAsn1)
// // 	assert.NilError(t, err, "unexpected error converting protoBytes to asn1Bytes")
// // 	assert.Assert(t, protoBytes != nil)
// // 	testAD := &e2smmet.E2SmMetActionDefinition{}
// // 	err = proto.Unmarshal(protoBytes, testAD)
// // 	assert.NilError(t, err)
// // 	t.Logf("Decoded message is \n%v", testAD)
// // 	assert.Equal(t, int32(12), testAD.GetRicStyleType().GetValue())
// // 	adf3 := testAD.GetActionDefinitionFormats().GetActionDefinitionFormat3()
// // 	assert.Equal(t, "onf", adf3.GetCellObjId().GetValue())
// // 	assert.Equal(t, int64(12345), adf3.GetSubscriptId().GetValue())
// // 	assert.Equal(t, int64(21), adf3.GetGranulPeriod().GetValue())
// // 	mcl := adf3.GetMeasCondList().GetValue()[0]
// // 	assert.Equal(t, "trial", mcl.GetMeasType().GetMeasName().GetValue())
// // 	mc := mcl.GetMatchingCond().GetValue()[0]
// // 	assert.Equal(t, int64(201), mc.GetTestCondInfo().GetTestValue().GetValueEnum())
// // }
