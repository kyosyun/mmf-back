package amazon

type PAAPIClient struct {
	AccessKey   string
	SecretKey   string
	PartnerTag  string
	Marketplace string
	ServiceHost string
}

type SearchItemsRequest struct {
	Keywords    string   `json:"Keywords,omitempty"`
	JanList     []string `json:"EANList,omitempty"`
	SearchIndex string   `json:"SearchIndex,omitempty"`
	Resources   []string `json:"Resources,omitempty"`
}

type SearchItemsResponse struct {
	Items []Item `json:"Items"`
}

type Item struct {
	ASIN          string     `json:"ASIN"`
	DetailPageURL string     `json:"DetailPageURL"`
	Images        ItemImages `json:"Images"`
}

type ItemImages struct {
	Primary ImageData `json:"Primary"`
}

type ImageData struct {
	Large  ImageDetails `json:"Large"`
	Medium ImageDetails `json:"Medium"`
	Small  ImageDetails `json:"Small"`
}

type ImageDetails struct {
	URL    string `json:"URL"`
	Height int    `json:"Height"`
	Width  int    `json:"Width"`
}

func NewPAAPIClient(accessKey, secretKey, partnerTag string) *PAAPIClient {
	return &PAAPIClient{
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		PartnerTag:  partnerTag,
		Marketplace: "www.amazon.co.jp",
		ServiceHost: "webservices.amazon.co.jp",
	}
}

func (c *PAAPIClient) SearchItemsByJAN(janCode string) (*SearchItemsResponse, error) {
	// endpoint := fmt.Sprintf("https://%s/paapi5/searchitems", c.ServiceHost)

	// request := SearchItemsRequest{
	// 	JanList:     []string{janCode},
	// 	SearchIndex: "All",
	// 	Resources: []string{
	// 		"Images.Primary.Large",
	// 		"Images.Primary.Medium",
	// 		"DetailPageURL",
	// 	},
	// }

	// リクエストの署名と送信処理
	// ... PA-APIの認証処理を実装

	return &SearchItemsResponse{}, nil
}
