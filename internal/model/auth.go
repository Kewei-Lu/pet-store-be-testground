package model

type JwtValidateInput struct {
	Token string
	Sig   []byte
}
type JwtIssueInput struct {
	UserName  string
	IssueTime string
	Sig       []byte
}

type JwtIssueOutPut struct {
	Token string
}

type JwtQueryInput struct {
	Token string
	Sig   []byte
}
type JwtQueryOutput struct {
	UserName  string
	IssueTime string
}
