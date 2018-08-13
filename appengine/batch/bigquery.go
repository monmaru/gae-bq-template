package batch

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type setting struct {
	projectID, datasetID, tableID string
}

func writeQueryResults(ctx context.Context, s setting, qs string) error {
	client, err := bigquery.NewClient(ctx, s.projectID)
	if err != nil {
		return err
	}

	// query settings & run
	q := client.Query(qs)
	q.Dst = client.Dataset(s.datasetID).Table(s.tableID)
	q.CreateDisposition = bigquery.CreateIfNeeded
	q.WriteDisposition = bigquery.WriteTruncate

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
