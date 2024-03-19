package assetinventory

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

const TemporaryAssetIdPrefix = "T:"

func GenerateTemporaryAssetIdFromSignedURL(mimeType string, urlStr string) (string, error) {
	bucketName, objectName, err := parseSignedURL(urlStr)
	if err != nil {
		return "", err
	}
	return GenerateTemporaryAssetId(mimeType, bucketName, objectName)
}

func GenerateTemporaryAssetId(mimeType, bucketName, objectName string) (string, error) {
	raw := map[string]string{
		"b": bucketName,
		"o": objectName,
		"m": mimeType,
	}
	data, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	if err := json.Compact(&buffer, data); err != nil {
		return "", err
	}
	assetID := TemporaryAssetIdPrefix + base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(buffer.Bytes())
	return assetID, nil
}

func ParseTemporaryAssetId(assetId string) (string, string, string, error) {
	var out map[string]string
	tempAssetId := strings.TrimPrefix(assetId, TemporaryAssetIdPrefix)
	if data, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(tempAssetId); err != nil {
		return "", "", "", err
	} else {
		if err := json.Unmarshal(data, &out); err != nil {
			return "", "", "", err
		}
	}
	var (
		bucketName string
		objectName string
		mimeType   string
	)
	if v, ok := out["b"]; ok {
		bucketName = v
	}
	if v, ok := out["o"]; ok {
		objectName = v
	}
	if v, ok := out["m"]; ok {
		mimeType = v
	}
	return bucketName, objectName, mimeType, nil
}

func parseSignedURL(urlStr string) (string, string, error) {
	u, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return "", "", err
	}
	if u.Host != "storage.googleapis.com" {
		return "", "", errors.New("invalid signed-url")
	}
	s := u.Path
	s = strings.TrimPrefix(s, "/")
	l := strings.SplitN(s, "/", 2)
	if len(l) < 2 {
		return "", "", errors.New("invalid signed-url")
	}
	return l[0], l[1], nil
}
