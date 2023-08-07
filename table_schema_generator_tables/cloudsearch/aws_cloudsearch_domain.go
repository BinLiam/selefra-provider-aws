package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"

	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-aws/aws_client"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsCloudsearchDomainGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudsearchDomainGenerator{}

func (x *TableAwsCloudsearchDomainGenerator) GetTableName() string {
	return "aws_cloudsearch_domain"
}

func (x *TableAwsCloudsearchDomainGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudsearchDomainGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudsearchDomainGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{"arn"},
	}
}

func (x *TableAwsCloudsearchDomainGenerator) GetColumns() []*schema.Column {
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
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Deleted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("processing").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("Processing")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requires_index_documents").ColumnType(schema.ColumnTypeBool).Extractor(column_value_extractor.StructSelector("RequiresIndexDocuments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("search_instance_count").ColumnType(schema.ColumnTypeBigInt).Extractor(column_value_extractor.StructSelector("SearchInstanceCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("search_instance_type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("SearchInstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("search_partition_count").ColumnType(schema.ColumnTypeInt).Extractor(column_value_extractor.StructSelector("SearchPartitionCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("doc_service").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("DocService")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("limits").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Limits")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("search_service").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("SearchService")).Build(),
	}
}

func (x *TableAwsCloudsearchDomainGenerator) GetSubTables() []*schema.Table {
	return nil
}

func (x *TableAwsCloudsearchDomainGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().CloudSearch
			input := &cloudsearch.ListDomainNamesInput{}
			resp, err := svc.ListDomainNames(ctx, input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for domain := range resp.DomainNames {
				params := &cloudsearch.DescribeDomainsInput{
					DomainNames: []string{domain},
				}
				item, err := svc.DescribeDomains(ctx, params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				if len(item.DomainStatusList) > 0 {
					resultChannel <- item.DomainStatusList[0]
				}
			}
			return nil
		},
	}
}

func (x *TableAwsCloudsearchDomainGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cloudsearch")
}
