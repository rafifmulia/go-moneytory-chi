package conf

var (
	preforkFlag = false // Run server on entirely different process
	childFlag   = false // Indicating if this process is running as spawned child or not in prefork mode. If not, then this will spawn childs in prefork mode.
	bindHost    = "127.0.0.1:8080"
)

func GetPreforkFlag() (bool, bool) {
	return preforkFlag, childFlag
}

func GetBindHost() string {
	return bindHost
}
