package conf

// These flag is intended for enable profiling.
// Either using "runtime/pprof" package, or with http access such as
// import _ "net/http/pprof",

var (
	cpuProfileFlag  string = ""    // Write CPU profile to this file.
	memProfileFlag  string = ""    // Write memory profile to this file.
	httpProfileFlag bool   = false // Live profiling with http access.
)

func GetMemProfileFlag() string {
	return memProfileFlag
}

func GetCpuProfileFlag() string {
	return cpuProfileFlag
}

func GetHttpProfileFlag() bool {
	return httpProfileFlag
}
