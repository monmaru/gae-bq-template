package batch

import (
	"context"
	"net/http"
)

func importHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) *appError {
	b := &importQueryBuilder{}
	return doWriteQueryResults(ctx, b)
}

func doWriteQueryResults(ctx context.Context, b queryBuilder) *appError {
	c := newClient(projectID)
	if err := c.writeQueryResults(ctx, b); err != nil {
		return appErrorf(err, http.StatusInternalServerError, "failed writeQueryResults")
	}
	return nil
}
