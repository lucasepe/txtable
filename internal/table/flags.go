package table

import (
	"fmt"
	"strings"
)

func (al *Align) String() string {
	for name, val := range alignLabels {
		if val == *al {
			return name
		}
	}
	return "unknown"
}

func (al *Align) Set(s string) error {
	s = strings.ToLower(s)
	val, ok := alignLabels[s]
	if !ok {
		return fmt.Errorf("invalid align format: %s", s)
	}
	*al = val
	return nil
}

var (
	alignLabels = map[string]Align{
		"center": AlignCenter,
		"left":   AlignLeft,
		"right":  AlignRight,
	}
)
