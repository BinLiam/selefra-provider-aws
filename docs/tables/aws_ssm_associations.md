# Table: aws_ssm_associations

## Primary Keys 

```
account_id, region, association_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| association_id | string | X | √ |  | 
| association_version | string | X | √ |  | 
| name | string | X | √ |  | 
| schedule_offset | big_int | X | √ |  | 
| targets | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| association_name | string | X | √ |  | 
| instance_id | string | X | √ |  | 
| last_execution_date | timestamp | X | √ |  | 
| schedule_expression | string | X | √ |  | 
| target_maps | json | X | √ |  | 
| document_version | string | X | √ |  | 
| overview | json | X | √ |  | 
| region | string | X | √ |  | 


