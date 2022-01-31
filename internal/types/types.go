package types

type Actions []map[string]string

type PrjktYAML struct {
	Welcome_message string
	Error_message   string
	Success_message string
	Actions         Actions
}
