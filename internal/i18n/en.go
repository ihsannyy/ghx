package i18n

var EnMessages = Messages{
	RootShort: "",
	RootLong:  "",

	LoginShort:      "Add or authenticate a GitHub account",
	LoginLong:       "Save account credentials to ghx and set it as the active account for gh CLI and git.",
	LoginPrompt:     "Enter your GitHub Personal Access Token: ",
	LoginSuccess:    "Successfully logged in as %s (%s). Account set as active.",
	LoginErrorFetch: "Failed to fetch account info using token. Check your token.",
	LoginErrorSave:  "Failed to save account configuration.",

	AddShort: "Add a GitHub account using token",
	AddLong:  "Save account credentials to ghx using a token and set it as active account.",

	SwitchShort:           "Switch to a saved GitHub account",
	SwitchLong:            "Switch active gh CLI credentials and global git identity to the specified account.",
	SwitchSuccess:         "Switched to account '%s' (%s).",
	SwitchAccountNotFound: "Account '%s' not found. Use 'ghx list' to view available accounts.",

	ListShort:          "List all saved GitHub accounts",
	ListLong:           "Display a list of all saved accounts in ghx with active status indicator.",
	ListNoAccounts:     "No accounts saved yet. Use 'ghx login' to add an account.",
	ListHeaderStatus:   "STATUS",
	ListHeaderUsername: "USERNAME",
	ListHeaderName:     "NAME",
	ListHeaderEmail:    "EMAIL",

	RemoveShort:           "Remove a saved GitHub account",
	RemoveLong:            "Remove the specified account from ghx saved accounts.",
	RemoveSuccess:         "Account '%s' removed successfully.",
	RemoveAccountNotFound: "Account '%s' not found. Use 'ghx list' to view available accounts.",

	EmailShort:   "View or set email for current account",
	EmailLong:    "View or update the email address for the active account and global git config.",
	EmailCurrent: "Current git email for '%s': %s",
	EmailUpdated: "Updated email for '%s' to '%s'.",

	CurrentShort:    "Display currently active account",
	CurrentLong:     "Show username, name, and email of the currently active account.",
	CurrentActive:   "Active account: %s (%s) <%s>",
	CurrentNoActive: "No account currently active. Use 'ghx switch <account>' or 'ghx login'.",

	LangShort:       "View or set active language",
	LangLong:        "Switch language between English (en) and Indonesian (id).",
	LangCurrent:     "Current language: %s",
	LangUpdated:     "Language changed to %s.",
	LangUnsupported: "Unsupported language '%s'. Supported options: en, id",

	ErrGHNotInstalled:  "GitHub CLI (gh) is not installed or not in PATH.",
	ErrGitNotInstalled: "Git is not installed or not in PATH.",
	ErrConfigLoad:      "Failed to load ghx config file.",
}
