package errcode

var (
	Success          = NewError(0, "Success")
	Fail             = NewError(10000000, "Internal Error")
	InvalidParams    = NewError(10000001, "Invalid Parameters")
	Unautorized      = NewError(10000002, "Unauthorized")
	NotFound         = NewError(10000003, "Not Found")
	Unknown          = NewError(10000004, "Unknown")
	DeadlineExceeded = NewError(10000005, "Deadline Exceeded")
	AccessDenied     = NewError(10000006, "Access Denied")
	LimitExceed      = NewError(10000007, "Limit Exceed")
	MethodNotAllow   = NewError(10000008, "Method Not Allow")
)
