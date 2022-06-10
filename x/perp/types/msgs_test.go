package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/testutil/sample"
)

func TestMsgOpenPosition_ValidateBasic(t *testing.T) {
	type test struct {
		msg     *MsgOpenPosition
		wantErr bool
	}

	cases := map[string]test{
		"ok": {
			msg: &MsgOpenPosition{
				Sender:               sample.AccAddress().String(),
				TokenPair:            "NIBI:USDN",
				Side:                 Side_BUY,
				QuoteAssetAmount:     sdk.NewInt(100),
				Leverage:             sdk.NewDec(10),
				BaseAssetAmountLimit: sdk.NewInt(100),
			},
			wantErr: false,
		},

		"invalid side": {
			msg: &MsgOpenPosition{
				Sender:               sample.AccAddress().String(),
				TokenPair:            "NIBI:USDN",
				Side:                 3,
				QuoteAssetAmount:     sdk.NewInt(100),
				Leverage:             sdk.NewDec(10),
				BaseAssetAmountLimit: sdk.NewInt(100),
			},
			wantErr: true,
		},
		"invalid side 2": {
			msg: &MsgOpenPosition{
				Sender:               sample.AccAddress().String(),
				TokenPair:            "NIBI:USDN",
				Side:                 Side_SIDE_UNSPECIFIED,
				QuoteAssetAmount:     sdk.NewInt(100),
				Leverage:             sdk.NewDec(10),
				BaseAssetAmountLimit: sdk.NewInt(100),
			},
			wantErr: true,
		},
		"invalid address": {
			msg: &MsgOpenPosition{
				Sender:               "",
				TokenPair:            "NIBI:USDN",
				Side:                 Side_SELL,
				QuoteAssetAmount:     sdk.NewInt(100),
				Leverage:             sdk.NewDec(10),
				BaseAssetAmountLimit: sdk.NewInt(100),
			},
			wantErr: true,
		},
		"invalid leverage": {
			msg: &MsgOpenPosition{
				Sender:               sample.AccAddress().String(),
				TokenPair:            "NIBI:USDN",
				Side:                 Side_BUY,
				QuoteAssetAmount:     sdk.NewInt(100),
				Leverage:             sdk.ZeroDec(),
				BaseAssetAmountLimit: sdk.NewInt(100),
			},
			wantErr: true,
		},
		"invalid quote asset amount": {
			msg: &MsgOpenPosition{
				Sender:               sample.AccAddress().String(),
				TokenPair:            "NIBI:USDN",
				Side:                 Side_BUY,
				QuoteAssetAmount:     sdk.NewInt(0),
				Leverage:             sdk.NewDec(10),
				BaseAssetAmountLimit: sdk.NewInt(100),
			},
			wantErr: true,
		},
		"invalid token pair": {
			msg: &MsgOpenPosition{
				Sender:               sample.AccAddress().String(),
				TokenPair:            "NIBI-USDN",
				Side:                 Side_BUY,
				QuoteAssetAmount:     sdk.NewInt(0),
				Leverage:             sdk.NewDec(10),
				BaseAssetAmountLimit: sdk.NewInt(100),
			},
			wantErr: true,
		},
		"invalid base asset amount limit": {
			msg: &MsgOpenPosition{
				Sender:               sample.AccAddress().String(),
				TokenPair:            "NIBI:USDN",
				Side:                 Side_BUY,
				QuoteAssetAmount:     sdk.NewInt(0),
				Leverage:             sdk.NewDec(10),
				BaseAssetAmountLimit: sdk.ZeroInt(),
			},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if err != nil && tc.wantErr == false {
				t.Fatalf("unexpected error: %s", err)
			}
			if err == nil && tc.wantErr == true {
				t.Fatalf("expected error: %s", err)
			}
		})
	}
}

func TestMsgLiquidate_ValidateBasic(t *testing.T) {
	type test struct {
		msg     *MsgLiquidate
		wantErr bool
	}

	cases := map[string]test{
		"ok": {
			msg: &MsgLiquidate{
				Sender:    sample.AccAddress().String(),
				TokenPair: "NIBI:USDN",
				Trader:    sample.AccAddress().String(),
			},
			wantErr: false,
		},
		"invalid pair": {
			msg: &MsgLiquidate{
				Sender:    sample.AccAddress().String(),
				TokenPair: "xxx:yyy:zzz",
				Trader:    sample.AccAddress().String(),
			},
			wantErr: true,
		},
		"invalid trader": {
			msg: &MsgLiquidate{
				Sender:    sample.AccAddress().String(),
				TokenPair: "NIBI:USDN",
				Trader:    "",
			},
			wantErr: true,
		},
		"invalid liquidator": {
			msg: &MsgLiquidate{
				Sender:    "",
				TokenPair: "NIBI:USDN",
				Trader:    sample.AccAddress().String(),
			},
			wantErr: true,
		},
	}

	for name, tc := range cases {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if err != nil && tc.wantErr == false {
				t.Fatalf("unexpected error: %s", err)
			}
			if err == nil && tc.wantErr == true {
				t.Fatalf("expected error: %s", err)
			}
		})
	}
}