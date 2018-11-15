package location

import mesg "github.com/mesg-foundation/go-service"

// output key for errors.
const errOutputKey = "error"

// errorOutput is the error output data.
type errorOutput struct {
	Message string `json:"message"`
}

// newErrorOutput returns a new error output from given err.
func newErrorOutput(err error) (outputKey string, outputData mesg.Data) {
	return errOutputKey, errorOutput{Message: err.Error()}
}
