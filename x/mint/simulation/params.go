package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

const (
	keyMaxMintableNanoseconds = "MaxMintableNanoseconds"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyMaxMintableNanoseconds,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%d\"", GenMaxMintableNanoseconds(r))
			},
		),
	}
}
