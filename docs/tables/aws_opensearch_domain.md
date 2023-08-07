# Table: aws_opensearch_domain

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| domain_name | string | X | √ |  | 
| domain_id | string | X | √ |  | 
| arn | string | X | √ |  | 
| access_policies | string | X | √ |  | 
| created | bool | X | √ |  | 
| deleted | bool | X | √ |  | 
| endpoint | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| processing | bool | X | √ |  | 
| upgrade_processing | bool | X | √ |  | 
| node_to_node_encryption_options_enabled | bool | X | √ |  | 
| advanced_options | json | X | √ |  | 
| advanced_security_options | json | X | √ |  | 
| auto_tune_options | json | X | √ |  | 
| cluster_config | json | X | √ |  | 
| cognito_options | json | X | √ |  | 
| domain_endpoint_options | json | X | √ |  | 
| ebs_options | json | X | √ |  | 
| encryption_at_rest_options | json | X | √ |  | 
| endpoints | json | X | √ |  | 
| log_publishing_options | json | X | √ |  | 
| service_software_options | json | X | √ |  | 
| snapshot_options | json | X | √ |  | 
| vpc_options | json | X | √ |  | 
| tags | json | X | √ |  | 


