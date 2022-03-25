// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRmetEventTriggerDefinition(rtPeriod int64) (*e2smrmet.E2SmRmetEventTriggerDefinition, error) {
	e2SmRmetPdu := e2smrmet.E2SmRmetEventTriggerDefinition{
		EventDefinitionFormats: &e2smrmet.EventTriggerDefinitionFormats{
			E2SmRmetEventTriggerDefinition: &e2smrmet.EventTriggerDefinitionFormats_EventDefinitionFormat1{
				EventDefinitionFormat1: &e2smrmet.E2SmRmetEventTriggerDefinitionFormat1{
					ReportingPeriod: rtPeriod,
				},
			},
		},
	}

	if err := e2SmRmetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRmetEventTriggerDefinition(): error validating E2SmRmetPDU %s", err.Error())
	}
	return &e2SmRmetPdu, nil
}
