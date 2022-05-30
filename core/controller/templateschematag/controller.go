package templateschematag

import (
	"context"

	"g.hz.netease.com/horizon/core/common"

	clustermanager "g.hz.netease.com/horizon/pkg/cluster/manager"
	"g.hz.netease.com/horizon/pkg/templateschematag/manager"
	"g.hz.netease.com/horizon/pkg/util/wlog"
)

type Controller interface {
	List(ctx context.Context, clusterID uint) (*ListResponse, error)
	Update(ctx context.Context, clusterID uint, r *UpdateRequest) error
}

type controller struct {
	clusterMgr          clustermanager.Manager
	clusterSchemaTagMgr manager.Manager
}

func NewController() Controller {
	return &controller{
		clusterMgr:          clustermanager.Mgr,
		clusterSchemaTagMgr: manager.Mgr,
	}
}

func (c *controller) List(ctx context.Context, clusterID uint) (_ *ListResponse, err error) {
	const op = "cluster template scheme tag controller: list"
	defer wlog.Start(ctx, op).StopPrint()

	tags, err := c.clusterSchemaTagMgr.ListByClusterID(ctx, clusterID)
	if err != nil {
		return nil, err
	}

	return ofClusterTemplateSchemaTags(tags), nil
}

func (c *controller) Update(ctx context.Context, clusterID uint, r *UpdateRequest) (err error) {
	const op = "cluster template scheme tag controller: update"
	defer wlog.Start(ctx, op).StopPrint()

	currentUser, err := common.FromContext(ctx)
	if err != nil {
		return err
	}

	clusterTemplateSchemaTags := r.toClusterTemplateSchemaTags(clusterID, currentUser)

	if err := manager.ValidateUpsert(clusterTemplateSchemaTags); err != nil {
		return err
	}

	cluster, err := c.clusterMgr.GetByID(ctx, clusterID)
	if err != nil {
		return err
	}

	return c.clusterSchemaTagMgr.UpsertByClusterID(ctx, cluster.ID, clusterTemplateSchemaTags)
}
