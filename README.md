# globalmutex

```go
package main

import (
    "github.com/tarmylan/globalmutex"
)

type User struct {
    Id  int
    Key string
    // ...
}

func (p *User) init() {
    p.Key = fmt.Sprintf("user_%v", p.Id)
}

func main() {
    user := &User{1}
    user.init()

    // lock the user and do sth. and unlock it after work done
    globalmutex.Lock(user.Key)
    // TODO ...
    globalmutex.Unlock(user.Key)

    // OR
    globalmutex.LockDo(user.Key, func() {
        // TODO ...
    })
}
```

ref [Multiple Lock Based on Input in Golang](https://medium.com/@kf99916/multiple-lock-based-on-input-in-golang-74931a3c8230)

