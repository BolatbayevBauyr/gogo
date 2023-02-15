package repository

import "testing"

func TestNewPostRepository(t *testing.T) {
	query := `select * from posts where title like '%` + "hello" + `%'`
	t.Logf("query : %s", query)
}
