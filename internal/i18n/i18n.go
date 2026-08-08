package i18n

type Messages struct {
	RootShort string
	RootLong  string

	LoginShort      string
	LoginLong       string
	LoginPrompt     string
	LoginSuccess    string
	LoginErrorFetch string
	LoginErrorSave  string
	AddShort        string
	AddLong         string

	SwitchShort           string
	SwitchLong            string
	SwitchSuccess         string
	SwitchAccountNotFound string

	ListShort          string
	ListLong           string
	ListNoAccounts     string
	ListHeaderStatus   string
	ListHeaderUsername string
	ListHeaderName     string
	ListHeaderEmail    string

	RemoveShort           string
	RemoveLong            string
	RemoveSuccess         string
	RemoveAccountNotFound string

	EmailShort   string
	EmailLong    string
	EmailCurrent string
	EmailUpdated string

	CurrentShort    string
	CurrentLong     string
	CurrentActive   string
	CurrentNoActive string

	LangShort       string
	LangLong        string
	LangCurrent     string
	LangUpdated     string
	LangUnsupported string

	DoctorShort         string
	DoctorLong          string
	DoctorHeader        string
	DoctorGHInstalled   string
	DoctorGHNotFound    string
	DoctorGitInstalled  string
	DoctorGitNotFound   string
	DoctorConfigValid   string
	DoctorConfigError   string
	DoctorActiveAccount string
	DoctorNoActive      string
	DoctorGHMatch       string
	DoctorGHMismatch    string
	DoctorGitMatch      string
	DoctorGitMismatch   string
	DoctorSSHMatch      string
	DoctorSSHMismatch   string
	DoctorSSHNotConfig  string
	DoctorPassed        string
	DoctorIssues        string

	ErrGHNotInstalled  string
	ErrGitNotInstalled string
	ErrConfigLoad      string
}

var (
	currentLang = "en"
	messages    map[string]Messages
)

func init() {
	messages = map[string]Messages{
		"en": EnMessages,
		"id": IdMessages,
	}
}

func SetLanguage(lang string) {
	if _, ok := messages[lang]; ok {
		currentLang = lang
	} else {
		currentLang = "en"
	}
}

func GetLanguage() string {
	return currentLang
}

func T() Messages {
	if msg, ok := messages[currentLang]; ok {
		return msg
	}
	return EnMessages
}
