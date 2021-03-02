package stargather

import (
	"fmt"
	"net/http"
	"path"

	"ktbs.dev/mubeng/pkg/mubeng"
)

// New to define stargather datas
func New(repository string, cookie string, proxy string) (*Data, error) {
	d := &Data{}

	repository = "https://" + path.Join("github.com", repository)
	req, err := http.NewRequest("GET", repository, nil)
	if err != nil {
		return d, err
	}

	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}

	if proxy != "" {
		client.Transport, err = mubeng.Transport(proxy)
		if err != nil {
			return d, err
		}

		d.Proxy = proxy
	}

	resp, err := client.Do(req)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return d, fmt.Errorf("error code %d", resp.StatusCode)
	}

	d.URL = repository
	d.Cookie = cookie

	return d, nil
}
