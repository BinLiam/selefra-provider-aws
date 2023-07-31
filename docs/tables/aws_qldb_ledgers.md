# Table: aws_qldb_ledgers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| creation_date_time | timestamp | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| encryption_description | json | X | √ |  | 
| state | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| tags | json | X | √ | `The tags associated with the pipeline.` | 
| permissions_mode | string | X | √ |  | 
| name | string | X | √ |  | 


