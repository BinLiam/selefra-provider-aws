# Table: aws_elasticsearch_domains

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| change_progress_details | json | X | √ |  | 
| region | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| vpc_options | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| access_policies | string | X | √ |  | 
| elasticsearch_version | string | X | √ |  | 
| service_software_options | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| advanced_security_options | json | X | √ |  | 
| domain_id | string | X | √ |  | 
| created | bool | X | √ |  | 
| encryption_at_rest_options | json | X | √ |  | 
| authorized_principals | json | X | √ |  | 
| arn | string | √ | √ |  | 
| ebs_options | json | X | √ |  | 
| processing | bool | X | √ |  | 
| log_publishing_options | json | X | √ |  | 
| elasticsearch_cluster_config | json | X | √ |  | 
| advanced_options | json | X | √ |  | 
| auto_tune_options | json | X | √ |  | 
| cognito_options | json | X | √ |  | 
| deleted | bool | X | √ |  | 
| endpoint | string | X | √ |  | 
| endpoints | json | X | √ |  | 
| snapshot_options | json | X | √ |  | 
| domain_endpoint_options | json | X | √ |  | 
| node_to_node_encryption_options | json | X | √ |  | 
| upgrade_processing | bool | X | √ |  | 


