package types

type Action map[string]string

type ActionsMap []Action

type PrjktYAML struct {
	Welcome_message    string
	Error_message      string
	Success_message    string
	No_actions_message string
	Actions            ActionsMap
}

type GitHubFilesListResp []struct {
	Name string
}
