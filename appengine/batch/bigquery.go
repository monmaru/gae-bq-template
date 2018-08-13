package batch

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type setting struct {
	projectID, datasetID, tableID string
}

type buildQueryFunc func() (string, error)

func writeQueryResults(ctx context.Context, s setting, qs string) error {
	client, err := bigquery.NewClient(ctx, s.projectID)
	if err != nil {
		return err
	}

	// query settings & run
	q := client.Query(qs)
	q.Dst = client.Dataset(s.datasetID).Table(s.tableID)
	q.CreateDisposition = "CREATE_IF_NEEDED"
	q.WriteDisposition = "WRITE_TRUNCATE"

	job, err := q.Run(ctx)
	if err != nil {
		return err
	}

	// Wait until async querying is done.
	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}

	return status.Err()
}
