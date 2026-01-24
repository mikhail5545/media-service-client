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
	"context"

	cldassetpbv1 "github.com/mikhail5545/media-service-client/pb/media_service/cloudinary/asset/v1"
	muxassetpbv1 "github.com/mikhail5545/media-service-client/pb/media_service/mux/asset/v1"
	"google.golang.org/grpc"
)

func (c *MuxAssetServiceClient) Get(ctx context.Context, in *muxassetpbv1.GetRequest, opts ...grpc.CallOption) (*muxassetpbv1.GetResponse, error) {
	return c.AssetServiceClient.Get(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) Get(ctx context.Context, in *cldassetpbv1.GetRequest, opts ...grpc.CallOption) (*cldassetpbv1.GetResponse, error) {
	return c.AssetServiceClient.Get(ctx, in, opts...)
}

func (c *MuxAssetServiceClient) GetWithArchived(ctx context.Context, in *muxassetpbv1.GetWithArchivedRequest, opts ...grpc.CallOption) (*muxassetpbv1.GetWithArchivedResponse, error) {
	return c.AssetServiceClient.GetWithArchived(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) GetWithArchived(ctx context.Context, in *cldassetpbv1.GetWithArchivedRequest, opts ...grpc.CallOption) (*cldassetpbv1.GetWithArchivedResponse, error) {
	return c.AssetServiceClient.GetWithArchived(ctx, in, opts...)
}

func (c *MuxAssetServiceClient) GetWithBroken(ctx context.Context, in *muxassetpbv1.GetWithBrokenRequest, opts ...grpc.CallOption) (*muxassetpbv1.GetWithBrokenResponse, error) {
	return c.AssetServiceClient.GetWithBroken(ctx, in, opts...)
}

func (c *CloudinaryAssetServiceClient) GetWithBroken(ctx context.Context, in *cldassetpbv1.GetWithBrokenRequest, opts ...grpc.CallOption) (*cldassetpbv1.GetWithBrokenResponse, error) {
	return c.AssetServiceClient.GetWithBroken(ctx, in, opts...)
}
