# Table: aws_rds_engine_versions

## Primary Keys 

```
account_id, region, _engine_version_hash
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| supported_engine_modes | string_array | X | √ |  | 
| supported_nchar_character_sets | json | X | √ |  | 
| supports_global_databases | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| database_installation_files_s3_bucket_name | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| supports_parallel_query | bool | X | √ |  | 
| supports_read_replica | bool | X | √ |  | 
| _engine_version_hash | string | X | √ |  | 
| default_character_set | json | X | √ |  | 
| supported_character_sets | json | X | √ |  | 
| image | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| supports_babelfish | bool | X | √ |  | 
| create_time | timestamp | X | √ |  | 
| db_engine_version_arn | string | X | √ |  | 
| supported_ca_certificate_identifiers | string_array | X | √ |  | 
| supported_feature_names | string_array | X | √ |  | 
| supports_certificate_rotation_without_restart | bool | X | √ |  | 
| tag_list | json | X | √ |  | 
| region | string | X | √ |  | 
| db_engine_media_type | string | X | √ |  | 
| db_engine_version_description | string | X | √ |  | 
| supports_log_exports_to_cloudwatch_logs | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| database_installation_files_s3_prefix | string | X | √ |  | 
| supported_timezones | json | X | √ |  | 
| valid_upgrade_target | json | X | √ |  | 
| custom_db_engine_version_manifest | string | X | √ |  | 
| status | string | X | √ |  | 
| db_engine_description | string | X | √ |  | 
| engine | string | X | √ |  | 
| major_engine_version | string | X | √ |  | 
| db_parameter_group_family | string | X | √ |  | 
| exportable_log_types | string_array | X | √ |  | 


