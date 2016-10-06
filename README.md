# alelo-go

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
  cards, err := client.Balance()
  if err != nil {
    log.Fatalln(err.Error())
  }
  log.Println(cards)
}
```

## Why?

Because I hate their Android app. So I ~hacked~ their website to
see how it works, created this lib and then used it to create
a [Telegram bot](https://github.com/caarlos0/alelobot),
so I can finally uninstall that piece of crappy software.
