package auth

import (
	"testing"
	"time"
)

func TestIssuerIssueAndParse(t *testing.T) {
	issuer := &Issuer{Secret: []byte("test-secret"), TTL: time.Hour}
	token, err := issuer.Issue(42, "member")
	if err != nil {
		t.Fatalf("Issue() error = %v", err)
	}

	id, role, err := issuer.Parse(token)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if id != 42 || role != "member" {
		t.Fatalf("Parse() = (%d, %q), want (42, %q)", id, role, "member")
	}
}
