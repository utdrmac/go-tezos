[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/DefinitelyNotAGoat/go-tezos)
# A Tezos Go Library

Go Tezos is a GoLang driven library for your Tezos node. 

## Installation

Get goTezos 
```
go get github.com/DefinitelyNotAGoat/go-tezos
```

## Getting Started 
Go Tezos is split into multiple services underneath to help organize it's functionality and also makes the library easier to maintain. 

To understand how Go Tezos works, take a look at the GoTezos Structure: 
```
type GoTezos struct {
	Client    tzc.TezosClient
	Constants network.Constants
	Block     block.TezosBlockService
	Snapshot  snapshot.TezosSnapshotService
	Cycle     cycle.TezosCycleService
	Account   account.TezosAccountService
	Delegate  delegate.TezosDelegateService
	Network   network.TezosNetworkService
	Operation operations.TezosOperationsService
	Contract  contracts.TezosContractsService
	Node      node.TezosNodeService
}
```
You can see GoTezos is a wrapper for several services such as `block`,  `Snapshot`, `Cycle`, `Account`, `Delegate`, `Network`, `Operation`, `Node`, and `Contract`.

The below examples assume you have go-tezos imported as follows:
```
import (
	goTezos "github.com/DefinitelyNotAGoat/go-tezos"
)
```
### Initializing Go Tezos
```
gt, err := NewGoTezos("127.0.0.1:8732")
if err != nil {
	fmt.Printf("could not connect to network: %v", err)
}

```

## Blocks
When getting a block, the Block structure will be returned, which contains all the normal fields as hitting the Tezos RPC directly. 

### Getting A Block By Level or Hash
This function will get a block by a level (int) or hash (string).
```
block, err := gt.Block.Get(1000)
if err != nil {
	fmt.Println(err)
}
fmt.Println(block)

block, err = gt.Block.Get("BKp1oX19NAAXdj3vU82sbpUrn5hAy51YnVtF5A5StsJKPKQGg3R")
if err != nil {
	fmt.Println(err)
}
```

### Getting The Head Block
This function will get the current head block of the chain.
```
block, err := gt.Block.GetHead()
if err != nil {
	fmt.Println(err)
}
fmt.Println(block)
```

## Snapshots
### Getting A Snapshot At A Specific Cycle
This function will get a snapshot at a specific cycle (int).
```
snapshot, err := gt.Snapshot.Get(50)
if err != nil {
	fmt.Println(err)
}
fmt.Println(snapshot)
```

## Accounts
### Getting the Current Balance of An Account
This function will get the balance of an address (string)
```
balance, err := gt.Account.GetBalance("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r")
if err != nil {
	fmt.Println(err)
}
fmt.Println(balance)
```

### Getting the Balance of An Account a Specific Block
This function will get the balance of an address (string) at a specific block (int).
```
balance, err := gt.Account.GetBalanceAtBlock("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r", 100000) //Can also be hash
if err != nil {
	fmt.Println(err)
}
fmt.Println(balance)
```

### Getting the Balance of An Account At a Specific Snapshot
This function will get the balance of an address (string) at a specific snapshot (cycle (int)).
```
balance, err := gt.Account.GetBalanceAtSnapshot("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r", 30)
if err != nil {
	fmt.Println(err)
}
fmt.Println(balance)
```

### Creating A Wallet 
This function will create a new wallet with a mnemonic (string) and password (string).
```
mnemonic := "normal dash crumble neutral reflect parrot know stairs culture fault check whale flock dog scout"
passwd := "PYh8nXDQLB"
email := "vksbjweo.qsrgfvbw@tezos.example.org"

wallet, err := gt.Account.CreateWallet(mnemonic, email+password)
if err != nil {
	fmt.Println(err)
}
fmt.Println(wallet)
```

### Importing A Wallet 
This function will import a wallet at with pkh (string) pk (string) sk (string)
```
pkh := "tz1fYvVTsSQWkt63P5V8nMjW764cSTrKoQKK"
pk := "edpkvH3h91QHjKtuR45X9BJRWJJmK7s8rWxiEPnNXmHK67EJYZF75G"
sk := "edskSA4oADtx6DTT6eXdBc6Pv5MoVBGXUzy8bBryi6D96RQNQYcRfVEXd2nuE2ZZPxs4YLZeM7KazUULFT1SfMDNyKFCUgk6vR"

wallet, err := gt.Account.ImportWallet(pkh, pk, sk)
if err != nil {
	fmt.Println(err)
}

fmt.Println(wallet)

```

### Import Encrypted Wallet 
This function will import an encrypted wallet with a password (string) and sk (string)
```
sk := "edesk1fddn27MaLcQVEdZpAYiyGQNm6UjtWiBfNP2ZenTy3CFsoSVJgeHM9pP9cvLJ2r5Xp2quQ5mYexW1LRKee2"
passwd := "password"

wallet, err := gt.Account.ImportEncryptedWallet(passwd, sk)
if err != nil {
	fmt.Println(err)
}

fmt.Println(wallet)

```

## Delegate 

