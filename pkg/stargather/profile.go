package stargather

// Profile defines user informations
type Profile struct {
	Organization string   `goquery:"ul.vcard-details li[itemprop='worksFor'] span,text"`
	Location     string   `goquery:"ul.vcard-details li[itemprop='homeLocation'] span,text"`
	Email        string   `goquery:"ul.vcard-details li[itemprop='email'] a,text"`
	Twitter      string   `goquery:"ul.vcard-details li[itemprop='twitter'] a,text"`
	Tabs         []string `goquery:"div.flex-order-1.flex-md-order-none.mt-2.mt-md-0 div a span,text"` // Followers, following, stars
	Repositories []string `goquery:"span.Counter,text"`                                                // 0 == 3 = Repositories count
}
