# Table: aws_lightsail_instance_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| location | json | X | √ |  | 
| size_in_gb | big_int | X | √ |  | 
| from_bundle_id | string | X | √ |  | 
| from_attached_disks | json | X | √ |  | 
| is_from_auto_snapshot | bool | X | √ |  | 
| progress | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| from_instance_name | string | X | √ |  | 
| from_blueprint_id | string | X | √ |  | 
| from_instance_arn | string | X | √ |  | 
| state | string | X | √ |  | 
| support_code | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 


