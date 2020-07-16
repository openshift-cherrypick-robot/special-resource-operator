package specialresource

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func init() {

}

// AddKubeClient Add a native non-caching client for advanced CRUD operations
func AddKubeClient(cfg *rest.Config) error {
	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return err
	}
	kubeclient = clientSet
	return nil
}

func checkNestedFields(found bool, err error) {
	if !found || err != nil {
		log.Error(err, "")
		os.Exit(1)
	}
}

func exitOnError(err error) {
	if err != nil {
		log.Info("Exiting On Error", "error", err)
		os.Exit(1)
	}
}
