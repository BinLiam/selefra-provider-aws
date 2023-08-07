package availability

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAvailabilityZoneGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAvailabilityZoneGenerator{}

func (x *TableAwsAvailabilityZoneGenerator) GetTableName() string {
	return "aws_availability_zone"
}

func (x *TableAwsAvailabilityZoneGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAvailabilityZoneGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAvailabilityZoneGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsAvailabilityZoneGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ZoneName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ZoneType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("opt_in_status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("OptInStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("GroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("RegionName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_zone_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ParentZoneName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_zone_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ParentZoneId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("messages").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Messages")).Build(),
	}
}

func (x *TableAwsAvailabilityZoneGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsAvailabilityZoneGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Ec2
			out, err := svc.DescribeAvailabilityZones(ctx, &ec2.DescribeAvailabilityZonesInput{
				AllAvailabilityZones: aws.Bool(true),
				Filters: []types.Filter{
					{
						Name:   aws.String("region-name"),
						Values: []string{cl.GetRegion()},
					},
				},
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- out.AvailabilityZones
			return nil
		},
	}
}

func (x *TableAwsAvailabilityZoneGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}
