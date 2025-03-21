package conf

var (
	debugFlag = false // Show debug message in http error response.
)

func GetDebugFlag() bool {
	return debugFlag
}
