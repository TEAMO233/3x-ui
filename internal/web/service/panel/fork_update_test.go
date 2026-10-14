package panel

import "testing"

func TestForkUpdaterUsesSelectedRelease(t *testing.T) {
	t.Setenv("XUI_REPO", "")
	t.Setenv("XUI_BRANCH", "")
	t.Setenv("XUI_UPDATER_URL", "")
	t.Setenv("XUI_RAW_BASE", "")
	for _, tag := range []string{"v3.9.0-fourgetu.1", "v3.8.5-fourgetu.1", "v3.9.0", devReleaseTag} {
		ref := tag
		if ref == devReleaseTag {
			ref = panelUpdateBranch()
		}
		want := "https://raw.githubusercontent.com/" + panelUpdateRepo() + "/" + ref + "/update.sh"
		got, err := panelUpdaterURL(tag)
		if err != nil || got != want {
			t.Fatalf("tag=%q URL=%q err=%v, want %q", tag, got, err, want)
		}
	}
	for _, tag := range []string{"", "../main", "v3.9.0/../../main", "v3.9.0;echo test"} {
		if _, err := panelUpdaterURL(tag); err == nil {
			t.Fatalf("accepted invalid release tag %q", tag)
		}
	}
}
