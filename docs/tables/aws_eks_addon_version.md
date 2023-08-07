# Table: aws_eks_addon_version

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_eks_clusters_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| addon_name | string | X | √ |  | 
| addon_version | string | X | √ |  | 
| type | string | X | √ |  | 
| addon_configuration | json | X | √ |  | 
| architecture | json | X | √ |  | 
| compatibilities | json | X | √ |  | 


