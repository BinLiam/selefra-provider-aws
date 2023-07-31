# Table: aws_mq_brokers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine_type | string | X | √ |  | 
| pending_authentication_strategy | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| ldap_server_metadata | json | X | √ |  | 
| logs | json | X | √ |  | 
| region | string | X | √ |  | 
| auto_minor_version_upgrade | bool | X | √ |  | 
| configurations | json | X | √ |  | 
| pending_ldap_server_metadata | json | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| users | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| broker_name | string | X | √ |  | 
| broker_state | string | X | √ |  | 
| encryption_options | json | X | √ |  | 
| host_instance_type | string | X | √ |  | 
| pending_engine_version | string | X | √ |  | 
| pending_security_groups | string_array | X | √ |  | 
| arn | string | √ | √ |  | 
| deployment_mode | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| maintenance_window_start_time | json | X | √ |  | 
| storage_type | string | X | √ |  | 
| broker_id | string | X | √ |  | 
| broker_instances | json | X | √ |  | 
| created | timestamp | X | √ |  | 
| authentication_strategy | string | X | √ |  | 
| broker_arn | string | X | √ |  | 
| pending_host_instance_type | string | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| actions_required | json | X | √ |  | 


