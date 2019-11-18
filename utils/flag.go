package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

var (
	TokenFlag = cli.StringFlag{
		Name:  "token",
		Usage: "erc20 token contract address",
	}

	FromFlag = cli.StringFlag{
		Name:  "from",
		Usage: "erc20 token sender",
	}

	ToFlag = cli.StringFlag{
		Name:  "to",
		Usage: "erc20 token receiver",
	}

	ValueFlag = cli.StringFlag{
		Name:  "value",
		Usage: "erc20 token amount",
	}

	NonceFlag = cli.Uint64Flag{
		Name:  "nonce",
		Usage: "multisig tx nonce",
	}

	SignerFlag = cli.StringFlag{
		Name:  "signer",
		Usage: "tx signer",
	}

	SenderFlag = cli.StringFlag{
		Name:  "sender",
		Usage: "tx sender",
	}

	KeyFlag = cli.StringFlag{
		Name:  "key",
		Usage: "private key",
	}

	KeystoreFlag = cli.StringFlag{
		Name:  "keystore",
		Usage: "keystore path",
	}

	PasswordFlag = cli.StringFlag{
		Name:  "password",
		Usage: "password for keystore",
	}

	SignaturesFlag = cli.StringSliceFlag{
		Name:  "signatures",
		Usage: "signature list",
	}
)

func FlagToAddress(ctx *cli.Context, flag *cli.StringFlag) (token common.Address, _ error) {
	if s := ctx.String(flag.Name); common.IsHexAddress(s) {
		token = common.HexToAddress(s)
		return token, nil
	} else {
		return token, fmt.Errorf("invalid %s address", flag.Name)
	}
}
