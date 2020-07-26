package command

import (
	"github.com/okeyaki/neo4j-query-log-analyzer/lib/service"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	recLoader := service.NewRecordLoader()
	repBuilder := service.NewReportBuilder()
	repFormatter := service.NewReportFormatter()

	cmd := &cobra.Command{
		Use: "neo4j-query-log-analyzer",
		RunE: func(cmd *cobra.Command, args []string) error {
			recs, err := recLoader.Run()
			if err != nil {
				return errors.WithStack(err)
			}

			rep, err := repBuilder.Run(recs)
			if err != nil {
				return errors.WithStack(err)
			}

			out, err := repFormatter.Run(rep)
			if err != nil {
				return errors.WithStack(err)
			}

			cmd.Println(out)

			return nil
		},
	}

	return cmd
}
