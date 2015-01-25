[![Build Status](https://travis-ci.org/ericaro/help.png?branch=master)](https://travis-ci.org/ericaro/help) [![GoDoc](https://godoc.org/github.com/ericaro/help?status.svg)](https://godoc.org/github.com/ericaro/help)


# Help

Package help allows you to define help sections for your command line interfaces. It extends the flag package to provide flag support for help section.

The package defines a subcommand "help" that you can integrate into you main CLI.

    $ sbr help file

The subcommand displays a summary of all the sections available, or the detail for any section.

*sections summary*

    $ sbr help
    Usage: sbr help  <section>
    
    where <section> is one of:
       1.file             help relative to file management

Sections are written in Markdown, and displayed in the terminal output using ansi escape code.

## registering sections

    help.Section("file", "help relative to file management", `
    # File Format
    
    this is Mardown *format*
    
    `)

this is enough to register the section "file", into the global var `help.Command`


## integration with http://github.com/rakyll/command

the package `github.com/ericaro/help` contains a `var Command` that implements Rakyll's `Cmd` interface

    Cmd interface {
        Flags(*flag.FlagSet) *flag.FlagSet
        Run(args []string)
    }

So you can just register the help subcommand:

    command.On("help","display help", help.Command, nil)


# License

help is available under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).

# Branches

master: [![Build Status](https://travis-ci.org/ericaro/help.png?branch=master)](https://travis-ci.org/ericaro/help) against go versions:

  - 1.2
  - 1.3
  - tip

dev: [![Build Status](https://travis-ci.org/ericaro/help.png?branch=dev)](https://travis-ci.org/ericaro/help) against go versions:

  - 1.2
  - 1.3
  - tip


