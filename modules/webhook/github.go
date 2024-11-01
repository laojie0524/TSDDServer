package webhook

import (
	"fmt"
	"io/ioutil"

	"github.com/laojie0524/TSDDServerLib/pkg/wkhttp"
)

func (w *Webhook) github(c *wkhttp.Context) {
	fmt.Println("github webhook-->", c.Params)

	result, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("github-result-->", result)
}
