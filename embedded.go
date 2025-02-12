package starboard

import (
	_ "embed"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/scheme"
)

var (
	//go:embed deploy/crd/vulnerabilityreports.crd.yaml
	vulnerabilityReportsCRD []byte
	//go:embed deploy/crd/clustervulnerabilityreports.crd.yaml
	clusterVulnerabilityReportsCRD []byte
	//go:embed deploy/crd/configauditreports.crd.yaml
	configAuditReportsCRD []byte
	//go:embed deploy/crd/clusterconfigauditreports.crd.yaml
	clusterConfigAuditReportsCRD []byte
	//go:embed deploy/crd/ciskubebenchreports.crd.yaml
	kubeBenchReportsCRD []byte
	//go:embed deploy/crd/kubehunterreports.crd.yaml
	kubeHunterReportsCRD []byte
)

func GetVulnerabilityReportsCRD() (apiextensionsv1.CustomResourceDefinition, error) {
	return getCRDFromBytes(vulnerabilityReportsCRD)
}

func GetClusterVulnerabilityReportsCRD() (apiextensionsv1.CustomResourceDefinition, error) {
	return getCRDFromBytes(clusterVulnerabilityReportsCRD)
}

func GetConfigAuditReportsCRD() (apiextensionsv1.CustomResourceDefinition, error) {
	return getCRDFromBytes(configAuditReportsCRD)
}

func GetClusterConfigAuditReportsCRD() (apiextensionsv1.CustomResourceDefinition, error) {
	return getCRDFromBytes(clusterConfigAuditReportsCRD)
}

func GetCISKubeBenchReportsCRD() (apiextensionsv1.CustomResourceDefinition, error) {
	return getCRDFromBytes(kubeBenchReportsCRD)
}

func GetKubeHunterReportsCRD() (apiextensionsv1.CustomResourceDefinition, error) {
	return getCRDFromBytes(kubeHunterReportsCRD)
}

func getCRDFromBytes(bytes []byte) (apiextensionsv1.CustomResourceDefinition, error) {
	var crd apiextensionsv1.CustomResourceDefinition
	_, _, err := scheme.Codecs.UniversalDecoder().Decode(bytes, nil, &crd)
	if err != nil {
		return apiextensionsv1.CustomResourceDefinition{}, err
	}
	return crd, nil
}
