package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/cli-runtime/pkg/printers"
	"k8s.io/kubectl/pkg/scheme"
)

func main() {

	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error ", err)
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode
	manifests := strings.Split(string(content), "---")
	printer := printers.YAMLPrinter{}

	for _, i := range manifests {
		if len(i) <= 0 {
			continue
		}

		runtime, gKV, _ := decode([]byte(i), nil, nil)

		if gKV.Kind == "Deployment" {
			deployment := runtime.(*appsv1.Deployment)
			annotations := map[string]string{"bnhp.co.il/my-custom-annotation": "hello"}
			deployment.GetObjectMeta().SetAnnotations(annotations)
			printer.PrintObj(deployment, os.Stdout)

		} else {
			printer.PrintObj(runtime, os.Stdout)
		}

	}
}
