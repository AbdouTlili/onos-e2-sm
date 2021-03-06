/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TEST-Choices.h"

static asn_TYPE_member_t asn_MBR_TEST_Choices_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_Choices, otherAttr),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"otherAttr"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_Choices, choice1),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_Choice1,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"choice1"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_Choices, choice2),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_Choice2,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"choice2"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_Choices, choice3),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_Choice3,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"choice3"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_Choices, choice4),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_Choice4,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"choice4"
		},
};
static const ber_tlv_tag_t asn_DEF_TEST_Choices_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_TEST_Choices_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* otherAttr */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* choice1 */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* choice2 */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* choice3 */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 } /* choice4 */
};
static asn_SEQUENCE_specifics_t asn_SPC_TEST_Choices_specs_1 = {
	sizeof(struct TEST_Choices),
	offsetof(struct TEST_Choices, _asn_ctx),
	asn_MAP_TEST_Choices_tag2el_1,
	5,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_TEST_Choices = {
	"TEST-Choices",
	"TEST-Choices",
	&asn_OP_SEQUENCE,
	asn_DEF_TEST_Choices_tags_1,
	sizeof(asn_DEF_TEST_Choices_tags_1)
		/sizeof(asn_DEF_TEST_Choices_tags_1[0]), /* 1 */
	asn_DEF_TEST_Choices_tags_1,	/* Same as above */
	sizeof(asn_DEF_TEST_Choices_tags_1)
		/sizeof(asn_DEF_TEST_Choices_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_TEST_Choices_1,
	5,	/* Elements count */
	&asn_SPC_TEST_Choices_specs_1	/* Additional specs */
};

