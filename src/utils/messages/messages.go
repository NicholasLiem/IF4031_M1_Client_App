package messages

/*
*
Error messages
*/
const (
	FailToParseID        = "Failed to parse data ID"
	InvalidRequestData   = "Failed to parse request data"
	FailToCreateData     = "Failed to create new data"
	FailToGetData        = "Failed to get data"
	FailToDeleteData     = "Failed to delete a data"
	FailToUpdateData     = "Failed to update data"
	UnsuccessfulLogin    = "Login attempt failed"
	FailToRegister       = "Registration attempt failed"
	JWTClaimError        = "JWT claim error"
	AllFieldMustBeFilled = "All field must be filled"
	AlreadyLoggedIn      = "Already logged in"
	FailToParseCookie    = "Fail to parse Cookie"
	SessionExpired       = "The session is already expired"
	FailToUnMarshalData  = "Fail to unmarshal data"
	FailToEncodeCookie   = "Fail to encode cookie"
)

/*
*
Success messages
*/
const (
	SuccessfulDataObtain   = "Successfully obtained data"
	SuccessfulDataCreation = "Successfully created a new"
	SuccessfulDataDeletion = "Successfully deleted a new"
	SuccessfulDataUpdate   = "Successfully updated a data"
	SuccessfulLogin        = "Successfully logged in"
	SuccessfulRegister     = "Successfully registered a user"
	SuccessfulLogout       = "Successfully logged out"
)
