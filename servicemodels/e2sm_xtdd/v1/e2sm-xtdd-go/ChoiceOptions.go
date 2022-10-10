package e2sm_xtdd_go

import "reflect"

var E2smXtddChoicemap = map[string]map[int]reflect.Type{
	// "measurement_record_item": {
	// 	1: reflect.TypeOf(MeasurementRecordItem_Integer{}),
	// 	2: reflect.TypeOf(MeasurementRecordItem_Str{}),
	// 	3: reflect.TypeOf(MeasurementRecordItem_NoValue{}),
	// },
	"e2_sm_xtdd_event_trigger_definition": {
		1: reflect.TypeOf(EventDefinitionformatsE2SmXTddEventTriggerDefinition_EventDefinitionFormat1{}),
	},
	"e2_sm_xtdd_action_definition": {
		1: reflect.TypeOf(ActionDefinitionformatsE2SmXTddActionDefinition_ActionDefinitionFormat1{}),
	},
	"e2_sm_xtdd_indication_header": {
		1: reflect.TypeOf(IndicationHeaderformatsE2SmXTddIndicationHeader_IndicationHeaderFormat1{}),
	},
	"e2_sm_xtdd_indication_message": {
		1: reflect.TypeOf(IndicationMessageformatsE2SmXTddIndicationMessage_IndicationMessageFormat1{}),
	},

	"e2_sm_xtdd_control_header": {
		1: reflect.TypeOf(E2SmXTddControlHeader_ControlHeaderFormat1{}),
	},
	"e2_sm_xtdd_control_message": {
		1: reflect.TypeOf(E2SmXTddControlMessage_ControlMessageFormat1{}),
	},
}
