# Table: aws_elasticache_reserved_cache_nodes_offerings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| duration | big_int | X | √ |  | 
| offering_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cache_node_type | string | X | √ |  | 
| fixed_price | float | X | √ |  | 
| product_description | string | X | √ |  | 
| reserved_cache_nodes_offering_id | string | X | √ |  | 
| usage_price | float | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 


