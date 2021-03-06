/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TEST-ListExtensible2.h"

#include "ItemExtensible.h"
static asn_per_constraints_t asn_PER_type_TEST_ListExtensible2_constr_1 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  3,  3,  0,  4 }	/* (SIZE(0..4,...)) */,
	0, 0	/* No PER value map */
};
static asn_TYPE_member_t asn_MBR_TEST_ListExtensible2_1[] = {
	{ ATF_POINTER, 0, 0,
		(ASN_TAG_CLASS_UNIVERSAL | (16 << 2)),
		0,
		&asn_DEF_ItemExtensible,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		""
		},
};
static const ber_tlv_tag_t asn_DEF_TEST_ListExtensible2_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static asn_SET_OF_specifics_t asn_SPC_TEST_ListExtensible2_specs_1 = {
	sizeof(struct TEST_ListExtensible2),
	offsetof(struct TEST_ListExtensible2, _asn_ctx),
	0,	/* XER encoding is XMLDelimitedItemList */
};
asn_TYPE_descriptor_t asn_DEF_TEST_ListExtensible2 = {
	"TEST-ListExtensible2",
	"TEST-ListExtensible2",
	&asn_OP_SEQUENCE_OF,
	asn_DEF_TEST_ListExtensible2_tags_1,
	sizeof(asn_DEF_TEST_ListExtensible2_tags_1)
		/sizeof(asn_DEF_TEST_ListExtensible2_tags_1[0]), /* 1 */
	asn_DEF_TEST_ListExtensible2_tags_1,	/* Same as above */
	sizeof(asn_DEF_TEST_ListExtensible2_tags_1)
		/sizeof(asn_DEF_TEST_ListExtensible2_tags_1[0]), /* 1 */
	{ 0, &asn_PER_type_TEST_ListExtensible2_constr_1, SEQUENCE_OF_constraint },
	asn_MBR_TEST_ListExtensible2_1,
	1,	/* Single element */
	&asn_SPC_TEST_ListExtensible2_specs_1	/* Additional specs */
};

