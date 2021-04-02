package cmd

import "errors"

func parseAmount(target []byte, asset string) (float64, error) {
	switch asset {
	case "CELO":
		output := ParseCmdOutput(target, "float", "CELO: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed output is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v CELO available to lock\n", result)
		return result, nil
	case "usd":
		output := ParseCmdOutput(target, "float", "usd: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed output is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v usd available to exchange\n", result)
		return result, nil
	case "lockedGold":
		output := ParseCmdOutput(target, "float", "lockedGold: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed output is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v lockedGold\n", result)
		return result, nil
	case "total":
		output := ParseCmdOutput(target, "float", "total: (\\d.\\d*.+)", 1)
		if output == nil {
			return 0, errors.New("parsed output is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v gold in total\n", result)
		return result, nil
	case "nonVotingLockedGold":
		output := ParseCmdOutput(target, "float", "nonvoting: (\\d.\\d*)", 1)
		if output == nil {
			return 0, errors.New("parsed output is nil")
		}
		result := output.(float64) / 1e18
		// fmt.Printf("\nYou have %v nonvoting lockedGold\n", result)
		return result, nil
	default:
		return 0, nil
	}
}
