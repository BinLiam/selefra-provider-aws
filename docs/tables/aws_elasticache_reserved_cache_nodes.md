# Table: aws_elasticache_reserved_cache_nodes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| reservation_arn | string | X | √ |  | 
| reserved_cache_node_id | string | X | √ |  | 
| reserved_cache_nodes_offering_id | string | X | √ |  | 
| start_time | timestamp | X | √ |  | 
| duration | big_int | X | √ |  | 
| offering_type | string | X | √ |  | 
| recurring_charges | json | X | √ |  | 
| state | string | X | √ |  | 
| cache_node_count | big_int | X | √ |  | 
| cache_node_type | string | X | √ |  | 
| fixed_price | float | X | √ |  | 
| product_description | string | X | √ |  | 
| usage_price | float | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 


