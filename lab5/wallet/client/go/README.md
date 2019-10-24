# Native go client

This example was built using the [go-ethereum](https://github.com/ethereum/go-ethereum) API and shows how to build a server side native DApp in Go.

## Go Bindings

The only thing needed to generate a Go binding to an Ethereum contract is the contract's ABI definition JSON file.
The go-ethereum has a source code generator (i.e. abigen) that can convert Ethereum ABI definitions to type-safe Go packages.

To build the go bindings for the MyWallet contract, you only need to have the [abigen](https://github.com/ethereum/go-ethereum#executables) installed in your system and run the following command:

```
abigen --sol ../../contracts/MyWallet.sol --pkg contract --out ./go-bindings/mywallet-contract/mywallet.go
```

Is also possible to use the [go generate](https://blog.golang.org/generate) command to perform the action above, by just including the following line in your go code, the `main.go` in our case:

```go
//go:generate abigen --sol ../../contracts/MyWallet.sol --pkg contract --out ./go-bindings/mywallet-contract/mywallet.go
```

And running the `generate` command:
```
go generate
```

But for the MyWallet example we provide a Makefile that simplify those operations.

## Build the example

To compile the code and generate the necessary bindings, just run:
```
make
```

The command above will create the `go-bindings/my-wallet-contract` directory and generate the `mywallet.go` binding on it, containing the code necessary to interact with the contract. 

## Running the example

Just execute the generate binary:
```
./wallet
```

## Making use of external libraries (optional)

In case that your contract makes use of an external solidity library, like [openzeppelin](https://github.com/OpenZeppelin/openzeppelin-contracts) you need to inform the solidity compiler (i.e. solc) about the new dependency to be compiled.
This can be done using the `--allow-paths` parameter.
And as stated [here](https://github.com/ethereum/go-ethereum/pull/16683) abigen doesn't have a way to include any imported path, thus to be able to tell to the abigen where the dependencies are, you need to generate a `combined.json` output from the solc, which will contain all the necessary information to the abigen to generate the correct binding and avoid problems of dependencies not found as shown below:

```
../../contracts/MyWallet.sol:3:1: Error: Source "../../node_modules/@openzeppelin/contractsmath/SafeMath.sol" not found: File not found.
import "@openzeppelin/contracts/math/SafeMath.sol";
```

You can take a look on the [Makefile](https://github.com/relab/blockchain_course/blob/1bec598be457aecb74431fc3b13888ab6a278bdc/blockchain/solutions/lab5/client/go/Makefile#L51) to see how this is done. But to compile the MyWallet contract using SafeMath you just need to:

1. Replace the following line in `main.go`:
```go
//go:generate abigen --sol ../../contracts/MyWallet.sol --pkg contract --out ./go-bindings/mywallet-contract/mywallet.go
```

by the line:
```go
//go:generate abigen --combined-json ../../build/combined.json --pkg contract --out ./go-bindings/mywallet-contract/mywallet.go
```

2. Use the `solc` to compile the contracts and create the `combined.json` output.

```
make solc
```

3. Compile the project
```
make
```

By default the make instruction will run the `make generate` and the  `make build` commands, which will generate the bindings and create the binary of the application.

*Note:*  Make sure that your contract is not importing the SafeMath library from the local path, by replacing the import statement in the `MyWallet.sol`:

From:
```javascript
import "./SafeMath.sol";
```

To:
```javascript
import "@openzeppelin/contracts/math/SafeMath.sol";
```