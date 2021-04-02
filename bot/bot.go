package bot

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jim380/Celory/cmd"
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/synced"),
		tgbotapi.NewKeyboardButton("/balance"),
		tgbotapi.NewKeyboardButton("/status"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/score"),
		tgbotapi.NewKeyboardButton("/signing"),
		tgbotapi.NewKeyboardButton("/exchange_rate"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/lockgold"),
		tgbotapi.NewKeyboardButton("/exchange"),
		tgbotapi.NewKeyboardButton("/vote"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/unlock"),
		tgbotapi.NewKeyboardButton("/reinvest"),
		tgbotapi.NewKeyboardButton("/close"),
	),
)

var lockGoldKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrLockGold"),
		tgbotapi.NewInlineKeyboardButtonData("Validator All", "valLockGold"),
	),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Amount", "valAmount"),
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Group Amount", "valGrAmount"),
	// ),
)

var transferKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("USD - Validator Group RG to Beneficiary", "valGrRGTransfer"),
		tgbotapi.NewInlineKeyboardButtonData("USD - Validator RG to Beneficiary", "valRGTransfer"),
		tgbotapi.NewInlineKeyboardButtonData("Celo - Validator Group Beneficiary to RG", "valGrBfTransfer"),
		tgbotapi.NewInlineKeyboardButtonData("Celo - Validator Beneficiary to RG", "valBfTransfer"),
	),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Amount", "valAmount"),
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Group Amount", "valGrAmount"),
	// ),
)

var unlockKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group", "valGrUnlock"),
		tgbotapi.NewInlineKeyboardButtonData("Validator", "valUnlock"),
	),
	// tgbotapi.NewInlineKeyboardRow(
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator", "valUnlock"),
	// 	tgbotapi.NewInlineKeyboardButtonData("Validator Voter", "valVoteUnlock"),
	// ),
)

var exchangeUsdKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group 25%", "valGrOneForthUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator Group 50%", "valGrHalfUsd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group 75%", "valGrFourThirdsUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrAllUsd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator 25%", "valOneFourthUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator 50%", "valHalfUsd"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator 75%", "valFourThirdsUsd"),
		tgbotapi.NewInlineKeyboardButtonData("Validator All", "valAllUsd"),
	),
)

var electionVoteKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Validator Group All", "valGrVote"),
		tgbotapi.NewInlineKeyboardButtonData("Validator All", "valVote"),
	),
)

// beneficiary accounts
type exchange interface {
	exchanegUSDToGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, perct uint16)
}

// beneficiary accounts
type transfer interface {
	transfer(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig)
}

// RG accounts
type goldManager interface {
	lockGold(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig)
}

// all accounts
type accountManager interface {
	getBalance() error
	// transfer
}

// account unlocker
type accountUnlocker interface {
	unlockAccount(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig)
}

// voter accounts
type voter interface {
	vote(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig)
}

type balance struct {
	gold       string
	usd        string
	lockedGold string
	nonVoting  string
	total      string
}

type validator struct {
	balance
	unlocked bool
}

type validatorGr struct {
	balance
	unlocked bool
}

type validatorRG struct {
	balance
	unlocked bool
}

type validatorGrRG struct {
	balance
	unlocked bool
}

type validatorBf struct {
	balance
	unlocked bool
}

type validatorGrBf struct {
	balance
	unlocked bool
}

