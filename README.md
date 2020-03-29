# ctlib

Consul Template hacked into a Go function

`go get github.com/pbar1/ctlib`

Execute populates a Consul Template template string, equivalent to running the command:

```sh
consul-template -template=<file containing the template string> -dry -once
```

Note: it logs to stderr, which is unavoidable due to the uninjectable logger used by Consul Template's library packages.
