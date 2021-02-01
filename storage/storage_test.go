package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	// keep track of original function
	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()

	var notifiedUser, notificationMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notificationMsg = user, msg
	}

	const user = "joe@example.com"
	usage[user] = 980000000 // simulate 980 MB used

	CheckQuota(user)
	if notifiedUser == "" && notificationMsg == "" {
		t.Errorf("notifyUser() not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notificationMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, want substring %s",
			notificationMsg, wantSubstring)
	}
}
