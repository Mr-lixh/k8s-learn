package deployment

import apps "k8s.io/api/apps/v1"

// DeploymentComplete considers a deployment to be complete once all of its desired replicas
// are updated and available, and no old pods are running.
// DeploymentComplete 用于判断 deployment 是否更新结束
func DeploymentComplete(deployment *apps.Deployment, newStatus *apps.DeploymentStatus) bool {
	return newStatus.UpdatedReplicas == *(deployment.Spec.Replicas) &&
		newStatus.Replicas == *(deployment.Spec.Replicas) &&
		newStatus.AvailableReplicas == *(deployment.Spec.Replicas) &&
		newStatus.ObservedGeneration >= deployment.Generation
}
