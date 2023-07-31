# Table: aws_lightsail_disks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| add_ons | json | X | √ |  | 
| attached_to | string | X | √ |  | 
| arn | string | √ | √ |  | 
| path | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| size_in_gb | big_int | X | √ |  | 
| state | string | X | √ |  | 
| is_attached | bool | X | √ |  | 
| gb_in_use | big_int | X | √ |  | 
| iops | big_int | X | √ |  | 
| is_system_disk | bool | X | √ |  | 
| location | json | X | √ |  | 
| support_code | string | X | √ |  | 
| region | string | X | √ |  | 
| attachment_state | string | X | √ |  | 


