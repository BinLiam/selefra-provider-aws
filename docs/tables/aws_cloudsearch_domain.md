# Table: aws_cloudsearch_domain

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
| domain_name | string | X | √ |  | 
| domain_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created | bool | X | √ |  | 
| deleted | bool | X | √ |  | 
| processing | bool | X | √ |  | 
| requires_index_documents | bool | X | √ |  | 
| search_instance_count | big_int | X | √ |  | 
| search_instance_type | string | X | √ |  | 
| search_partition_count | int | X | √ |  | 
| doc_service | json | X | √ |  | 
| limits | json | X | √ |  | 
| search_service | json | X | √ |  | 