### Getting Delegations
This function will get all the current delegations for an address (string).
```
delegations, err := gt.Delegate.GetDelegations("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r")
if err != nil {
	fmt.Println(err)
}

fmt.Println(delegations)

```

### Getting Delegations At a Specific Cycle
This function will get all the current delegations for an address (string) at a specific cycle (int).
```
delegations, err := gt.Delegate.GetDelegationsAtCycle("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r", 30)
if err != nil {
	fmt.Println(err)
}

fmt.Println(delegations)

```

### Getting A Rewards Report 
This is a helper function that returns a report of a delegates rewards for a cycle and each delegations share. This function takes in a
delegate address (string), cycle (int), and fee (0.05).
```
report, err := gt.Delegate.GetReport("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r", 30, 0.05) // fee
if err != nil {
	fmt.Println(err)
}

fmt.Println(delegations)
```

### Getting Payments
This function is called by a delegate report and generates payments, this payment structure can be used for batch payments.
```
payments, err := report.GetPayments() 
if err != nil {
	fmt.Println(err)
}

fmt.Println(payments)

```

### Getting Rewards 
This function gets the rewards for a delegate (string) at a specific cycle (int).
```
rewards, err := gt.Delegate.GetRewards("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r", 30) 
if err != nil {
	fmt.Println(err)
}

fmt.Println(rewards)
```

### Getting a Delegate
This function gets general information about a delegate (string).
```
delegate, err := gt.Delegate.GetDelegate("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r") 
if err != nil {
	fmt.Println(err)
}

fmt.Println(delegate)
```

### Getting All Delegates
This function gets all delegates known to the network.  
```
delegates, err := gt.Delegate.GetAllDelegates() 
if err != nil {
	fmt.Println(err)
}

fmt.Println(delegates)

```

### Getting All Delegates At a Specific Hash
This function gets all delegates known to the network at a specific hash (string). 
```
delegates, err := gt.Delegate.GetAllDelegatesByHash("BKp1oX19NAAXdj3vU82sbpUrn5hAy51YnVtF5A5StsJKPKQGg3R") 
if err != nil {
	fmt.Println(err)
}

fmt.Println(delegates)

```

### Getting the Current Staking Balance For a Delegate 
This function will get the current staking balance for a delegate (string)
```
balance, err := gt.Delegate.GetStakingBalance("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r") 
if err != nil {
	fmt.Println(err)
}

fmt.Println(balance)
```

### Getting the Staking Balance At a Cycle 
This function will get the staking balance for a delegate (string) at a specific cycle (int). 
```
balance, err := gt.Delegate.GetStakingBalanceAtCycle("tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r", 30) 
if err != nil {
	fmt.Println(err)
}

fmt.Println(balance)
```

### Getting Baking Rights For a Cycle
This funciton gets the baking rights for the entire network at a specific cycle (int). 
```
rights, err := gt.Delegate.GetBakingRights(30) 
if err != nil {
	fmt.Println(err)
}

fmt.Println(rights)
```

### Getting Baking Rights For a Specific Delegate
This function will get the baking rights for a specific delegate (string) at a specific cycle (int), with priority (int). 
```
rights, err := gt.Delegate.GetBakingRightsForDelegate(30, "tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r", 1) // 1 is priority
if err != nil {
	fmt.Println(err)
}

fmt.Println(rights)
```

### Getting Endorsing Rights For a Cycle
This funciton gets the endorsing rights for the entire network at a specific cycle (int). 
```
rights, err := gt.Delegate.GetEndorsingRights(30) 
if err != nil {
	fmt.Println(err)
}

fmt.Println(rights)
```

### Getting Endorsing Rights For a Specific Delegate
This function will get the endorsing rights for a specific delegate (string) at a specific cycle (int). 
```
rights, err := gt.Delegate.GetEndorsingRightsForDelegate(30, "tz3UoffC7FG7zfpmvmjUmUeAaHvzdcUvAj6r") 
if err != nil {
	fmt.Println(err)
}

fmt.Println(rights)
```

## Operations
### Creating A Batch Payment
This function will create return signed operation strings of batch payments. The function takes in
payments (goTezos.Payments), wallet (goTezos.Wallet), network fee (int), gas limit (fee).

```
ops, err := gt.Operation.CreateBatchPayment(payments, wallet, networkFee, gasLimit) 
if err != nil {
	fmt.Println(err)
}

fmt.Println(ops)

```

### Injecting An Operation
This function will take in an operation (string) and inject it into the Tezos network. It will return the operation hash if successful.
```
hash, err := gt.Operation.InjectOperation(ops) 
if err != nil {
	fmt.Println(err)
}

fmt.Println(hash)
```

## Network
### Getting Network Constants
This function will get the network contstant. 
```
constants, err := gt.Network.GetConstants() 
if err != nil {
	fmt.Println(err)
}

fmt.Println(constants)
```

### Getting Network Version
This function will get the network versions. 
```
versions, err := gt.Network.GetVersions() 
if err != nil {
	fmt.Println(err)
}

fmt.Println(versions)
```