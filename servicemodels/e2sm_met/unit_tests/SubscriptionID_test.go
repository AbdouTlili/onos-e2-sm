package met

import (
	"encoding/hex"
	"fmt"

	e2smmet "github.com/AbdouTlili/onos-e2-sm/servicemodels/e2sm_met/v1/e2sm-met-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	// hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerSubID = "0x00,0x09"

func Test_perEncodingSubscriptionID(t *testing.T) {

	subID := &e2smmet.SubscriptionId{
		Value: 10,
	}

	per, err := aper.Marshal(subID, nil, nil)
	fmt.Printf("%#v--", per)
	assert.NilError(t, err)
	t.Logf("SubscriptionID PER\n%v", hex.Dump(per))

	result := e2smmet.SubscriptionId{}
	err = aper.Unmarshal(per, &result, nil, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("SubscriptionID PER - decoded\n%v", &result)
	assert.Equal(t, subID.GetValue(), result.GetValue())
}

func Test_perSubscriptionIDCompareBytes(t *testing.T) {

	subID := &e2smmet.SubscriptionId{
		Value: 10,
	}

	per, err := aper.Marshal(subID, nil, nil)
	assert.NilError(t, err)
	t.Logf("SubscriptionID PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	// perRefBytes, err := hexlib.DumpToByte(refPerSubID)
	perRefBytes := []byte{0x00, 0x09}
	t.Logf("%#v", perRefBytes)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
