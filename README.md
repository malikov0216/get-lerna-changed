# Get lerna changed

This library helps to get changed lerna bundle. Actually, it runs:

```bash
lerna changed
```
And gets changed project name
## Installation

Use the "go get".

```bash
go get github.com/malikov0216/get-lerna-changed
```

## Usage

```go
import (
   "fmt"
   getChanged "github.com/malikov0216/get-lerna-changed"
)

func main () {
   projectName, err := getChanged.ExecuteLerna()
   if err != nil {
      fmt.Println(err)
   }
   fmt.Println(projectName)
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)