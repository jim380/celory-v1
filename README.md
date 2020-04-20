# Celory

A wrapper for [celocli](https://www.npmjs.com/package/@celo/celocli) - a tool for interacting with the Celo Protocol.

### Instructions
1. `$ git clone https://github.com/jim380/Celory.git` into your `$GOPATH/src/github.com`
2. `$ go get`
3. `$ touch config.env` , then fill in your environment variables in it. Check `config.env.example` for reference.
4. `$ go build`
5. `$ ./celory`

#### Run Node Set-Up
- `$ ./celory`

##### Sample Outputs
* [Account](https://pastebin.com/aRW2Txv3)
* [Validator](https://pastebin.com/Zq6PrFg5)
* [Proxy](https://pastebin.com/Bbg8NnFh)
* Attestation --- TODO

#### Run Celocli Wrapper
- `$ ./celory -cmd`

- Menu:
    1. Election Show
        - `election:show`
    2. Account Balance
        - `account:balance`
        - `lockedgold:lock`
        - `exchange:dollars`
    3. Account Show
        - `account:show`
    4. Lockgold Show
        - `lockedgold:show`
        - `election:vote`
    5. Validator Show
        - `validator:show`
    6. Validator Status
        - `validator:status`
    7. Get Metadata
        - `account:get-metadata`
    8. Node Synced
        - `node:synced`
    9. More to add
##### Sample Outputs
* [Wrapper](https://pastebin.com/6nt2XGCD)

#### Run Telegram Bot
- `$ ./celory -bot`
- Menu
    1. Balance (`account:balance`)
    2. Synced (`node:synced`)
    3. Status (`validator:status`)
    4. Score (`lockedgold:show`)
    5. Signing (`validator:signed-blocks`)
    6. Exchange (`exchange:dollars`)
    7. Lockgold (`lockedgold:lock`)
    8. Vote (`election:vote`)
    9. Exchange Rate (`exchange:show`)
    10. Unlock (`account:unlock`)
    11. Transfer (`releasegold:transfer-dollars`)

![image](./assets/telebot.png)