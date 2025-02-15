package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsEc2NetworkAclsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2NetworkAclsGenerator{}

func (x *TableAwsEc2NetworkAclsGenerator) GetTableName() string {
	return "aws_ec2_network_acls"
}

func (x *TableAwsEc2NetworkAclsGenerator) GetTableDescription() string {
	return "https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAcl.html"
}

func (x *TableAwsEc2NetworkAclsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2NetworkAclsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsEc2NetworkAclsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ec2.DescribeNetworkAclsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Ec2
			for {
				output, err := svc.DescribeNetworkAcls(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.NetworkAcls
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEc2NetworkAclsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2NetworkAclsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.NetworkAcl)
					a := arn.ARN{
						Partition: cl.Partition,
						Service:   "ec2",
						Region:    cl.Region,
						AccountID: cl.AccountID,
						Resource:  "network-acl/" + aws.ToString(item.NetworkAclId),
					}
					return a.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Associations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_acl_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkAclId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnerId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entries").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Entries")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsDefault")).Build(),
	}
}

func (x *TableAwsEc2NetworkAclsGenerator) GetSubTables() []*schema.Table {
	return nil
}
