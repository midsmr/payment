package application

var CoreVersion = "dev"

var LastCommit = "unknown"

var BuildDate = "unknown"

func InitServer() error {

	if err := initWorkdir(); err != nil {
		return err
	}

	if err := initLogger(); err != nil {
		return err
	}

	if err := initConfig(); err != nil {
		return err
	}

	return nil
}
