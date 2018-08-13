package batch

import (
	"bytes"
	"context"
	"net/http"
	"text/template"
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

	d := make(map[string]string)
	d["table"] = "my.table"
	const ts = "SELECT * FROM [{{.table}}]"

	qs, err := queryFromTemplate(ts, d)
	if err != nil {
		return appErrorf(err, http.StatusInternalServerError, "template error")
	}

	if err := writeQueryResults(ctx, s, qs); err != nil {
		return appErrorf(err, http.StatusInternalServerError, "failed importHandler")
	}
	return nil
}

func queryFromTemplate(ts string, d interface{}) (string, error) {
	tmpl, err := template.New("buildquery").Parse(ts)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, d); err != nil {
		return "", err
	}

	return buf.String(), nil
}
