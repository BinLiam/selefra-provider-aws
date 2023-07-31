# Table: aws_lightsail_disk_snapshots

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| location | json | X | √ |  | 
| name | string | X | √ |  | 
| support_code | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| disk_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| from_disk_arn | string | X | √ |  | 
| from_disk_name | string | X | √ |  | 
| from_instance_name | string | X | √ |  | 
| is_from_auto_snapshot | bool | X | √ |  | 
| progress | string | X | √ |  | 
| size_in_gb | big_int | X | √ |  | 
| aws_lightsail_disks_selefra_id | string | X | X | fk to aws_lightsail_disks.selefra_id | 
| created_at | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| from_instance_arn | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | X | √ |  | 


