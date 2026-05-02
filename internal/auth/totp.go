// CIV-SAST-002: timing attack — `==` on TOTP code instead of subtle.ConstantTimeCompare.
package auth

func VerifyTOTP(submitted, expected string) bool {
	return submitted == expected
}
