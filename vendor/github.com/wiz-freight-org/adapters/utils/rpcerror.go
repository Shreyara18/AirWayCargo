package utils

import "strings"

func RPCError(err error) string {
	if err == nil || err.Error() == "" {
		return ""
	}
	contains := strings.Contains(err.Error(), "rpc error: ")
	if contains {
		newMessage := strings.Replace(err.Error(), "rpc error: ", "", 1)
		message := strings.Split(newMessage, " ")
		desc := message[5:]
		code := message[2]
		errorMessage := strings.Join(desc, " ")
		if code != "Unknown" {
			return err.Error()
		} else {
			return errorMessage
		}
	} else {
		return err.Error()
	}
}
