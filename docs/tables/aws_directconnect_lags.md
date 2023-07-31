# Table: aws_directconnect_lags

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| allows_hosted_connections | bool | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| connections | json | X | √ |  | 
| has_logical_redundancy | string | X | √ |  | 
| mac_sec_keys | json | X | √ |  | 
| minimum_links | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_device | string | X | √ |  | 
| aws_logical_device_id | string | X | √ |  | 
| mac_sec_capable | bool | X | √ |  | 
| provider_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| tags | json | X | √ |  | 
| connections_bandwidth | string | X | √ |  | 
| encryption_mode | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| number_of_connections | big_int | X | √ |  | 
| owner_account | string | X | √ |  | 
| region | string | X | √ |  | 
| lag_id | string | X | √ |  | 
| lag_name | string | X | √ |  | 
| lag_state | string | X | √ |  | 
| location | string | X | √ |  | 


