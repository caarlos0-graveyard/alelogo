# alelo-go [![Build Status](https://travis-ci.org/caarlos0/alelogo.svg?branch=master)](https://travis-ci.org/caarlos0/alelogo) [![SayThanks.io](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/caarlos0)

An unofficial Alelo API implementation to get Card's balances.

```go
import (
  "log"

  "github.com/caarlos0/alelogo"
)

func main() {
  cpf := "123456789-10"
  pwd := "s3cr3t"
  client, err := alelogo.New(cpf, pwd)
	if err != nil {
		log.Fatalln(err.Error())
	}
	cards, err := client.Cards()
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, card := range cards {
		result, err := client.Details(card)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(result)
	}
}
```

## Why?

Because I hate their Android app. So I ~hacked~ their website to
see how it works, created this lib and then used it to create
a [Telegram bot](https://github.com/caarlos0/alelobot),
so I can finally uninstall that piece of crappy software.
