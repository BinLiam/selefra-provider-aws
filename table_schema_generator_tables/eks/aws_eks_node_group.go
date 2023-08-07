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

type TableAwsEksNodeGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEksNodeGroupGenerator{}

func (x *TableAwsEksNodeGroupGenerator) GetTableName() string {
	return "aws_eks_node_group"
}

func (x *TableAwsEksNodeGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEksNodeGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEksNodeGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEksNodeGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_eks_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("nodegroup_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("NodegroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("NodegroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_name").Extractor(column_value_extractor.StructSelector("ClusterName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ami_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AmiType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("CapacityType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_size").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("DiskSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("ModifiedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_role").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("NodeRole")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("release_version").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ReleaseVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Health")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_types").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("InstanceTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_template").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("LaunchTemplate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remote_access").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("RemoteAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Resources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ScalingConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnets").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Subnets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("taints").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Taints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("UpdateConfig")).Build(),
	}
}

func (x *TableAwsEksNodeGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsEksNodeGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*types.Cluster)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Eks
			var nextToken *string
			input := eks.ListNodegroupsInput{
				ClusterName: r.Name,
			}
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListNodegroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, nodegroup := range output.Nodegroups {
					params := &eks.DescribeNodegroupInput{
						ClusterName:   r.Name,
						NodegroupName: &nodegroup,
					}
					op, err := svc.DescribeNodegroup(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- op.Nodegroup
				}
			}
		},
	}
}

func (x *TableAwsEksNodeGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("eks")
}
