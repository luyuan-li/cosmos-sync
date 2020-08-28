package distribution

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/models"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DocTxMsgSetWithdrawAddress struct {
	DelegatorAddress string `bson:"delegator_address"`
	WithdrawAddress  string `bson:"withdraw_address"`
}

func (doctx *DocTxMsgSetWithdrawAddress) GetType() string {
	return TxTypeSetWithdrawAddress
}

func (doctx *DocTxMsgSetWithdrawAddress) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgStakeSetWithdrawAddress)
	doctx.DelegatorAddress = msg.DelegatorAddress.String()
	doctx.WithdrawAddress = msg.WithdrawAddress.String()
}

func (m *DocTxMsgSetWithdrawAddress) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.DelegatorAddress, m.WithdrawAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}

// msg struct for delegation withdraw from a single validator
type DocTxMsgWithdrawDelegatorReward struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
}

func (doctx *DocTxMsgWithdrawDelegatorReward) GetType() string {
	return TxTypeWithdrawDelegatorReward
}

func (doctx *DocTxMsgWithdrawDelegatorReward) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgWithdrawDelegatorReward)
	doctx.DelegatorAddress = msg.DelegatorAddress.String()
	doctx.ValidatorAddress = msg.ValidatorAddress.String()
}
func (m *DocTxMsgWithdrawDelegatorReward) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.DelegatorAddress, m.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}

// msg struct for delegation withdraw for all of the delegator's delegations
type DocTxMsgFundCommunityPool struct {
	Amount    []models.Coin `bson:"amount"`
	Depositor string        `bson:"depositor"`
}

func (doctx *DocTxMsgFundCommunityPool) GetType() string {
	return TxTypeMsgFundCommunityPool
}

func (doctx *DocTxMsgFundCommunityPool) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgFundCommunityPool)
	doctx.Depositor = msg.Depositor.String()
	doctx.Amount = models.BuildDocCoins(msg.Amount)
}
func (m *DocTxMsgFundCommunityPool) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
// msg struct for validator withdraw
type DocTxMsgWithdrawValidatorCommission struct {
	ValidatorAddress string `bson:"validator_address"`
}

func (doctx *DocTxMsgWithdrawValidatorCommission) GetType() string {
	return TxTypeMsgWithdrawValidatorCommission
}

func (doctx *DocTxMsgWithdrawValidatorCommission) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgWithdrawValidatorCommission)
	doctx.ValidatorAddress = msg.ValidatorAddress.String()
}

func (m *DocTxMsgWithdrawValidatorCommission) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}