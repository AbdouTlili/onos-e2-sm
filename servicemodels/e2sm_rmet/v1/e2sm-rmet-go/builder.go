package e2sm_rmet_go

func (s *Snssai) SetSliceID(sd []byte) *Snssai {
	s.SD = sd
	return s
}

func (ih *E2SmRmetIndicationHeader) SetFileFormatVersion(ffv string) *E2SmRmetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmRmetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().FileFormatversion = &ffv
	default:
		return ih
	}
	return ih
}

func (ih *E2SmRmetIndicationHeader) SetSenderName(sn string) *E2SmRmetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmRmetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().SenderName = &sn
	default:
		return ih
	}
	return ih
}

func (ih *E2SmRmetIndicationHeader) SetSenderType(st string) *E2SmRmetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmRmetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().SenderType = &st
	default:
		return ih
	}
	return ih
}

func (ih *E2SmRmetIndicationHeader) SetVendorName(vn string) *E2SmRmetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmRmetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().VendorName = &vn
	default:
		return ih
	}
	return ih
}

func (ih *E2SmRmetIndicationHeader) SetGlobalRmetnodeID(gknID *GlobalRmetnodeId) *E2SmRmetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmRmetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().RmetNodeId = gknID
	default:
		return ih
	}
	return ih
}

func (im *E2SmRmetIndicationMessage) SetMeasInfoList(measInfoList *MeasurementInfoList) *E2SmRmetIndicationMessage {
	switch im.IndicationMessageFormats.E2SmRmetIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().MeasInfoList = measInfoList
	// Left for future extensions
	//case *IndicationMessageFormats_IndicationMessageFormat2:
	//	im.GetIndicationMessageFormats().GetIndicationMessageFormat2().MeasInfoList = measInfoList
	default:
		return im
	}

	return im
}

func (im *E2SmRmetIndicationMessage) SetGranularityPeriod(gp int64) *E2SmRmetIndicationMessage {
	switch im.IndicationMessageFormats.E2SmRmetIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GranulPeriod = &GranularityPeriod{
			Value: gp,
		}
	// case *IndicationMessageFormats_IndicationMessageFormat2:
	// 	im.GetIndicationMessageFormats().GetIndicationMessageFormat2().GranulPeriod = &GranularityPeriod{
	// 		Value: gp,
	// 	}
	default:
		return im
	}

	return im
}

func (im *E2SmRmetIndicationMessage) SetCellObjectID(cellObjID string) *E2SmRmetIndicationMessage {
	switch im.IndicationMessageFormats.E2SmRmetIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().CellObjId = &CellObjectId{
			Value: cellObjID,
		}
	// case *IndicationMessageFormats_IndicationMessageFormat2:
	// 	im.GetIndicationMessageFormats().GetIndicationMessageFormat2().CellObjId = &CellObjectId{
	// 		Value: cellObjID,
	// 	}
	default:
		return im
	}

	return im
}

func (rfd *E2SmRmetRanfunctionDescription) SetRanFunctionInstance(rfi int32) *E2SmRmetRanfunctionDescription {
	rfd.GetRanFunctionName().RanFunctionInstance = &rfi
	return rfd
}

func (rfd *E2SmRmetRanfunctionDescription) SetRicRmetNodeList(rknl []*RicRmetnodeItem) *E2SmRmetRanfunctionDescription {
	rfd.RicRmetNodeList = rknl
	return rfd
}

func (rfd *E2SmRmetRanfunctionDescription) SetRicEventTriggerStyleList(retsl []*RicEventTriggerStyleItem) *E2SmRmetRanfunctionDescription {
	rfd.RicEventTriggerStyleList = retsl
	return rfd
}

func (rfd *E2SmRmetRanfunctionDescription) SetRicReportStyleList(rrsl []*RicReportStyleItem) *E2SmRmetRanfunctionDescription {
	rfd.RicReportStyleList = rrsl
	return rfd
}

func (rkni *RicRmetnodeItem) SetCellMeasurementObjectList(cmol []*CellMeasurementObjectItem) *RicRmetnodeItem {
	rkni.CellMeasurementObjectList = cmol
	return rkni
}
