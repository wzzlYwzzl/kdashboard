package volume

import (
	"os"
	"testing"

	"github.com/kubernetes/dashboard/resource/user"
)

func TestCreateRealPath(t *testing.T) {
	cases := []VolumeRelatedInfo{
		{
			U: user.User{
				Name:     "test-name",
				Password: "test-password",
			},
			CetusfsVolPath: "/tmp",
		},
	}

	for _, c := range cases {
		realPath, _ := CreateRealPath("/test/abc/cde", c)
		if realPath != "/tmp/fe10ac20544bf8024a6a3ee603188662/test/abc/cde" {
			t.Errorf("The realPath %s is not equal /tmp/fe10ac20544bf8024a6a3ee603188662/test/abc/cde", realPath)
		}

		fi, err := os.Stat(realPath)
		if err != nil {
			if !os.IsExist(err) {
				t.Errorf("Create directory error")
			}
		}
	}
}
