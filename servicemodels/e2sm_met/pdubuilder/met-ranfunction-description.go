// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMetRanfunctionDescription(rfSn string, rfE2SMoid string, rfd string) (*e2smmet.E2SmMetRanfunctionDescription, error) {

	e2SmMetPdu := e2smmet.E2SmMetRanfunctionDescription{
		RanFunctionName: &e2smmet.RanfunctionName{
			RanFunctionShortName:   rfSn,
			RanFunctionE2SmOid:     rfE2SMoid,
			RanFunctionDescription: rfd,
		},
	}

	if err := e2SmMetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMetRanfunctionDescription(): error validating E2SmMetRanfunctionDescription %s", err.Error())
	}
	return &e2SmMetPdu, nil
}
