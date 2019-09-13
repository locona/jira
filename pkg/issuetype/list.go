package issuetype

const (
	AD  = "AD"
	REC = "REC"
)

const (
	TASK    = "タスク"
	EPIC    = "エピック"
	BUG     = "バグ"
	STORY   = "ストーリー"
	SUBTASK = "サブタスク"
	QA      = "QA"
)

const (
	AD_TASK  = "10031"
	AD_EPIC  = "10032"
	AD_BUG   = "10037"
	AD_STORY = "10038"
)

const (
	REC_TASK    = "10033"
	REC_EPIC    = "10034"
	REC_BUG     = "10036"
	REC_STORY   = "10035"
	REC_SUBTASK = "10058"
	REC_QA      = "10059"
)

var AD_ISSUE_TYPE = map[string]string{
	TASK:  AD_TASK,
	EPIC:  AD_EPIC,
	BUG:   AD_BUG,
	STORY: AD_STORY,
}

var REC_ISSUE_TYPE = map[string]string{
	TASK:    REC_TASK,
	EPIC:    REC_EPIC,
	BUG:     REC_BUG,
	STORY:   REC_STORY,
	SUBTASK: REC_SUBTASK,
	QA:      REC_QA,
}

var ISSUE_TYPE = map[string]map[string]string{
	AD:  AD_ISSUE_TYPE,
	REC: REC_ISSUE_TYPE,
}

func IssueType(projectID string, _type string) string {
	return ISSUE_TYPE[projectID][_type]
}
