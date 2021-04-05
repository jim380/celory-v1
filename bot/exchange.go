package bot

import (
	"fmt"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jim380/Celory/cmd"
	"github.com/shopspring/decimal"
)

func (b *beneficiaries) exchanegUSDToGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, perct uint16) {
	msg.ParseMode = "Markdown"
	var valGrBf validatorGrBf
	if reflect.TypeOf(b) == reflect.TypeOf(valGrBf) {
		// validator group
		botSendMsg(bot, msg, boldText("Exchange USD to CELO from validator group was requested"))
		usdAvailable, _ := decimal.NewFromString(b.validatorGrBf.usd)
		if perct == 100 {
			msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
			botSendMsg(bot, msg, msg.Text)
		} else if perct == 50 {
			dividend, _ := decimal.NewFromString("2")
			usdAvailable = usdAvailable.DivRound(dividend, 18)
			msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
			botSendMsg(bot, msg, msg.Text)
		} else if perct == 25 {
			dividend, _ := decimal.NewFromString("4")
			usdAvailable = usdAvailable.DivRound(dividend, 18)
			msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
			botSendMsg(bot, msg, msg.Text)
		} else if perct == 75 {
			dividend, _ := decimal.NewFromString("1.333333")
			usdAvailable = usdAvailable.DivRound(dividend, 18)
			msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
			botSendMsg(bot, msg, msg.Text)
		}

	}
	// validator
	botSendMsg(bot, msg, boldText("Exchange USD to CELO from validator was requested"))
	usdAvailable, _ := decimal.NewFromString(b.validatorGrBf.usd)
	if perct == 100 {
		msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 50 {
		dividend, _ := decimal.NewFromString("2")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 25 {
		dividend, _ := decimal.NewFromString("4")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	} else if perct == 75 {
		dividend, _ := decimal.NewFromString("1.333333")
		usdAvailable = usdAvailable.DivRound(dividend, 18)
		msg.Text = b.validatorGrBf.exchangeDollars(bot, msg, usdAvailable)
		botSendMsg(bot, msg, msg.Text)
	}

}

func (v *validatorGrBf) exchangeDollars(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount decimal.Decimal) string {
	zeroValue, _ := decimal.NewFromString("0")
	if amount.Cmp(zeroValue) == 1 {
		toExchange := amount.String()
		botSendMsg(bot, msg, boldText("Exchanging "+toExchange+" usd from validator group"))
		output, _ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_GROUP_RELEASE_GOLD_BENEFICIARY_ADDRESS --value " + toExchange)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			return successText("Success")
		}
		return errText(fmt.Sprintf("%v", outputParsed))
	}
	return warnText("Don't bite more than you can chew! You only have " + amount.String() + " usd available")
}

func (v *validatorBf) exchangeDollars(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, amount decimal.Decimal) string {
	zeroValue, _ := decimal.NewFromString("0")
	if amount.Cmp(zeroValue) == 1 {
		toExchange := amount.String()
		botSendMsg(bot, msg, boldText("Exchanging "+toExchange+" usd from validator group"))
		output, _ := botExecCmdOut("celocli exchange:dollars --from $CELO_VALIDATOR_RELEASE_GOLD_BENEFICIARY_ADDRESS --value " + toExchange)
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			return successText("Success")
		}
		return errText(fmt.Sprintf("%v", outputParsed))
	}
	return warnText("Don't bite more than you can chew! You only have " + amount.String() + " usd available")
}

func getExchangeRate(msg tgbotapi.MessageConfig) (string, error) {
	output, err := botExecCmdOut("celocli exchange:show")
	if err != nil {
		return "", err
	}
	cGold := cmd.ParseCmdOutput(output, "string", "(\\d.*) CELO =>", 1)
	cUsd := cmd.ParseCmdOutput(output, "string", "=> (\\d.*) cUSD", 1)
	cGoldDecimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cGold))
	cUsdDecimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cUsd))
	goldToUsdRatio := cUsdDecimal.DivRound(cGoldDecimal, 18)

	cUsd2 := cmd.ParseCmdOutput(output, "string", "(\\d.*) CELO =>", 1)
	cGold2 := cmd.ParseCmdOutput(output, "string", "=> (\\d.*) CELO", 1)
	cUsd2Decimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cUsd2))
	cGold2Decimal, _ := decimal.NewFromString(fmt.Sprintf("%v", cGold2))
	usdToGoldRatio := cGold2Decimal.DivRound(cUsd2Decimal, 18)
	msgPiece1 := "1 CELO = " + goldToUsdRatio.String() + " cUSD\n"
	msgPiece2 := "1 cUSD = " + usdToGoldRatio.String() + " CELO"
	return msgPiece1 + msgPiece2, nil
}

func exchangeUSDToGoldRun(e exchange, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, perct uint16) {
	e.exchanegUSDToGold(bot, msg, perct)
}
