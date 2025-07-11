// Copyright 2025 snowy-jaguar
// Contact: @snowyjaguar (Discord)
// Contact: contact@snowyjaguar.xyz (Email)
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"go.uber.org/zap"

	"github.com/snowy-jaguar/adguardhomesync-swarm/pkg/client/model"
	"github.com/snowy-jaguar/adguardhomesync-swarm/pkg/log"
	"github.com/snowy-jaguar/adguardhomesync-swarm/pkg/types"
)

var l = log.GetLogger("client")

// New create a new api client.
func New(config types.AdGuardInstance) (Client, error) {
	var apiURL string
	if config.APIPath == "" {
		apiURL = config.URL + "/control"
	} else {
		apiURL = fmt.Sprintf("%s/%s", config.URL, config.APIPath)
	}
	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Clean(u.Path)

	httpClient := &http.Client{
		Transport: &http.Transport{
			// #nosec G402 has to be explicitly enabled
			TLSClientConfig: &tls.Config{InsecureSkipVerify: config.InsecureSkipVerify},
		},
	}

	aghClient, err := model.NewClient(u.String(), func(client *model.AdguardHomeClient) error {
		client.Client = httpClient
		client.RequestEditors = append(client.RequestEditors, func(ctx context.Context, req *http.Request) error {
			if config.Username != "" && config.Password != "" {
				req.Header.Add("Authorization", "Basic "+basicAuth(config.Username, config.Password))
			}
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &apiClient{
		host:   u.Host,
		client: aghClient,
		log:    l.With("host", u.Host),
	}, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type apiClient struct {
	host   string
	client *model.AdguardHomeClient
	log    *zap.SugaredLogger
}

func (a apiClient) Host(context.Context) string {
	return a.host
}

func (a apiClient) GetServerStatus(ctx context.Context) (*model.ServerStatus, error) {
	sr, err := read(ctx, a.client.Status, model.ParseStatusResp)
	if err != nil {
		return nil, err
	}
	return sr.JSON200, nil
}

func (a apiClient) GetFilteringStatus(ctx context.Context) (*model.FilterStatus, error) {
	sr, err := read(ctx, a.client.FilteringStatus, model.ParseFilteringStatusResp)
	if err != nil {
		return nil, err
	}
	return sr.JSON200, nil
}

func (a apiClient) SetFilteringConfig(ctx context.Context, config model.FilterConfig) error {
	return write(ctx, config, a.client.FilteringConfig)
}

func write[B any](
	ctx context.Context,
	body B,
	req func(ctx context.Context, body B, reqEditors ...model.RequestEditorFn) (*http.Response, error),
) error {
	resp, err := req(ctx, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return detailedError(resp)
	}
	return nil
}

func read[I any](
	ctx context.Context,
	req func(ctx context.Context, reqEditors ...model.RequestEditorFn) (*http.Response, error),
	parse func(rsp *http.Response) (*I, error),
) (*I, error) {
	resp, err := req(ctx)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, detailedError(resp)
	}
	return parse(resp)
}

func detailedError(resp *http.Response) error {
	e := resp.Status

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(body) > 0 {
		e += fmt.Sprintf("(%s)", string(body))
	}
	return errors.New(e)
}
