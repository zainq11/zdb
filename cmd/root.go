package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var cfgFile string

var config *Config

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Client struct {
		ServerAddress string `yaml:"address"`
	} `yaml:"client"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zdb",
	Short: "Key-value store in go",
	Long:  `A simple kv store written in go`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	//rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")
	serverCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "./kvconfig.yaml", "path to config file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func LoadConfig(cfgFile string) (config *Config) {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Could not read config file %v", err)
	}

	log.Printf("Read Config file ... %s", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&config); err != nil {
		log.Panicf("Could not unmarshall configuration data %v", err)
	}

	return
}

func initConfig() {
	config = LoadConfig(cfgFile)
}
