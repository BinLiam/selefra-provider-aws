# Table: aws_directconnect_connections

## Primary Keys 

```
arn, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| mac_sec_keys | json | X | √ |  | 
| region | string | X | √ |  | 
| vlan | big_int | X | √ |  | 
| connection_state | string | X | √ |  | 
| aws_logical_device_id | string | X | √ |  | 
| encryption_mode | string | X | √ |  | 
| jumbo_frame_capable | bool | X | √ |  | 
| location | string | X | √ |  | 
| owner_account | string | X | √ |  | 
| provider_name | string | X | √ |  | 
| tags | json | X | √ |  | 
| aws_device_v2 | string | X | √ |  | 
| account_id | string | X | √ |  | 
| has_logical_redundancy | string | X | √ |  | 
| aws_device | string | X | √ |  | 
| connection_id | string | X | √ |  | 
| connection_name | string | X | √ |  | 
| lag_id | string | X | √ |  | 
| loa_issue_time | timestamp | X | √ |  | 
| mac_sec_capable | bool | X | √ |  | 
| partner_name | string | X | √ |  | 
| port_encryption_status | string | X | √ |  | 
| bandwidth | string | X | √ |  | 
| id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | X | √ |  | 


