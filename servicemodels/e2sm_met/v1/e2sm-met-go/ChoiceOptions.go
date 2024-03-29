package e2smmet

import "reflect"

var E2smMetChoicemap = map[string]map[int]reflect.Type{
	// "measurement_record_item": {
	// 	1: reflect.TypeOf(MeasurementRecordItem_Integer{}),
	// 	2: reflect.TypeOf(MeasurementRecordItem_Str{}),
	// 	3: reflect.TypeOf(MeasurementRecordItem_NoValue{}),
	// },
	"e2_sm_met_event_trigger_definition": {
		1: reflect.TypeOf(EventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_met_action_definition": {
		1: reflect.TypeOf(ActionDefinitionFormats_ActionDefinitionFormat1{}),
	},
	"e2_sm_met_indication_header": {
		1: reflect.TypeOf(IndicationHeaderFormats_IndicationHeaderFormat1{}),
	},
	"e2_sm_met_indication_message": {
		1: reflect.TypeOf(IndicationMessageFormats_IndicationMessageFormat1{}),
	},
}
