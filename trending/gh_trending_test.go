package trending

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQuery(t *testing.T) {
	// 创建一个模拟的HTTP服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<div class="Box-row">
				<h2><a href="/repo1">Repo 1</a></h2>
				<p>Description 1</p>
				<p>Description 2</p>
				<div>
					<a>100</a>
					<a>200</a>
					<span>1</span>
					<span>2</span>
					<span>300 stars today</span>
				</div>
			</div>
		`))
	}))
	defer server.Close()

	// 创建一个 GithubTrending 实例
	gh := &GithubTrending{
		Host:        server.URL,
		RequestPath: "trending",
		Language:    "go",
		DataRange:   DAILY,
	}

	// 调用 Query 方法
	trends, err := gh.Query()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 验证返回的结果
	if len(trends) != 1 {
		t.Fatalf("expected 1 trend, got %d", len(trends))
	}

	if trends[0].Title != "Repo1" {
		t.Errorf("expected title 'Repo1', got %s", trends[0].Title)
	}

	if trends[0].Description != "Description 2" {
		t.Errorf("expected description 'Description 2', got %s", trends[0].Description)
	}

	if trends[0].Forks != "200" {
		t.Errorf("expected forks '200', got %s", trends[0].Forks)
	}

	if trends[0].Stars != "100" {
		t.Errorf("expected stars '100', got %s", trends[0].Stars)
	}

	if trends[0].StarsDay != "300" {
		t.Errorf("expected stars today '300', got %s", trends[0].StarsDay)
	}
}
