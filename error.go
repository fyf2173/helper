package helper

var lastError error = nil

func GetLastError() error {
	return lastError
}

func SetLastError(err error) {
	lastError = err
	return
}