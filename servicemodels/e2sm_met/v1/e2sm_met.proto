////////////////////// e2sm-met-ies.proto //////////////////////
// Protobuf generated from /e2sm-met.asn1 by asn1c-0.9.29
// E2SM-MET-IEs { iso(1) identified-organization(3) dod(6) internet(1) private(4) enterprise(1) oran(53148) e2(1) version2(2) e2sm(2) e2sm-METMON-IEs(98) }

syntax = "proto3";

package e2sm_met.v1;
option go_package = "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go";

import "validate/v1/validate.proto";
import "asn1/v1/asn1.proto";

// range of Integer from e2sm-met.asn1:16
// {CellObjectID}
message CellObjectId {
    int32 value = 1 [(validate.v1.rules).int32 = {gte: 0, lte: 255}, json_name="value"];
};

// range of Integer from e2sm-met.asn1:18
// {GranularityPeriod}
message GranularityPeriod {
    // @inject_tag: aper:"valueLB:1,valueUB:4294967295"
    int64 value = 1 [(validate.v1.rules).int64 = {gte: 1, lte: 4294967295}, json_name = "value"];
};

// range of Integer from e2sm-met.asn1:20
// {SubscriptionID}
message SubscriptionId {
    // @inject_tag: aper:"valueLB:1,valueUB:4294967295"
    int64 value = 1 [(validate.v1.rules).int64 = {gte: 1, lte: 4294967295}, json_name = "value"];
};

// range of Integer from e2sm-met.asn1:24
// {GlobalMETnode-ID}
message GlobalMetnodeId {
    // @inject_tag: aper:"valueLB:1,valueUB:4294967295"
    int64 value = 1 [(validate.v1.rules).int64 = {gte: 1, lte: 4294967295}, json_name = "value"];
};

// sequence from e2sm-met.asn1:29
// {RANfunction-Name}
message RanfunctionName {
    // @inject_tag: aper:"sizeExt,sizeLB:1,sizeUB:150"
    string ran_function_short_name = 1 [(validate.v1.rules).string = {min_len: 1, max_len: 150}, json_name="ranFunction-ShortName"];
    // @inject_tag: aper:"sizeExt,sizeLB:1,sizeUB:1000"
    string ran_function_e2_sm_oid = 2 [(validate.v1.rules).string = {min_len: 1, max_len: 1000}, json_name="ranFunction-E2SM-OID"];
    // @inject_tag: aper:"sizeExt,sizeLB:1,sizeUB:150"
    string ran_function_description = 3 [(validate.v1.rules).string = {min_len: 1, max_len: 150}, json_name="ranFunction-Description"];
};

// constant Integer from e2sm-met.asn1:39
// {-}
message MaxnoofMeasurementInfo {
    // @inject_tag: aper:"valueLB:65536,valueUB:65536,"
    int32 value = 1 [(validate.v1.rules).int32.const = 65536, json_name="value"];
};

// constant Integer from e2sm-met.asn1:40
// {-}
message MaxnoofMeasurementRecord {
    // @inject_tag: aper:"valueLB:65536,valueUB:65536,"
    int32 value = 1 [(validate.v1.rules).int32.const = 65536, json_name="value"];
};

// constant Integer from e2sm-met.asn1:41
// {-}
message MaxnoofMeasurementValue {
    // @inject_tag: aper:"valueLB:65536,valueUB:65536,"
    int32 value = 1 [(validate.v1.rules).int32.const = 65536, json_name="value"];
};

// constant Integer from e2sm-met.asn1:42
// {-}
message MaxnoofRicstyles {
    // @inject_tag: aper:"valueLB:63,valueUB:63,"
    int32 value = 1 [(validate.v1.rules).int32.const = 63, json_name="value"];
};

// constant Integer from e2sm-met.asn1:47
// {-}
message MaxofUe {
    // @inject_tag: aper:"valueLB:65536,valueUB:65536,"
    int32 value = 1 [(validate.v1.rules).int32.const = 65536, json_name="value"];
};

