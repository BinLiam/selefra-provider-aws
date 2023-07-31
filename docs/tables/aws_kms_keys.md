# Table: aws_kms_keys

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rotation_enabled | bool | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| cloud_hsm_cluster_id | string | X | √ |  | 
| encryption_algorithms | string_array | X | √ |  | 
| pending_deletion_window_in_days | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| multi_region | bool | X | √ |  | 
| xks_key_configuration | json | X | √ |  | 
| enabled | bool | X | √ |  | 
| multi_region_configuration | json | X | √ |  | 
| signing_algorithms | string_array | X | √ |  | 
| origin | string | X | √ |  | 
| key_spec | string | X | √ |  | 
| key_state | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| custom_key_store_id | string | X | √ |  | 
| customer_master_key_spec | string | X | √ |  | 
| description | string | X | √ |  | 
| expiration_model | string | X | √ |  | 
| key_manager | string | X | √ |  | 
| valid_to | timestamp | X | √ |  | 
| replica_keys | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| key_id | string | X | √ |  | 
| aws_account_id | string | X | √ |  | 
| deletion_date | timestamp | X | √ |  | 
| key_usage | string | X | √ |  | 
| mac_algorithms | string_array | X | √ |  | 


