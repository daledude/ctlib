# ctlib

Consul Template hacked into a Go function

Execute populates a Consul Template template string, equivalent to running the command:
```consul-template -template=<file containing the template string> -dry -once```

Note: it logs to stderr, which is unavoidable due to the uninjectable logger used by Consul Template's library packages.
