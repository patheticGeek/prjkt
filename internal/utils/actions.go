package utils

import (
	"fmt"
	types "internal/types"
)

func ExecAction(action types.Action, destination string) error {
	out, _, err := Shellout(action["run"])
	fmt.Println(out)

	if err != nil {
		return err
	}

	return nil
}
