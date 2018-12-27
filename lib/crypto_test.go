package lib

import "testing"

func TestEncryptSha256(t *testing.T) {
	actual := EncryptSha256("hogehoge")
	expected := "4c716d4cf211c7b7d2f3233c941771ad0507ea5bacf93b492766aa41ae9f720d"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
