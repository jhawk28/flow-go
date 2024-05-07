package level1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/utils/unittest"
)

const BootstrapPath = "../testdata/level1/data/"
const ExpectedOutputPath = "../testdata/level1/expected"

// TestGenerateBootstrap_DataTable validates that a root snapshot fixture produces the
// expected output template. If this test fails, it is likely because the underlying
// Snapshot model has changed. In that case, you can generate a new fixture file with
// a script like:
//
//	participants := unittest.IdentityListFixture(10, unittest.WithAllRoles())
//	snapshot := unittest.RootSnapshotFixture(participants)
//	json.NewEncoder(os.Stdout).Encode(snapshot.Encodable())
func TestGenerateBootstrap_DataTable(t *testing.T) {
	testDataMap := map[string]testData{
		"10 nodes": {
			bootstrapPath:  filepath.Join(BootstrapPath, "snapshot-fixture1.json"),
			expectedOutput: filepath.Join(ExpectedOutputPath, "template-fixture1.json"),
			dockerTag:      "v0.27.6",
			dockerRegistry: "gcr.io/flow-container-registry/",
		},
	}

	for i, testData := range testDataMap {
		t.Run(i, func(t *testing.T) {
			// generate template data file from bootstrap file
			bootstrap := NewBootstrap(testData.bootstrapPath)
			actualNodeData := bootstrap.GenTemplateData(false, testData.dockerTag, testData.dockerRegistry)

			// load expected template data file
			var expectedNodeData []NodeData
			expectedDataBytes, err := os.ReadFile(testData.expectedOutput)
			require.NoError(t, err)
			err = json.Unmarshal(expectedDataBytes, &expectedNodeData)
			require.Nil(t, err)

			// check generated template data file is correct
			require.Equal(t, len(expectedNodeData), len(actualNodeData))
			require.ElementsMatch(t, expectedNodeData, actualNodeData)
		})
	}
}

type testData struct {
	bootstrapPath  string
	expectedOutput string
	dockerTag      string
	dockerRegistry string
}

func TestGenerateTestData(t *testing.T) {
	// 		"2 AN, 6 LN, 3 CN, 2 EN, 1 VN": {
	participants := unittest.IdentityListFixture(10, unittest.WithAllRoles())
	unittest.RootSnapshotFixture(participants)

	json.NewEncoder(os.Stdout)
}
