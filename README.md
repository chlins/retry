## Retry

It's a very simple retry package for golang, just for fun.

## Usage

```go
import (
    "github.com/chlins/retry"
)

func main() {
    fn := func() error {
        return nil
    } 
    retry.Do(context.Background(), fn, 10)
}
```