// Run instantiates the bot
func Run() {
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// ignore any non-Message Updates
		// if update.Message == nil {
		// 	continue
		// }

		// ignore any non-command Messages
		// if !update.Message.IsCommand() {
		// 	continue
		// }

		if update.CallbackQuery != nil {
			// msg := update.CallbackQuery.Message
			data := update.CallbackQuery.Data
			chatID := update.CallbackQuery.Message.Chat.ID
			msg := tgbotapi.NewMessage(chatID, "")
			msg.ParseMode = "Markdown"
			switch data {
			case "valGrUnlock":
				var valGr validatorGr
				unlock(&valGr, bot, msg)
				if valGr.unlocked {
					msg.Text = successText("unlocked!")
				} else {
					msg.Text = errText("failed!")
				}
				// break
			case "valUnlock":
				var val validator
				unlock(&val, bot, msg)
				if val.unlocked {
					msg.Text = successText("unlocked!")
				} else {
					msg.Text = errText("failed!")
				}
				// break
			case "valGrRGTransfer":
				var valGrRG validatorGrRG
				xfer(&valGrRG, bot, msg)
				updateBalance(&valGrRG)
				msgPiece := `usd: ` + valGrRG.balance.usd
				msg.Text = boldText("Validator Group RG Balance After Tranferring") + "\n\n" + msgPiece
				// break
			case "valRGTransfer":
				var valRG validatorRG
				xfer(&valRG, bot, msg)
				updateBalance(&valRG)
				msgPiece := `usd: ` + valRG.balance.usd
				msg.Text = boldText("Validator RG Balance After Tranferring") + "\n\n" + msgPiece
				// break
			case "valGrBfTransfer":
				var valGrBf validatorGrBf
				xfer(&valGrBf, bot, msg)
				updateBalance(&valGrBf)
				msgPiece := `gold: ` + valGrBf.balance.gold
				msg.Text = boldText("Validator Group Beneficiary Balance After Tranferring") + "\n\n" + msgPiece
				// break
			case "valBfTransfer":
				var valBf validatorBf
				xfer(&valBf, bot, msg)
				updateBalance(&valBf)
				msgPiece := `gold: ` + valBf.balance.gold
				msg.Text = boldText("Validator Beneficiary Balance After Tranferring") + "\n\n" + msgPiece
				// break
			case "valGrLockGold":
				var valGr validatorGr
				updateBalance(&valGr) // update balance before locking
				lockGoldRun(&valGr, bot, msg)
				updateBalance(&valGr) // update balance after locking
				msgPiece := `gold: ` + valGr.balance.gold + "\n" + `lockedGold: ` + valGr.balance.lockedGold
				msg.Text = boldText("Validator Group Balance After Locking") + "\n\n" + msgPiece
				// break
			case "valLockGold":
				var val validator
				updateBalance(&val) // update balance before locking
				lockGoldRun(&val, bot, msg)
				updateBalance(&val) // update balance after locking
				msgPiece := `gold: ` + val.balance.gold + "\n" + `lockedGold: ` + val.balance.lockedGold
				msg.Text = boldText("Validator Balance After Locking") + "\n\n" + msgPiece
				// break
			case "valAmount":
				msg.Text = "Locking a specific amount from validator was requested"
				// break
			case "valGrAmount":
				msg.Text = "Locking a specific amount from validator group was requested"
				// break
			case "valGrAllUsd":
				var valGr validatorGr
				updateBalance(&valGr) // update balance before exchange
				exchangeUSDToGoldRun(&valGr, bot, msg, 100)
				updateBalance(&valGr) // update balance after exchange
				msgPiece := `gold: ` + valGr.balance.gold + "\n" + `usd: ` + valGr.balance.usd
				msg.Text = boldText("Validator Group Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valGrHalfUsd":
				var valGr validatorGr
				updateBalance(&valGr) // update balance before exchange
				exchangeUSDToGoldRun(&valGr, bot, msg, 50)
				updateBalance(&valGr) // update balance after exchange
				msgPiece := `gold: ` + valGr.balance.gold + "\n" + `usd: ` + valGr.balance.usd
				msg.Text = boldText("Validator Group Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valGrOneForthUsd":
				var valGr validatorGr
				updateBalance(&valGr) // update balance before exchange
				exchangeUSDToGoldRun(&valGr, bot, msg, 25)
				updateBalance(&valGr) // update balance after exchange
				msgPiece := `gold: ` + valGr.balance.gold + "\n" + `usd: ` + valGr.balance.usd
				msg.Text = boldText("Validator Group Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valGrFourThirdsUsd":
				var valGr validatorGr
				updateBalance(&valGr) // update balance before exchange
				exchangeUSDToGoldRun(&valGr, bot, msg, 75)
				updateBalance(&valGr) // update balance after exchange
				msgPiece := `gold: ` + valGr.balance.gold + "\n" + `usd: ` + valGr.balance.usd
				msg.Text = boldText("Validator Group Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valAllUsd":
				var val validator
				updateBalance(&val) // update balance before exchange
				exchangeUSDToGoldRun(&val, bot, msg, 100)
				updateBalance(&val) // update balance after exchange
				msgPiece := `gold: ` + val.balance.gold + "\n" + `usd: ` + val.balance.usd
				msg.Text = boldText("Validator Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valHalfUsd":
				var val validator
				updateBalance(&val) // update balance before exchange
				exchangeUSDToGoldRun(&val, bot, msg, 50)
				updateBalance(&val) // update balance after exchange
				msgPiece := `gold: ` + val.balance.gold + "\n" + `usd: ` + val.balance.usd
				msg.Text = boldText("Validator Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valOneForthUsd":
				var val validator
				updateBalance(&val) // update balance before exchange
				exchangeUSDToGoldRun(&val, bot, msg, 25)
				updateBalance(&val) // update balance after exchange
				msgPiece := `gold: ` + val.balance.gold + "\n" + `usd: ` + val.balance.usd
				msg.Text = boldText("Validator Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valFourThirdsUsd":
				var val validator
				updateBalance(&val) // update balance before exchange
				exchangeUSDToGoldRun(&val, bot, msg, 75)
				updateBalance(&val) // update balance after exchange
				msgPiece := `gold: ` + val.balance.gold + "\n" + `usd: ` + val.balance.usd
				msg.Text = boldText("Validator Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valGrVote":
				var valGr validatorGr
				updateBalance(&valGr) // update balance before voting
				voteRun(&valGr, bot, msg)
				updateBalance(&valGr) // update balance after voting
				msgPiece := `Non-voting: ` + valGr.balance.nonVoting
				msg.Text = boldText("Validator Group Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "valVote":
				var val validator
				updateBalance(&val) // update balance before voting
				voteRun(&val, bot, msg)
				updateBalance(&val) // update balance after voting
				msgPiece := `Non-voting: ` + val.balance.nonVoting
				msg.Text = boldText("Validator Balance After Exhchange") + "\n\n" + msgPiece
				// break
			case "cancel":
				msg.Text = "Back to back menu"
				// break
			}

			// send final message out
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}
		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		chatID := update.Message.Chat.ID
		msg := tgbotapi.NewMessage(chatID, "")
		msg.ParseMode = "Markdown"
		// msg.ReplyToMessageID = update.Message.MessageID
		// bot.Send(msg)
		switch update.Message.Command() {
		case "help":
			msg.Text = "type /balance or /status."
		case "test":
			msg.Text = "I'm ok."
		case "open":
			msg.Text = "What would you like to query?"
			msg.ReplyMarkup = mainKeyboard
		case "close":
			msg.Text = "keyboard closed. Type /open to reopen"
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "synced":
			command, err := botExecCmdOut("celocli node:synced")
			if err != nil {
				msg.Text = err.Error()
				break
			}
			output := cmd.ParseCmdOutput(command, "string", "^.*\\r?(.*)", 0)
			msg.Text = fmt.Sprintf("%v", output)
		case "unlock":
			msg.Text = "Which account would like to unlock?"
			msg.ReplyMarkup = unlockKeyboard
		case "balance":
			var valGr validatorGrRG
			var val validatorRG
			if err := updateBalance(&valGr); err != nil {
				log.Fatal(err)
			}
			if err := updateBalance(&val); err != nil {
				log.Fatal(err)
			}
			// updateBalance(&val)
			msgPiece1 := `*gold*: ` + valGr.balance.gold + "\n" + `*lockedGold*: ` + valGr.balance.lockedGold + "\n" + `*usd*: ` + valGr.balance.usd + "\n" + `*non-voting*: ` + valGr.balance.nonVoting + "\n" + `*total*: ` + valGr.balance.total + "\n"
			msgPiece2 := `*gold*: ` + val.balance.gold + "\n" + `*lockedGold*: ` + val.balance.lockedGold + "\n" + `*usd*: ` + val.balance.usd + "\n" + `*non-voting*: ` + val.balance.nonVoting + "\n" + `*total*: ` + val.balance.total + "\n"
			msg.Text = "Validator Group\n\n" + msgPiece1 + "--------------\n" + "Validator\n\n" + msgPiece2
		case "status":
			command, err := botExecCmdOut("celocli validator:status --validator $CELO_VALIDATOR_RG_ADDRESS")
			if err != nil {
				msg.Text = err.Error()
				break
			}
			words := cmd.ParseCmdOutput(command, "string", "(true|false)\\s*(true|false)\\s*(\\d*)\\s*(\\d*.*)", 0)
			wordsSplit := strings.Fields(fmt.Sprintf("%v", words))
			// fmt.Println("TEST PRINT: ", &wordsSplit)
			ifElected := wordsSplit[0] + "\n"
			ifFrontRunner := wordsSplit[1] + "\n"
			// numProposed := wordsSplit[2] + "\n"
			perctSigned := wordsSplit[2] + "\n"
			message := `*Elected*: ` + ifElected + `*Frontrunner*: ` + ifFrontRunner + `*Signatures*: ` + perctSigned
			msg.Text = message
		case "score":
			command, err := botExecCmdOut("celocli validator:show $CELO_VALIDATOR_RG_ADDRESS")
			if err != nil {
				msg.Text = err.Error()
				break
			}
			words := cmd.ParseCmdOutput(command, "string", "score: (\\S*)", 1)
			msg.Text = `*Score: *` + fmt.Sprintf("%v", words)
		case "transfer":
			var valGrRG validatorGrRG
			var valRG validatorRG
			var valGrBf validatorGrBf
			var valBf validatorBf
			updateBalance(&valGrRG)
			updateBalance(&valRG)
			updateBalance(&valGrBf)
			updateBalance(&valBf)
			msgPiece1 := `*gold*: ` + valGrRG.balance.gold + "\n" + `*usd*: ` + valGrRG.balance.usd + "\n"
			msgPiece2 := `*gold*: ` + valRG.balance.gold + `*usd*: ` + valRG.balance.usd + "\n"
			msgPiece3 := `*gold*: ` + valGrBf.balance.gold + "\n" + `*usd*: ` + valGrBf.balance.usd + "\n"
			msgPiece4 := `*gold*: ` + valBf.balance.gold + `*usd*: ` + valBf.balance.usd + "\n"
			msg.Text = "Validator Group RG\n\n" + msgPiece1 + "--------------\n" + "Validator RG\n\n" + msgPiece2 + "--------------\n" + "Validator Group Beneficiary\n\n" + msgPiece3 + "--------------\n" + "Validator Beneficiary\n\n" + msgPiece4
			msg.ReplyMarkup = transferKeyboard
		case "lockgold":
			// update balance before locking
			var valGr validatorGr
			var val validator
			updateBalance(&valGr)
			updateBalance(&val)
			msgPiece1 := boldText("Gold Available\n") + "Validator Group: " + valGr.balance.gold + "\n"
			msgPiece2 := "Validator: " + val.balance.gold + "\n"
			msgPiece3 := "\nHow much would you like to lock?"
			msg.Text = msgPiece1 + msgPiece2 + msgPiece3
			msg.ReplyMarkup = lockGoldKeyboard
		case "exchange":
			var valGr validatorGr
			var val validator
			updateBalance(&valGr)
			updateBalance(&val)
			msgPiece1 := boldText("USD Available\n") + "Validator Group: " + valGr.balance.usd + "\n"
			msgPiece2 := "Validator: " + val.balance.usd + "\n"
			msgPiece3 := "\nHow much would you like to exchange?\n"
			msg.Text = msgPiece1 + msgPiece2 + msgPiece3
			msg.ReplyMarkup = exchangeUsdKeyboard
		case "vote":
			var valGr validatorGr
			var val validator
			updateBalance(&valGr)
			updateBalance(&val)
			if valGr.balance.nonVoting == "" && val.balance.nonVoting == "" {
				msg.Text = "You have no non-voting lockedGold available"
			} else {
				msgPiece1 := boldText("Non-voting Locked Gold Available\n") + "Validator Group: " + valGr.balance.nonVoting + "\n"
				msgPiece2 := "Validator: " + val.balance.nonVoting + "\n"
				msgPiece3 := "\nHow much would you like to cast?\n"
				msg.Text = msgPiece1 + msgPiece2 + msgPiece3
				msg.ReplyMarkup = electionVoteKeyboard
			}
		case "signing":
			command, err := botExecCmdOut("celocli validator:signed-blocks --signer $CELO_VALIDATOR_SIGNER_ADDRESS")
			if err != nil {
				msg.Text = err.Error()
				break
			}
			output := cmd.ParseCmdOutput(command, "string", "^.*\\r?\\n\\n(.*\\n.*\\n.*\\n.*)", 1)
			msg.Text = fmt.Sprintf("%v", output)
		case "exchange_rate":
			msg.Text = getExchangeRate(msg)
		default:
			msg.Text = "Command not yet supported"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// botExecCmdOut executes commands and returns command outputs
func botExecCmdOut(cmd string) ([]byte, error) {
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		// msg.Text = fmt.Sprint(err) + ": " + string(output)
		return nil, err
	}

	// if string(output) != "" {

	// 	msg.Text = string(output)

	// }
	return output, nil
}
