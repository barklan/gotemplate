package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (c *PublicCtrl) helloHandler(w http.ResponseWriter, r *http.Request) { //nolint:unparam
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := c.db.Exec(
		ctx, `
		CREATE TABLE IF NOT EXISTS client(
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			active BOOLEAN NOT NULL DEFAULT TRUE,
			name VARCHAR (100) NOT NULL
		);
		`,
	)
	if err != nil {
		c.internalError(w, "failed to exec query", err)

		return
	}

	resp, err := json.Marshal(map[string]string{"hello": "world"})
	if err != nil {
		http.Error(w, "failed to marshal response", 500)

		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
