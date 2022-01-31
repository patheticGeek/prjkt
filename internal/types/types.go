package types

type Action map[string]string

type ActionsMap []Action

type PrjktYAML struct {
	Welcome_message string
	Error_message   string
	Success_message string
	Actions         ActionsMap
}
