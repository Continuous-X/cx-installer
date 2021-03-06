/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	error2 "cx/pkg/error"
	"cx/pkg/versions"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/version"
	"runtime"
)

// versionCmd represents the versions command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
/*		message := fmt.Sprintf("use %s", cmd.Name())
		output.PrintCliInfo(message)
*/
		versionInfo := get()
		marshalled, err := json.MarshalIndent(&versionInfo, "", "  ")
		if err != nil {
			error2.FailHandleCommand(err)
		}
		fmt.Println(string(marshalled))

		//printVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printVersion() {
	fmt.Printf("Operating System: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Version: %s\n", versions.VersionFromGit)
	fmt.Printf("BuildDate: %s\n", versions.BuildDate)
	fmt.Printf("SHA: %s\n", versions.CommitFromGit)
}

func get() version.Info {
	return version.Info{
		Major:        versions.MajorFromGit,
		Minor:        versions.MinorFromGit,
		GitCommit:    versions.CommitFromGit,
		GitVersion:   versions.VersionFromGit,
		GitTreeState: versions.GitTreeState,
		BuildDate:    versions.BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
