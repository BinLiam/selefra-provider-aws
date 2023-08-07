# Table: aws_eks_node_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_eks_clusters_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| nodegroup_name | string | X | √ |  | 
| arn | string | X | √ |  | 
| cluster_name | not_assign | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| ami_type | string | X | √ |  | 
| capacity_type | string | X | √ |  | 
| disk_size | int | X | √ |  | 
| modified_at | timestamp | X | √ |  | 
| node_role | string | X | √ |  | 
| release_version | string | X | √ |  | 
| version | string | X | √ |  | 
| health | json | X | √ |  | 
| instance_types | json | X | √ |  | 
| labels | json | X | √ |  | 
| launch_template | json | X | √ |  | 
| remote_access | json | X | √ |  | 
| resources | json | X | √ |  | 
| scaling_config | json | X | √ |  | 
| subnets | json | X | √ |  | 
| tags | json | X | √ |  | 
| taints | json | X | √ |  | 
| update_config | json | X | √ |  | 


