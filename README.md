# globalmutex

```go
package main

import (
  log "github.com/tarmylan/globalmutex"
)

type User struct {
    Id int
    // ...
}

func main() {
    user := &User{1}

    // lock the user and do sth.
    globalmutex.Lock(user.Id)
    defer globalmutex.Unlock(user.Id)
    // TODO ...

    // OR
    globalmutex.LockDo(user.Id, func() {
        // TODO ...
    })
}
```

ref [Multiple Lock Based on Input in Golang](https://medium.com/@kf99916/multiple-lock-based-on-input-in-golang-74931a3c8230)

