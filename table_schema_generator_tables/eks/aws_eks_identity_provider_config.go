package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableAwsEksIdentityProviderConfigGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEksIdentityProviderConfigGenerator{}

func (x *TableAwsEksIdentityProviderConfigGenerator) GetTableName() string {
	return "aws_eks_identity_provider_config"
}

func (x *TableAwsEksIdentityProviderConfigGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEksIdentityProviderConfigGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEksIdentityProviderConfigGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEksIdentityProviderConfigGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_eks_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),

		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_id").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ClientId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_name").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("ClusterName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("IdentityProviderConfigArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups_claim").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("GroupsClaim")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups_prefix").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("GroupsPrefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuer_url").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("IssuerUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username_claim").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("UsernameClaim")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username_prefix").ColumnType(schema.ColumnTypeString).Extractor(column_value_extractor.StructSelector("UsernamePrefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("required_claims").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("RequiredClaims")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Extractor(column_value_extractor.StructSelector("Tags")).Build(),
	}
}

func (x *TableAwsEksIdentityProviderConfigGenerator) GetSubTables() []*schema.Table {
	return nil
}

type warpEKSIdentityProviderConfig struct {
	Name *string
	Type *string
	*types.OidcIdentityProviderConfig
}

func (x *TableAwsEksIdentityProviderConfigGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*types.Cluster)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Eks
			var nextToken *string
			input := eks.ListIdentityProviderConfigsInput{
				ClusterName: r.Name,
			}
			for {
				if nextToken != nil {
					input.NextToken = nextToken
				}
				output, err := svc.ListIdentityProviderConfigs(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, config := range output.IdentityProviderConfigs {
					params := &eks.DescribeIdentityProviderConfigInput{
						ClusterName: r.Name,
						IdentityProviderConfig: &types.IdentityProviderConfig{
							Name: config.Name,
							Type: config.Type,
						},
					}
					op, err := svc.DescribeIdentityProviderConfig(ctx, params)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)
					}
					resultChannel <- warpEKSIdentityProviderConfig{
						Name:                       config.Name,
						Type:                       config.Type,
						OidcIdentityProviderConfig: op.IdentityProviderConfig.Oidc,
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

func (x *TableAwsEksIdentityProviderConfigGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("eks")
}
