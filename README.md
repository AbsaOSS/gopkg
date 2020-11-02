# gopkg

Absa Go package library

## Project Health

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Absa Go package library

## Absa library

- [env](#env)
- [shell](#shell)

### env
Environment variable helper extending standard [os.Getenv](https://golang.org/pkg/os/#Getenv).
#### Overview
Package contains several functions which returns the typed env variable for the given 
key and falls back to the given defaultValue if not set.
#### Usage
```go
noProxy := env.GetEnvAsStringOrFallback("NO_PROXY", "*.foo.bar.com")
 
runOnStart := env.GetEnvAsStringOrFallback("RUN_ON_START", true)
```

### shell
Shell command runner 
#### Overview
Package shell is used for running executables and returns command output for further processing. It is a synchronized 
wrapper around the standard Go [command](https://golang.org/pkg/os/exec/#Cmd) package, which allows you to define environment variables in a native way 
and returns the output as a return value.
#### Usage
```go
cmd := shell.Command{
	Command: "sh",
	Args: []string{"-c", "terraform apply -auto-approve tfplan"},
	Env: map[string]string{"name":"test"},
}
o, _ = shell.Execute(cmd)
```
