package format_test

import (
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/umats/gomarkdoc/format"
	"github.com/umats/gomarkdoc/lang"
)

func TestGitLabFlavoredMarkdownLocalHref(t *testing.T) {
	is := is.New(t)

	var f format.GitLabFlavoredMarkdown
	res, err := f.LocalHref("GitLab -- heading!")
	is.NoErr(err)
	is.Equal(res, "#gitlab-heading")
}

func TestGitLabFlavoredMarkdownCodeHref(t *testing.T) {
	is := is.New(t)

	wd, err := filepath.Abs(".")
	is.NoErr(err)

	var f format.GitLabFlavoredMarkdown
	res, err := f.CodeHref(lang.Location{
		Start:    lang.Position{Line: 12},
		End:      lang.Position{Line: 14},
		Filepath: filepath.Join(wd, "subdir", "file.go"),
		WorkDir:  wd,
		Repo: &lang.Repo{
			Remote:        "https://gitlab.example.com/group/project",
			DefaultBranch: "main",
			PathFromRoot:  "/",
		},
	})
	is.NoErr(err)
	is.Equal(res, "https://gitlab.example.com/group/project/-/blob/main/subdir/file.go#L12-14")
}
