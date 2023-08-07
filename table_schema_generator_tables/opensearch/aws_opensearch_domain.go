package opensearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"

	"github.com/aws/aws-sdk-go-v2/service/opensearch"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsOpensearchDomainGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsOpensearchDomainGenerator{}

func (x *TableAwsOpensearchDomainGenerator) GetTableName() string {
	return "aws_opensearch_domain"
}

func (x *TableAwsOpensearchDomainGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsOpensearchDomainGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsOpensearchDomainGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsOpensearchDomainGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("DomainId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_policies").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("AccessPolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Deleted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("EngineVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("processing").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Processing")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("upgrade_processing").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("UpgradeProcessing")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_to_node_encryption_options_enabled").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("NodeToNodeEncryptionOptions.Enabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("AdvancedOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_security_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("AdvancedSecurityOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_tune_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("AutoTuneOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_config").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ClusterConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cognito_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("CognitoOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_endpoint_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("DomainEndpointOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ebs_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("EBSOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_at_rest_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("EncryptionAtRestOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoints").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Endpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_publishing_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("LogPublishingOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_software_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("ServiceSoftwareOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("SnapshotOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_options").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("VPCOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAwsOpensearchDomainGenerator) GetSubTables() []*schema.Table {
	return nil
}

type warpOpensearchDomain struct {
	*types.DomainStatus
	Tags map[string]string
}

func (x *TableAwsOpensearchDomainGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Opensearch

			input := &opensearch.ListDomainNamesInput{}
			op, err := svc.ListDomainNames(ctx, input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, name := range op.DomainNames {
				params := &opensearch.DescribeDomainInput{
					DomainName: name.DomainName,
				}
				data, err := svc.DescribeDomain(ctx, params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				wp := warpOpensearchDomain{
					DomainStatus: data.DomainStatus,
					Tags:         map[string]string{},
				}
				func() {
					ps := &opensearch.ListTagsInput{
						ARN: data.DomainStatus.ARN,
					}
					ops, err := svc.ListTags(ctx, ps)
					if err != nil {
						return
					}
					for _, tag := range ops.TagList {
						wp.Tags[*tag.Key] = *tag.Value
					}
				}()
				resultChannel <- wp
			}
			return nil
		},
	}
}

func (x *TableAwsOpensearchDomainGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}
