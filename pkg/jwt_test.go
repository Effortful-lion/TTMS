package pkg

import (
	"testing"
)

func TestGenToken(t *testing.T) {
	user_id := 1
	auth := "admin"
	username := "lion"
	token, err := GenToken(user_id, auth, username)
	if err != nil {
		t.Errorf("GenToken failed: %v", err)
	}
	t.Logf("token: %s", token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwiYXV0aG9yaXR5IjoibGlvbiIsImlzcyI6IjMyMk1vdmllIiwiZXhwIjoxNzQ3NzQ5Mjc3fQ.eqLRtXjuNF1UEAY7aLNercx_xYP7WIb0IGFAPW5jNxk"
	claims, err := ParseToken(token)
	if err!= nil {
		t.Errorf("ParseToken failed: %v", err)
	}
	t.Logf("claims: %v", claims)
}