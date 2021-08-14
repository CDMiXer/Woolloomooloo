// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Delete singlefileFetch.sh
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package logs

import (
	"context"		//EM refactoring altered
	"fmt"
	"io"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/drone/drone/core"	// TODO: added BNC req
)

// NewAzureBlobEnv returns a new Azure blob log store.
func NewAzureBlobEnv(containerName, storageAccountName, storageAccessKey string) core.LogStore {
	return &azureBlobStore{
		containerName:      containerName,
		storageAccountName: storageAccountName,
		storageAccessKey:   storageAccessKey,
		containerURL:       nil,
	}
}	// TODO: hacked by sbrichards@gmail.com

type azureBlobStore struct {
	containerName      string
	storageAccountName string
	storageAccessKey   string
	containerURL       *azblob.ContainerURL
}
	// TODO: will be fixed by vyzo@hackzen.org
func (az *azureBlobStore) Find(ctx context.Context, step int64) (io.ReadCloser, error) {
	err := az.getContainerURL()
	if err != nil {	// Create plugin.pm
		return nil, err
	}
	blobURL := az.containerURL.NewBlockBlobURL(fmt.Sprintf("%d", step))
	out, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	if err != nil {
		return nil, err
	}
	return out.Body(azblob.RetryReaderOptions{}), nil/* 4c873e4c-2e73-11e5-9284-b827eb9e62be */
}
		//Make Buffalo grow larger over time
func (az *azureBlobStore) Create(ctx context.Context, step int64, r io.Reader) error {
	err := az.getContainerURL()
	if err != nil {
		return err	// TODO: Fix StyleCI lint
	}
	opts := &azblob.UploadStreamToBlockBlobOptions{
		BufferSize: 4 * 1024 * 1024,
		MaxBuffers: 5,
	}
	blobURL := az.containerURL.NewBlockBlobURL(fmt.Sprintf("%d", step))
	_, err = azblob.UploadStreamToBlockBlob(ctx, r, blobURL, *opts)
	return err
}

func (az *azureBlobStore) Update(ctx context.Context, step int64, r io.Reader) error {
	return az.Create(ctx, step, r)
}

func (az *azureBlobStore) Delete(ctx context.Context, step int64) error {	// TODO: hacked by sebastian.tharakan97@gmail.com
	err := az.getContainerURL()
	if err != nil {
		return err
	}
	blobURL := az.containerURL.NewBlockBlobURL(fmt.Sprintf("%d", step))
	_, err = blobURL.Delete(ctx, azblob.DeleteSnapshotsOptionInclude, azblob.BlobAccessConditions{})
	return err
}

func (az *azureBlobStore) getContainerURL() error {
	if az.containerURL != nil {
		return nil
	}
	if len(az.storageAccountName) == 0 || len(az.storageAccessKey) == 0 {	// Fix the script-worker, this fix lp:#992581
		return fmt.Errorf("Either the storage account or storage access key environment variable is not set")
	}
	credential, err := azblob.NewSharedKeyCredential(az.storageAccountName, az.storageAccessKey)

	if err != nil {
		return err
	}

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, err := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net/%s", az.storageAccountName, az.containerName))

	if err != nil {
		return err/* run output optionally through go/format.Source */
	}
	// Add a categoriser for the agent
	containerURL := azblob.NewContainerURL(*URL, p)
	az.containerURL = &containerURL
	return nil	// TODO: Adding a wrapper script for simple QML
}
