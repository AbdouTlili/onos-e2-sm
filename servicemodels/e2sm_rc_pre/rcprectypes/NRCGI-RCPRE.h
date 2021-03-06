/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "../v2/e2sm-rc-pre_v2_rsys.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_NRCGI_RCPRE_H_
#define	_NRCGI_RCPRE_H_


#include "asn_application.h"

/* Including external dependencies */
#include "PLMN-Identity-RCPRE.h"
#include "NRCellIdentity-RCPRE.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* NRCGI-RCPRE */
typedef struct NRCGI_RCPRE {
	PLMN_Identity_RCPRE_t	 pLMN_Identity;
	NRCellIdentity_RCPRE_t	 nRCellIdentity;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} NRCGI_RCPRE_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_NRCGI_RCPRE;
extern asn_SEQUENCE_specifics_t asn_SPC_NRCGI_RCPRE_specs_1;
extern asn_TYPE_member_t asn_MBR_NRCGI_RCPRE_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _NRCGI_RCPRE_H_ */
#include "asn_internal.h"
