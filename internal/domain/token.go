package domain

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	UserID  string
	TokenID string
	Type    string
}
