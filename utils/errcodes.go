package wzlib_utils

const (
	EX_OK          = 0  // successful termination
	EX_GENERIC     = 1  // generic error
	EX_USAGE       = 64 // command line usage error
	EX_NOUSER      = 67 // addressee unknown
	EX_UNAVAILABLE = 69 // service unavailable
	EX_SOFTWARE    = 70 // internal software error
	EX_CANTCREAT   = 73 // can't create (user) output file
	EX_TEMPFAIL    = 75 // temp failure; user is invited to retry
	EX_NOPERM      = 77 // permission denied
)
