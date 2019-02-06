package common

import (
	"encoding/json"
	"fmt"
)

// Convenience struct for CLI output
type PrettyParams struct {
	CommunityTax        json.RawMessage `json:"community_tax"`
	ProposerReward      json.RawMessage `json:"proposer_reward"`
	SignerReward        json.RawMessage `json:"signer_reward"`
	WithdrawAddrEnabled json.RawMessage `json:"withdraw_addr_enabled"`
}

// Construct a new PrettyParams
func NewPrettyParams(communityTax json.RawMessage, proposerReward json.RawMessage, signerReward json.RawMessage, withdrawAddrEnabled json.RawMessage) PrettyParams {
	return PrettyParams{
		CommunityTax:        communityTax,
		ProposerReward:      proposerReward,
		SignerReward:        signerReward,
		WithdrawAddrEnabled: withdrawAddrEnabled,
	}
}

func (pp PrettyParams) String() string {
	return fmt.Sprintf(`Distribution Params:
  Community Tax:          %s
  Proposer Reward:        %s
  Signer Reward:          %s
  Withdraw Addr Enabled:  %s`, pp.CommunityTax,
		pp.ProposerReward, pp.SignerReward, pp.WithdrawAddrEnabled)

}