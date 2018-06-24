package compiler

import (
	"flag"

	_ "github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/koki/concerto/pkg/manager"
)

var (
	RootCmd = &cobra.Command{
		Use:   "concerto",
		Short: "DevOps Programming Language",
		Long: `DevOps Programming Language

Full documentation available at https://docs.koki.io/concerto
`,
		RunE: func(c *cobra.Command, args []string) error {
			return concerto(c, args)
		},
		SilenceUsage: true,
		Example: `

  # Compile a concerto program
  concerto redis.concerto
`,
	}
)

func init() {
	// parse the go default flagset to get flags for glog and other packages in future
	RootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	// defaulting this to true so that logs are printed to console
	flag.Set("logtostderr", "true")

	//suppress the incorrect prefix in glog output
	flag.CommandLine.Parse([]string{})

	RootCmd.AddCommand(versionCmd)
}

func concerto(c *cobra.Command, args []string) error {
	return manager.Compile(args)
}
