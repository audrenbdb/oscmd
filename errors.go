package oscmd

const (
	cmdNotFound = cmdErr("command not found in path")
)


//error handling
type cmdErr string
func (e cmdErr) Error() string {return string(e)}

