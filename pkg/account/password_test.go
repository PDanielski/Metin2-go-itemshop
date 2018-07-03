package account

import "testing"

func TestNewPasswordWithPlainPassword(t *testing.T) {
	password := NewPassword("ciao123", false)
	if password != Password("*E68E43CEF23812E87E81EDC9F5945C55DB24416D") {
		t.Errorf("The password must be '*E68E43CEF23812E87E81EDC9F5945C55DB24416D'\nBut is '%s'\n", password)
	}
}

func TestNewPasswordWithEncryptedPassword(t *testing.T) {
	password := NewPassword("*E68E43CEF23812E87E81EDC9F5945C55DB24416D", true)
	if password != Password("*E68E43CEF23812E87E81EDC9F5945C55DB24416D") {
		t.Error("The password must be '*E68E43CEF23812E87E81EDC9F5945C55DB24416D'")
	}
}
