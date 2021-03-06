# AtomicFS Wallet Command Line Interface


## Create new wallet

```
afs wallet new

2022-05-16T14:15:32.857+0200    INFO    Successfully loaded configuration file. {"root": "atomicfs", "path": "/home/cortex/atomicfs/config.yaml"}
Please enter wallet password: ****************
New account created: 0x210a23c2e142c4dc39b628aca6E4F99F6F4c37B1
Saved in: /home/cortex/atomicfs/keystore
```

## Dump the contents of the keystore

An example keystore file 

```
afs wallet dump --keystore ~/atomicfs/keystore/UTC--2022-05-16T11-51-01.022642972Z--e43049f09ca93406bf1bcc6646a1b69eae77df05
```

```
2022-05-16T14:05:52.966+0200    INFO    Successfully loaded configuration file. {"root": "atomicfs", "path": "/home/cortex/atomicfs/config.yaml"}
Please enter a password to decrypt the wallet: ****************
(*keystore.Key)(0xc00019e360)({
 Id: (uuid.UUID) (len=16 cap=16) 48ef9059-d2fe-4c67-a970-3ba107699238,
 Address: (common.Address) (len=20 cap=20) 0xe43049F09CA93406Bf1bcC6646a1b69eae77DF05,
 PrivateKey: (*ecdsa.PrivateKey)(0xc00019e300)({
  PublicKey: (ecdsa.PublicKey) {
   Curve: (*secp256k1.BitCurve)(0xc00011b7a0)({
    P: (*big.Int)(0xc00014e4e0)(115792089237316195423570985008687907853269984665640564039457584007908834671663),
    N: (*big.Int)(0xc00014e520)(115792089237316195423570985008687907852837564279074904382605163141518161494337),
    B: (*big.Int)(0xc00014e560)(7),
    Gx: (*big.Int)(0xc00014e5a0)(55066263022277343669578718895168534326250603453777594175500187360389116729240),
    Gy: (*big.Int)(0xc00014e5e0)(32670510020758816978083085130507043184471273380659243275938904335757337482424),
    BitSize: (int) 256
   }),
   X: (*big.Int)(0xc00014e040)(46817351676451284027004570633162974802343842193961792855937379496502088483823),
   Y: (*big.Int)(0xc00014e120)(85935256213564115532912944951268975492949542864702092253687836764285194091120)
  },
  D: (*big.Int)(0xc00014e020)(114559874429000718450040373015896173838529453334519861524768539435157947912923)
 })
})
```