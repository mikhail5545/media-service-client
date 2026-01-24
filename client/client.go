/*
 * Copyright (c) 2026. Mikhail Kulik.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package client

import (
	cldassetpbv1 "github.com/mikhail5545/media-service-client/pb/media_service/cloudinary/asset/v1"
	muxassetpbv1 "github.com/mikhail5545/media-service-client/pb/media_service/mux/asset/v1"
	"google.golang.org/grpc"
	"time"
)

type config struct {
	timeout time.Duration
}

type MuxAssetServiceClient struct {
	muxassetpbv1.AssetServiceClient
	conn   *grpc.ClientConn
	config *config
}

type CloudinaryAssetServiceClient struct {
	cldassetpbv1.AssetServiceClient
	conn   *grpc.ClientConn
	config *config
}

var (
	_ muxassetpbv1.AssetServiceClient = (*MuxAssetServiceClient)(nil)
	_ cldassetpbv1.AssetServiceClient = (*CloudinaryAssetServiceClient)(nil)
)

func NewMuxAssetServiceClient(opts ...Option) *MuxAssetServiceClient {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &MuxAssetServiceClient{
		config: cfg,
	}
}

func NewCloudinaryAssetServiceClient(opts ...Option) *CloudinaryAssetServiceClient {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &CloudinaryAssetServiceClient{
		config: cfg,
	}
}
