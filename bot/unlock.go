package bot

import (
	//"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jim380/Celory/cmd"
)

func (v *validator) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("celocli account:unlock $CELO_VALIDATOR_ADDRESS --password=$PASSWORD")
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		v.unlocked = true
		return
	}
	v.unlocked = false
}

func (v *validatorGr) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("celocli account:unlock $CELO_VALIDATOR_GROUP_ADDRESS --password=$PASSWORD")
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		v.unlocked = true
		return
	}
	v.unlocked = false
}

func unlock(acc accountUnlocker, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	acc.unlockAccount(bot, msg)
}
