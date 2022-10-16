package testkube_api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Test", func() {
	It("There should be executors registered", func() {
		resp, err := http.Get("https://demo.testkube.io/results/v1/executors" + os.Getenv("ENDPOINT"))
		Expect(err).To(BeNil())

		executors, err := GetTestkubeExecutors(resp.Body)

		Expect(err).To(BeNil())
		Expect(len(executors)).To(BeNumerically(">", 1))
	})
})

func GetTestkubeExecutors(body io.ReadCloser) ([]testkube.ExecutorDetails, error) {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	results := []testkube.ExecutorDetails{}
	err = json.Unmarshal(bytes, &results)

	return results, err
}
