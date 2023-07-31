# Table: aws_ec2_reserved_instances

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| instance_count | big_int | X | √ |  | 
| instance_tenancy | string | X | √ |  | 
| offering_class | string | X | √ |  | 
| reserved_instances_id | string | X | √ |  | 
| duration | big_int | X | √ |  | 
| end | timestamp | X | √ |  | 
| usage_price | float | X | √ |  | 
| scope | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| availability_zone | string | X | √ |  | 
| currency_code | string | X | √ |  | 
| offering_type | string | X | √ |  | 
| product_description | string | X | √ |  | 
| tags | json | X | √ |  | 
| fixed_price | float | X | √ |  | 
| instance_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| start | timestamp | X | √ |  | 


