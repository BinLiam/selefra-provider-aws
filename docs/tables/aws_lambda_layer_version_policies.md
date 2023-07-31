# Table: aws_lambda_layer_version_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| layer_version | int | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| policy | string | X | √ |  | 
| revision_id | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| layer_version_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_lambda_layer_versions_selefra_id | string | X | X | fk to aws_lambda_layer_versions.selefra_id | 


