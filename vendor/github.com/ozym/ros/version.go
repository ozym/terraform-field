package ros

import (
	"strconv"
	"strings"
)

func RouterOsVersion(version string) (int, int) {
	var major, minor int

	if split := strings.Split(version, " "); len(split) > 0 {
		parts := strings.Split(split[0], ".")
		if len(parts) > 0 {
			if n, err := strconv.Atoi(parts[0]); err == nil {
				major = n
			}
		}
		if len(parts) > 1 {
			if n, err := strconv.Atoi(parts[1]); err == nil {
				minor = n
			}
		}
	}

	return major, minor
}
