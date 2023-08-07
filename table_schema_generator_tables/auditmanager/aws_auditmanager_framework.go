package auditmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAuditmanagerFrameworkGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAuditmanagerFrameworkGenerator{}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetTableName() string {
	return "aws_auditmanager_framework"
}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetVersion() uint64 {
	return 0
}

type warpFramework struct {
	types.AssessmentFrameworkMetadata
	CreatedBy      *string
	ControlSources *string
	LastUpdatedBy  *string
	ControlSets    interface{}
}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("CreatedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compliance_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ComplianceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("controls_count").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("ControlsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_sets_count").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("ControlSetsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_sources").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ControlSources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("LastUpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_by").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("LastUpdatedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logo").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Logo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_sets").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ControlSets")).Build(),
	}
}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{PrimaryKeys: []string{"id"}}
}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Auditmanager
			var input auditmanager.ListAssessmentFrameworksInput
			var nextToken *string
			for {
				input.NextToken = nextToken
				output, err := svc.ListAssessmentFrameworks(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, framework := range output.FrameworkMetadataList {
					params := &auditmanager.GetAssessmentFrameworkInput{
						FrameworkId: framework.Id,
					}
					op, err := svc.GetAssessmentFramework(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- &warpFramework{
						AssessmentFrameworkMetadata: framework,
						CreatedBy:                   op.Framework.CreatedBy,
						ControlSources:              op.Framework.ControlSources,
						LastUpdatedBy:               op.Framework.LastUpdatedBy,
						ControlSets:                 op.Framework.ControlSets,
					}
				}
				if output.NextToken == nil {
					break
				}
				nextToken = output.NextToken
			}
			return nil
		}}
}

func (x *TableAwsAuditmanagerFrameworkGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("auditmanager")
}
