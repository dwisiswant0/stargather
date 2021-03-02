# Stargather

Stargather is fast GitHub repository stargazers information gathering tool that can scrapes:
- **Organization**,
- **Location**,
- **Email**,
- **Twitter**,
- **Followers**,
- **Following**,
- **Stars**, and
- **Repositories** count.

## Installation

It's fairly simple, you will need [Go 1.15+](https://golang.org/doc/install) compiled: `go get dw1.io/stargather`.

## Usage

```console
$ stargather -h
Usage of stargather:
  -c string
        GitHub cookies (optional)
  -d string
        Data delimiter (default ",")
  -o string
        Output data file
  -r string
        Repository (format: owner/name)
  -x string
        Proxy URL (HTTP/SOCKS5)
```

<table>
	<td>
		<b>NOTE:</b>
		<ul>
			<li>Cookies are only needed if you want to scrape email information <i>(or private repository)</i>.</li>
			<li>If you encounter <code>error code 429</code>, it means you're rate-limited by GitHub. Try it in a few seconds to initiate again, or if you hate to wait, use a proxy with <code>-x</code> flag.</li>
		</ul>
	</td>
</table>

## License

`stargather` is distributed under MIT.