package bootstrap

import (
	"github.com/kubeflow/manifests/tests"
	"testing"
)

func TestKustomize(t *testing.T) {
	testCase := &tests.KustomizeTestCase{
		Package:  "../../../../../stacks/ibm/application/bootstrap",
		Expected: "test_data/expected",
	}

	tests.RunTestCase(t, testCase)
}