package mint

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	cmn "github.com/cosmos/evm/precompiles/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// EventTypeMint defines the event type for the mint MintMethod transaction
	EventTypeMint = "Mint"
)

// EmitMintEvent creates a new event emitted on a Mint transaction.
func (p Precompile) EmitMintEvent(ctx sdk.Context, stateDB vm.StateDB, to common.Address, token string, value *big.Int) error {
	// Prepare the event topics
	event := p.Events[EventTypeMint]
	topics := make([]common.Hash, 3)

	// The first topic is always the signature of the event
	topics[0] = event.ID

	var err error
	topics[1], err = cmn.MakeTopic(to)
	if err != nil {
		return err
	}

	topics[2], err = cmn.MakeTopic(token)
	if err != nil {
		return err
	}

	arguments := abi.Arguments{event.Inputs[2]}
	packed, err := arguments.Pack(value)
	if err != nil {
		return err
	}

	stateDB.AddLog(&ethtypes.Log{
		Address:     p.Address(),
		Topics:      topics,
		Data:        packed,
		BlockNumber: uint64(ctx.BlockHeight()), //nolint:gosec // G115
	})

	return nil
}
