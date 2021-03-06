/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "../v2/e2sm-rc-pre_v2_rsys.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RANparameter_Type_RCPRE_H_
#define	_RANparameter_Type_RCPRE_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum RANparameter_Type_RCPRE {
	RANparameter_Type_RCPRE_integer	= 0,
	RANparameter_Type_RCPRE_enumerated	= 1,
	RANparameter_Type_RCPRE_boolean	= 2,
	RANparameter_Type_RCPRE_bit_string	= 3,
	RANparameter_Type_RCPRE_octet_string	= 4,
	RANparameter_Type_RCPRE_printable_string	= 5
	/*
	 * Enumeration is extensible
	 */
} e_RANparameter_Type_RCPRE;

/* RANparameter-Type-RCPRE */
typedef long	 RANparameter_Type_RCPRE_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_RANparameter_Type_RCPRE_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_RANparameter_Type_RCPRE;
extern const asn_INTEGER_specifics_t asn_SPC_RANparameter_Type_RCPRE_specs_1;
asn_struct_free_f RANparameter_Type_RCPRE_free;
asn_struct_print_f RANparameter_Type_RCPRE_print;
asn_constr_check_f RANparameter_Type_RCPRE_constraint;
ber_type_decoder_f RANparameter_Type_RCPRE_decode_ber;
der_type_encoder_f RANparameter_Type_RCPRE_encode_der;
xer_type_decoder_f RANparameter_Type_RCPRE_decode_xer;
xer_type_encoder_f RANparameter_Type_RCPRE_encode_xer;
per_type_decoder_f RANparameter_Type_RCPRE_decode_uper;
per_type_encoder_f RANparameter_Type_RCPRE_encode_uper;
per_type_decoder_f RANparameter_Type_RCPRE_decode_aper;
per_type_encoder_f RANparameter_Type_RCPRE_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RANparameter_Type_RCPRE_H_ */
#include "asn_internal.h"
