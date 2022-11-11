package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyBlockReward = []byte("BlockReward")
	// TODO: Determine the default value
	DefaultBlockReward uint64 = 0
)

var (
	KeyEarnDenom = []byte("EarnDenom")
	// TODO: Determine the default value
	DefaultEarnDenom string = "earn_denom"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	blockReward uint64,
	earnDenom string,
) Params {
	return Params{
		BlockReward: blockReward,
		EarnDenom:   earnDenom,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultBlockReward,
		DefaultEarnDenom,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBlockReward, &p.BlockReward, validateBlockReward),
		paramtypes.NewParamSetPair(KeyEarnDenom, &p.EarnDenom, validateEarnDenom),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBlockReward(p.BlockReward); err != nil {
		return err
	}

	if err := validateEarnDenom(p.EarnDenom); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateBlockReward validates the BlockReward param
func validateBlockReward(v interface{}) error {
	blockReward, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = blockReward

	return nil
}

// validateEarnDenom validates the EarnDenom param
func validateEarnDenom(v interface{}) error {
	earnDenom, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = earnDenom

	return nil
}
