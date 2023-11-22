package http_test

import (
	"context"
	"github.com/WalterPaes/rest-api-http-client/pkg/http"
	httpClient "net/http"
	"testing"
	"time"
)

func TestClient_Post(t *testing.T) {
	c := http.New(
		http.WithBaseURL("http://localhost:3000"),
		http.WithTimeout(time.Second*5),
	)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := map[string]any{
		"id":       "08f807b7-87f2-4990-bc61-56d69627147e",
		"replicas": 1,
		"image":    "golang",
		"labels": map[string]string{
			"monitor-label":  "service-y",
			"deployment-tag": "xpto",
		},
		"ports": []map[string]any{
			{
				"name": "http",
				"port": 65,
			},
		},
	}

	r, e := c.Post(context.Background(), "deployments", headers, body)

	if e != nil {
		t.Error(e)
	}

	if r.StatusCode != httpClient.StatusCreated {
		t.Errorf("Was expected '%d' but got '%d'", httpClient.StatusCreated, r.StatusCode)
	}
}

func TestClient_Get(t *testing.T) {
	c := http.New(
		http.WithBaseURL("http://localhost:3000"),
		http.WithTimeout(time.Second*5),
	)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	r, e := c.Get(context.Background(), "deployments/08f807b7-87f2-4990-bc61-56d69627147e", headers)

	if e != nil {
		t.Error(e)
	}

	if r.StatusCode != httpClient.StatusOK {
		t.Errorf("Was expected '%d' but got '%d'", httpClient.StatusOK, r.StatusCode)
	}
}

func TestClient_Delete(t *testing.T) {
	c := http.New(
		http.WithBaseURL("http://localhost:3000"),
		http.WithTimeout(time.Second*5),
	)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	r, e := c.Delete(context.Background(), "deployments/08f807b7-87f2-4990-bc61-56d69627147e", headers)

	if e != nil {
		t.Error(e)
	}

	if r.StatusCode != httpClient.StatusNoContent {
		t.Errorf("Was expected '%d' but got '%d'", httpClient.StatusNoContent, r.StatusCode)
	}
}
