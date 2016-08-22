package volume

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"

	"github.com/kubernetes/dashboard/resource/user"
)

type VolumeRelatedInfo struct {
	U              user.User
	CetusfsVolPath string
}

//Create the real path according to the user and the virtual hostpath
func CreateRealPath(virtPath string, v *VolumeRelatedInfo) (string, error) {
	h := md5.New()
	io.WriteString(h, v.U.Name)
	io.WriteString(h, v.U.Password)

	userHash := hex.EncodeToString(h.Sum(nil))
	realPath := path.Join(v.CetusfsVolPath, userHash, virtPath)

	err := os.MkdirAll(realPath, 0777)
	if err != nil {
		log.Println("mkdir error : ", err)
	}

	return realPath, err
}
