package auditmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAuditmanagerEvidenceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAuditmanagerEvidenceGenerator{}

func (x *TableAwsAuditmanagerEvidenceGenerator) GetTableName() string {
	return "aws_auditmanager_evidence"
}

func (x *TableAwsAuditmanagerEvidenceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAuditmanagerEvidenceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAuditmanagerEvidenceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"id"},
	}
}

func (x *TableAwsAuditmanagerEvidenceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.AssessmentEvidenceFolder)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Auditmanager
			params := auditmanager.GetEvidenceByEvidenceFolderInput{
				AssessmentId:     r.AssessmentId,
				ControlSetId:     r.ControlSetId,
				EvidenceFolderId: r.Id,
			}
			var nextToken *string
			for {
				if nextToken != nil {
					params.NextToken = nextToken
				}
				output, err := svc.GetEvidenceByEvidenceFolder(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, evidence := range output.Evidence {
					resultChannel <- evidence
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

func (x *TableAwsAuditmanagerEvidenceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_auditmanager_evidence_folder_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
			task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

			extractor := func() (any, error) {
				return client.(*aws_client.Client).ARN("auditmanager", "evidence", *result.(types.Evidence).Id), nil
			}
			extractResultValue, err := extractor()
			if err != nil {
				return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
			} else {
				return extractResultValue, nil
			}
		})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assessment_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.ParentColumnValue("assessment_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("control_set_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.ParentColumnValue("control_set_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_folder_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("EvidenceFolderId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assessment_report_selection").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AssessmentReportSelection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_account_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AwsAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_organization").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AwsOrganization")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compliance_check").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ComplianceCheck")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_source").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DataSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("EventName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_source").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("EventSource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_aws_account_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("EvidenceAwsAccountId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evidence_by_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("EvidenceByType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iam_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("IamId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("Time")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Attributes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources_included").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ResourcesIncluded")).Build(),
	}
}

func (x *TableAwsAuditmanagerEvidenceGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsAuditmanagerEvidenceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("auditmanager")
}
