// +build !linux, !darwin, !windows

package du

import "errors"

func Get(path string) (Usage, error) {
	return Usage{}, errors.New("Usage data is unavailable on this platform")
}
