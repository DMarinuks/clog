# clog
Is a color log for golang with different log levels. 

## Description
Log levels:
1. Error 
2. Info
3. Warn 
4. Debug

Available methods:
1. InitLogger( io.writer, prefix string, flag int, clog.LogLevel )
2. Error( args ...interface{} )
3. Info( args ...interface{} )
4. Warn( args ...interface{} )
5. Debug( args ...interface{} )
6. Fatal( args ...interface{} )


## Example

```go
package main

import (
    "os"
    "github.com/DMarinuks/clog"
)

func main() {
  // init logger
  clog.InitLogger(os.Stdout, "", 0, clog.DebugLevel)
  
  clog.Info("Foo", "Bar")

}
```
