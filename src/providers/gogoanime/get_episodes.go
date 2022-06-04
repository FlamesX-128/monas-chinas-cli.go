package gogoanime

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FlamesX-128/monas-chinas-cli/src/tools"
	"github.com/gocolly/colly"
)

func (p *provider) GetEpisodes(u string) (e []string) {
	q := "ul#episode_page a[href]"

	f := func(h *colly.HTMLElement) string {

		return h.Attr("ep_end")
	}

	n, _ := strconv.Atoi(tools.NewScraper((url + u), q, f)[0])

	for i := 1; i <= n; i++ {
		fmt.Println(i, n)

		e = append(e,
			(url + u[strings.LastIndex(u, "/")+1:] + "-episode-" + strconv.Itoa(i)),
		)

	}

	return
}
