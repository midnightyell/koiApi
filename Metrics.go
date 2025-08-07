package koiApi

import "fmt"

// Metrics represents a map of metrics data.
type Metrics string // Read-only

func (m *Metrics) Summary() string {
	return fmt.Sprintf("%-40s %s", string(*m), "")
}