// range of Integer from e2sm-met.asn1:49
// {UEID}
message Ueid {
    // @inject_tag: aper:"valueLB:1,valueUB:4294967295"
    int64 value = 1 [(validate.v1.rules).int64 = {gte: 1, lte: 4294967295}, json_name = "value"];
};

// sequence from e2sm_met.asn1:52
// {UEItem}
message Ueitem {
    Ueid ue_id = 1 [ json_name="ueID"];
    // @inject_tag: aper:"optional"
    optional Uetag ue_tag = 2 [ json_name="ueTag"];
};

// sequence from e2sm-met.asn1:63
// {MeasurementInfoList}
message MeasurementInfoList {
    // @inject_tag: aper:"valueExt,sizeLB:1,sizeUB:65535"
    repeated MeasurementInfoItem value = 1 [ json_name="value"];
};

// sequence from e2sm-met.asn1:64
// {MeasurementInfoItem}
message MeasurementInfoItem {
    string meas_type = 1 [(validate.v1.rules).string = {min_len: 0, max_len: 15}, json_name="measType"];
};

// sequence from e2sm-met.asn1:70
// {MeasurementData}
message MeasurementData {
    // @inject_tag: aper:"valueExt,sizeLB:1,sizeUB:65535"
    repeated MeasurementRecord value = 1 [ json_name="value"];
};

// sequence from e2sm-met.asn1:71
// {MeasurementRecord}
message MeasurementRecord {
    Ueid ue_id = 1 [ json_name="ueID"];
    Uetag ue_tag = 2 [ json_name="ueTag"];
    // @inject_tag: aper:"choiceExt,sizeLB:0,sizeUB:65536"
    repeated MeasurementRecordItem meas_record_item = 3 [(validate.v1.rules).repeated = {min_items: 0, max_items: 65536}, json_name="measRecordItem"];
};

// sequence from e2sm-met.asn1:78
// {MeasurementRecordItem}
message MeasurementRecordItem {
    // choice from e2sm_met.asn1:76
    oneof measurement_record_item {
        // @inject_tag: aper:"choiceIdx:1,valueLB:0,valueUB:4294967295"
        int64 integer = 1 [json_name = "integer"];
        // @inject_tag: aper:"choiceIdx:2"
        double real = 2 [json_name = "real"];
        // @inject_tag: aper:"choiceIdx:3,valueLB:0,valueUB:0"
        int32 no_value = 3 [json_name = "noValue"];
    }
};

// sequence from e2sm-met.asn1:86
// {MeasurementInfo-Action-List}
message MeasurementInfoActionList {
    // @inject_tag: aper:"valueExt,sizeLB:1,sizeUB:65535"
    repeated MeasurementInfoActionItem value = 1 [json_name = "value"];
};

// sequence from e2sm-met.asn1:87
// {MeasurementInfo-Action-Item}
message MeasurementInfoActionItem {
    MeasurementTypeName meas_name = 1 [json_name = "measName"];
    // @inject_tag: aper:"optional"
    optional MeasurementTypeId meas_id = 2 [json_name = "measID"];
};

// range of Integer from e2sm-met.asn1:94
// {MeasurementTypeID}
message MeasurementTypeId {
    // @inject_tag: aper:"valueExt,valueLB:1,valueUB:65536"
    int32 value = 1 [(validate.v1.rules).int32 = {gte: 1, lte: 65536}, json_name = "value"];
};

// sequence from e2sm-met.asn1:108
// {E2SM-MET-EventTriggerDefinition}
message E2SmMetEventTriggerDefinition {
    // choice from e2sm-met-ies_v2.asn1:337
    // @inject_tag: aper:"choiceExt"
    EventTriggerDefinitionFormats event_definition_formats = 1 [json_name = "eventDefinition-formats"];
};

message EventTriggerDefinitionFormats {
    oneof e2_sm_met_event_trigger_definition {
        // @inject_tag: aper:"choiceIdx:1,valueExt"
        E2SmMetEventTriggerDefinitionFormat1 event_definition_format1 = 1 [json_name = "eventDefinition_Format1"];
    }
}

