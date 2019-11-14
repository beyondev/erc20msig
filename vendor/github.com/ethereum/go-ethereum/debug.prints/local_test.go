package debugutils_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

const (
	ERC20Transfer = "transfer(address,uint256)"
)

func sigHash(fn string, token, from, to common.Address, value, nonce *big.Int) (common.Hash, error) {
	var h common.Hash
	fnId, err := getFnId(fn)
	if err != nil {
		return h, err
	}

	var b []byte
	//b = append(b, fnId...)
	b = append(b, common.LeftPadBytes(fnId, 4)...)
	b = append(b, token.Bytes()...)
	//b = append(b, common.LeftPadBytes(token.Bytes(), 32)...)
	b = append(b, from.Bytes()...)
	//b = append(b, common.LeftPadBytes(from.Bytes(), 32)...)
	b = append(b, to.Bytes()...)
	//b = append(b, common.LeftPadBytes(to.Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(value.Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(nonce.Bytes(), 32)...)

	hasher := sha3.NewLegacyKeccak256()
	if _, err := hasher.Write(b); err != nil {
		return h, err
	}

	hasher.Sum(h[:0])
	return h, nil
}

func getFnId(fn string) ([]byte, error) {
	hw := sha3.NewLegacyKeccak256()
	if _, err := hw.Write([]byte(fn)); err != nil {
		return nil, err
	}
	return hw.Sum(nil)[:4], nil
}

func sigValues(sig []byte) (r, s, v *big.Int, err error) {
	if len(sig) != 65 {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v, nil
}

func sign(hash common.Hash, dir string, addr common.Address, pass string) ([]byte, error) {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	acct := accounts.Account{Address: addr}
	if err := ks.Unlock(acct, pass); err != nil {
		return nil, err
	}
	return ks.SignHash(acct, hash.Bytes())
}

func Test_signHash(t *testing.T) {
	hash, _ := sigHash(ERC20Transfer,
		common.HexToAddress("0x3390a97486c4771aD0895b123c4bda6F4f956e10"),
		common.HexToAddress("0x7eE9cb50561f45646F1113A18f60f884FB2c45cf"),
		common.HexToAddress("0xd0f113e0b5639945e24b2f00856b5285af06d33d"),
		big.NewInt(10), big.NewInt(0))
	fmt.Println(hash.String())
}

func Test_getFnId(t *testing.T) {
	fnId, _ := getFnId(ERC20Transfer)
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

	hash, err := sigHash(ERC20Transfer, token, from, to, value, nonce)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("hash:", hash.String())

	sig, err := sign(hash, "./keystore", common.HexToAddress("0xd339985aca76fd1af57556ed37b090f543851837"), "1")
	//sig, err := sign(hash, "./keystore", common.HexToAddress("0x7e844D1F96af3dAcFEBCF545750778eb8b6d762e"), "1")
	if err != nil {
		t.Fatal(err)
	}

	hex := common.Bytes2Hex(sig)
	fmt.Println("sighex:", hex)

	r, s, v, _ := sigValues(sig)

	fmt.Printf("v: %d\nr: %s\ns: %s\n", v.Uint64(), hexutil.Encode(r.Bytes()), hexutil.Encode(s.Bytes()))

}
