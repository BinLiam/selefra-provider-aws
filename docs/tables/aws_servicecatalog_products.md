# Table: aws_servicecatalog_products

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| product_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_time | timestamp | X | √ |  | 
| product_view_summary | json | X | √ |  | 
| source_connection | json | X | √ |  | 


