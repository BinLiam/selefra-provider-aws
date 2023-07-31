# Table: aws_ec2_route_tables

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| associations | json | X | √ |  | 
| propagating_vgws | json | X | √ |  | 
| route_table_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| owner_id | string | X | √ |  | 
| routes | json | X | √ |  | 
| tags | json | X | √ |  | 
| vpc_id | string | X | √ |  | 
| account_id | string | X | √ |  | 