// sequence from e2sm_met_v2.0.3-radisys.asn:342
// {E2SM-MET-EventTriggerDefinition-Format1}
message E2SmMetEventTriggerDefinitionFormat1 {
    // @inject_tag: aper:"valueLB:1,valueUB:4294967295"
    int64 reporting_period = 1 [json_name = "reportingPeriod"];
};

// sequence from e2sm-met.asn1:122
// {E2SM-MET-ActionDefinition}
message E2SmMetActionDefinition {
    RicStyleType ric_style_type = 1 [json_name = "ric-Style-Type"];
    // @inject_tag: aper:"choiceExt"
    ActionDefinitionFormats action_definition_formats = 2 [json_name = "actionDefinition-formats"];
};

message ActionDefinitionFormats {
    oneof e2_sm_met_action_definition {
        // @inject_tag: aper:"valueExt,choiceIdx:1"
        E2SmMetActionDefinitionFormat1 action_definition_format1 = 2 [json_name = "actionDefinition_Format1"];
    }
}

// sequence from e2sm_met_v2.0.3-radisys.asn:362
// {E2SM-MET-ActionDefinition-Format1}
message E2SmMetActionDefinitionFormat1 {
    CellObjectId cell_obj_id = 1 [json_name = "cellObjID"];
    MeasurementInfoList meas_info_list = 2 [json_name = "measInfoList"];
    GranularityPeriod granul_period = 3 [json_name = "granulPeriod"];
    SubscriptionId subscript_id = 4 [json_name = "subscriptID"];
};

// sequence from e2sm-met.asn1:146
// {E2SM-MET-IndicationHeader}
message E2SmMetIndicationHeader {
    // @inject_tag: aper:"choiceExt"
    IndicationHeaderFormats indication_header_formats = 1 [json_name = "indicationHeader-formats"];
  };
  
  message IndicationHeaderFormats {
      oneof e2_sm_met_indication_header {
          // @inject_tag: aper:"valueExt,choiceIdx:1"
          E2SmMetIndicationHeaderFormat1 indication_header_format1 = 1 [json_name = "indicationHeader_Format1"];
      }
  }
  
  // sequence from e2sm_met.asn1:120
  // {E2SM-MET-IndicationHeader-Format1}
  message E2SmMetIndicationHeaderFormat1 {
      TimeStamp collet_start_time = 1 [ json_name="colletStartTime"];
      // @inject_tag: aper:"optional,sizeExt,sizeLB:0,sizeUB:15"
      optional string file_formatversion = 2 [(validate.v1.rules).string = {min_len: 0, max_len: 15}, json_name="fileFormatversion"];
      // @inject_tag: aper:"optional,sizeExt,sizeLB:0,sizeUB:400"
      optional string sender_name = 3 [(validate.v1.rules).string = {min_len: 0, max_len: 400}, json_name="senderName"];
      // @inject_tag: aper:"optional"
      optional GlobalMetnodeId met_node_id = 4 [ json_name="metNodeID"];
  };
  

// sequence from e2sm-met.asn1:166
// {E2SM-MET-IndicationMessage}
message E2SmMetIndicationMessage {
    // @inject_tag: aper:"choiceExt"
    IndicationMessageFormats indication_message_formats = 1 [json_name = "indicationMessage-formats"];
 };
 
 message IndicationMessageFormats {
     oneof e2_sm_met_indication_message {
         // @inject_tag: aper:"valueExt,choiceIdx:1"
         E2SmMetIndicationMessageFormat1 indication_message_format1 = 1 [json_name = "indicationMessage_Format1"];
     }
 }
 // sequence from e2sm_met.asn1:140
 // {E2SM-MET-IndicationMessage-Format1}
 message E2SmMetIndicationMessageFormat1 {
     SubscriptionId subscript_id = 1 [ json_name="subscriptID"];
     // @inject_tag: aper:"optional"
     optional CellObjectId cell_obj_id = 2 [ json_name="cellObjID"];
     // @inject_tag: aper:"optional"
     optional GranularityPeriod granul_period = 3 [ json_name="granulPeriod"];
     // @inject_tag: aper:"optional"
     optional MeasurementInfoList meas_info_list = 4 [ json_name="measInfoList"];
     MeasurementData meas_data = 5 [ json_name="measData"];
 };

