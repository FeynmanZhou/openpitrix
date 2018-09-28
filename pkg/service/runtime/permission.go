// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Auto generated by 'go run gen_helper.go', DO NOT EDIT.

package runtime

import (
	"context"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/senderutil"
)

func CheckRuntimesPermission(ctx context.Context, resourceIds []string) ([]*models.Runtime, error) {
	if len(resourceIds) == 0 {
		return nil, nil
	}
	var sender = senderutil.GetSenderFromContext(ctx)
	var runtimes []*models.Runtime
	_, err := pi.Global().DB(ctx).
		Select(models.RuntimeColumns...).
		From(constants.TableRuntime).
		Where(db.Eq(constants.ColumnRuntimeId, resourceIds)).Load(&runtimes)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil && !sender.IsGlobalAdmin() {
		for _, runtime := range runtimes {
			if runtime.Owner != sender.UserId {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, runtime.RuntimeId)
			}
		}
	}
	if len(runtimes) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceIds)
	}
	return runtimes, nil
}

func CheckRuntimePermission(ctx context.Context, resourceId string) (*models.Runtime, error) {
	if len(resourceId) == 0 {
		return nil, nil
	}
	var sender = senderutil.GetSenderFromContext(ctx)
	var runtimes []*models.Runtime
	_, err := pi.Global().DB(ctx).
		Select(models.RuntimeColumns...).
		From(constants.TableRuntime).
		Where(db.Eq(constants.ColumnRuntimeId, resourceId)).Load(&runtimes)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil && !sender.IsGlobalAdmin() {
		for _, runtime := range runtimes {
			if runtime.Owner != sender.UserId {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, runtime.RuntimeId)
			}
		}
	}
	if len(runtimes) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceId)
	}
	return runtimes[0], nil
}