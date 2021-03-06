/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_TEST_BitString_H_
#define	_TEST_BitString_H_


#include "asn_application.h"

/* Including external dependencies */
#include "BIT_STRING.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* TEST-BitString */
typedef struct TEST_BitString {
	BIT_STRING_t	 attrBs1;
	BIT_STRING_t	 attrBs2;
	BIT_STRING_t	 attrBs3;
	BIT_STRING_t	 attrBs4;
	BIT_STRING_t	 attrBs5;
	BIT_STRING_t	 attrBs6;
	BIT_STRING_t	*attrBs7;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TEST_BitString_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TEST_BitString;
extern asn_SEQUENCE_specifics_t asn_SPC_TEST_BitString_specs_1;
extern asn_TYPE_member_t asn_MBR_TEST_BitString_1[7];

#ifdef __cplusplus
}
#endif

#endif	/* _TEST_BitString_H_ */
#include "asn_internal.h"
