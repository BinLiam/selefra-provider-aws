package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsBackupReportPlanGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupReportPlanGenerator{}

func (x *TableAwsBackupReportPlanGenerator) GetTableName() string {
	return "aws_backup_report_plan"
}

func (x *TableAwsBackupReportPlanGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupReportPlanGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupReportPlanGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"arn"},
	}
}

func (x *TableAwsBackupReportPlanGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ReportPlanArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("report_plan_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ReportPlanName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ReportPlanDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DeploymentStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_attempted_execution_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("LastAttemptedExecutionTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_successful_execution_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("LastSuccessfulExecutionTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("report_delivery_channel").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ReportDeliveryChannel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("report_setting").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ReportSetting")).Build(),
	}
}

func (x *TableAwsBackupReportPlanGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsBackupReportPlanGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Backup
			var nextToken *string
			var input backup.ListReportPlansInput
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListReportPlans(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, reportPlan := range output.ReportPlans {
					resultChannel <- reportPlan
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

func (x *TableAwsBackupReportPlanGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}
