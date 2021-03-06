package content

import (
	"bytes"
	"gocms/pkg/files"
	"gocms/pkg/page"
	"net/http"
	"strings"
	"testing"
)

func TestReplaceTokens(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	html := "#TITLE# #HOME# #CONTENT#"
	page := page.Info{Title: "Test", Content: "Content"}
	req, _ := http.NewRequest("GET", "", bytes.NewBufferString(""))
	req.Host = "test"
	got := ReplaceTokens(req, html, page, []string{})
	want := `Test <a href="http://test/">Home Page</a> Content`
	if got != want {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestGetBreadcrumbLinks(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	p := page.GetByID("test", 8, false)
	want := `<a href="base.com/">Home Page</a> > <a href="base.com/0-1">Page 0-1</a> > Page 0-1-0`
	req, _ := http.NewRequest("GET", p.Route, bytes.NewBufferString(""))
	req.Host = "base.com"
	req.RequestURI = "/0-1-0"
	got := GetBreadcrumbLinks(req, p, "base.com")
	if got != want {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestGetDayArchiveLinks(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	var got string
	for _, day := range getDayArchiveLinks(2015, 2, "test.com") {
		println(day)
		got = day
	}
	want := `<li><a href="test.com/blog/3-0-3`
	if !strings.HasPrefix(got, want) {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestGetMonthArchiveLinks(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	var got int
	for _, month := range getMonthArchiveLinks(2015, "test.com") {
		println(month)
		got++
	}
	want := 36
	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}

func TestGetYearArchiveLinks(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	var got string
	for _, month := range getYearArchiveLinks("test.com") {
		println(month)
		got = month
	}
	want := `</ul>`
	if !strings.HasPrefix(got, want) {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestGetInt(t *testing.T) {
	want := 0
	got := getInt("a")
	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
	want = 2015
	got = getInt("2015")
	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}

func TestGenerateArchive(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	req, _ := http.NewRequest("GET", "test.com", bytes.NewBufferString(""))
	req.Host = "test"
	baseURL := "test.com"
	p := page.Info{}
	arch := generateArchive(baseURL, &p, []string{})
	got := len(arch)
	want := 7992
	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
	arch = generateArchive(baseURL, &p, []string{"2015"})
	gots := p.Title
	wants := "Archives for 2015"
	if gots != wants {
		t.Errorf("Want %s got %s", wants, gots)
	}
	got = len(arch)
	want = 1914
	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
	arch = generateArchive(baseURL, &p, []string{"2015", "01"})
	got = len(arch)
	want = 67
	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}

func TestAddPosts(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	got := addPosts("test", "abc #POSTS_1#", "test.com")
	want := "Read more"
	if !strings.Contains(got, "Read more") {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestGetPagesInCategory(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	pages := getPagesInCategory(page.Info{ID: 601}, "test.com")
	got := len(pages)
	want := 125
	if got != want {
		t.Errorf("Want %d got %d", want, got)
	}
}

func TestAddMenus(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	got := addMenus("aaa #SIDE_MENU# bbb", "test.com", "-")
	want := "aaa <li><a"
	if !strings.HasPrefix(got, want) {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestAddPageLinks(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	got := addPageLinks("aaa #PAGES_1# bbb", "test.com", page.Info{Parent: 1})
	want := "aaa <li><a"
	if !strings.HasPrefix(got, want) {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestAddCategoryLinks(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	got := addCategoryLinks("aaa #CATEGORIES_1# bbb", "test.com", page.Info{ID: 600})
	want := "aaa <li><a"
	if !strings.HasPrefix(got, want) {
		t.Errorf("Want %s got %s", want, got)
	}
}

func TestAddRecentPostsLinks(t *testing.T) {
	files.Root = "../files/"
	page.LoadData("test")
	got := addRecentPostsLinks("aaa #RECENT_4# bbb", "test.com")
	want := "aaa <li><a"
	if !strings.HasPrefix(got, want) {
		t.Errorf("Want %s got %s", want, got)
	}
}
