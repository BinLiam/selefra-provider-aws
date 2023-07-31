# Table: aws_lightsail_databases

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| latest_restorable_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| pending_modified_values | json | X | √ |  | 
| relational_database_blueprint_id | string | X | √ |  | 
| backup_retention_enabled | bool | X | √ |  | 
| hardware | json | X | √ |  | 
| preferred_maintenance_window | string | X | √ |  | 
| secondary_availability_zone | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| engine | string | X | √ |  | 
| preferred_backup_window | string | X | √ |  | 
| ca_certificate_identifier | string | X | √ |  | 
| master_endpoint | json | X | √ |  | 
| state | string | X | √ |  | 
| support_code | string | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| publicly_accessible | bool | X | √ |  | 
| master_username | string | X | √ |  | 
| location | json | X | √ |  | 
| master_database_name | string | X | √ |  | 
| parameter_apply_status | string | X | √ |  | 
| pending_maintenance_actions | json | X | √ |  | 
| relational_database_bundle_id | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| engine_version | string | X | √ |  | 


