package types

import (
	"testing"

	"github.com/PavelTabacu/auction/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgInitiateAuction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgInitiateAuction
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgInitiateAuction{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgInitiateAuction{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
