/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/mark-ignacio/http2sqlite/ent"
	"github.com/mark-ignacio/http2sqlite/internal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "http2sqlite",
	Short: "Starts a server that logs HTTP requests to SQLite.",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			dbPath           = viper.GetString("db")
			address          = viper.GetString("address")
			code             = viper.GetInt("response-code")
			filter           = viper.GetString("filter")
			requestsToStderr = viper.GetBool("requests-to-stderr")
			ctx              = context.Background()
			db               = mustOpenDB(dbPath)
			filterProgram    cel.Program
		)
		defer db.Close()
		if filter != "" {
			program, err := internal.CreateCELProgram(filter)
			if err != nil {
				log.Fatal().Err(err).Str("filter", filter).Msg("error evaluating filter")
			}
			filterProgram = program
		} else {
			log.Info().Msg("no filter specified, will log everything")
		}
		log.Info().
			Str("path", dbPath).
			Str("address", address).
			Msg("now listening")
		err := http.ListenAndServe(address, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Err(err).Msg("error reading from request body")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if filterProgram != nil {
				val, _, err := filterProgram.Eval(map[string]interface{}{
					"host":           r.Host,
					"path":           r.URL.Path,
					"method":         r.Method,
					"source_address": r.RemoteAddr,
					"header":         underscoredHeader(r.Header),
				})
				if err != nil {
					log.Fatal().Err(err).Msg("error evaluating CEL expression")
				}
				iface, err := val.ConvertToNative(reflect.TypeOf(true))
				if err != nil {
					log.Fatal().Err(err).Msg("error converting CEL expression result to native bool")
				}
				if !iface.(bool) {
					log.Info().Msg("ignoring request")
					w.WriteHeader(code)
					return
				}
			}
			saved, err := db.HTTPRequest.Create().
				SetHost(r.Host).
				SetPath(r.URL.Path).
				SetMethod(r.Method).
				SetHeader(r.Header).
				SetBody(body).
				Save(ctx)
			if err != nil {
				log.Err(err).Msg("error saving request")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if requestsToStderr {
				log.Info().Interface("saved", saved).Msg("saved request")
			} else {
				log.Info().Int("id", saved.ID).Msg("saved request")
			}
			w.WriteHeader(code)
		}))
		if err != nil {
			log.Fatal().Err(err).Msg("error listening")
		}
	},
}

func underscoredHeader(header http.Header) map[string]string {
	output := map[string]string{}
	replacer := strings.NewReplacer("-", "_")
	for k, vs := range header {
		output[replacer.Replace(k)] = vs[0]
	}
	return output
}

func mustOpenDB(path string) *ent.Client {
	var (
		ctx     = context.Background()
		connStr = fmt.Sprintf("file:%s?cache=shared&_fk=1", path)
	)
	client, err := ent.Open("sqlite3", connStr)
	if err != nil {
		panic(err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		panic(fmt.Errorf("error migrating schema resources: %w", err))
	}
	return client
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.http2sqlite.yaml)")
	flags := rootCmd.Flags()
	flags.String("db", "db.sqlite3", "path to SQLite database")
	flags.String("address", ":4676", "HTTP listen address")
	flags.String("filter", "", "CEL filter")
	flags.Bool("requests-to-stderr", false, "Log matching request bodies to stderr")
	flags.Int("response-code", 200, "Server-wide HTTP response code, if nothing goes wrong.")
	viper.BindPFlags(flags)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".http2sqlite" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".http2sqlite")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
