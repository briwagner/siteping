package siteping

import "testing"

func Test_PingSite(t *testing.T) {
	res, err := PingSite("https://brianwagner.org", "")
	if err != nil {
		t.Errorf("PingSite failed. Got %s", err)
	}
	if !res {
		t.Errorf("PingSite failed to reach site.")
	}

	content := "Brian Wagner | Blog"
	res, err = PingSite("https://brianwagner.org", content)
	if !res {
		t.Errorf("PingSite failed to confirm content. Want %s", content)
	}
}
