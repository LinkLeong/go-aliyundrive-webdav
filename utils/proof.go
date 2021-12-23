package utils

import (
	"encoding/base64"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func CalcOffSet(token string, size string) (int64, error) {
	cmd := exec.Command("./bn_v1.0.8_macos", "--token="+token, "--size="+size)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	offset := string(out)
	if strings.HasSuffix(offset, "\n") {
		offset = offset[:len(offset)-1]
	}

	return strconv.ParseInt(offset, 10, 64)
}

func GetProof(proofBytes []byte) string {
	return base64.StdEncoding.EncodeToString(proofBytes)
}
