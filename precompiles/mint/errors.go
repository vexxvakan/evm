package mint

import errorsmod "cosmossdk.io/errors"

var (
	// ErrInvalidReceiver is raised when the receiver address is not valid
	ErrInvalidReceiver = errorsmod.Register(MintPrecompileName, 1, "invalid address: ")
	// ErrInvalidDenom is raised when a denom is not valid or is not registered in the bank module
	ErrInvalidDenom = errorsmod.Register(MintPrecompileName, 2, "invalid denom: ")
	// ErrInvalidAmount is raised when an amount value is not valid
	ErrInvalidAmount = errorsmod.Register(MintPrecompileName, 3, "invalid denom: ")
)
