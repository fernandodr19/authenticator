package vos

type (
	AccessToken  string
	RefreshToken string
)

func (a AccessToken) String() string {
	return string(a)
}

func (r RefreshToken) String() string {
	return string(r)
}

type Tokens struct {
	AccessToken  AccessToken
	RefreshToken RefreshToken
}
