# Table: aws_eks_addon

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_eks_clusters_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| addon_name | string | X | √ |  | 
| arn | string | X | √ |  | 
| cluster_name | string | X | √ |  | 
| addon_version | string | X | √ |  | 
| status | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| modified_at | timestamp | X | √ |  | 
| service_account_role_arn | string | X | √ |  | 
| health_issues | json | X | √ |  | 


