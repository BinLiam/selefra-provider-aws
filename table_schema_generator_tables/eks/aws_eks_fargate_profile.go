package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsEksFargateProfileGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEksFargateProfileGenerator{}

func (x *TableAwsEksFargateProfileGenerator) GetTableName() string {
	return "aws_eks_fargate_profile"
}

func (x *TableAwsEksFargateProfileGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEksFargateProfileGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEksFargateProfileGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEksFargateProfileGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_eks_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("fargate_profile_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FargateProfileName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ClusterName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fargate_profile_arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FargateProfileArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pod_execution_role_arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("PodExecutionRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selectors").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Selectors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnets").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Subnets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAwsEksFargateProfileGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsEksFargateProfileGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*types.Cluster)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Eks
			var nextToken *string
			input := eks.ListFargateProfilesInput{
				ClusterName: r.Name,
			}
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListFargateProfiles(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, name := range output.FargateProfileNames {
					params := &eks.DescribeFargateProfileInput{
						ClusterName:        r.Name,
						FargateProfileName: &name,
					}
					op, err := svc.DescribeFargateProfile(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- op.FargateProfile

				}
				if output.NextToken == nil {
					break
				}
				nextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEksFargateProfileGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("eks")
}
