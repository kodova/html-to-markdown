package main

import (
	"fmt"
	md "github.com/kodova/html-to-markdown"
	"github.com/kodova/html-to-markdown/plugin"
	"log"

)

func main() {
	html := `
	<ul>
		<li><input type=checkbox checked>Checked!</li>
		<li><input type=checkbox>Check Me!</li>
	</ul>
	`
	/*
		- [x] Checked!
		- [ ] Check Me!
	*/

	conv := md.NewConverter("", true, nil)

	// Use the `GitHubFlavored` plugin from the `plugin` package.
	conv.Use(plugin.GitHubFlavored())

	markdown, err := conv.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(markdown)
}
