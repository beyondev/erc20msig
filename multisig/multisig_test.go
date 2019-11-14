package multisig

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_encodeHash(t *testing.T) {
	hash, _ := encodeHash(ERC20Transfer,
		common.HexToAddress("0x3390a97486c4771aD0895b123c4bda6F4f956e10"),
		common.HexToAddress("0x7eE9cb50561f45646F1113A18f60f884FB2c45cf"),
		common.HexToAddress("0xd0f113e0b5639945e24b2f00856b5285af06d33d"),
		big.NewInt(10), big.NewInt(0))
	fmt.Println(hash.String())
}

func Test_getFnId(t *testing.T) {
	fnId, _ := calFnId(ERC20Transfer)
	if "0xa9059cbb" != hexutil.Encode(fnId) {
		t.Fatalf("want 0xa9059cbb, got %s", hexutil.Encode(fnId))
	}
}

func Test_sign(t *testing.T) {
	token := common.HexToAddress("0xcb732820DA130e3AC8f9E00588caC657628FABFE")
	from := common.HexToAddress("0xb3778d4a40959f9a26C5C42FD1Ac952Fafb9C900")
	to := common.HexToAddress("0xd0f113e0b5639945e24b2f00856b5285af06d33d")
	value := big.NewInt(15)
	nonce := big.NewInt(1)

	hash, err := encodeHash(ERC20Transfer, token, from, to, value, nonce)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("hash:", hash.String())

	sig, err := sigByKeyStore(hash, "./keystore", common.HexToAddress("0xd339985aca76fd1af57556ed37b090f543851837"), "1")
	//sig, err := sign(hash, "./keystore", common.HexToAddress("0x7e844D1F96af3dAcFEBCF545750778eb8b6d762e"), "1")
	if err != nil {
		t.Fatal(err)
	}

	r, s, v, _ := sigValues(sig)

	fmt.Printf("v: %d\nr: %s\ns: %s\n", v.Uint64(), hexutil.Encode(r.Bytes()), hexutil.Encode(s.Bytes()))

}
