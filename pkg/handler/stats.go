// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the Apache-2.0 license, a copy of
// which may be found at https://github.com/silverton-io/buz/blob/main/LICENSE

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/silverton-io/buz/pkg/meta"
	"github.com/silverton-io/buz/pkg/stats"
)

type StatsResponse struct {
	CollectorMeta *meta.CollectorMeta  `json:"collectorMeta"`
	Stats         *stats.ProtocolStats `json:"stats"`
}

func StatsHandler(m *meta.CollectorMeta, s *stats.ProtocolStats) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		resp := StatsResponse{
			CollectorMeta: m,
			Stats:         s,
		}
		c.JSON(200, resp)
	}
	return gin.HandlerFunc(fn)
}
