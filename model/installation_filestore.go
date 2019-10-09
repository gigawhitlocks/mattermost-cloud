package model

const (
	// InstallationFilestoreOperator is a filestore hosted in kubernetes via the operator.
	InstallationFilestoreOperator = "operator"
	// InstallationFilestoreS3 is a filestore hosted via Amazon S3.
	InstallationFilestoreS3 = "s3"
)

// InternalFilestore returns true if the installation's filestore is internal
// to the kubernetes cluster it is running on.
func (i *Installation) InternalFilestore() bool {
	return i.Filestore == InstallationFilestoreOperator
}

// IsSupportedFilestore returns true if the given filestore string is supported.
func IsSupportedFilestore(filestore string) bool {
	return filestore == InstallationFilestoreOperator || filestore == InstallationFilestoreS3
}
