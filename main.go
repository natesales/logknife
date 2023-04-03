package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var version = "dev" // Set by build process

// Flags
var (
	noIPs            bool
	noUUIDs          bool
	redact           bool
	entropyThreshold int
	redactionPattern string
	verbose          bool
	showVersion      bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&noIPs, "no-ips", "i", false, "Don't match IP addresses")
	rootCmd.PersistentFlags().BoolVarP(&noUUIDs, "no-uuids", "u", false, "Don't match UUIDs")
	rootCmd.PersistentFlags().BoolVarP(&redact, "redact", "r", false, "Replace matches with redaction pattern instead of innocuous substitutes")
	rootCmd.PersistentFlags().IntVarP(&entropyThreshold, "entropy-threshold", "t", 50, "Minimum entropy threshold")
	rootCmd.PersistentFlags().StringVar(&redactionPattern, "redaction-pattern", "********", "Redaction pattern")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "V", false, "Show version")
}

var rootCmd = &cobra.Command{
	Use:   "logknife",
	Short: "Remove sensitive information from your logs.",
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.SetLevel(log.DebugLevel)
		}

		if showVersion {
			fmt.Printf("logknife version %s https://github.com/natesales/logknife", version)
			os.Exit(0)
		}

		if len(args) == 0 {
			log.Fatal("No input file specified, use - for stdin")
		}
		file := args[0]

		log.Debugf("Redact: %s, IPs: %v, UUIDs: %v, File: %s", redact, !noUUIDs, !noUUIDs, file)

		rand.Seed(time.Now().UnixNano())

		// Open input file
		var r io.Reader
		if file == "-" {
			r = os.Stdin
		} else {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			r = f
		}
		Knife(r, !noIPs, !noUUIDs)
	},
}

func main() {
	rootCmd.Execute()
}
