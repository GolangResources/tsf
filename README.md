# tsf
Telegram Simple Functions (Bot)

## Examples

### Message
```
package main

import (
        "log"
        "io/ioutil"
        "github.com/GolangResources/tsf/tsf"
)

func main() {
        var tg tsf.TGInfo
        var dest []string
        dest = append(dest, "TELEGRAM-DEST-ID")
        tg.Token = "TELEGRAM BOT TOKEN"
        tsf.SendMSG(tg, dest, "Message here `inline markdown`")
}
```

### Message with Photo
```
package main

import (
        "log"
        "io/ioutil"
        "github.com/GolangResources/tsf/tsf"
)

func main() {
        var tg tsf.TGInfo
        var dest []string
        dest = append(dest, "TELEGRAM-DEST-ID")
        tg.Token = "TELEGRAM BOT TOKEN"
        img, err := ioutil.ReadFile("focus.jpg")
        if err != nil {
                log.Println("ERROR: ", err)
        }
        tsf.SendIMG(tg, dest, "Message here `inline markdown`", img)
}
```
