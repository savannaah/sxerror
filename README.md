# Savannaah error
This package extends the error interface to include code, title and description of the error.

## Error code description
Error code has been designed to pin point the what and where the error has occurred. To simply the error code can be broken down as follows:
    
    regex error code pattern: SX[DIWSB]XXX[0-9]{6}$
    
There are total of 12 characters in the error code: 6 alphabets and 6 digits

	First two character is reserved for SAVANNAAH. i.e. SX
	Third character is for error level:
	D = debug
	I = info
    W = warning
    S = severe - internal errors
    B = business - error made by users
    
    Last three characters is used to identify the micro service
    Last six digits is used to identify specific kind of error

Characters to identify micro service:

| Service | Characters |
| :------ | :--------- |
| session-server | SES |
| master-server | MAS |
| relatedrecord-server | REL |
| profile-server | PRO |
| organisation-server | ORG |

Digits to identify the exact error:

| Digits | Error title | Error description |
| :----- | :---------- | :---------------- |
| 000001 | failed to connect to the database | NA |
| 000002 | failed to ping the database | NA |
| 000003 | failed to query the database | NA |

| Digits | Error title | Error description |
| :----- | :---------- | :---------------- |
| 001001 | database returned no data | NA |
| 001002 | unacceptable parameter used in the function call | Generally, this happens if the parameter supplied to call a function is not acceptable. |
| 001003 | failed to marshal golang data structure into json | This happens if the golang data structure member type or value are not supported. |
| 001004 | failed to unmarshal json into golang data structure | This happens if the type or value in json are not supported by golang data structure. |
| 001005 | token seems to be corrupted | This happens if the token is corrupted and is not acceptable. |
| 001006 | base64 encoded string seems to be corrupted | NA |

| Digits | Error title | Error description |
| :----- | :---------- | :---------------- |
| 100001 | $record does not exits in the database | This happens if the database cannot find the record. |
| 100002 | unauthorised to access the requested data | This happens if the user does not have read access to the query or record(s). |
| 100003 | unauthorised to update the record or records | NA |
| 100004 | unauthorised to activate the record or records | NA |
| 100005 | unauthorised to edit the record or records | NA |
| 100006 | unauthorised to deactivate the record or records | NA |
| 100007 | unauthorised to delete the record or records | NA |
| 100008 | $record is not ready to be $action | This happens if the record is not in correct status to perform the requested action. |
| 100009 | no record or records are selected | This happens if no record ids is supplied |
| 100010 | one ore more mandatory fields are missing | NA |
| 100011 | record already exists with $field $value | This happens if unique field value already exits |
| 100012 | character limit is violated for the $field | NA