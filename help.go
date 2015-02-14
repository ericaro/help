// Package help allows you to define help sections
// for your command line interfaces. It extends the flag package
// to provide flag support for help section.
//
// The package defines a subcommand "help" that you can integrate into you main CLI.
//
// For example:
//
//      $ sbr help file
//
// The subcommand displays a summary of all the sections available or the detail for any section.
//
// Sections summary example:
//
//     $ sbr help
//     Usage: sbr help  <section>
//
//     where <section> is one of:
//        1.file             help relative to file management
//
// Sections details:
//     $ sbr help file dir
//     File
//
//     This document describe the file format
//     ...
//
// Sections are written using Markdown, and displayed in the terminal output using ansi escape code.
package help

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/ericaro/ansifmt/ansiblackfriday"
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

//HelpCommand contains a slice of sections representing the help subsystem.
//
// HelpCommand can parse "args"
//
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

//Flags is to update a FlagSet with the options for the helpCommand.
//
//Currently, there are no flags.
func (h *HelpCommand) Flags(fs *flag.FlagSet) *flag.FlagSet { return fs }

//Run execute the HelpCommand:
//
// with no args: prints the section summary
//
// with arguments print out every section found in the args, or an error message
func (h *HelpCommand) Run(args []string) {
	if len(args) == 0 {
		prog, helps := os.Args[0], os.Args[1]
		fmt.Fprintf(os.Stderr, "Usage: %s %s  <section>\n\n", prog, helps)
		fmt.Fprintf(os.Stderr, "where <section> is one of:\n")

		for _, s := range h.sections {
			fmt.Fprintf(os.Stderr, "  %-15s %s\n", s.name, s.description)
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

//Section create a new section in 'Command'.
//
// a section's name is used to identify it
//
// a section's description is used in the summary
//
// a section's content is a Markdown description.
func (h *HelpCommand) Section(name, description, content string) {
	h.sections = append(h.sections, &section{name, description, content})
}

//Section create a new section in 'Command'.
//
// a section's name is used to identify it
//
// a section's description is used in the summary
//
// a section's content is a Markdown description.
func Section(name, description, content string) {
	Command.Section(name, description, content)
}

type section struct{ name, description, content string }

//Print out the section help
func (s *section) Print() {
	renderer := ansiblackfriday.NewAnsiRenderer()

	buf := new(bytes.Buffer)
	_, err := buf.Write(blackfriday.Markdown(([]byte)(s.content), renderer, commonExtensions))
	if err != nil {
		fmt.Fprintf(os.Stderr, "buffer error: %v", err)
		return
	}
	fmt.Fprintln(os.Stderr, buf.String())
}
