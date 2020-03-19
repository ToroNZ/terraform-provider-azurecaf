package caf

const (
	// ConventionCafClassic applies the CAF recommended naming convention
	ConventionCafClassic string = "cafclassic"
	// ConventionCafRandom defines the CAF random naming convention
	ConventionCafRandom string = "cafrandom"
	// ConventionRandom applies a random naming convention based on the max length of the resource
	ConventionRandom string = "random"
	// ConventionPassThrough defines the CAF random naming convention
	ConventionPassThrough string = "passthrough"
)

const (
	alphanum            string = "[^0-9A-Za-z]"
	alphanumh           string = "/[^0-9A-Za-z,-]/"
	alphanumu           string = "/[^0-9A-Za-z,_]/"
	alphanumhu          string = "/[^0-9A-Za-z,_,-]/"
	alphanumhup         string = "/[^0-9A-Za-z,_,.,-]/"
	alphanumStartletter string = "/\\A[^a-z][^0-9A-Za-z]/"
	unicode             string = "^[-\\w\\._\\(\\)]"
)

// ResourceStructure stores the CafPrefix and the MaxLength of an azure resource
type ResourceStructure struct {
	// Resource type name
	ResourceTypeName string
	// Resource prefix as defined in the Azure Cloud Adoption Framework
	CafPrefix string
	// MaxLength attribute define the maximum length of the name
	MinLength int
	// MaxLength attribute define the maximum length of the name
	MaxLength int
	// Regular expression to apply to the resource type
	RegEx string
	// the Regular expression to validate the generated string
	ValidationRegExp string
}

// Resources currently supported
var Resources = map[string]ResourceStructure{
	"aaa":  {"azure automation account", "aaa-", 6, 50, alphanumh, "^[a-zA-Z][^0-9A-Za-z,-]{5,49}$"},
	"acr":  {"azure container registry", "acr-", 5, 50, alphanum, "^[0-9A-Za-z]{5,50}$"},
	"rg":   {"resource group", "rg-", 1, 90, unicode, "^[-\\w\\._\\(\\)]{1,90}$"},
	"st":   {"storage account", "st-", 3, 24, alphanum, "^[0-9A-Za-z]{3,24}$"},
	"afw":  {"azure firewall", "afw-", 1, 80, alphanumhup, "^[a-zA-Z][^0-9A-Za-z,_,.,-]{0,79}$"},
	"asr":  {"azure site recovery", "asr-", 2, 50, alphanumh, "^[a-zA-Z][^0-9A-Za-z,-]{1,49}$"},
	"evh":  {"event hub", "evh-", 1, 50, alphanumh, "^[a-zA-Z][^0-9A-Za-z,-]{0,48}[0-9a-zA-Z]{0,1}$"},
	"kv":   {"keyvault", "kv-", 3, 24, alphanumh, "^[a-zA-Z][^0-9A-Za-z,-]{0,21}[0-9a-zA-Z]{0,1}$"},
	"la":   {"loganalytics", "la-", 4, 63, alphanumh, "^[0-9a-zA-Z][^0-9A-Za-z,-]{3,62}[0-9a-zA-Z]{0,1}$"},
	"nic":  {"network interface card", "nic-", 1, 80, alphanumhup, "^[0-9a-zA-Z][^0-9A-Za-z,_,.,-]{0,79}[0-9a-zA-Z,_]{0,1}$"},
	"nsg":  {"network security group", "nsg-", 1, 80, alphanumhup, "^[0-9a-zA-Z][^0-9A-Za-z,_,.,-]{0,79}[0-9a-zA-Z,_]{0,1}$"},
	"pip":  {"public ip address", "pip-", 1, 80, alphanumhup, "^[0-9a-zA-Z][^0-9A-Za-z,_,.,-]{0,79}[0-9a-zA-Z,_]{0,1}$"},
	"snet": {"virtual network subnet", "snet-", 1, 80, alphanumhup, "^[0-9a-zA-Z][^0-9A-Za-z,_,.,-]{0,79}[0-9a-zA-Z,_]{0,1}$"},
	"vnet": {"virtual network", "snet-", 1, 80, alphanumhup, "^[0-9a-zA-Z][^0-9A-Za-z,_,.,-]{0,79}[0-9a-zA-Z,_]{0,1}$"},
	"vmw":  {"virtual machine (windows)", "vmw-", 1, 15, alphanumh, "^[0-9a-zA-Z][^0-9A-Za-z,_,-]{0,13}[0-9a-zA-Z,_]{0,1}$"},
	"vml":  {"virtual machine (linux)", "vml-", 1, 64, alphanumh, "^[0-9a-zA-Z][^0-9A-Za-z,_,-]{0,63}[0-9a-zA-Z,_]{0,1}$"},
}
