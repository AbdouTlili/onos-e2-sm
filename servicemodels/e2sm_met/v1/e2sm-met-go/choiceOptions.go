// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2sm_met_go

import (
	"reflect"
)

var E2smMetChoicemap = map[string]map[int]reflect.Type{
	"measurement_record_item": {
		1: reflect.TypeOf(MeasurementRecordItem_Integer{}),
		2: reflect.TypeOf(MeasurementRecordItem_Real{}),
		3: reflect.TypeOf(MeasurementRecordItem_NoValue{}),
	},
	"e2_sm_met_event_trigger_definition": {
		1: reflect.TypeOf(EventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	},
	"e2_sm_met_indication_header": {
		1: reflect.TypeOf(IndicationHeaderFormats_IndicationHeaderFormat1{}),
	},
	"e2_sm_met_indication_message": {
		1: reflect.TypeOf(IndicationMessageFormats_IndicationMessageFormat1{}),
	},
}
