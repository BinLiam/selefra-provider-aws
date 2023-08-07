package cloudfront

import (
	"context"

	"github.com/songzhibin97/go-ognl"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsCloudfrontOriginAccessIdentityGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudfrontOriginAccessIdentityGenerator{}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetTableName() string {
	return "aws_cloudfront_origin_access_identity"
}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

type warpAccessIdentity struct {
	Id                *string
	Arn               string
	S3CanonicalUserId *string
	ETag              *string
	CallerReference   interface{}
	Comment           interface{}
}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("s3_canonical_user_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("S3CanonicalUserId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("caller_reference").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("CallerReference")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Comment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ETag")).Build(),
	}
}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudfront
			var nextMarker *string
			input := cloudfront.ListCloudFrontOriginAccessIdentitiesInput{}
			for {
				if nextMarker != nil {
					input.Marker = nextMarker
				}
				output, err := svc.ListCloudFrontOriginAccessIdentities(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, v := range output.CloudFrontOriginAccessIdentityList.Items {
					wp := warpAccessIdentity{
						Id:                v.Id,
						Arn:               c.ARN("cloudfront", "origin-access-identity", *v.Id),
						S3CanonicalUserId: v.S3CanonicalUserId,
						Comment:           v.Comment,
					}
					func() {
						params := &cloudfront.GetCloudFrontOriginAccessIdentityInput{
							Id: v.Id,
						}
						op, err := svc.GetCloudFrontOriginAccessIdentity(ctx, params)
						if err != nil {
							return
						}
						wp.ETag = op.ETag
						if wp.S3CanonicalUserId == nil {
							wp.S3CanonicalUserId = op.CloudFrontOriginAccessIdentity.S3CanonicalUserId
						}
						wp.CallerReference = ognl.Get(op, "CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig.CallerReference").Value()
						wp.Comment = ognl.Get(op, "CloudFrontOriginAccessIdentity.CloudFrontOriginAccessIdentityConfig.Comment").Value()
					}()

					resultChannel <- wp
				}
				if output.CloudFrontOriginAccessIdentityList.NextMarker == nil {
					break
				}
				nextMarker = output.CloudFrontOriginAccessIdentityList.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsCloudfrontOriginAccessIdentityGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}
