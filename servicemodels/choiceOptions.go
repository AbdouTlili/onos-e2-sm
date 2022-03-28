// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package choiceOptions

import (
	e2smrmetv1 "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"reflect"
)

var E2smRmetChoicemap = map[string]map[int]reflect.Type{
	"cell_global_id": {
		1: reflect.TypeOf(e2smrmetv1.CellGlobalId_NrCgi{}),
		2: reflect.TypeOf(e2smrmetv1.CellGlobalId_EUtraCgi{}),
	},
	"measurement_type": {
		1: reflect.TypeOf(e2smrmetv1.MeasurementType_MeasName{}),
		2: reflect.TypeOf(e2smrmetv1.MeasurementType_MeasId{}),
	},
	"global_rmetnode_id": {
		1: reflect.TypeOf(e2smrmetv1.GlobalRmetnodeId_GNb{}),
		2: reflect.TypeOf(e2smrmetv1.GlobalRmetnodeId_EnGNb{}),
		3: reflect.TypeOf(e2smrmetv1.GlobalRmetnodeId_NgENb{}),
		4: reflect.TypeOf(e2smrmetv1.GlobalRmetnodeId_ENb{}),
	},
	"gnb_id_choice": {
		1: reflect.TypeOf(e2smrmetv1.GnbIdChoice_GnbId{}),
	},
	"engnb_id": {
		1: reflect.TypeOf(e2smrmetv1.EngnbId_GNbId{}),
	},
	"enb_id_choice": {
		1: reflect.TypeOf(e2smrmetv1.EnbIdChoice_EnbIdMacro{}),
		2: reflect.TypeOf(e2smrmetv1.EnbIdChoice_EnbIdShortmacro{}),
		3: reflect.TypeOf(e2smrmetv1.EnbIdChoice_EnbIdLongmacro{}),
	},
	"enb_id": {
		1: reflect.TypeOf(e2smrmetv1.EnbId_MacroENbId{}),
		2: reflect.TypeOf(e2smrmetv1.EnbId_HomeENbId{}),
	},
	"measurement_record_item": {
		1: reflect.TypeOf(e2smrmetv1.MeasurementRecordItem_Integer{}),
		2: reflect.TypeOf(e2smrmetv1.MeasurementRecordItem_Real{}),
		3: reflect.TypeOf(e2smrmetv1.MeasurementRecordItem_NoValue{}),
	},
	"e2_sm_rmet_event_trigger_definition": {
		1: reflect.TypeOf(e2smrmetv1.EventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_rmet_action_definition": {
		1: reflect.TypeOf(e2smrmetv1.ActionDefinitionFormats_ActionDefinitionFormat1{}),
	},
	"e2_sm_rmet_indication_header": {
		1: reflect.TypeOf(e2smrmetv1.IndicationHeaderFormats_IndicationHeaderFormat1{}),
	},
	"e2_sm_rmet_indication_message": {
		1: reflect.TypeOf(e2smrmetv1.IndicationMessageFormats_IndicationMessageFormat1{}),
	},
}
