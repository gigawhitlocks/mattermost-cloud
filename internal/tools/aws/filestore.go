package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/arn"
	log "github.com/sirupsen/logrus"
)

func (a *Client) S3FilestoreProvision(installationID string, logger log.FieldLogger) error {
	awsID := CloudID(installationID)

	user, err := a.iamEnsureUserCreated(awsID, logger)
	if err != nil {
		return err
	}

	// The IAM policy lookup requires the AWS account ID for the ARN. The user
	// object contains this ID so we will user that.
	arn, err := arn.Parse(*user.Arn)
	if err != nil {
		return err
	}
	policyARN := fmt.Sprintf("arn:aws:iam::%s:policy/%s", arn.AccountID, awsID)
	policy, err := a.iamEnsurePolicyCreated(awsID, policyARN, logger)
	if err != nil {
		return err
	}
	err = a.iamEnsurePolicyAttached(awsID, policyARN)
	if err != nil {
		return err
	}
	logger.WithFields(log.Fields{
		"iam-policy-name": *policy.PolicyName,
		"iam-user-name":   *user.UserName,
	}).Info("AWS IAM policy attached to user")

	err = a.s3EnsureBucketCreated(awsID)
	if err != nil {
		return err
	}
	logger.WithField("s3-bucket-name", awsID).Info("AWS S3 bucket created")

	ak, err := a.iamEnsureAccessKeyCreated(awsID, logger)
	if err != nil {
		return err
	}
	logger.WithField("iam-user-name", *user.UserName).Info("AWS IAM user access key created")

	err = a.secretsManagerEnsureIAMAccessKeySecretCreated(awsID, ak)
	if err != nil {
		return err
	}
	logger.WithField("iam-user-name", *user.UserName).Info("AWS secrets manager secret created")

	return nil
}

func (a *Client) S3FilestoreTeardown(installationID string, keepBucket bool, logger log.FieldLogger) error {
	awsID := CloudID(installationID)

	err := a.iamEnsureUserDeleted(awsID, logger)
	if err != nil {
		return err
	}
	err = a.secretsManagerEnsureIAMAccessKeySecretDeleted(awsID, logger)
	if err != nil {
		return err
	}
	logger.WithField("iam-user-name", awsID).Info("AWS secrets manager secret created")

	if !keepBucket {
		err = a.s3EnsureBucketDeleted(awsID)
		if err != nil {
			return err
		}
		logger.WithField("s3-bucket-name", awsID).Info("AWS S3 bucket deleted")
	}

	return nil
}
