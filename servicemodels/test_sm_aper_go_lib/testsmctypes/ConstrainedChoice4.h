/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_ConstrainedChoice4_H_
#define	_ConstrainedChoice4_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ConstrainedChoice4_PR {
	ConstrainedChoice4_PR_NOTHING,	/* No components present */
	ConstrainedChoice4_PR_constrainedChoice4A
	/* Extensions may appear below */
	
} ConstrainedChoice4_PR;

/* ConstrainedChoice4 */
typedef struct ConstrainedChoice4 {
	ConstrainedChoice4_PR present;
	union ConstrainedChoice4_u {
		long	 constrainedChoice4A;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ConstrainedChoice4_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ConstrainedChoice4;
extern asn_CHOICE_specifics_t asn_SPC_ConstrainedChoice4_specs_1;
extern asn_TYPE_member_t asn_MBR_ConstrainedChoice4_1[1];
extern asn_per_constraints_t asn_PER_type_ConstrainedChoice4_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _ConstrainedChoice4_H_ */
#include "asn_internal.h"
