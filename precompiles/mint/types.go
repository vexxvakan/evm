package mint

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/core/address"
	sdkerrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cmn "github.com/cosmos/evm/precompiles/common"
)

const (
	MintPrecompileName = "mint_precompile"
	// MintPrecompileAddress defines the hex address of the smart contract for the x/mint precompile
	MintPrecompileAddress = "0x0000000000000000000000000000000000001111"
	// MintMethod defines the ABI method name for the Mint method
	MintMethod = "mint"
)

// EventMint defines the event data structure for the Mint events
type EventMint struct {
	To    common.Address
	Token string
	Value *big.Int
}

// ValidateMint validates a mint request and constructs a correctly formatted (sdk.AccAddress, sdk.Coins) tuple
// args: [to cmn.Address, token string, value *big.Int]
func ValidateMint(ctx context.Context, addrCdc address.Codec, bk cmn.BankKeeper, args []interface{}) (to common.Address, addr sdk.AccAddress, coins sdk.Coins, err error) {
	if len(args) != 3 {
		return common.Address{}, nil, nil, fmt.Errorf("invalid number of arguments; expected 3; got: %d", len(args))
	}

	to, ok := args[0].(common.Address)
	if !ok {
		return common.Address{}, nil, nil, sdkerrors.Wrapf(ErrInvalidReceiver, "%s is invalid or empty", args[0])
	}
	toStr, err := addrCdc.BytesToString(to.Bytes())
	if err != nil {
		return common.Address{}, nil, nil, sdkerrors.Wrapf(ErrInvalidReceiver, "failed to decode address: %s", to)
	}
	addr, err = sdk.AccAddressFromBech32(toStr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	token, ok := args[1].(string)
	if !ok || len(token) == 0 {
		return common.Address{}, nil, nil, sdkerrors.Wrapf(ErrInvalidDenom, "%s is invalid or empty", args[1])
	}
	ok = bk.HasDenomMetaData(ctx, token)
	if !ok {
		return common.Address{}, nil, nil, sdkerrors.Wrapf(ErrInvalidDenom, "%s is not registered in the bank module", token)
	}

	value, ok := args[2].(*big.Int)
	if !ok {
		return common.Address{}, nil, nil, sdkerrors.Wrapf(ErrInvalidAmount, "%s is not a valid *big.Int", args[2])
	}
	ok = value.Sign() > 0
	if !ok {
		return common.Address{}, nil, nil, sdkerrors.Wrapf(ErrInvalidAmount, "%s must be greater than 0", value)
	}

	tmpCoins := []cmn.Coin{{Denom: token, Amount: value}}
	coins, err = cmn.NewSdkCoinsFromCoins(tmpCoins)
	if err != nil {
		return common.Address{}, nil, nil, sdkerrors.Wrap(ErrInvalidAmount, "failed to convert coins to sdk coin format")
	}

	return to, addr, coins, nil
}
