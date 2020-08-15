package bot

import (
	"github.com/jim380/Celory/cmd"
)

func (v *validator) getBalance() {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS")
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func (v *validatorRG) getBalance() {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_RG_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_RG_ADDRESS")
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func (v *validatorBf) getBalance() {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func (v *validatorGr) getBalance() {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS")
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func (v *validatorGrRG) getBalance() {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_RG_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_RG_ADDRESS")
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func (v *validatorGrBf) getBalance() {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	gold := cmd.ParseAmount(target, "gold")
	goldVal := isZero(gold, "goldVal")
	usd := cmd.ParseAmount(target, "usd")
	usdVal := isZero(usd, "usdVal")
	lockedGold := cmd.ParseAmount(target, "lockedGold")
	lockedGoldVal := isZero(lockedGold, "lockedGoldVal")
	nonVotingLockedGold := cmd.ParseAmount(target1, "nonVotingLockedGold")
	nonVotingLockedGoldVal := isZero(nonVotingLockedGold, "nonVotingLockedGoldVal")
	total := cmd.ParseAmount(target, "total")
	totalVal := isZero(total, "totalVal")
	v.balance.gold = goldVal
	v.balance.usd = usdVal
	v.balance.lockedGold = lockedGoldVal
	v.balance.nonVoting = nonVotingLockedGoldVal
	v.balance.total = totalVal
}

func updateBalance(acc accountManager) {
	acc.getBalance()
}
