package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func push(imageName string, authConfig types.AuthConfig) (reader io.ReadCloser, err error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return nil, err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	reader, err = cli.ImagePush(ctx, imageName, types.ImagePushOptions{RegistryAuth: authStr})
	return reader, err
}
