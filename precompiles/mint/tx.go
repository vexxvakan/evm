package mint

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Mint defines a method to mint native tokens using the x/mint module.
func (p *Precompile) Mint(
	ctx sdk.Context,
	contract *vm.Contract,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	to, toAddr, coins, err := ValidateMint(ctx, p.addrCdc, p.bankKeeper, args)
	if err != nil {
		return nil, err
	}

	err = p.bankKeeper.MintCoins(ctx, banktypes.ModuleName, coins)
	if err != nil {
		return nil, err
	}

	err = p.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		banktypes.ModuleName,
		toAddr,
		coins,
	)
	if err != nil {
		return nil, err
	}

	if err = p.EmitMintEvent(ctx, stateDB, to, coins); err != nil {
		return nil, err
	}

	// Returning `true` here because our abi.json specifies "internalType": "bool" as only output
	return method.Outputs.Pack(true)
}
