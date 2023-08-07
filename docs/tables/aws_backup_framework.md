# Table: aws_backup_framework

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| framework_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| framework_description | string | X | √ |  | 
| deployment_status | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| number_of_controls | int | X | √ |  | 
| framework_status | string | X | √ |  | 
| framework_controls | json | X | √ |  | 
| tags | json | X | √ |  | 


