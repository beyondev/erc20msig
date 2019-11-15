# erc20msig
##部署合约msig.sol
+ `M`: 多签需要的签名数量
+ `__members`: 多签地址列表

##多重签名
```
./erc20msig sign --token=0xcb732820DA130e3AC8f9E00588caC657628FABFE \
--from=0xb3778d4a40959f9a26C5C42FD1Ac952Fafb9C900 --to=0xd0f113e0b5639945e24b2f00856b5285af06d33d \
--value=100 --nonce=0 --signer 0xd339985aca76fd1af57556ed37b090f543851837 \
--keystore=/Users/yuanchao/eth/testdata/keystore --password=1 
```
+ `token`: erc20合约地址
+ `from`: 多签合约账户地址
+ `to`: 转账收款人
+ `value`: 转账金额
+ `nonce`: 交易nonce，多签需要指定相同的nonce，转账完成后自动递增(初始为0)
+ `signer`: 本次签名的账户
+ `keystore`: keystore目录
+ `password`: keystore账户解锁密码

##发送多签后的交易

+ `v`: 签名V值列表
+ `r`: 签名V值列表
+ `s`: 签名V值列表
+ `_token`: erc20合约地址
+ `_to`: 转账收款人
+ `_value`: 转账金额

```
[27],["0x09461a45c4ba81b811875908cf11a5e82bcd29603733e0717d69e091bc9de9c2"],["0x54beccdbf4082c265b2525edff9194faa601a286b995cc5f68244d42ab6d46fb"],"0xcb732820DA130e3AC8f9E00588caC657628FABFE","0xd0f113e0b5639945e24b2f00856b5285af06d33d","100"
```