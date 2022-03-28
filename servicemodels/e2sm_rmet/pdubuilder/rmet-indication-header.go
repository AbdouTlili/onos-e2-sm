// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
package pdubuilder

import (
	e2smrmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_rmet/v1/e2sm-rmet-go"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func CreateE2SmRmetIndicationHeader(timeStamp []byte) (*e2smrmet.E2SmRmetIndicationHeader, error) {

	if len(timeStamp) != 4 {
		return nil, errors.NewInvalid("TimeStamp should be 4 chars")
	}

	e2SmRmetPdu := e2smrmet.E2SmRmetIndicationHeader{
		IndicationHeaderFormats: &e2smrmet.IndicationHeaderFormats{
			E2SmRmetIndicationHeader: &e2smrmet.IndicationHeaderFormats_IndicationHeaderFormat1{
				IndicationHeaderFormat1: &e2smrmet.E2SmRmetIndicationHeaderFormat1{
					ColletStartTime: &e2smrmet.TimeStamp{
						Value: timeStamp,
					},
				},
			},
		},
	}

	if err := e2SmRmetPdu.Validate(); err != nil {
		return nil, errors.NewInvalid("CreateE2SmRmetIndicationHeader(): error validating E2SmRmetPDU %s", err.Error())
	}
	return &e2SmRmetPdu, nil
}

func CreateGlobalRmetnodeIDgNBID(bs *asn1.BitString, plmnID []byte) (*e2smrmet.GlobalRmetnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if bs.GetLen() < 22 || bs.GetLen() > 32 {
		return nil, errors.NewInvalid("expecting GNbID length in range 22 to 32 bits, got %d", bs.GetLen())
	}

	return &e2smrmet.GlobalRmetnodeId{
		GlobalRmetnodeId: &e2smrmet.GlobalRmetnodeId_GNb{
			GNb: &e2smrmet.GlobalRmetnodeGnbId{
				GlobalGNbId: &e2smrmet.GlobalgNbId{
					GnbId: &e2smrmet.GnbIdChoice{
						GnbIdChoice: &e2smrmet.GnbIdChoice_GnbId{
							GnbId: bs,
						},
					},
					PlmnId: &e2smrmet.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateGlobalRmetnodeIDenGNbID(bsValue []byte, bsLen uint32, plmnID []byte) (*e2smrmet.GlobalRmetnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	if bsLen < 22 || bsLen > 32 {
		return nil, errors.NewInvalid("expecting GNbID length in range 22 to 32 bits, got %d", bsLen)
	}

	return &e2smrmet.GlobalRmetnodeId{
		GlobalRmetnodeId: &e2smrmet.GlobalRmetnodeId_EnGNb{
			EnGNb: &e2smrmet.GlobalRmetnodeEnGnbId{
				GlobalGNbId: &e2smrmet.GlobalenGnbId{
					GNbId: &e2smrmet.EngnbId{
						EngnbId: &e2smrmet.EngnbId_GNbId{
							GNbId: &asn1.BitString{
								Value: bsValue,
								Len:   bsLen, // should be 22 to 32
							},
						},
					},
					PLmnIdentity: &e2smrmet.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateGlobalRmetnodeIDngENbID(enbID *e2smrmet.EnbIdChoice, plmnID []byte, shortMacroEnbID *asn1.BitString,
	longMacroEnbID *asn1.BitString) (*e2smrmet.GlobalRmetnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	return &e2smrmet.GlobalRmetnodeId{
		GlobalRmetnodeId: &e2smrmet.GlobalRmetnodeId_NgENb{
			NgENb: &e2smrmet.GlobalRmetnodeNgEnbId{
				GlobalNgENbId: &e2smrmet.GlobalngeNbId{
					EnbId: enbID,
					PlmnId: &e2smrmet.PlmnIdentity{
						Value: plmnID,
					},
					ShortMacroENbId: shortMacroEnbID,
					LongMacroENbId:  longMacroEnbID,
				},
			},
		},
	}, nil
}

func CreateEnbIDchoiceMacro(bs *asn1.BitString) (*e2smrmet.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2smrmet.EnbIdChoice{
		EnbIdChoice: &e2smrmet.EnbIdChoice_EnbIdMacro{
			EnbIdMacro: bs,
		},
	}, nil
}

func CreateEnbIDchoiceShortMacro(bs *asn1.BitString) (*e2smrmet.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x3f > 0 {
		return nil, errors.NewInvalid("expected last 6 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2smrmet.EnbIdChoice{
		EnbIdChoice: &e2smrmet.EnbIdChoice_EnbIdShortmacro{
			EnbIdShortmacro: bs,
		},
	}, nil
}

func CreateEnbIDchoiceLongMacro(bs *asn1.BitString) (*e2smrmet.EnbIdChoice, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x07 > 0 {
		return nil, errors.NewInvalid("expected last 3 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}
	return &e2smrmet.EnbIdChoice{
		EnbIdChoice: &e2smrmet.EnbIdChoice_EnbIdLongmacro{
			EnbIdLongmacro: bs,
		},
	}, nil
}

func CreateGlobalRmetnodeIDeNBID(enbID *e2smrmet.EnbId, plmnID []byte) (*e2smrmet.GlobalRmetnodeId, error) {

	if len(plmnID) != 3 {
		return nil, errors.NewInvalid("PlmnID should be 3 chars")
	}

	return &e2smrmet.GlobalRmetnodeId{
		GlobalRmetnodeId: &e2smrmet.GlobalRmetnodeId_ENb{
			ENb: &e2smrmet.GlobalRmetnodeEnbId{
				GlobalENbId: &e2smrmet.GlobalEnbId{
					ENbId: enbID,
					PLmnIdentity: &e2smrmet.PlmnIdentity{
						Value: plmnID,
					},
				},
			},
		},
	}, nil
}

func CreateHomeEnbID(bs *asn1.BitString) (*e2smrmet.EnbId, error) {

	if len(bs.GetValue()) != 4 {
		return nil, errors.NewInvalid("expecting length to be exactly 4 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[3]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[3])
	}

	return &e2smrmet.EnbId{
		EnbId: &e2smrmet.EnbId_HomeENbId{
			HomeENbId: bs,
		},
	}, nil
}

func CreateMacroEnbID(bs *asn1.BitString) (*e2smrmet.EnbId, error) {

	if len(bs.GetValue()) != 3 {
		return nil, errors.NewInvalid("expecting length to be exactly 3 bytes, got %d", len(bs.GetValue()))
	}
	if bs.GetValue()[2]&0x0f > 0 {
		return nil, errors.NewInvalid("expected last 4 bits of byte array to be unused, and to contain only trailing zeroes. %b", bs.GetValue()[2])
	}

	return &e2smrmet.EnbId{
		EnbId: &e2smrmet.EnbId_MacroENbId{
			MacroENbId: bs,
		},
	}, nil
}
