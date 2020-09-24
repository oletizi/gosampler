package util

func HandleFatal(err error) {
	logger := GetLogger()
	if err != nil {
		logger.Fatal(err)
	}
}

func HandleWarning(err error) {
	logger := GetLogger()
	if err != nil {
		logger.Error(err)
	}
}
