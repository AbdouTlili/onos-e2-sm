// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	"fmt"
	e2sm_rc_pre_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v2/e2sm-rc-pre-v2"
)

func CreateE2SmRcPreEventTriggerDefinitionPeriodic(rtPeriod uint32) (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition, error) {

	if rtPeriod > 4294967295 {
		return nil, fmt.Errorf("RT-Period should be within range 0 to 4294967295")
	}

	eventDefinitionFormat1 := &e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1{
		TriggerType:       e2sm_rc_pre_v2.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_PERIODIC,
		ReportingPeriodMs: rtPeriod,
	}

	E2SmRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition{
		E2SmRcPreEventTriggerDefinitionEventDefinitionFormats: &e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: eventDefinitionFormat1,
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmRcPrePDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}

func CreateE2SmRcPreEventTriggerDefinitionUponChange() (*e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition, error) {

	eventDefinitionFormat1 := &e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinitionFormat1{
		TriggerType:       e2sm_rc_pre_v2.RcPreTriggerType_RC_PRE_TRIGGER_TYPE_UPON_CHANGE,
		ReportingPeriodMs: 0,
	}

	E2SmRcPrePdu := e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition{
		E2SmRcPreEventTriggerDefinitionEventDefinitionFormats: &e2sm_rc_pre_v2.E2SmRcPreEventTriggerDefinition_EventDefinitionFormat1{
			EventDefinitionFormat1: eventDefinitionFormat1,
		},
	}

	if err := E2SmRcPrePdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmRcPrePDU %s", err.Error())
	}
	return &E2SmRcPrePdu, nil
}
