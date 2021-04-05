package bot

import (
	"fmt"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jim380/Celory/cmd"
	"github.com/shopspring/decimal"
)

func (v *validatorRG) transfer(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"

	botSendMsg(bot, msg, boldText("Transferring dollars from validator releaseGold contract was requested"))
	usdAvailable, _ := decimal.NewFromString(v.usd)
	msg.Text = v.rgXferDollars(bot, msg, usdAvailable)
	botSendMsg(bot, msg, msg.Text)
}

func (v *validatorGrRG) transfer(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"

	botSendMsg(bot, msg, boldText("Transferring dollars from validator group releaseGold contract was requested"))
	usdAvailable, _ := decimal.NewFromString(v.usd)
	msg.Text = v.rgXferDollars(bot, msg, usdAvailable)
	botSendMsg(bot, msg, msg.Text)
}

func (b *beneficiaries) transfer(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"
	var valGrBf validatorGrBf
	if reflect.TypeOf(b) == reflect.TypeOf(valGrBf) {
		botSendMsg(bot, msg, boldText("Transferring gold from validator group releaseGold beneficiary address was requested"))
		goldAvailable, _ := decimal.NewFromString(b.validatorGrBf.usd)
		msg.Text = b.xferGold(bot, msg, goldAvailable)
		botSendMsg(bot, msg, msg.Text)
	}

	botSendMsg(bot, msg, boldText("Transferring gold from validator releaseGold beneficiary address was requested"))
	goldAvailable, _ := decimal.NewFromString(b.validatorBf.usd)
	msg.Text = b.xferGold(bot, msg, goldAvailable)
	botSendMsg(bot, msg, msg.Text)
}

func (v *validatorRG) rgXferDollars(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount decimal.Decimal) string {
	zeroValue, _ := decimal.NewFromString("0")
	if amount.Cmp(zeroValue) == 1 {
		toXfer := amount.String()
		botSendMsg(bot, msg, boldText("Transferring "+toXfer+" usd from validator releaseGold address"))
		output, _ := botExecCmdOut("celocli releasegold:transfer-dollars --contract $CELO_VALIDATOR_RG_ADDRESS --to $CELO_VALIDATOR_RELEASE_GOLD_BENEFICIARY_ADDRESS --value " + toXfer)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			return successText("Success")
		}
		return errText(fmt.Sprintf("%v", outputParsed))
	}
	return warnText("Don't bite more than you can chew! You only have " + amount.String() + " usd available")
}

func (b *beneficiaries) xferGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount decimal.Decimal) string {
	zeroValue, _ := decimal.NewFromString("0")
	if amount.Cmp(zeroValue) == 1 {
		toXfer := amount.String()
		var valGrBf validatorGrBf
		if reflect.TypeOf(b) == reflect.TypeOf(valGrBf) {
			botSendMsg(bot, msg, boldText("Transferring "+toXfer+" gold from validator group releaseGold beneficiary address"))
			output, _ := botExecCmdOut("celocli transfer:gold --from $CELO_VALIDATOR_GROUP_RELEASE_GOLD_BENEFICIARY_ADDRESS --to CELO_VALIDATOR_GROUP_RG_ADDRESS --value " + toXfer)
			outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
			if outputParsed == nil {
				return successText("Success")
			}
			return errText(fmt.Sprintf("%v", outputParsed))
		}

		botSendMsg(bot, msg, boldText("Transferring "+toXfer+" gold from validator releaseGold beneficiary address"))
		output, _ := botExecCmdOut("celocli transfer:gold --from $CELO_VALIDATOR_RELEASE_GOLD_BENEFICIARY_ADDRESS --to CELO_VALIDATOR_RG_ADDRESS --value " + toXfer)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			return successText("Success")
		}
		return errText(fmt.Sprintf("%v", outputParsed))
	}
	return warnText("Don't bite more than you can chew! You only have " + amount.String() + " gold available")
}

func (v *validatorGrRG) rgXferDollars(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount decimal.Decimal) string {
	zeroValue, _ := decimal.NewFromString("0")
	if amount.Cmp(zeroValue) == 1 {
		toXfer := amount.String()
		botSendMsg(bot, msg, boldText("Transferring "+toXfer+" usd from validator group releaseGold address"))
		output, _ := botExecCmdOut("celocli releasegold:transfer-dollars --contract $CELO_VALIDATOR_GROUP_RG_ADDRESS --to $CELO_VALIDATOR_GROUP_RELEASE_GOLD_BENEFICIARY_ADDRESS --value " + toXfer)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			return successText("Success")
		}
		return errText(fmt.Sprintf("%v", outputParsed))
	}
	return warnText("Don't bite more than you can chew! You only have " + amount.String() + " usd available")
}

func xfer(x transfer, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	x.transfer(bot, msg)
}
