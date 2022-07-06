package my_errors

const (
	ErrorsBasePath                  string = "failed to initialize project root directory"
	ErrorsContainerKeyAlreadyExists string = "The key is already registered in the container"
	ErrorsConfigInitFail            string = "Error initializing configuration file"
	ErrorsFuncEventAlreadyExists    string = ""
	ErrorsFuncEventNotCall          string = ""
	ErrorsFuncEventNotRegister      string = ""
	ErrorsConfigYamlNotExists       string = ""
	ErrorsPublicNotExists           string = ""
	ErrorsSoftLinkDeleteFail        string = ""
	ErrorsSoftLinkCreateFail        string = ""
	ErrorsStorageLogsNotExists      string = ""
)
