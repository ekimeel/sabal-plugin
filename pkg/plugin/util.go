package plugin

import (
	"github.com/ekimeel/sabal-pb/pb"
	"sort"
)

func GroupAndSortMetrics(metrics []*pb.Metric) map[uint32][]*pb.Metric {
	groupedMetrics := make(map[uint32][]*pb.Metric)

	// Group by PointId
	for _, metric := range metrics {
		groupedMetrics[metric.PointId] = append(groupedMetrics[metric.PointId], metric)
	}

	// Sort each group by Ts
	for pointId := range groupedMetrics {
		group := groupedMetrics[pointId]
		sort.Slice(group, func(i, j int) bool {
			return group[i].Timestamp.AsTime().Before(group[j].Timestamp.AsTime())
		})
		groupedMetrics[pointId] = group // Re-assign after sorting
	}

	return groupedMetrics
}
