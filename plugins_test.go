package fabric_plugins

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestPlugins(t *testing.T) {
	gt := NewGomegaWithT(t)
	// Compile plugins
	compilePlugin(gt, "endorsement")
	compilePlugin(gt, "validation")
	compilePlugin(gt, "consensus")
}

// compilePlugin compiles the plugin of the given type and returns the path for the plugin file
func compilePlugin(gt *GomegaWithT, pluginType string) string {
	pluginFilePath := filepath.Join("output", pluginType, "plugin.so")
	cmd := exec.Command(
		"go", "build", "-buildmode=plugin",
		"-o", pluginFilePath,
		fmt.Sprintf("./plugins/%s", pluginType),
	)
	sess, err := gexec.Start(cmd, nil, nil)
	gt.Expect(err).NotTo(HaveOccurred())
	defer sess.Kill()
	gt.Eventually(sess, time.Minute).Should(gexec.Exit(0))

	return pluginFilePath
}
