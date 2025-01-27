package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgResetStore = "reset_store"

var _ sdk.Msg = &MsgResetStore{}

func NewMsgResetStore(creator string) *MsgResetStore {
	return &MsgResetStore{
		Creator: creator,
	}
}

func (msg *MsgResetStore) Route() string {
	return RouterKey
}

func (msg *MsgResetStore) Type() string {
	return TypeMsgResetStore
}

func (msg *MsgResetStore) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgResetStore) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResetStore) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
