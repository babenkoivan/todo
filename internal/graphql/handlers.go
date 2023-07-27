package graphql

import (
	"context"
	"encoding/json"
	"github.com/babenkoivan/todo/internal/users"
	"github.com/graphql-go/graphql"
	"net/http"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operationName"`
	Variables map[string]interface{} `json:"variables"`
}

func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	var p postData
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if token := r.Header.Get("Authorization"); token != "" {
		if user, err := users.Authenticate(token); err == nil {
			ctx = setUser(ctx, user)
		}
	}

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		Context:        ctx,
	})

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(result); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
