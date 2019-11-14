package multisig

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	"golang.org/x/crypto/sha3"
)

const (
	ERC20Transfer = "transfer(address,uint256)"
)

func encodeHash(fn string, token, from, to common.Address, value, nonce *big.Int) (common.Hash, error) {
	var h common.Hash
	fnId, err := calFnId(fn)
	if err != nil {
		return h, err
	}

	var b []byte
	b = append(b, common.LeftPadBytes(fnId, 4)...)
	b = append(b, token.Bytes()...)
	b = append(b, from.Bytes()...)
	b = append(b, to.Bytes()...)
	b = append(b, common.LeftPadBytes(value.Bytes(), 32)...)
	b = append(b, common.LeftPadBytes(nonce.Bytes(), 32)...)

	hasher := sha3.NewLegacyKeccak256()
	if _, err := hasher.Write(b); err != nil {
		return h, err
	}

	hasher.Sum(h[:0])
	return h, nil
}

func calFnId(fn string) ([]byte, error) {
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

func sigByKeyStore(hash common.Hash, dir string, addr common.Address, pass string) ([]byte, error) {
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	acct := accounts.Account{Address: addr}
	if err := ks.Unlock(acct, pass); err != nil {
		return nil, err
	}
	return ks.SignHash(acct, hash.Bytes())
}
