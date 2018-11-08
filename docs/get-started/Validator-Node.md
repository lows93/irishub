# Running a Validator Node

Before setting up your validator node, make sure you've already installed  **Iris** by this [guide](Full-Node.md)

Validators are responsible for committing new blocks to the blockchain through consensus. A validator's stake will be slashed if they become unavailable, double sign a transaction, or don't cast their votes. Please read about Sentry Node Architecture to protect your node from DDOS attacks and to ensure high-availability.

## Get IRIS Token

### Create Account

You need to get `iris` and `iriscli` installed first. Then, follow the instructions below to create a new account:

```
iriscli keys add <NAME_OF_KEY>
```

Then, you should set a password of at least 8 characters.

The output will look like the following:
```
NAME:	TYPE:	ADDRESS:						PUBKEY:
tom	local	faa1arlugktm7p64uylcmh6w0g5m09ptvklxm5k69x	fap1addwnpepqvlmtpv7tke2k93vlyfpy2sxup93jfulll6r3jty695dkh09tekrzagazek
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

blast change tumble toddler rival ordinary chicken dirt physical club few language noise oak moment consider enemy claim elephant cruel people adult peanut garden
```

You could see the address and public key of this account. Please node that account address in IRISnet will start with `faa` and public key of account will start with `fap`.

The seed phrase of this account will also be displayed. You could use these 24 phrases to recover this account in another server. The recover command is:
```
iriscli keys add <NAME_OF_KEY> --recover
```


### Claim tokens

You can always get some `IRIS`  by using the [Faucet](https://testnet.irisplorer.io/#/faucet). The faucet will send you 10IRIS every request, Please don't abuse it.

Once you have created your own address, please  then you could use this　account to stake as a validatord. The following command is used to check the balance of your account:
```
iriscli bank account <ACCOUNT> --node=http://localhost:26657
```

## Create Validator

### Confirm Your Validator is Synced

Your validator is active if the following command returns anything:

```
iriscli status --node=tcp://localhost:26657 
```

You should also be able to see `catching_up` is `false`. 

You need to get the public key of your node before upgrade your node to a validator node. The public key of your node starts with `fvp`, 
it can be used to create a new validator by staking tokens. To understand more about the address encoding in IRISHub, 
please read this [doc](Bech32-on-IRISnet.md)

You can find your validator's pubkey by running:

```
iris tendermint show_validator --home=<IRIS-HOME>
```
Example output:
```
fvp1zcjduepqv7z2kgussh7ufe8e0prupwcm7l9jcn2fp90yeupaszmqjk73rjxq8yzw85
```
Next, use the output as  `<pubkey>` field for `iriscli stake create-validator` command:


```
iriscli stake create-validator --amount=XXiris --pubkey=<pubkey>  --moniker=<moniker> --fee=0.05iris  --gas=2000000  --chain-id=fuxi-4000  --node=http://localhost:26657
```
Please note the **fee** can be the **decimal** of IRIS token, like `0.01iris`. And you could also use other coin-type like `iris-milli`

To read more about fee mechanism in IRISHub, go to this [doc](../modules/fee-token/Fee.md)


In this way, to stake 1IRIS, you need to do:

```
iriscli stake create-validator --pubkey=pubkey  --fee=0.05iris  --gas=2000000 --from=<name> --chain-id=fuxi-4000   --node=tcp://localhost:26657  --amount=1iris
```
Don't forget the `fee` and `gas` field.  To read more about coin-type in IRISHub, you should read [this](../zh/modules/coin/README.md)



### View Validator Info

View the validator's information with this command:

```
iriscli stake validator <val-address-operator>  --chain-id=fuxi-4000 --node=tcp://localhost:26657 
```

The `<val-address-operator>` is your account address that starts with 'faa1'


### Confirm Your Validator is Running

Your validator is active if the following command returns anything:

```
iriscli status --node=tcp://localhost:26657 
```

You should also be able to see your power is above 0 if your bonded toke is in top 100. Also, you should see validator on the [Explorer](https://testnet.irisplorer.io).


### Edit Validator Description

You can edit your validator's public description. This info is to identify your validator, and will be relied on by delegators to decide which validators to stake to. Make sure to provide input for every flag below, otherwise the field will default to empty (`--moniker`defaults to the machine name).

You should put your name of your team in `details`. 

```
iriscli stake edit-validator --from= < name >  --moniker="choose a moniker"  --website="https://irisnet.org"  --details="team" --chain-id=fuxi-4000 
  --details="details"--node=tcp://localhost:26657 --fee=0.04iris  --gas=2000000
```
### View Validator Description

View the validator's information with this command:

```
iriscli stake validator <val-address> --chain-id=fuxi-4000
```

### Use IRISPlorer

You should also be able to see your validator on the [Explorer](https://testnet.irisplorer.io). You are looking for the `bech32` encoded `address` in the `~/.iris/config/priv_validator.json` file.

