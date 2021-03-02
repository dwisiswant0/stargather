package stargather

import (
	"fmt"
	"net/http"
	"strings"

	"astuart.co/goq"
	"ktbs.dev/mubeng/pkg/mubeng"
)

// Gather stargazers of GitHub repository
func (d *Data) Gather() (*Data, error) {
	if len(d.Button) > 0 {
		d.URL = d.Button[0]
		d.Button = make([]string, 0)

		if strings.Contains(d.URL, "before") {
			d.End = true
			return d, nil
		}
	} else {
		d.URL += "/stargazers"
	}

	req, err := d.get(d.URL)
	if err != nil {
		return d, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return d, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return d, fmt.Errorf("error code %d", resp.StatusCode)
	}

	if err := goq.NewDecoder(resp.Body).Decode(&d); err != nil {
		return d, err
	}

	if len(d.Button) == 2 {
		d.Button = d.Button[1:]
		d.End = false
	}

	return d, nil
}

// Extract GitHub user informations
func (d *Data) Extract(username string) (*Profile, error) {
	p := &Profile{}

	req, err := d.get("https://github.com/" + username)
	if err != nil {
		return p, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return p, fmt.Errorf("error code %d", resp.StatusCode)
	}

	if err := goq.NewDecoder(resp.Body).Decode(p); err != nil {
		return p, err
	}

	return p, nil
}

func (d *Data) get(URL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return req, err
	}

	if d.Cookie != "" {
		req.Header.Set("Cookie", d.Cookie)
	}

	if d.Proxy != "" {
		client.Transport, err = mubeng.Transport(d.Proxy)
		if err != nil {
			return req, err
		}
	}

	return req, nil
}
