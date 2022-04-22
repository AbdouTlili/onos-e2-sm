package e2sm_met_go

func (ih *E2SmMetIndicationHeader) SetFileFormatVersion(ffv string) *E2SmMetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmMetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().FileFormatversion = &ffv
	default:
		return ih
	}
	return ih
}

func (ih *E2SmMetIndicationHeader) SetSenderName(sn string) *E2SmMetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmMetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().SenderName = &sn
	default:
		return ih
	}
	return ih
}

func (ih *E2SmMetIndicationHeader) SetGlobalMETnodeID(gknID *GlobalMetnodeId) *E2SmMetIndicationHeader {
	switch ih.GetIndicationHeaderFormats().E2SmMetIndicationHeader.(type) {
	case *IndicationHeaderFormats_IndicationHeaderFormat1:
		ih.GetIndicationHeaderFormats().GetIndicationHeaderFormat1().MetNodeId = gknID
	default:
		return ih
	}
	return ih
}

func (im *E2SmMetIndicationMessage) SetMeasInfoList(measInfoList *MeasurementInfoList) *E2SmMetIndicationMessage {
	switch im.IndicationMessageFormats.E2SmMetIndicationMessage.(type) {
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

func (im *E2SmMetIndicationMessage) SetGranularityPeriod(gp int64) *E2SmMetIndicationMessage {
	switch im.IndicationMessageFormats.E2SmMetIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().GranulPeriod = &GranularityPeriod{
			Value: gp,
		}
	default:
		return im
	}

	return im
}

func (im *E2SmMetIndicationMessage) SetCellObjectID(cellObjID int32) *E2SmMetIndicationMessage {
	switch im.IndicationMessageFormats.E2SmMetIndicationMessage.(type) {
	case *IndicationMessageFormats_IndicationMessageFormat1:
		im.GetIndicationMessageFormats().GetIndicationMessageFormat1().CellObjId = &CellObjectId{
			Value: cellObjID,
		}

	default:
		return im
	}

	return im
}
