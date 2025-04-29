package main

import (
	"github.com/aobco/log"
	"time"
	"xsky.com/ocpf/console/service/download/downloader"
	"xsky.com/ocpf/pkg/datasource/model"
)

func main() {
	fileDownloader := new(downloader.SmbDownloader)
	if err := fileDownloader.Setup(&model.DetailCredentialVO{
		Datasource: &model.Datasource{
			Id:               0,
			Name:             "smb",
			Description:      "",
			Protocol:         "SMB/CIFS",
			ProtocolVersion:  "v2.1",
			Provider:         "",
			Status:           "",
			StatusMsg:        nil,
			Endpoint:         "\\\\10.16.12.233\\dir2_smb",
			Extra:            nil,
			Owner:            0,
			Deleted:          false,
			Creator:          "",
			CreatedTime:      time.Time{},
			UpdatedTime:      time.Time{},
			EndpointMappings: nil,
		},
		Credential: &model.DsCredential{
			Id:           100000,
			Name:         "",
			DatasourceId: 0,
			Description:  "",
			Info: &model.DsCredentialInfo{
				S3Credential: model.S3Credential{},
				FsCredential: model.FsCredential{},
			},
			Owner:       0,
			Creator:     "",
			CreatedTime: time.Time{},
			UpdatedTime: time.Time{},
		},
		ResourceName:       "",
		ArchiveClusterInfo: nil,
		BlurayClusterInfo:  nil,
		Actions:            nil,
		User:               nil,
	}); err != nil {
		log.Panic(err)
	}
	file, err := fileDownloader.Open("bag/apache-hive-3.1.3-bin.tar.gz", "", "")
	if err != nil {
		log.Panic(err)
	}
	log.Infof("open file success")
	defer file.Close()
}
