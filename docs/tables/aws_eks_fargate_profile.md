# Table: aws_eks_fargate_profile

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_eks_clusters_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| fargate_profile_name | string | X | √ |  | 
| cluster_name | string | X | √ |  | 
| fargate_profile_arn | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| pod_execution_role_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| selectors | json | X | √ |  | 
| subnets | json | X | √ |  | 
| tags | json | X | √ |  | 


