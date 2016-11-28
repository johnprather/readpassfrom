# readpassfrom

A small package to handle reading a password from specified file.

It basically wraps file open/close around terminal.ReadPassword() which takes a file descriptor as an argument.

The purpose of this is to be able to easily read a password from /etc/tty instead of an open (stdin) file descriptor.  This is handy when
stdin is in use, and therefore, cannot be used to read user input.  The effect is similar to when a user redirects
data to ssh, and ssh prompts for a password on /etc/tty because its stdin is receiving redirected data.

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

