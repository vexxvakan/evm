# Bank Precompile

## Address

`0x0000000000000000000000000000000000001111`

## Description

The Mint precompile provides access to the Cosmos SDK `x/mint` module mint functionality through an EVM-compatible interface.
This enables smart contracts to mint native tokens for accounts using either their `hex` or their `bech32` address.

## Interface

### Methods

#### mint

```solidity
function mint(address to, string token, uint256 value) external returns (bool success)
```

Retrieves all native token balances for the specified account.
Each balance includes the ERC-20 contract address and amount in the token's original precision.

**Parameters:**

- `to`: The account address to mint the token to.
- `token`: The token denom minted. (e.g. `uatom`).
- `value`: The amount of tokens to mint as base denomination (e.g. if token decimal is 6 then 1 ATOM would be `1000000`)

**Returns:**

- Bool indicating success of the mint operation

TODO: **Gas Cost:**

### Data Structures

```solidity
TODO: Spec out
```

## Implementation Details

TODO: Spec out
