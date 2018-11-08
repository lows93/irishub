# Changelog

## 0.6.1

*November 6th, 2018*

FEATURES:

- [iservice] Add the feature of iService binding. And need to get the feature by software upgrade process.

## 0.6.0

*November 1st, 2018*

- Use --def-chain-id flag to reference the blockchain defined of the iService
- Fix some bugs about iservice definition and record
- Add cli and lcd test for record module
- Update the user doc of iservice definition and record

## 0.6.0-rc0

*October 24th, 2018*

BREAKING CHANGES:

- [monitor] Use new executable binary in monitor

FEATURES:

- [record] Add the record module of the data certification on blockchain
- [iservice] Add the feature of iService definition
- [cli] Add the example description in the cli help
- [test] Add Cli/LCD/Sim auto-test

BUG FIXES:

- Fix software upgrade issue caused by tx fee
- Report Panic when building the lcd proof
- Fix bugs in converting validator power to byte array
- Fix panic bug in wrong account number


## 0.5.0-rc1

*October 11th, 2018*

FEATURES:

- Make all the gov and upgrade parameters can be configured in the genesis.json

BUG FIXES

- Add check for iavl proof and value before building multistore proof


## 0.5.0-rc0

*September 30th, 2018*

BREAKING CHANGES:

- [cointype] Introduce the cointype of iris:
  - 1 iris = 10^18 iris-atto
  - 1 iris-milli = 10^15 iris-atto
  - 1 iris-micro = 10^12 iris-atto
  - 1 iris-nano = 10^9 iris-atto
  - 1 iris-pico = 10^6 iris-atto
  - 1 iris-femto = 10^3 iris-atto

FEATURES:

- [tendermint] Upgrade to Tendermint v0.23.1-rc0
- [cosmos-sdk] Upgrade to cosmos-sdk v0.24.2
    - Move the previous irisnet changeset about cosmos-sdk into irishub
- [irisdebug] Add irisdebug tool
- [LCD/cli] Add the proof verification to the LCD and CLI
- [iparam] Support the modification of governance parameters of complex data type through governance, and the submission of modified proposals through json config files
- [software-upgrade] Software upgrade solutions of the irisnet


## 0.4.2

*September 22th, 2018*

BUG FIXES

- Fix consensus failure due to the double sign evidence be broadcasted before the genesis block

## 0.4.1

*September 12th, 2018*

BUG FIXES

- Missing to set validator intraTxCount in stake genesis init


## 0.4.0

*September 6th, 2018*

BREAKING CHANGES:

- [cosmos-sdk] Upgrade to cosmos-sdk v0.23.0
    - Change the address prefix format:
      - cosmosaccaddr --> faa
      - cosmosaccpub --> fap
      - cosmosvaladdr --> fva
      - cosmosvalpub --> fvp
    - Adjust the Route & rootMultiStore Commit for software upgrade
    - Must specify gas and fee in both command line and rest api
    - The fee should be iris token and the token amount should be no less than 2*(10^10)*gas

FEATURES:

- [tendermint] Upgrade to Tendermint v0.22.6
    - Store the pre-state to support the replay function
- [cosmos-sdk] Upgrade to cosmos-sdk v0.23.0
    - Add the paramProposal and softwareUpgradeProposal in gov module
    - Improve fee token mechanism to more reasonably deduct transaction fee and achieve more ability to defent DDOS attack.
    - Introduce the global parameter module

BUG FIXES

- Default account balance in genesis
- Fix iris version issue
- Fix the unit conflict issue in slashing
- Check the voting power when create validator
- Fix evidence amimo register issue


## 0.4.0-rc2

*Sep 5th, 2018*

BUG FIXES

- Fix evidence amimo register issue


## 0.4.0-rc1

*Aug 27th, 2018*

BUG FIXES

- Default account balance in genesis
- iris version issue
- Fix the unit conflict issue in slashing
- Check the voting power when create validator


## 0.3.0

*July 30th, 2018*

BREAKING CHANGES:

- [tendermint] Upgrade to Tendermint v0.22.2
    - Default ports changed from 466xx to 266xx
    - ED25519 addresses are the first 20-bytes of the SHA256 of the raw 32-byte pubkey (Instead of RIPEMD160)
- [cosmos-sdk] Upgrade to cosmos-sdk v0.22.0
- [monitor] Move `iriscli monitor` subcommand to `iris monitor`

FEATURES:

- [lcd] /tx/send is now the only endpoint for posing transaction to irishub; aminofied all transaction messages 
- [monitor] Improve the metrics for iris-monitor 

BUG FIXES

- [cli] solve the issue of iriscli stake sign-info

##

## 0.2.0

*July 19th, 2018*

BREAKING CHANGES:

- [tendermint] Upgrade to Tendermint v0.21.0
- [cosmos-sdk] Upgrade to cosmos-sdk v0.19.1-rc1

FEATURES:

- [lcd] code refactor

- [cli] improve sendingand querying the  transactions 

- [monitor]export data which is collected by Prometheus Server

  ​

##  