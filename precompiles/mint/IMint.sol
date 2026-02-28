/// SPDX-License-Identifier: LGPL-3.0-only
pragma solidity >=0.8.17;

import "../common/Types.sol";

/// @dev The IMint contract's address.
address constant MINT_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000001111;

/// @dev The IMint contract's instance.
IMint constant MINT_CONTRACT = IMint(MINT_PRECOMPILE_ADDRESS);

/// @author Marius "Vexx" Modlich
/// @title Mint Precompile Contract
/// @dev The interface through which solidity contracts will interact with the native x/mint module
interface IMint {
    /// @dev Mint defines an Event, emitted when the authority address successfully mints a token
    /// @param to the address receiving the minted tokens
    /// @param token the native token denom minted
    /// @param value the token amount in base denom minted
    event Mint(address indexed to, string token, uint256 value);

    /// TRANSACTIONS
    /// @dev mint defines a method to mint native tokens to a receiver address.
    /// @param to the address to mint the tokens to
    /// @param token the native token denom to mint
    /// @param value the token amount in base denom
    /// @return success Whether the transaction was successful or not
    function mint(
        address to,
        string memory token,
        uint64 value
    ) external returns (bool success);
}
