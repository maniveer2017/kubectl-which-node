package cmd

import (
	""
)

func findNodes(kind string, object string) {
	
	klog.Info("Kind:%v Object:%v",kind, object)
	klog.V(3).Info("Finding API resources")
	_ = findApiResources()
}
