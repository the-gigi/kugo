package kugo

import (
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

// Run runs a raw kubectl command and returns stdout + stderr as a combined string
//
// It accepts a slice of arguments or a single string with space separated arguments
func Run(args ...string) (combinedOutput string, err error) {
	if len(args) == 0 {
		err = errors.New("Run() requires at least one argument")
		return
	}

	if len(args) == 1 {
		args = strings.Split(args[0], " ")
	}

	bytes, err := exec.Command("kubectl", args...).CombinedOutput()
	combinedOutput = string(bytes)
	return
}

// Get gets resources of a specific kind
//
// If a namespace is not provided it gets resources in all namespaces
// If output is not provided the efault is json (unlike kubectl's yaml)
// Extra args can be provided via r.ExtraArgs
func Get(r GetRequest) (result string, err error) {
	if r.Kind == "" {
		err = errors.New("Must specify Kind filed")
		return
	}
	args := []string{r.Kind}

	handleCommonArgs(args, r.BaseRequest)

	output := "json"
	if r.Output != "" {
		output = r.Output
	}
	args = append(args, "-o", output)

	if r.Label != "" {
		args = append(args, "-l", r.Label)
	}
	if len(r.FieldSelectors) > 0 {
		fieldSelectors := strings.Join(r.FieldSelectors, ",")
		args = append(args, "--field-selector", "'" + fieldSelectors + "'")
	}

	return Run(args...)
}

// Exec executes a command in a pod
//
// The target pod can specified by name or an arbitrary pod
// from a deployment or service.
//
// If the pod has multiple containers you can choose which
// container to run the command in
func Exec(r ExecRequest) (result string, err error) {
	if r.Command == "" {
		err = errors.New("Must specify Command filed")
		return
	}

	if r.Target == "" {
		err = errors.New("Must specify Target filed")
		return
	}

	args := []string{"exec", r.Target}
	if r.Container != "" {
		args = append(args, "-c", r.Container)
	}

	handleCommonArgs(args, r.BaseRequest)

	args = append(args, "--", r.Command)

	return Run(args...)
}

func handleCommonArgs(args []string, r BaseRequest) {
	if r.KubeContext != "" {
		args = append(args, "--context", r.KubeContext)
	}
	if r.Namespace != "" {
		args = append(args, "-n", r.Namespace)
	} else {
		args = append(args, "-A")
	}

	if len(r.ExtraArgs) > 0 {
		args = append(args, r.ExtraArgs...)
	}
	return
}