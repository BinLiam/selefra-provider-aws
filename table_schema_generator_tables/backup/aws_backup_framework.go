package backup

import (
	"context"
	"errors"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/service/backup"

	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsBackupFrameworkGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupFrameworkGenerator{}

func (x *TableAwsBackupFrameworkGenerator) GetTableName() string {
	return "aws_backup_framework"
}

func (x *TableAwsBackupFrameworkGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupFrameworkGenerator) GetVersion() uint64 {
	return 0
}

type warpFramework struct {
	types.Framework
	FrameworkStatus   *string
	FrameworkControls []types.FrameworkControl
	Tags              map[string]string
}

func (x *TableAwsBackupFrameworkGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("framework_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FrameworkName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FrameworkArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("framework_description").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FrameworkDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DeploymentStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_controls").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("NumberOfControls")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("framework_status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("FrameworkStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("framework_controls").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("FrameworkControls")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAwsBackupFrameworkGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsBackupFrameworkGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"arn"},
	}
}

func (x *TableAwsBackupFrameworkGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Backup
			var nextToken *string
			var input backup.ListFrameworksInput
			for {
				input.NextToken = nextToken
				output, err := svc.ListFrameworks(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, framework := range output.Frameworks {
					wp := &warpFramework{
						Framework: framework,
						Tags:      make(map[string]string),
					}
					err = func() error {
						if c.GetRegion() == "ap-northeast-3" {
							return nil
						}
						params := &backup.DescribeFrameworkInput{
							FrameworkName: framework.FrameworkName,
						}
						op, err := svc.DescribeFramework(ctx, params)
						if err != nil {
							var ae smithy.APIError
							if errors.As(err, &ae) {
								if ae.ErrorCode() == "ResourceNotFoundException" {
									return nil
								}
							}
							return err
						}
						wp.FrameworkStatus = op.FrameworkStatus
						wp.FrameworkControls = op.FrameworkControls

						var nextTagsToken *string
						nextTagInput := backup.ListTagsInput{
							ResourceArn: framework.FrameworkArn,
						}
						for {
							nextTagInput.NextToken = nextTagsToken
							opt, err := svc.ListTags(ctx, &nextTagInput)
							if err != nil {
								return err
							}
							for k, v := range opt.Tags {
								wp.Tags[k] = v
							}
							if opt.NextToken == nil {
								break
							}
							nextTagsToken = opt.NextToken
						}

						return nil
					}()
					resultChannel <- wp
				}
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
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

func (x *TableAwsBackupFrameworkGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}
