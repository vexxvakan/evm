package mint

import (
	sdkerrors "cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	cmn "github.com/cosmos/evm/precompiles/common"
)

// Mint defines a method to mint native tokens using the x/mint module.
func (p *Precompile) Mint(
	ctx sdk.Context,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	to, toAddr, token, value, err := ValidateMint(ctx, p.addrCdc, p.bankKeeper, args)
	if err != nil {
		return nil, err
	}

	tmpCoins := []cmn.Coin{{Denom: token, Amount: value}}
	coins, err := cmn.NewSdkCoinsFromCoins(tmpCoins)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInvalidAmount, "failed to convert coins to sdk coin format")
	}

	err = p.bankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = p.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		minttypes.ModuleName,
		toAddr,
		coins,
	)
	if err != nil {
		return nil, err
	}

	if err = p.EmitMintEvent(ctx, stateDB, to, token, value); err != nil {
		return nil, err
	}

	// Returning `true` here because our abi.json specifies "internalType": "bool" as only output
	return method.Outputs.Pack(true)
}
