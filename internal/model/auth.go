package model

type JWTValidateInput struct {
	Token string
	Sig   []byte
}
type RefreshTokenIssueInput struct {
	UserName        string
	IssueTime       string
	RefreshTokenSig []byte
}

type RefreshTokenIssueOutput struct {
	Token string
}
type AccessTokenIssueInput struct {
	UserName       string
	IssueTime      string
	AccessTokenSig []byte
}

type AccessTokenIssueOutput struct {
	Token string
}

type AccessTokenQueryInput struct {
	Token          string
	AccessTokenSig []byte
}
type AccessTokenQueryOutput struct {
	UserName  string
	IssueTime string
}
type RefreshTokenQueryInput struct {
	Token           string
	RefreshTokenSig []byte
}
type RefreshTokenQueryOutput struct {
	UserName  string
	IssueTime string
}
