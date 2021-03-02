package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"dw1.io/stargather/pkg/stargather"
	"github.com/briandowns/spinner"
)

var (
	err  error
	file *os.File
	wg   sync.WaitGroup
	data *stargather.Data

	repo, cookie, delim, output, proxy string
)

func init() {
	flag.StringVar(&repo, "r", "", "Repository (format: owner/name)")
	flag.StringVar(&cookie, "c", "", "GitHub cookies (optional)")
	flag.StringVar(&delim, "d", ",", "Data delimiter")
	flag.StringVar(&output, "o", "", "Output data file")
	flag.StringVar(&proxy, "x", "", "Proxy URL (HTTP/SOCKS5)")
	flag.Parse()

	if repo == "" {
		log.Fatal("flag -r required")
	}

	if output != "" {
		file, err = os.OpenFile(output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	spin := spinner.New(spinner.CharSets[11], 90*time.Millisecond, spinner.WithWriter(os.Stderr))
	if err := spin.Color("cyan"); err != nil {
		log.Fatal(err)
	}

	data, err = stargather.New(repo, cookie, proxy)
	if err != nil {
		log.Fatal(err)
	}

	spin.Start()
	spin.Suffix = " Collecting stargazers"

	for {
		for {
			data, err = data.Gather()
			if err == nil {
				break
			}
		}

		spin.Suffix = fmt.Sprintf(" %d stargazers collected", len(data.Stars))

		if data.End {
			break
		}
	}
	spin.Stop()

	for _, user := range data.Stars {
		wg.Add(1)
		for {
			go func(user string) {
				defer wg.Done()

				info, err := data.Extract(user)
				if err != nil {
					return
				}

				write := fmt.Sprintf(
					"%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n",
					user, delim, info.Organization, delim,
					info.Location, delim, info.Email, delim,
					info.Twitter, delim, info.Tabs[0], delim,
					info.Tabs[1], delim, info.Tabs[2], delim,
					info.Repositories[0],
				)

				fmt.Print(write)
				if file != nil {
					fmt.Fprint(file, write)
				}
			}(user)

			if err == nil {
				break
			}
		}
	}

	wg.Wait()
	file.Close()
}
