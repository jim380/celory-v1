package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jim380/Celory/cmd"
)

func (v *validator) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("echo $PASSWORD|celocli account:unlock $CELO_VALIDATOR_ADDRESS", msg)
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		botSendMsg(bot, msg, successText("Success"))
	}
	botSendMsg(bot, msg, errText(fmt.Sprintf("%v", outputParsed)))
}

func (v *validatorGr) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("echo $PASSWORD|celocli account:unlock $CELO_VALIDATOR_GROUP_ADDRESS", msg)
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		botSendMsg(bot, msg, successText("Success"))
	}
	botSendMsg(bot, msg, errText(fmt.Sprintf("%v", outputParsed)))
}

func (v *validatorRG) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("echo $PASSWORD|celocli account:unlock $CELO_VALIDATOR_RG_ADDRESS", msg)
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		botSendMsg(bot, msg, successText("Success"))
	}
	botSendMsg(bot, msg, errText(fmt.Sprintf("%v", outputParsed)))
}

func (v *validatorGrRG) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("echo $PASSWORD|celocli account:unlock $CELO_VALIDATOR_GROUP_RG_ADDRESS", msg)
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		botSendMsg(bot, msg, successText("Success"))
	}
	botSendMsg(bot, msg, errText(fmt.Sprintf("%v", outputParsed)))
}

func Unlock(acc accountManager, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	acc.unlockAccount(bot, msg)
}
