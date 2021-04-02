package bot

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jim380/Celory/cmd"
)

func (v *validator) getBalance() error {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_ADDRESS")
	bal, err := parseBalances(target, target1)
	if err != nil {
		return err
	}
	v.balance.gold = bal.gold
	v.balance.usd = bal.usd
	v.balance.lockedGold = bal.lockedGold
	v.balance.nonVoting = bal.nonVoting
	v.balance.total = bal.total
	return nil
}

func (v *validatorRG) getBalance() error {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_RG_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_RG_ADDRESS")
	bal, err := parseBalances(target, target1)
	fmt.Println("BALANCE IS: ", bal)
	if err != nil {
		return err
	}
	v.balance.gold = bal.gold
	v.balance.usd = bal.usd
	v.balance.lockedGold = bal.lockedGold
	v.balance.nonVoting = bal.nonVoting
	v.balance.total = bal.total
	return nil
}

func (v *validatorBf) getBalance() error {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	bal, err := parseBalances(target, target1)
	if err != nil {
		return err
	}
	v.balance.gold = bal.gold
	v.balance.usd = bal.usd
	v.balance.lockedGold = bal.lockedGold
	v.balance.nonVoting = bal.nonVoting
	v.balance.total = bal.total
	return nil
}

func (v *validatorGr) getBalance() error {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_ADDRESS")
	bal, err := parseBalances(target, target1)
	if err != nil {
		return err
	}
	v.balance.gold = bal.gold
	v.balance.usd = bal.usd
	v.balance.lockedGold = bal.lockedGold
	v.balance.nonVoting = bal.nonVoting
	v.balance.total = bal.total
	return nil
}

func (v *validatorGrRG) getBalance() error {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_RG_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_RG_ADDRESS")
	bal, err := parseBalances(target, target1)
	if err != nil {
		return err
	}
	v.balance.gold = bal.gold
	v.balance.usd = bal.usd
	v.balance.lockedGold = bal.lockedGold
	v.balance.nonVoting = bal.nonVoting
	v.balance.total = bal.total
	return nil
}

func (v *validatorGrBf) getBalance() error {
	target, _ := botExecCmdOut("celocli account:balance $CELO_VALIDATOR_GROUP_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	target1, _ := botExecCmdOut("celocli lockedgold:show $CELO_VALIDATOR_GROUP_RELEASE_GOLD_BENEFICIARY_ADDRESS")
	bal, err := parseBalances(target, target1)
	if err != nil {
		return err
	}
	v.balance.gold = bal.gold
	v.balance.usd = bal.usd
	v.balance.lockedGold = bal.lockedGold
	v.balance.nonVoting = bal.nonVoting
	v.balance.total = bal.total
	return nil
}

func parseBalances(acc, lockedCELO []byte) (balance, error) {
	gold, err := parseAmount(acc, "gold")
	if err != nil {
		return balance{}, err
	}
	goldVal := toString(gold)
	usd, err := parseAmount(acc, "usd")
	if err != nil {
		return balance{}, err
	}
	usdVal := toString(usd)
	lockedGold, err := parseAmount(acc, "lockedGold")
	if err != nil {
		return balance{}, err
	}
	lockedGoldVal := toString(lockedGold)
	nonVotingLockedGold, err := parseAmount(lockedCELO, "nonVotingLockedGold")
	if err != nil {
		return balance{}, err
	}
	nonVotingLockedGoldVal := toString(nonVotingLockedGold)
	total, err := parseAmount(lockedCELO, "total")
	if err != nil {
		return balance{}, err
	}
	totalVal := toString(total)
	return balance{
		gold:       goldVal,
		usd:        usdVal,
		lockedGold: lockedGoldVal,
		nonVoting:  nonVotingLockedGoldVal,
		total:      totalVal,
	}, nil
}

func parseAmount(target []byte, asset string) (float64, error) {
	switch asset {
	case "CELO":
		output := cmd.ParseCmdOutput(target, "float", "CELO: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed CELO is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v CELO available to lock\n", result)
		return result, nil
	case "usd":
		output := cmd.ParseCmdOutput(target, "float", "cUSD: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed cUSD is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v usd available to exchange\n", result)
		return result, nil
	case "lockedGold":
		output := cmd.ParseCmdOutput(target, "float", "lockedCELO: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed lockedCELO is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v lockedGold\n", result)
		return result, nil
	case "total":
		output := cmd.ParseCmdOutput(target, "string", "total: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed Total is nil")
		}
		// result := output.(float64) / 1e18
		result, _ := strconv.ParseFloat(fmt.Sprintf("%v", output), 64)
		// fmt.Printf("\nYou have %v gold in total\n", result)
		return result, nil
	case "nonVotingLockedGold":
		output := cmd.ParseCmdOutput(target, "string", "nonvoting: (\\d.\\d*)", 1)
		if output == nil {
			return 0, errors.New("parsed nonVotingCELO is nil")
		}
		// result := output.(float64) / 1e18
		result, _ := strconv.ParseFloat(fmt.Sprintf("%v", output), 64)
		// fmt.Printf("\nYou have %v nonvoting lockedGold\n", result)
		return result, nil
	default:
		return 0, nil
	}
}

func updateBalance(acc accountManager) error {
	return acc.getBalance()
}
