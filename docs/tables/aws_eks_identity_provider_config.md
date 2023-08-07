# Table: aws_eks_identity_provider_config

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_eks_clusters_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| client_id | string | X | √ |  | 
| cluster_name | string | X | √ |  | 
| arn | string | X | √ |  | 
| groups_claim | string | X | √ |  | 
| groups_prefix | string | X | √ |  | 
| issuer_url | string | X | √ |  | 
| username_claim | string | X | √ |  | 
| status | string | X | √ |  | 
| username_prefix | string | X | √ |  | 
| required_claims | json | X | √ |  | 
| tags | json | X | √ |  | 


