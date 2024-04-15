package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ProfoundNetworks/gpnutil"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"display verbose debug output"`
	Stdin   bool   `short:"i" long:"stdin" description:"read from stdin instead of args"`
	Args    struct {
		Hostnames []string `description:"hostnames to domainify"`
	} `positional-args:"yes"`
}

func domainify(hostname string) {
	domain, err := gpnutil.GetEntityDomain(hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
	} else {
		fmt.Println(domain)
	}
}

func runCLI(opts Options) error {
	if opts.Stdin {
		if len(opts.Args.Hostnames) > 0 {
			return fmt.Errorf("cannot specify hostnames with --stdin")
		}

		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			hostname := scanner.Text()
			domainify(hostname)
		}
	} else {
		// Read from opts.Args.Hostnames
		for _, hostname := range opts.Args.Hostnames {
			domainify(hostname)
		}
	}

	return nil
}

func main() {
	// Parse default options are HelpFlag | PrintErrors | PassDoubleDash
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		if flags.WroteHelp(err) {
			os.Exit(0)
		}

		// Does PrintErrors work? Is it not set?
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err.Error())
		parser.WriteHelp(os.Stderr)
		os.Exit(2)
	}

	err = runCLI(opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		os.Exit(2)
	}
}
