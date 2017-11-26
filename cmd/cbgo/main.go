package main

import (
	"os"
	"path/filepath"

	"github.com/coinbase-go/client/cmd/cbgo/subcommand"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

)

var cmd = &cobra.Command{
	Use:   "cbgo",
	Short: "cbgo - coinbase go client.",
	Long:  `cbgo - An opensource coinbase go client.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("Ran successfully.......")
	},
}

func init() {
	log.SetLevel(log.DebugLevel)
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)

	cmd.AddCommand(subcommand.VersionCmd)

	viper.SetEnvPrefix("cbgo")
	viper.SetConfigName("cbgo")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("current directory unknown")
	}

	viper.AddConfigPath(dir)

	viper.AutomaticEnv()
	viper.SetConfigType("json")

	flags := cmd.Flags()
	flags.String("started-by", "uruddarraju", "Testing the command line feature using viper....")

	viper.BindPFlag("started-by", flags.Lookup("started-by"))

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading thyra configuration: %v", err)
	}
	log.Debug("Successfully loaded thyra config")
}

func main() {
	log.Debug("Starting up coinbase client command, logging for debugging purposes......")
	cmd.Execute()
	return
}