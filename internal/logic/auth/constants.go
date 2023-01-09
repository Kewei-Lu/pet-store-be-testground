package auth

var (
	JWT_SIG []byte = []byte("NSWE-PET-STORE")
)

type JWT_PAYLOAD struct {
	UserName  string
	IssueTime string
}
