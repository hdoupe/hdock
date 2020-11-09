package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

var buildPath string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [buildPath]",
	Short: "Run a docker command mounted on some dir.",
	Long:  `Run a docker command mounted on some dir.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("either no arguments or one argument must be set")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Got buildPath", buildPath)

		if len(args) == 1 {
			buildPath = args[0]
		} else {
			buildPath = "."
		}

		fmt.Println("maybe gotta buildPath...", buildPath)

		fmt.Println("image tag", tag)

		if tag == "" {
			var dirname string
			if buildPath == "." {
				workingdir, err := os.Getwd()
				if err != nil {
					panic(err)
				}
				pieces := strings.Split(workingdir, "/")
				dirname = strings.ToLower(pieces[len(pieces)-1])
			} else {
				pieces := strings.Split(buildPath, "/")
				dirname = strings.ToLower(pieces[len(pieces)-1])
			}
			tag = dirname + ":latest"
		} else if strings.Contains(tag, ":") {
			tag = tag + ":latest"
		}

		fmt.Println("tag", tag)

		Build(buildPath, tag)

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

// Build docker image using buildPath as context.
// TODO: handle dockerignore
func Build(buildPath string, tag string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	buildContext, err := archive.TarWithOptions(buildPath, &archive.TarOptions{
		// ExcludePatterns: [],
	})
	if err != nil {
		panic(err)
	}

	buildOpts := types.ImageBuildOptions{Tags: []string{tag}}
	imageBuildResponse, err := cli.ImageBuild(context.Background(), buildContext, buildOpts)
	if err != nil {
		panic(err)
	}
	defer imageBuildResponse.Body.Close()

	type streamMsg struct {
		Stream string
	}

	scanner := bufio.NewScanner(imageBuildResponse.Body)
	for scanner.Scan() {
		var msg streamMsg
		json.Unmarshal(scanner.Bytes(), &msg)
		fmt.Print(msg.Stream)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// Run command in specified docker image
func Run(tag string, cmd string) {}

// func Exists(name string) (bool, error) {
// 	_, err := os.Stat(name)
// 	if os.IsNotExist(err) {
// 		return false, nil
// 	}
// 	return err != nil, err
// }
