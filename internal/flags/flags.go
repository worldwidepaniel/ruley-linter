package flags

import (
	"flag"
	"fmt"
	"os"
)

func Init() (*string, *string) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <input-file>\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Arguments:")
		fmt.Fprintln(os.Stderr, "  <input-file>  The file you want the program to parse")
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}
	configPath := flag.String("config", ".ruley.config.json", "Path to ruley config json file")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Error: input file is required")
		flag.Usage()
		os.Exit(1)
	}

	return configPath, &args[0]
}
