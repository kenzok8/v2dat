package unpack

import (
	"github.com/kenzok8/v2dat/cmd"
	"github.com/spf13/cobra"
)

var unpack = &cobra.Command{
	Use:   "unpack",
	Short: "unpack geosite and geoip to text files.",
}

func init() {
	unpack.AddCommand(newGeoSiteCmd(), newGeoIPCmd())
	cmd.RootCmd.AddCommand(unpack)
}
