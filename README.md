# gopkg

Absa Go package library

## Absa library

- [controller](#controller)
- [dns](#dns)
- [env](#env)
- [k8s](#k8s)
- [reflect](#reflect)  
- [shell](#shell)
- [strings](#strings)

### controller
package controller extends sigs.k8s.io/controller-runtime [os.Getenv](https://golang.org/pkg/os/#Getenv).
#### Overview
Package contains `ReconcileResult` which provides abstraction over reconciliation loop management.

#### Usage
```go
const reconcileSeconds = 10
result := utils.NewReconcileResult(time.Duration(reconcileSeconds) * time.Second)
...
func (r *MyReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	...
	if finish {
		// sucesfuly stop reconciliation loop
		return result.Stop()
	}
	if err := doSomething(); err != nil {
		// requeue loop immediately with error
		return result.RequeueError(err)
	}
	// requeue loop after reconcileSeconds 
	return result.Requeue()
}
```

### dns
DNS contains helper functions related to DNS.
#### Overview
Package currently contains Dig function retrieving 
slice of tuples <IP address, A record > for specific FQDN.
#### Usage
```go
	edgeDNSServer := "8.8.8.8"
	fqdn := "google.com"
	result, err := Dig(edgeDNSServer, fqdn)
```

### env
Environment variable helper extending standard [os.Getenv](https://golang.org/pkg/os/#Getenv).
#### Overview
Package contains several functions which returns the typed env variable for the given 
key and falls back to the default value if not set.
#### Usage
```go
noProxy := env.GetEnvAsStringOrFallback("NO_PROXY", "*.foo.bar.com")
 
runOnStart := env.GetEnvAsStringOrFallback("RUN_ON_START", true)
```

### k8s
Package k8s provides extensions for k8s apimachinery/pkg/apis
#### Overview
Package contains following functions:
 - `MergeAnnotations(target *metav1.ObjectMeta, source *metav1.ObjectMeta)` adds or updates annotations from 
   defaultSource to defaultTarget
 - `ContainsAnnotations(target *metav1.ObjectMeta, source *metav1.ObjectMeta) bool` checks if defaultTarget 
   contains all annotations of defaultSource
Any kubernetes resource having ObjectMeta can be passed.  
#### Usage
```go
if !k8s.ContainsAnnotations(currentIngress.ObjectMeta, expectedIngress.ObjectMeta) {
    MergeAnnotations(currentIngress.ObjectMeta, expectedIngress.ObjectMeta)
}
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

### strings
Package strings provide helper functions related to string
#### Overview
Package strings contains extensions of standard strings and formatting functions:
- `Format(v interface{}) string` Format converts type to formatted string. If value is struct, function returns formatted JSON.

#### Usage
```go
log.Debug().Msgf("current config: %s",utils.ToString(config))
```
