package command

import (
	log "github.com/Sirupsen/logrus"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	// ConfigFile The path to the configuration file.
	ConfigFile string

	// CacheDir The path to the cache directory for skel templates.
	CacheDir string

	// Verbose The global value if verbosity was enabled.
	Verbose bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "skel",
	Short: "A brief description of your application",
}

// Execute the entrypoint for executing the command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&ConfigFile, "config", "", "config file (default is $HOME/.skel.toml)")
	RootCmd.PersistentFlags().StringVar(&CacheDir, "cacheDir", "", "Filesystem path to the cache directory.")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Use verbose output.")
	//RootCmd.PersistentFlags().BoolVar(&ignoreCache, "ignoreCache", false, "Ignores the cache directory.")

	log.SetOutput(os.Stderr)

	if Verbose == true {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	}

	viper.AutomaticEnv()
	viper.SetConfigType("toml")
	viper.SetConfigName("skel")
	viper.AddConfigPath(".")

	dir, err := homedir.Dir()
	if err != nil {
		log.Warn("Unable to detect user home directory.")
	} else {
		log.Debug("Using home directory: ", dir)
		viper.AddConfigPath(dir)

		log.Debug("Using cache directory: ", filepath.Join(dir, ".skel"))
		viper.SetDefault("CacheDir", filepath.Join(dir, ".skel"))
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Failed reading config file: ", viper.ConfigFileUsed())
		os.Exit(1)

	}

	log.Debug("Using config file: ", viper.ConfigFileUsed())
}
