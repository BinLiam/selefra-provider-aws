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

type TableAwsAuditmanagerControlGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAuditmanagerControlGenerator{}

func (x *TableAwsAuditmanagerControlGenerator) GetTableName() string {
	return "aws_auditmanager_control"
}

func (x *TableAwsAuditmanagerControlGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAuditmanagerControlGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAuditmanagerControlGenerator) GetColumns() []*schema.Column {
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
		table_schema_generator.NewColumnBuilder().ColumnName("action_plan_title").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ActionPlanTitle")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("action_plan_instructions").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ActionPlanInstructions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_sources").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ControlSources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_at").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("LastUpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_by").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("LastUpdatedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("testing_information").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("TestingInformation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_mapping_sources").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ControlMappingSources")).Build(),
	}
}

type warpControl struct {
	types.ControlMetadata
	Type                   types.ControlType
	CreatedBy              *string
	ActionPlanInstructions *string
	ActionPlanTitle        *string
	Description            *string
	LastUpdatedBy          *string
	TestingInformation     *string
	ControlSources         *string
	ControlMappingSources  interface{}
}

func (x *TableAwsAuditmanagerControlGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsAuditmanagerControlGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{PrimaryKeys: []string{"id"}}
}

func (x *TableAwsAuditmanagerControlGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Auditmanager
			var input auditmanager.ListControlsInput
			var nextToken *string
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListControls(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, v := range output.ControlMetadataList {

					params := &auditmanager.GetControlInput{ControlId: v.Id}
					op, err := svc.GetControl(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}

					resultChannel <- &warpControl{
						ControlMetadata:        v,
						Type:                   op.Control.Type,
						CreatedBy:              op.Control.CreatedBy,
						ActionPlanInstructions: op.Control.ActionPlanInstructions,
						ActionPlanTitle:        op.Control.ActionPlanTitle,
						Description:            op.Control.Description,
						LastUpdatedBy:          op.Control.LastUpdatedBy,
						TestingInformation:     op.Control.TestingInformation,
						ControlSources:         op.Control.ControlSources,
						ControlMappingSources:  op.Control.ControlMappingSources,
					}
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

func (x *TableAwsAuditmanagerControlGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("auditmanager")
}
