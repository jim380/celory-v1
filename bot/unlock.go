package bot

import (
	//"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jim380/Celory/cmd"
)

func (v *validator) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("celocli account:unlock $CELO_VALIDATOR_ADDRESS --password=$PASSWORD", msg)
	// if err != nil {
	// 	msg.Text = err.Error()
	// 	botSendMsg(bot, msg, errText(err.Error()))
	// }
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		// botSendMsg(bot, msg, successText("Success"))
		v.unlocked = true
		return
	}
	// botSendMsg(bot, msg, errText(fmt.Sprintf("%v", outputParsed)))
	v.unlocked = false
}

func (v *validatorGr) unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	output, _ := botExecCmdOut("celocli account:unlock $CELO_VALIDATOR_GROUP_ADDRESS --password=$PASSWORD", msg)
	outputParsed := cmd.ParseCmdOutput(output, "string", "Error: Returned (.*)", 1)
	if outputParsed == nil {
		// botSendMsg(bot, msg, successText("Success"))
		v.unlocked = true
		return
	}
	// botSendMsg(bot, msg, errText(fmt.Sprintf("%v", outputParsed)))
	v.unlocked = false
}

func Unlock(acc accountUnlocker, bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	acc.unlockAccount(bot, msg)
}
