package e2sm_rmet_go

import (
	"reflect"
)

var E2smRmetChoicemap = map[string]map[int]reflect.Type{
	"cell_global_id": {
		1: reflect.TypeOf(CellGlobalId_NrCgi{}),
		2: reflect.TypeOf(CellGlobalId_EUtraCgi{}),
	},
	"measurement_type": {
		1: reflect.TypeOf(MeasurementType_MeasName{}),
		2: reflect.TypeOf(MeasurementType_MeasId{}),
	},
	"global_rmetnode_id": {
		1: reflect.TypeOf(GlobalRmetnodeId_GNb{}),
		2: reflect.TypeOf(GlobalRmetnodeId_EnGNb{}),
		3: reflect.TypeOf(GlobalRmetnodeId_NgENb{}),
		4: reflect.TypeOf(GlobalRmetnodeId_ENb{}),
	},
	"gnb_id_choice": {
		1: reflect.TypeOf(GnbIdChoice_GnbId{}),
	},
	"engnb_id": {
		1: reflect.TypeOf(EngnbId_GNbId{}),
	},
	"enb_id_choice": {
		1: reflect.TypeOf(EnbIdChoice_EnbIdMacro{}),
		2: reflect.TypeOf(EnbIdChoice_EnbIdShortmacro{}),
		3: reflect.TypeOf(EnbIdChoice_EnbIdLongmacro{}),
	},
	"enb_id": {
		1: reflect.TypeOf(EnbId_MacroENbId{}),
		2: reflect.TypeOf(EnbId_HomeENbId{}),
	},
	"measurement_record_item": {
		1: reflect.TypeOf(MeasurementRecordItem_Integer{}),
		2: reflect.TypeOf(MeasurementRecordItem_Real{}),
		3: reflect.TypeOf(MeasurementRecordItem_NoValue{}),
	},
	"e2_sm_rmet_event_trigger_definition": {
		1: reflect.TypeOf(EventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_rmet_action_definition": {
		1: reflect.TypeOf(ActionDefinitionFormats_ActionDefinitionFormat1{}),
	},
	"e2_sm_rmet_indication_header": {
		1: reflect.TypeOf(IndicationHeaderFormats_IndicationHeaderFormat1{}),
	},
	"e2_sm_rmet_indication_message": {
		1: reflect.TypeOf(IndicationMessageFormats_IndicationMessageFormat1{}),
	},
}
