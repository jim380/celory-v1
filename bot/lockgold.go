package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jim380/Celory/cmd"
	"github.com/shopspring/decimal"
)

func (v *validatorGr) lockGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Locking all gold from validator group"))
	goldAvailable, _ := decimal.NewFromString(v.gold)
	reserve, _ := decimal.NewFromString("1000000000000000000")
	goldAvailable = goldAvailable.Sub(reserve)
	zeroValue, _ := decimal.NewFromString("0")
	if goldAvailable.Cmp(zeroValue) == -1 {
		msg.Text = warnText("Not enough gold to set aside 1 gold for fees. Must have at least 1 gold reserved for paying gas.")
	} else {
		botSendMsg(bot, msg, boldText("Locking "+goldAvailable.String()+" gold from validator group"))
		output, _ := botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_GROUP_ADDRESS --value " + goldAvailable.String())
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
			botSendMsg(bot, msg, msg.Text)
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
			botSendMsg(bot, msg, msg.Text)
		}
	}
}

func (v *validatorGrRG) lockGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Locking all gold from validator group releaseGold contract"))
	goldAvailable, _ := decimal.NewFromString(v.gold)
	reserve, _ := decimal.NewFromString("1000000000000000000")
	goldAvailable = goldAvailable.Sub(reserve)
	zeroValue, _ := decimal.NewFromString("0")
	if goldAvailable.Cmp(zeroValue) == -1 {
		msg.Text = warnText("Not enough gold to set aside 1 gold for fees. Must have at least 1 gold reserved for paying gas.")
	} else {
		botSendMsg(bot, msg, boldText("Locking "+goldAvailable.String()+" gold from validator group releaseGold contract"))
		output, _ := botExecCmdOut("celocli releasegold:locked-gold --contract CELO_VALIDATOR_GROUP_RG_ADDRESS --action lock --value " + goldAvailable.String())
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
			botSendMsg(bot, msg, msg.Text)
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
			botSendMsg(bot, msg, msg.Text)
		}
	}
}

func (v *validator) lockGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Locking all gold from validator"))
	goldAvailable, _ := decimal.NewFromString(v.gold)
	reserve, _ := decimal.NewFromString("1000000000000000000")
	goldAvailable = goldAvailable.Sub(reserve)
	zeroValue, _ := decimal.NewFromString("0")
	if goldAvailable.Cmp(zeroValue) == -1 {
		msg.Text = warnText("Not enough gold to set aside 1 gold for fees. Must have at least 1 gold reserved for paying gas.")
	} else {
		botSendMsg(bot, msg, boldText("Locking "+goldAvailable.String()+" gold from validator"))
		output, _ := botExecCmdOut("celocli lockedgold:lock --from $CELO_VALIDATOR_ADDRESS --value " + goldAvailable.String())
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
			botSendMsg(bot, msg, msg.Text)
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
			botSendMsg(bot, msg, msg.Text)
		}
	}
}

func (v *validatorRG) lockGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	msg.ParseMode = "Markdown"
	botSendMsg(bot, msg, boldText("Locking all gold from validator releaseGold contract"))
	goldAvailable, _ := decimal.NewFromString(v.gold)
	reserve, _ := decimal.NewFromString("1000000000000000000")
	goldAvailable = goldAvailable.Sub(reserve)
	zeroValue, _ := decimal.NewFromString("0")
	if goldAvailable.Cmp(zeroValue) == -1 {
		msg.Text = warnText("Not enough gold to set aside 1 gold for fees. Must have at least 1 gold reserved for paying gas.")
	} else {
		botSendMsg(bot, msg, boldText("Locking "+goldAvailable.String()+" gold from validator releaseGold contract"))
		output, _ := botExecCmdOut("celocli releasegold:locked-gold --contract CELO_VALIDATOR_RG_ADDRESS --action lock --value " + goldAvailable.String())
		outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
		if outputParsed == nil {
			msg.Text = successText("Success")
			botSendMsg(bot, msg, msg.Text)
		} else {
			msg.Text = errText(fmt.Sprintf("%v", outputParsed))
			botSendMsg(bot, msg, msg.Text)
		}
	}
}

func lockGoldRun(g goldManager, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	g.lockGold(bot, msg)
}
