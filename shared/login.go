package shared

type Login struct {
	UID      int
	Username string
	Passwd   string
	Result   string
}

func (t *Login) ModelId() string {
	return ""
}

func (t *Login) RootURL() string {
	return "/api/login"
}
