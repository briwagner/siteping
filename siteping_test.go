package siteping

import "testing"

func Test_PingSite(t *testing.T) {
	site := Site{
		Url:     "https://brianwagner.org",
		Content: "",
	}
	res, err := PingSite(site.Url, site.Content)
	if err != nil {
		t.Errorf("PingSite failed. Got %s", err)
	}
	if !res {
		t.Errorf("PingSite failed to reach site.")
	}

	site.Content = "Brian Wagner | Blog"
	res, _ = PingSite("https://brianwagner.org", site.Content)
	if !res {
		t.Errorf("PingSite failed to confirm content. Want %s", site.Content)
	}
}
