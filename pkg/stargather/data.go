package stargather

// Data defines required options
type Data struct {
	URL    string
	End    bool
	Proxy  string
	Stars  []string `goquery:"h3 span a[data-hovercard-type='user'],text"`
	Button []string `goquery:"div.BtnGroup a.btn.btn-outline.BtnGroup-item,[href]"`
	Cookie string
}
