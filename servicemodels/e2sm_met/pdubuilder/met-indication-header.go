// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmMetIndicationHeader(timeStamp []byte, mil *e2smmet.MeasurementInfoList) (*e2smmet.E2SmMetIndicationHeader, error) {

	if len(timeStamp) != 4 {
		return nil, errors.NewInvalid("TimeStamp should be 4 chars")
	}

	e2SmMetPdu := e2smmet.E2SmMetIndicationHeader{
		IndicationHeaderFormats: &e2smmet.IndicationHeaderFormats{
			E2SmMetIndicationHeader: &e2smmet.IndicationHeaderFormats_IndicationHeaderFormat1{
				IndicationHeaderFormat1: &e2smmet.E2SmMetIndicationHeaderFormat1{
					ColletStartTime: &e2smmet.TimeStamp{
						Value: timeStamp,
					},
					MeasInfoList: mil,
				},
			},
		},
	}

	if err := e2SmMetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMetIndicationHeader(): error validating E2SmMetPDU %s", err.Error())
	}
	return &e2SmMetPdu, nil
}

func CreateGlobalMetnodeID(globalKpmNodeID int64) (*e2smmet.GlobalMetnodeId, error) {
	globalMetnodeIDObj := &e2smmet.GlobalMetnodeId{
		Value: globalKpmNodeID,
	}

	if err := globalMetnodeIDObj.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmMetIndicationHeader(): error validating globalMetnodeIDObj %s", err.Error())
	}
	return globalMetnodeIDObj, nil
}
