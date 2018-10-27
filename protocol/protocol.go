package protocol

import "encoding/json"

// Message is the message structure
type Message struct {
	ID             int32       `json:"id"`     // ID is the id of the request
	Channel        string      `json:"chan"`   // Channel is the channel name
	Message        interface{} `json:"msg"`    // Message is the message body
	HasReturnValue bool        `json:"return"` // HasReturnValue describes if it has return values
}

// CommandMessage is the structure of command message body
type CommandMessage struct {
	Command    string            `json:"cmd"`   // The command name
	Parameters map[string]string `json:"param"` // The parameters
}

// ParseMessagePackage parses message package
func ParseMessagePackage(msg []byte, body interface{}) (*Message, error) {
	rslt := new(Message)
	rslt.Message = body
	err := json.Unmarshal(msg, rslt)
	if err != nil {
		return nil, err
	}
	return rslt, nil
}
