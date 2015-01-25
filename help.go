// Package help allows you to define help sections
// for your command line interfaces. It extends the flag package
// to provide flag support for help section.
package help

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/ericaro/ansifmt/renderer"
	"github.com/russross/blackfriday"
)

var (
	Command          = New()
	commonExtensions = 0 |
		blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
		blackfriday.EXTENSION_TABLES |
		blackfriday.EXTENSION_FENCED_CODE |
		blackfriday.EXTENSION_AUTOLINK |
		blackfriday.EXTENSION_STRIKETHROUGH |
		blackfriday.EXTENSION_SPACE_HEADERS |
		blackfriday.EXTENSION_HEADER_IDS
)

//HelpCommand contains a bunch of sections to be displayed. It can parse "args"
// and it is compatible with rakyll command package.
type HelpCommand struct {
	sections []*section
}

//New creates a new empty HelpCommand
func New() *HelpCommand {
	return &HelpCommand{
		sections: make([]*section, 0, 10),
	}
}

// unchanged flags parsing

//Flags is to update 'fs' with current flags. Currently, there is no flags.
func (h *HelpCommand) Flags(fs *flag.FlagSet) *flag.FlagSet { return fs }

//Run execute the HelpCommand:
// with no args: prints the section summary
// with arguments print out every section found in the args, or an error message
func (h *HelpCommand) Run(args []string) {
	if len(args) == 0 {
		prog, helps := os.Args[0], os.Args[1]
		fmt.Fprintf(os.Stderr, "Usage: %s %s  <section>\n\n", prog, helps)
		fmt.Fprintf(os.Stderr, "where <section> is one of:\n")

		for i, s := range h.sections {
			fmt.Fprintf(os.Stderr, "  %v.%-15s %s\n", i+1, s.name, s.description)
		}
		fmt.Fprintln(os.Stderr)

	} else {
		for _, name := range args { // for each arg, that is suppose to be a "section"
			// section lookup
			exists := false
			var s *section
			for _, s = range h.sections {
				if s.name == name {
					exists = true
					break
				}
			}

			if !exists {
				fmt.Printf("No help found for %q\n", name)
			} else {
				s.Print()
			}
		}
	}
}

//Section add a section to the HelpCommand
func (h *HelpCommand) Section(name, description, content string) {
	h.sections = append(h.sections, &section{name, description, content})
}

//Section adds a section on 'Command'
func Section(name, description, content string) {
	Command.Section(name, description, content)
}

type section struct{ name, description, content string }

//Print out the section help
func (s *section) Print() {
	renderer := renderer.NewAnsiRenderer()

	buf := new(bytes.Buffer)
	_, err := buf.Write(blackfriday.Markdown(([]byte)(s.content), renderer, commonExtensions))
	if err != nil {
		fmt.Fprintf(os.Stderr, "buffer error: %v", err)
		return
	}
	fmt.Fprintln(os.Stderr, buf.String())
}
