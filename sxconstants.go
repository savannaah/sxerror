package sxerror

var errorTitle = map[string]string {
	"000001": "failed to connect to the database",
	"000002": "failed to ping the database",
	"000003": "failed to query the database",

	"001001": "database returned no data",
	"001002": "unacceptable parameter or parameters for the function call",
	"001003": "failed to marshal golang data structure into json",
	"001004": "failed to unmarshal json into golang data structure",
	"001005": "token seems to be corrupted",
	"001006": "base64 encoded string seems to be corrupted",

	"100001": "record does not exits in the database",
	"100002": "unauthorised to create a record",
	"100003": "unauthorised to access the requested data",
	"100004": "unauthorised to update the record",
	"100005": "unauthorised to activate the record",
	"100006": "unauthorised to edit the record",
	"100007": "unauthorised to deactivate the record",
	"100008": "unauthorised to void the record",
	"100009": "unauthorised to delete the record or records",
	"100010": "record cannot be activated",
	"100011": "record cannot be updated",
	"100012": "record cannot be edited",
	"100013": "record cannot be deactivated",
	"100014": "record cannot be voided",
	"100015": "one or more records cannot be deleted",
	"100016": "no record or records are selected",
	"100017": "mandatory field violation",
	"100018": "unique field violation",
	"100019": "character limit violation",
}