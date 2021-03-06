/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_TEST_Enumerated_H_
#define	_TEST_Enumerated_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum TEST_Enumerated {
	TEST_Enumerated_enum1	= 0,
	TEST_Enumerated_enum2	= 1,
	TEST_Enumerated_enum3	= 2,
	TEST_Enumerated_enum4	= 3,
	TEST_Enumerated_enum5	= 4,
	TEST_Enumerated_enum6	= 5
} e_TEST_Enumerated;

/* TEST-Enumerated */
typedef long	 TEST_Enumerated_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TEST_Enumerated;
asn_struct_free_f TEST_Enumerated_free;
asn_struct_print_f TEST_Enumerated_print;
asn_constr_check_f TEST_Enumerated_constraint;
ber_type_decoder_f TEST_Enumerated_decode_ber;
der_type_encoder_f TEST_Enumerated_encode_der;
xer_type_decoder_f TEST_Enumerated_decode_xer;
xer_type_encoder_f TEST_Enumerated_encode_xer;
per_type_decoder_f TEST_Enumerated_decode_uper;
per_type_encoder_f TEST_Enumerated_encode_uper;
per_type_decoder_f TEST_Enumerated_decode_aper;
per_type_encoder_f TEST_Enumerated_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _TEST_Enumerated_H_ */
#include "asn_internal.h"
