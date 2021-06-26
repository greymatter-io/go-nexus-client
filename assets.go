package client

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	assetsAPIEndpoint = "service/rest/v1/assets"
)

type AssetResponse struct {
	Items             []Asset `json:"items"`
	ContinuationToken string `json:"continuationToken"`
}

type Asset struct {
	DownloadURL    string `json:"downloadUrl"`
	Path           string `json:"path"`
	ID             string `json:"id"`
	Repository     string `json:"repository"`
	Format         string `json:"format"`
	Checksum       Checksum  `json:"checksum"`
	ContentType    string `json:"contentType"`
	LastModified   time.Time `json:"lastModified"`
	BlobCreated    time.Time `json:"blobCreated"`
	LastDownloaded time.Time `json:"lastDownloaded"`
}

type Checksum struct {
	SHA1   string `json:"sha1"`
	SHA256 string `json:"sha256"`
}

func (c client) AssetList(repository, continuationToken string) (*AssetResponse, error) {

	url := assetsAPIEndpoint+"?repository="+repository
	if continuationToken != "" {
		url += "&continuationToken"+continuationToken
	}
	body, resp, err := c.Get(url, nil)
	if err != nil {
		return nil, fmt.Errorf("status code:%v %v", resp.StatusCode, err)
	}
	var ar AssetResponse
	if err := json.Unmarshal(body, &ar); err != nil {
		return nil, err
	}

	return &ar, nil
}

/*
  "items" : [ {
    "downloadUrl" : "https://nexus.greymatter.io/repository/docker/v2/-/blobs/sha256:c46b5fa4d940569e49988515c1ea0295f56d0a16228d8f854e27613f467ec892",
    "path" : "v2/-/blobs/sha256:c46b5fa4d940569e49988515c1ea0295f56d0a16228d8f854e27613f467ec892",
    "id" : "ZG9ja2VyOmM3NTQ1NTc5YTdhMTUzOTA0ZmRlODEwZjY4OTc5Yjkw",
    "repository" : "docker",
    "format" : "docker",
    "checksum" : {
      "sha1" : "8e0d539ecd39cd48eda037001f9d5335aa106c21",
      "sha256" : "c46b5fa4d940569e49988515c1ea0295f56d0a16228d8f854e27613f467ec892"
    },
    "contentType" : "application/vnd.docker.image.rootfs.diff.tar.gzip",
    "lastModified" : "2020-07-20T21:15:09.694+00:00",
    "blobCreated" : "2020-07-20T21:15:09.694+00:00",
    "lastDownloaded" : "2021-01-29T18:41:56.110+00:00"
  },

*/
