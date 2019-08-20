package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cosmos/cosmos-sdk/x/slashing/internal/types"
)

// Simulation parameter constants
const (
	SignedBlocksWindow      = "signed_blocks_window"
	MinSignedPerWindow      = "min_signed_per_window"
	DowntimeJailDuration    = "downtime_jail_duration"
	SlashFractionDoubleSign = "slash_fraction_double_sign"
	SlashFractionDowntime   = "slash_fraction_downtime"
)

// GenSignedBlocksWindow randomized SignedBlocksWindow
func GenSignedBlocksWindow(r *rand.Rand) int64 {
	return int64(simulation.RandIntBetween(r, 10, 1000))
}

// GenMinSignedPerWindow randomized MinSignedPerWindow
func GenMinSignedPerWindow(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(10)), 1)
}

// GenDowntimeJailDuration randomized DowntimeJailDuration
func GenDowntimeJailDuration(r *rand.Rand) time.Duration {
	return time.Duration(simulation.RandIntBetween(r, 60, 60*60*24)) * time.Second
}

// GenSlashFractionDoubleSign randomized SlashFractionDoubleSign
func GenSlashFractionDoubleSign(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(1).Quo(sdk.NewDec(int64(r.Intn(50) + 1)))
}

// GenSlashFractionDowntime randomized SlashFractionDowntime
func GenSlashFractionDowntime(r *rand.Rand) sdk.Dec {
	return sdk.NewDec(1).Quo(sdk.NewDec(int64(r.Intn(200) + 1)))
}

// RandomizedGenState generates a random GenesisState for slashing
func RandomizedGenState(input *module.GeneratorInput) {

	var (
		signedBlocksWindow      int64
		minSignedPerWindow      sdk.Dec
		downtimeJailDuration    time.Duration
		slashFractionDoubleSign sdk.Dec
		slashFractionDowntime   sdk.Dec
	)

	input.AppParams.GetOrGenerate(input.Cdc, SignedBlocksWindow, &signedBlocksWindow, input.R,
		func(r *rand.Rand) { signedBlocksWindow = GenSignedBlocksWindow(input.R) })

	input.AppParams.GetOrGenerate(input.Cdc, MinSignedPerWindow, &minSignedPerWindow, input.R,
		func(r *rand.Rand) { minSignedPerWindow = GenMinSignedPerWindow(input.R) })

	input.AppParams.GetOrGenerate(input.Cdc, DowntimeJailDuration, &downtimeJailDuration, input.R,
		func(r *rand.Rand) { downtimeJailDuration = GenDowntimeJailDuration(input.R) })

	input.AppParams.GetOrGenerate(input.Cdc, SlashFractionDoubleSign, &slashFractionDoubleSign, input.R,
		func(r *rand.Rand) { slashFractionDoubleSign = GenSlashFractionDoubleSign(input.R) })

	input.AppParams.GetOrGenerate(input.Cdc, SlashFractionDowntime, &slashFractionDowntime, input.R,
		func(r *rand.Rand) { slashFractionDowntime = GenSlashFractionDowntime(input.R) })

	params := types.NewParams(input.UnbondTime, signedBlocksWindow, minSignedPerWindow,
		downtimeJailDuration, slashFractionDoubleSign, slashFractionDowntime)

	slashingGenesis := types.NewGenesisState(params, nil, nil)

	fmt.Printf("Selected randomly generated slashing parameters:\n%s\n", codec.MustMarshalJSONIndent(input.Cdc, slashingGenesis.Params))
	input.GenState[types.ModuleName] = input.Cdc.MustMarshalJSON(slashingGenesis)
}