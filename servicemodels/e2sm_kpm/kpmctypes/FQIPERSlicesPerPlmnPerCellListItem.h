/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "e2sm-kpm-v01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_FQIPERSlicesPerPlmnPerCellListItem_H_
#define	_FQIPERSlicesPerPlmnPerCellListItem_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* FQIPERSlicesPerPlmnPerCellListItem */
typedef struct FQIPERSlicesPerPlmnPerCellListItem {
	long	 fiveQI;
	long	*dl_PRBUsage;	/* OPTIONAL */
	long	*ul_PRBUsage;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} FQIPERSlicesPerPlmnPerCellListItem_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_FQIPERSlicesPerPlmnPerCellListItem;
extern asn_SEQUENCE_specifics_t asn_SPC_FQIPERSlicesPerPlmnPerCellListItem_specs_1;
extern asn_TYPE_member_t asn_MBR_FQIPERSlicesPerPlmnPerCellListItem_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _FQIPERSlicesPerPlmnPerCellListItem_H_ */
#include "asn_internal.h"
