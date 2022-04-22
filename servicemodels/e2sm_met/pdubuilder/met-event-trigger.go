package pdubuilder

import (
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMetEventTriggerDefinition(rtPeriod int64) (*e2smmet.E2SmMetEventTriggerDefinition, error) {
	e2SmMetPdu := e2smmet.E2SmMetEventTriggerDefinition{
		EventDefinitionFormats: &e2smmet.EventTriggerDefinitionFormats{
			E2SmMetEventTriggerDefinition: &e2smmet.EventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: &e2smmet.E2SmMetEventTriggerDefinitionFormat1{
					ReportingPeriod: rtPeriod,
				},
			},
		},
	}

	if err := e2SmMetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMetEventTriggerDefinition(): error validating E2SmMetPDU %s", err.Error())
	}
	return &e2SmMetPdu, nil
}