// sequence from e2sm-met.asn1:184
// {E2SM-MET-RANfunction-Description}
message E2SmMetRanfunctionDescription {
    // @inject_tag: aper:"valueExt"
    RanfunctionName ran_function_name = 1 [json_name = "ranFunction-Name"];
    // @inject_tag: aper:"optional,valueExt,sizeLB:1,sizeUB:63"
    repeated RicEventTriggerStyleItem ric_event_trigger_style_list = 3 [(validate.v1.rules).repeated = {min_items: 0, max_items: 63}, json_name = "ric-EventTriggerStyle-List"];
    // @inject_tag: aper:"optional,valueExt,sizeLB:1,sizeUB:63"
    repeated RicReportStyleItem ric_report_style_list = 4 [(validate.v1.rules).repeated = {min_items: 0, max_items: 63}, json_name = "ric-ReportStyle-List"];
};
// range of Integer from e2sm-met.asn1:192
// {RIC-Style-Type}
message RicStyleType {
    // @inject_tag: aper:"valueExt,valueLB:0,valueUB:255"
    int32 value = 1 [(validate.v1.rules).int32 = {gte: 0, lte: 255}, json_name = "value"];
};

// range of Integer from e2sm-met.asn1:196
// {RIC-Format-Type}
message RicFormatType {
    // @inject_tag: aper:"valueExt,valueLB:0,valueUB:255"
    int32 value = 1 [(validate.v1.rules).int32 = {gte: 0, lte: 255}, json_name = "value"];
};

// sequence from e2sm-met.asn1:197
// {RIC-EventTriggerStyle-Item}
message RicEventTriggerStyleItem {
    RicStyleType ric_event_trigger_style_type = 1 [json_name = "ric-EventTriggerStyle-Type"];
    RicStyleName ric_event_trigger_style_name = 2 [json_name = "ric-EventTriggerStyle-Name"];
    RicFormatType ric_event_trigger_format_type = 3 [json_name = "ric-EventTriggerFormat-Type"];
};

// sequence from e2sm-met.asn1:204
// {RIC-ReportStyle-Item}
message RicReportStyleItem {
    RicStyleType ric_report_style_type = 1 [json_name = "ric-ReportStyle-Type"];
    RicStyleName ric_report_style_name = 2 [json_name = "ric-ReportStyle-Name"];
    RicFormatType ric_action_format_type = 3 [json_name = "ric-ActionFormat-Type"];
    MeasurementInfoActionList meas_info_action_list = 4 [json_name = "measInfo-Action-List"];
    RicFormatType ric_indication_header_format_type = 5 [json_name = "ric-IndicationHeaderFormat-Type"];
    RicFormatType ric_indication_message_format_type = 6 [json_name = "ric-IndicationMessageFormat-Type"];
};

message RicStyleName {
    // @inject_tag: aper:"sizeExt,sizeLB:1,sizeUB:150"
    string value = 1 [(validate.v1.rules).string = {min_len: 1, max_len: 150}, json_name = "value"];
}

//{TimeStamp}
message TimeStamp {
    // @inject_tag: aper:"sizeLB:4,sizeUB:4"
    bytes value = 1 [(validate.v1.rules).bytes.len = 4, json_name = "value"];
}


//{CellObjectID}
message Uetag {
    // @inject_tag: aper:"sizeExt,sizeLB:0,sizeUB:400"
    string value = 1 [(validate.v1.rules).string = {min_len: 0, max_len: 150}, json_name = "value"];
}

//{MeasurementTypeName}
message MeasurementTypeName {
    // @inject_tag: aper:"sizeExt,sizeLB:1,sizeUB:150"
    string value = 1 [(validate.v1.rules).string = {min_len: 1, max_len: 150}, json_name = "value"];
}