package format

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/umats/gomarkdoc/format/formatcore"
	"github.com/umats/gomarkdoc/lang"
)

// GitLabFlavoredMarkdown provides a Format compatible with GitLab Flavored
// Markdown's syntax and semantics.
type GitLabFlavoredMarkdown struct {
	GitHubFlavoredMarkdown
}

// LocalHref generates a link to a heading within the same GitLab document.
func (f *GitLabFlavoredMarkdown) LocalHref(headerText string) (string, error) {
	result := formatcore.PlainText(headerText)
	result = strings.ToLower(result)
	result = strings.TrimSpace(result)
	result = gfmWhitespaceRegex.ReplaceAllString(result, "-")
	result = gfmRemoveRegex.ReplaceAllString(result, "")
	for strings.Contains(result, "--") {
		result = strings.ReplaceAll(result, "--", "-")
	}

	return fmt.Sprintf("#%s", result), nil
}

// CodeHref generates a GitLab source link for the provided code entry.
func (f *GitLabFlavoredMarkdown) CodeHref(loc lang.Location) (string, error) {
	if loc.Repo == nil {
		return "", nil
	}

	relative := loc.Filepath
	if filepath.IsAbs(loc.Filepath) {
		var err error
		relative, err = filepath.Rel(loc.WorkDir, loc.Filepath)
		if err != nil {
			return "", fmt.Errorf("relative file path: %w", err)
		}
	}

	full := filepath.Join(loc.Repo.PathFromRoot, relative)
	path, err := filepath.Rel(string(filepath.Separator), full)
	if err != nil {
		return "", fmt.Errorf("repository-relative path: %w", err)
	}

	line := fmt.Sprintf("L%d", loc.Start.Line)
	if loc.Start.Line != loc.End.Line {
		line = fmt.Sprintf("L%d-%d", loc.Start.Line, loc.End.Line)
	}

	return fmt.Sprintf(
		"%s/-/blob/%s/%s#%s",
		loc.Repo.Remote,
		loc.Repo.DefaultBranch,
		filepath.ToSlash(path),
		line,
	), nil
}
