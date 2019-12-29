package codec

import (
	amino "github.com/tendermint/go-amino"
	"inside.omertex.com/txgen/poll"
	"inside.omertex.com/txgen/session"
	"inside.omertex.com/txgen/stdTx"
)

var (
	Codec = amino.NewCodec()
)

func RegisterCodec(cdc *amino.Codec) {
	Codec.RegisterConcrete(stdTx.Tx{}, "cosmos-sdk/StdTx", nil)
	cdc.RegisterConcrete(session.MsgAd{}, "melodia/Ad", nil)
	cdc.RegisterConcrete(session.MsgRegisterListeningSession{}, "melodia/ListeningSession", nil)
	cdc.RegisterConcrete(session.MsgIntegrationData{}, "melodia/IntegrationData", nil)
	cdc.RegisterConcrete(session.MsgDistributeRewards{}, "melodia/DistributeRewards", nil)
	Codec.RegisterConcrete(poll.MsgRegisterPoll{}, "melodia/RegisterPoll", nil)
	Codec.RegisterConcrete(poll.MsgPollSubmission{}, "melodia/SubmitPoll", nil)
	cdc.RegisterInterface((*stdTx.Msg)(nil), nil)
}
