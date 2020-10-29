package shell

// Command is a synchronized wrapper around the standard Go command, which allows you to define environment variables
// in a native way and returns the output as a return value
type Command struct {
	Command           string            // The command to run
	Args              []string          // The args to pass to the command
	WorkingDir        string            // The working directory
	Env               map[string]string // Additional environment variables to set
	OutputMaxLineSize int               // The max line size of stdout and stderr (in bytes)
}
