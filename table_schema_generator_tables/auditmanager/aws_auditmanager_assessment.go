package auditmanager

import (
	"context"

	"github.com/songzhibin97/go-ognl"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsAuditmanagerAssessmentGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAuditmanagerAssessmentGenerator{}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetTableName() string {
	return "aws_auditmanager_assessment"
}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Auditmanager
			var input auditmanager.ListAssessmentsInput
			var nextToken *string
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListAssessments(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, assessment := range output.AssessmentMetadata {

					opt, err := svc.GetAssessment(ctx, &auditmanager.GetAssessmentInput{AssessmentId: assessment.Id})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}

					resultChannel <- warpAssessment{
						AssessmentMetadataItem:          assessment,
						Arn:                             ognl.Get(opt, "Assessment.Arn"),
						AssessmentReportDestination:     ognl.Get(opt, "Assessment.Metadata.AssessmentReportsDestination.Destination"),
						AssessmentReportDestinationType: ognl.Get(opt, "Assessment.Metadata.AssessmentReportsDestination.DestinationType"),
						Description:                     ognl.Get(opt, "Assessment.Metadata.Description"),
						Framework:                       ognl.Get(opt, "Assessment.Framework"),
						Scope:                           ognl.Get(opt, "Assessment.Metadata.Scope"),
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

type warpAssessment struct {
	types.AssessmentMetadataItem
	Arn                             interface{}
	AssessmentReportDestination     interface{}
	AssessmentReportDestinationType interface{}
	Description                     interface{}
	Framework                       interface{}
	Scope                           interface{}
}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compliance_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ComplianceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assessment_report_destination").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AssessmentReportDestination")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assessment_report_destination_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AssessmentReportDestinationType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("LastUpdated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delegations").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Delegations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("framework").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Framework")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scope").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Scope")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("roles").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Roles")).Build(),
	}
}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("auditmanager")
}

func (x *TableAwsAuditmanagerAssessmentGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAuditmanagerEvidenceFolderGenerator{}),
	}
}
