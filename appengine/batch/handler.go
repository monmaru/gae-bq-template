package batch

import (
	"context"
	"net/http"
)

const (
	projectID = "__PROJECT_ID__"
	datasetID = "__DATASET_ID__"
	tableID   = "__TABLE_ID__"
)

func importHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) *appError {
	s := setting{
		projectID: projectID,
		datasetID: datasetID,
		tableID:   tableID,
	}

	err := writeQueryResults(ctx, s, func() (string, error) {
		return "your query", nil
	})

	if err != nil {
		return appErrorf(err, http.StatusInternalServerError, "failed importHandler")
	}
	return nil
}
