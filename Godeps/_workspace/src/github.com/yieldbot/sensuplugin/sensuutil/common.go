// Library for all global variables used by the Yieldbot
// Infrastructure teams in sensu
//
// LICENSE:
//   Copyright 2015 Yieldbot. <devops@yieldbot.com>
//   Released under the MIT License; see LICENSE
//   for details.

package sensuutil

// MonitoringErrorCodes provides a standard set of error codes to use.
// Please use the below codes instead of random non-zero so that monitoring can
// utilize existing maps for alerting and help avoid unnecessary noise.
var MonitoringErrorCodes = map[string]int{
	"GENERALGOLANGERROR": 129, // internal script error
	"CONFIGERROR":        127, // unix config error, not enough parms, etc
	"PERMISSIONERROR":    126, // not executable, etc
	"RUNTIMEERROR":       42,  // self explantory
	"DEBUG":              37,  // exit w/ debugging output
	"OK":                 0,   // “I'd far rather be happy than right any day.”
	"WARNING":            1,   // this kinda sucks but don't get out of bed to deal with it
	"CRITICAL":           2,   // “The ships hung in the sky in much the same way that bricks don't.”
	"UNKNOWN":            3,   // Would it save you a lot of time if I just gave up and went mad now?
}
