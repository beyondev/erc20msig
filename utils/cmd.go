package utils

import (
	"fmt"
	"math/big"

	"github.com/Beyond-simplechain/erc20msig/multisig"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli"
)

var (
	SignCommand = cli.Command{
		Action: func(ctx *cli.Context) error {
			token, err := FlagToAddress(ctx, &TokenFlag)
			from, err := FlagToAddress(ctx, &FromFlag)
			to, err := FlagToAddress(ctx, &ToFlag)
			if err != nil {
				return err
			}

			value, ok := new(big.Int).SetString(ctx.String(ValueFlag.Name), 10)
			if !ok {
				return fmt.Errorf("invalid value: %s", ctx.String(ValueFlag.Name))
			}
			nonce := new(big.Int).SetUint64(ctx.Uint64(NonceFlag.Name))

			key := ctx.String(KeyFlag.Name)
			if key != "" {
				pk, err := crypto.HexToECDSA(key)
				if err != nil {
					return err
				}

				return multisig.Sign(token, from, to, value, nonce, common.Address{}, "", "", pk)
			}

			signer, err := FlagToAddress(ctx, &SignerFlag)
			if err != nil {
				return err
			}
			keystore := ctx.String(KeystoreFlag.Name)
			if keystore == "" {
				return fmt.Errorf("invalid keystore path %s", keystore)
			}
			password := ctx.String(PasswordFlag.Name)

			return multisig.Sign(token, from, to, value, nonce, signer, keystore, password, nil)
		},
		Name:  "sign",
		Usage: "sign transaction by address",
		Flags: []cli.Flag{
			TokenFlag, FromFlag, ToFlag, ValueFlag, NonceFlag, SignerFlag, KeyFlag, KeystoreFlag, PasswordFlag,
		},
	}

	SendCommand = cli.Command{
		Action: func(ctx *cli.Context) error {
			token, err := FlagToAddress(ctx, &TokenFlag)
			from, err := FlagToAddress(ctx, &FromFlag)
			to, err := FlagToAddress(ctx, &ToFlag)
			if err != nil {
				return err
			}

			value, ok := new(big.Int).SetString(ctx.String(ValueFlag.Name), 10)
			if !ok {
				return fmt.Errorf("invalid value: %s", ctx.String(ValueFlag.Name))
			}

			signatures := ctx.StringSlice(SignaturesFlag.Name)

			sender, err := FlagToAddress(ctx, &SignerFlag)
			if err != nil {
				return err
			}
			keystore := ctx.String(KeystoreFlag.Name)
			if keystore == "" {
				return fmt.Errorf("invalid keystore path %s", keystore)
			}
			password := ctx.String(PasswordFlag.Name)

			return multisig.Send("", token, from, to, value, signatures, sender, keystore, password)
		},
		Name:  "sign",
		Usage: "sign transaction by address",
		Flags: []cli.Flag{
			TokenFlag, FromFlag, ToFlag, ValueFlag, NonceFlag, SignerFlag, KeystoreFlag, PasswordFlag,
		},
	}
)
