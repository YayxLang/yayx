package main

type Action struct {
	actionType string   // Action Type
	body       []string // ActIon Body
}

func createAction(actionType string, body []string) Action {
	var action Action

	action.actionType = actionType
	action.body = body

	return action
}
