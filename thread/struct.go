package thread

type thread struct {
	Variables threadVariables `json:"Variables"`
}

type threadVariables struct {
	Postlist []threadVariablesPostlist `json:"postlist"`
}

type threadVariablesPostlist struct {
	Adminid      string `json:"adminid"`
	Anonymous    string `json:"anonymous"`
	Attachment   string `json:"attachment"`
	Author       string `json:"author"`
	Authorid     string `json:"authorid"`
	Dateline     string `json:"dateline"`
	Dbdateline   string `json:"dbdateline"`
	First        string `json:"first"`
	Groupiconid  string `json:"groupiconid"`
	Groupid      string `json:"groupid"`
	Memberstatus string `json:"memberstatus"`
	Message      string `json:"message"`
	Number       string `json:"number"`
	Pid          string `json:"pid"`
	Position     string `json:"position"`
	Replycredit  string `json:"replycredit"`
	Status       string `json:"status"`
	Tid          string `json:"tid"`
	Username     string `json:"username"`
}
