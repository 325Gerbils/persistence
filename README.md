# persistence
Go persistent key-value storage, unoptimized

### Installation:

```
go get -u github.com/325gerbils/persistence
```

### Usage:

```go

import "github.com/325gerbils/persistence"

data := "value"
persistence.Store("key", data)




// Exit and restart program...

import (
    "fmt"
    
    "github.com/325gerbils/persistence"
)

data := persistence.Get("key")

// Prints "value"
fmt.Println(data)

```

### Todo:

* Use a faster method of data storage than `json.Marshal()` and `json.Unmarshal()`
* Write to a common directory instead of *projectFolder*/data
