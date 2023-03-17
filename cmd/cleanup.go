package cmd

import (
	"github.com/spf13/cobra"
)

type cleanupOptions struct {
	bucket string
}

func newCleanupCmd(globalOpts *globalOptions) *cobra.Command {
	opts := &cleanupOptions{}
	cmd := &cobra.Command{
		Use:           "cleanup",
		Short:         "",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfni, err := newCFNI(globalOpts)
			if err != nil {
				return err
			}

			return cfni.Cleanup(opts.bucket)
		},
	}

	cmd.Flags().StringVarP(&opts.bucket, "bucket", "b", "", "bucket name (required)")
	_ = cmd.MarkPersistentFlagRequired("bucket")

	return cmd
}
