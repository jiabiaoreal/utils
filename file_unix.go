// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.
// +build !windows

package utils

import (
	"os"
)

// Replace atomically replaces the destination file or directory with the source.
func Replace(source, destination string) error {
	return os.Rename(source, destination)
}
