// Package plugin contains all the rules that are not
// part of Commonmark like GitHub Flavored Markdown.
package plugin

import (
	md "github.com/kodova/html-to-markdown"
)

// GitHubFlavored is GitHub's Flavored Markdown
func GitHubFlavored() md.Plugin {
	return func(c *md.Converter) (rules []md.Rule) {
		rules = append(rules, Strikethrough("")(c)...)
		rules = append(rules, EXPERIMENTAL_Table...)
		rules = append(rules, TaskListItems()(c)...)
		return
	}
}
