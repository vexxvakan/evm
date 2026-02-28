# Bank Precompile

## Address

`0x0000000000000000000000000000000000001111`

## Description

The Mint precompile provides access to the Cosmos SDK `x/bank` module MintCoins functionality through an EVM-compatible interface.
This enables smart contracts to mint native tokens for accounts.

## Interface

### Methods

#### mint

```solidity
function mint(address to, string token, uint256 value) external returns (bool success)
```

Mints `value` amount of `token` to the address of `to`

**Parameters:**

- `to`: The account address to mint the token to.
- `token`: The token denom minted. (e.g. `uatom`).
- `value`: The amount of tokens to mint as base denomination (e.g. if token decimal is 6 then 1 ATOM would be `1000000`)

**Returns:**

- Bool indicating success of the mint operation

## Demo

There is a simple demo script located at ./precompiles/mint/demo.sh.
It will run an account balance query, then mint a token, then query if the mint succeeded.
The script requires a running local node. Easiest way is to run ./local_node.sh
