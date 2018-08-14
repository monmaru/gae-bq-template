package batch

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type client struct {
	projectID string
}

func newClient(projectID string) *client {
	return &client{projectID: projectID}
}

type queryBuilder interface {
	build(client *bigquery.Client) *bigquery.Query
}

type importQueryBuilder struct{}

func (b *importQueryBuilder) build(client *bigquery.Client) *bigquery.Query {
	q := client.Query(`SELECT word, word_count
FROM ` + "`bigquery-public-data.samples.shakespeare`" + `
WHERE corpus = @corpus
AND word_count >= @min_word_count
ORDER BY word_count DESC;`)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "corpus",
			Value: "romeoandjuliet",
		},
		{
			Name:  "min_word_count",
			Value: 250,
		},
	}
	q.Dst = client.Dataset(datasetID).Table(tableID)
	q.CreateDisposition = bigquery.CreateIfNeeded
	q.WriteDisposition = bigquery.WriteTruncate
	return q
}

func (c *client) writeQueryResults(ctx context.Context, b queryBuilder) error {
	client, err := bigquery.NewClient(ctx, c.projectID)
	if err != nil {
		return err
	}

	q := b.build(client)
	job, err := q.Run(ctx)
	if err != nil {
		return err
	}

	status, err := job.Wait(ctx)
	if err != nil {
		return err
	}

	return status.Err()
}
