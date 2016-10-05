# readpassfrom

A small package to handle reading a password from specified file.

It basically wraps file open/close around terminal.ReadPassword() which takes a file descriptor as an argument.

# usage

Import the package:

```
import (
  ...

  rpf "github.com/johnprather/readpassfrom"
)
```

Use ReadPasswordFrom():

```
func main() {
  var pass *string
  var err error
  pass, err = rpf.ReadPasswordFrom("/dev/tty", "Password: ")
  if err != nil {
    ...
  }
}
```
