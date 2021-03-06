package main

import (
	"path/filepath"

	"github.com/mattermost/mattermost-load-test/ltops"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var loadTest = &cobra.Command{
	Use:   "loadtest -- [args...]",
	Short: "Runs a mattermost-load-test command against the given cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		loadtestOptions := &ltops.LoadTestOptions{}
		clusterName, _ := cmd.Flags().GetString("cluster")
		loadtestOptions.ForceBulkLoad, _ = cmd.Flags().GetBool("force-bulk-load")

		//config, _ := cmd.Flags().GetString("config")

		workingDir, err := defaultWorkingDirectory()
		if err != nil {
			return err
		}

		cluster, err := LoadCluster(filepath.Join(workingDir, clusterName))
		if err != nil {
			return errors.Wrap(err, "Couldn't load cluster")
		}

		return cluster.Loadtest(loadtestOptions)
	},
}

func init() {
	loadTest.Flags().StringP("cluster", "c", "", "cluster name (required)")
	loadTest.MarkFlagRequired("cluster")

	loadTest.Flags().BoolP("force-bulk-load", "", false, "force bulk load even if bulk loading already complete")

	// TODO: Implement
	//loadTest.Flags().StringP("config", "f", "", "a config file to use instead of the default (the ConnectionConfiguration section is mostly ignored)")

	rootCmd.AddCommand(loadTest)
}
