package conf

import (
	"bufio"
	"os"
)

// Path is relative.
// If path is "", it will not read the env file.
func ReadEnvFile(path string) bool {
	if path == "" {
		return false
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Read the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()
		isVal := false
		for i := range line {
			if isVal {
				os.Setenv(string(line[:i-1]), string(line[i:]))
				isVal = false
				break
			}
			if line[i] == 0x3d {
				isVal = true
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return true
}
