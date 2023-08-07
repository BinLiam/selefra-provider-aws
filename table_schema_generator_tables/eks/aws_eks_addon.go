package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"

	"github.com/aws/aws-sdk-go-v2/service/eks/types"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsEksAddonGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEksAddonGenerator{}

func (x *TableAwsEksAddonGenerator) GetTableName() string {
	return "aws_eks_addon"
}

func (x *TableAwsEksAddonGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEksAddonGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEksAddonGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEksAddonGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_eks_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("addon_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AddonName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AddonArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ClusterName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addon_version").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AddonVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("ModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account_role_arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ServiceAccountRoleArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_issues").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Health.Issues")).Build(),
	}
}

func (x *TableAwsEksAddonGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsEksAddonGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*types.Cluster)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Eks
			var nextToken *string
			input := eks.ListAddonsInput{
				ClusterName: r.Name,
			}
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListAddons(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, addon := range output.Addons {
					params := &eks.DescribeAddonInput{
						AddonName:   &addon,
						ClusterName: r.Name,
					}
					op, err := svc.DescribeAddon(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- op.Addon

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

func (x *TableAwsEksAddonGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("eks")
}
