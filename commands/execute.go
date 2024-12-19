package commands

import "fmt"

func ExecuteCommand(cmd interface{}) error {
	var err error

	switch cmd := cmd.(type) {
	case *CreateCustomMacaroonCommand:
		err = ExecuteCustomMacaroonCommand(cmd)
	case *WatchLndCommand:
		err = ExecuteWatchLndCommand(cmd)
	case *CreateWatcherMacaroonCommand:
		err = ExecuteCreateWatcherMacaroonCommand(cmd)
	default:
		err = fmt.Errorf("unknown command type")
	}

	if err != nil {
		return fmt.Errorf("cannot execute command %T: %w", cmd, err)
	}

	return err
}
