package models

import (
	"net"
	"net/netip"
	"time"
)

// printer is a set of methods for printers to implement.
// Printers should NOT modify any existing data nor do any calculations.
// They should only perform visual operations on given data.
type Printer interface {
	// printStart should print the first message, after the program starts.
	// This message is printed only once, at the very beginning.
	printStart(hostname string, port uint16)

	// printProbeSuccess should print a message after each successful probe.
	// hostname could be empty, meaning it's pinging an address.
	// streak is the number of successful consecutive probes.
	printProbeSuccess(localAddr string, UserInput UserInput, streak uint, rtt float32)

	// printProbeFail should print a message after each failed probe.
	// hostname could be empty, meaning it's pinging an address.
	// streak is the number of successful consecutive probes.
	printProbeFail(UserInput UserInput, streak uint)

	// printRetryingToResolve should print a message with the hostname
	// it is trying to resolve an ip for.
	//
	// This is only being printed when the -r flag is applied.
	printRetryingToResolve(hostname string)

	// printTotalDownTime should print a downtime duration.
	//
	// This is being called when host was unavailable for some time
	// but the latest probe was successful (became available).
	printTotalDownTime(downtime time.Duration)

	// printStatistics should print a message with
	// helpful statistics information.
	//
	// This is being called on exit and when user hits "Enter".
	printStatistics(s tcping)

	// printVersion should print the current version.
	printVersion()

	// printInfo should a message, which is not directly related
	// to the pinging and serves as a helpful information.
	//
	// Example of such: new version with -u flag.
	printInfo(format string, args ...any)

	// printError should print an error message.
	// Printer should also apply \n to the given string, if needed.
	printError(format string, args ...any)
}

type UserInput struct {
	ip                       netip.Addr
	hostname                 string
	networkInterface         NetworkInterface
	retryHostnameLookupAfter uint // Retry resolving target's hostname after a certain number of failed requests
	probesBeforeQuit         uint
	timeout                  time.Duration
	intervalBetweenProbes    time.Duration
	port                     uint16
	useIPv4                  bool
	useIPv6                  bool
	shouldRetryResolve       bool
	showFailuresOnly         bool
	showLocalAddress         bool
}

type genericUserInputArgs struct {
	retryResolve         *uint
	probesBeforeQuit     *uint
	timeout              *float64
	secondsBetweenProbes *float64
	intName              *string
	showFailuresOnly     *bool
	showLocalAddress     *bool
	args                 []string
}

type NetworkInterface struct {
	remoteAddr *net.TCPAddr
	dialer     net.Dialer
	use        bool
}

type LongestTime struct {
	start    time.Time
	end      time.Time
	duration time.Duration
}

type rttResult struct {
	min        float32
	max        float32
	average    float32
	hasResults bool
}

type HostnameChange struct {
	Addr netip.Addr `json:"addr,omitempty"`
	When time.Time  `json:"when,omitempty"`
}
