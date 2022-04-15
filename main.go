package main

import (
	"context"
	"log"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

var (
	pod       = os.Getenv("POD_NAME")
	namespace = os.Getenv("NAMESPACE")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to get config, are you on Kubernetes? %v", err)
	}

	leaseLock := &resourcelock.LeaseLock{
		Client: clientset.NewForConfigOrDie(config).CoordinationV1(),
		LeaseMeta: metav1.ObjectMeta{
			Name:      "lock",
			Namespace: namespace,
		},
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: pod,
		},
	}
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            leaseLock,
		ReleaseOnCancel: true,
		LeaseDuration:   10 * time.Second,
		RenewDeadline:   5 * time.Second,
		RetryPeriod:     1 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				log.Println("Starting to lead...")
			},
			OnStoppedLeading: func() {
				log.Println("No longer the leader...")
			},
			OnNewLeader: func(leader string) {
				log.Printf("New leader is %q", leader)
			},
		},
	})
}
