// Package tags contains code to demonstrate usage of build tags.
//
//nolint:mnd // This legacy implementation intentionally relies on these patterns.
package tags

// Untagged is visible without tags.
func Untagged() int {
	return 6
}
