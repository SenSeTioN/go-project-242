package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetSize(path string, all bool, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64
	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if recursive && entry.IsDir() {
			subSize, err := GetSize(filepath.Join(path, entry.Name()), all, recursive)
			if err != nil {
				return 0, err
			}
			total += subSize
		} else {
			entryInfo, err := entry.Info()
			if err != nil {
				return 0, err
			}
			total += entryInfo.Size()
		}
	}

	return total, nil
}

func FormatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	fSize := float64(size)
	i := 0

	for fSize >= 1024 && i < len(units)-1 {
		fSize /= 1024
		i++
	}

	if i == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", fSize, units[i])
}
