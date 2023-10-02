package main

import (
	"github.com/emporous/uor-zot/iac/cert-manager/resources"
	"github.com/emporous/uor-zot/iac/cert-manager/utils"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Initialize Kubernetes Provider
		kubeProvider, err := utils.InitKubeProvider(ctx)
		if err != nil {
			return err
		}

		// Create Namespace
		ns, err := resources.CreateNamespace(ctx, "my-namespace", kubeProvider)
		if err != nil {
			return err
		}

		// Create Cert Manager and Issuers
		issuerResource, err := resources.CreateCertManager(ctx, kubeProvider)
		if err != nil {
			return err
		}

		// Create Certificates
		err = resources.CreateCertificates(ctx, ns.Metadata.Name().Elem(), kubeProvider, issuerResource)
		if err != nil {
			return err
		}

		// Additional resources can be added here

		return nil
	})
}
