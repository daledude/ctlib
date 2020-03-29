# ctlib

`go get github.com/pbar1/ctlib`

Populates a [Consul Template](https://github.com/hashicorp/consul-template) template string, equivalent to running the command:

```sh
consul-template -template=<file containing the template string> -dry -once
```

Note: it logs to stderr, which is unavoidable due to the uninjectable logger used by Consul Template's library packages.
