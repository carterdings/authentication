package dao

import (
	"testing"

	"github.com/patrickmn/go-cache"
)

// Test_Client ...
func Test_Client(t *testing.T) {
	Set("foo", "bar", cache.NoExpiration)

	v, ok := Get("foo")
	if ok {
		t.Logf("val: %s", v.(string))
	}
	Delete("foo")

	_, ok = Get("foo")
	t.Logf("ok: %v", ok)

}
