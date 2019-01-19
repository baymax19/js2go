# js2go
Implement golang cosmos-sdk txns using js with help of [gopherjs](github.com/gopherjs/gopherjs) library


#Pre-requirments

**golang**  
**dep**  
**nodejs**

#install
```$xslt
go get -u github.com/baymax19/js2go

git checkout baymax19

dep ensure
```


build the main.go using
```$xslt
gopherjs build main.go 
```
you will get **main.js** file in the directory

#Transactions
- createKey
```$xslt
parameters:
   name : string
   password : string
```
- sendCoins
```$xslt
 parameter:
 from : string
 to : string 
 amount : string
 seed : string  // from address seed to sign the txn
```
#example

**createKey**

in bash 
```
node

> var main = require("./main")
undefinded  // main contain the functions of golang (createKey, sendCoins)
> main.createKey("name","password")

{ address: 'cosmos143570cxhpkwh47elnlfuau2hzjg75axup9qz00',
  pub_key: 'cosmospub1addwnpepq08vmcafz30kjr7yxul0u8tz86p9kck603up6t6a4a69w3mvgcsv5rxm5wa',
  name: 'name',
  seed: 'view address symptom girl steel feature silver item quiz canal ugly march example antique submit barrel dizzy mix man easily stadium area border dwarf' }
```

**sendCoins**
```
> var seed = "view address symptom girl steel feature silver item quiz canal ugly march example antique submit barrel dizzy mix man easily stadium area border dwarf"
> main.sendCoins("cosmos1v0m40792sx0cf69elugcqqxmqg3rdy7ra0j9kl","cosmos143570cxhpkwh47elnlfuau2hzjg75axup9qz00","1STAKE",seed)

0AHwYl3uCkwqLIf6CiIKFGP3V/iqgZ+E6Ln/EYAA2wIiNpPDEgoKBVNUQUtFEgExEiIKFKxp5+DXDZ16+z+f087xVxSR6nTcEgoKBVNUQUtFEgExEhAKCgoFU1RBS0USATAQwJoMGmoKJuta6YchAiHNyQjIUA99h2twp8nljvkuEQ72MYFaHl76aRrB4KAsEkBCg8ThLKXpDHe6yiRiui4cp6lyPXxcB2Oo/8TAd2I65A7sPei3D0iHjJLpnRLDnuWp3mTswQuLgAizGFB7dVdY  //txbytes

```
generated txbytes are send throw axios post method in [tendermint_rpc](https://github.com/tendermint/tendermint/wiki/RPC)
```$xslt
> var a = require("axios"); 
> var data = null ; 
> a.post('http://localhost:26657', 
    { 
    "method" : "broadcast_tx_sync", 
    "jsonrpc" : "2.0", 
    "params":["0AHwYl3uCkwqLIf6CiIKFGP3V/iqgZ+E6Ln/EYAA2wIiNpPDEgoKBVNUQUtFEgExEiIKFL+dpB7HaoCf5hMRAdDR4mfaJi0NEgoKBVNUQUtFEgExEhAKCgoFU1RBS0USATEQwJoMGmoKJuta6YchAiHNyQjIUA99h2twp8nljvkuEQ72MYFaHl76aRrB4KAsEkDX9z7SriUkBMBtAo3BaQoqn7qO9ngslOHF9Lsi1RWIfmIRKWdKGr80xEZZHVmlYNlWPiuXfAW2SJSJAfay4x1f"],
    "id":"dontcare"
    }).then((res)=>data = res.data)

> data

{ jsonrpc: '2.0',
  id: 'dontcare',
  result: 
   { code: 0,
     data: '',
     log: '',
     hash: '972AE23D40F9A518ABA3708A9597A9AC7F3A73B93F10E924661E35BE4EFB1DAE' } }

```

#Note
for these I am used cosmos-sdk v0.29.1  
before post txbytes using axios, start the gaiad using 
``` gaiad start```  
this project totally based on the reference of [cosmos-sdk](github.com/cosmos/cosmos-sdk) and [tendermint](github.com/tendermint),  
it will useful to run the golang code in webbrowser using js