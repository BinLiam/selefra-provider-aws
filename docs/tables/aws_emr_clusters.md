# Table: aws_emr_clusters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kerberos_attributes | json | X | √ |  | 
| arn | string | √ | √ |  | 
| log_encryption_kms_key_id | string | X | √ |  | 
| normalized_instance_hours | big_int | X | √ |  | 
| running_ami_version | string | X | √ |  | 
| visible_to_all_users | bool | X | √ |  | 
| region | string | X | √ |  | 
| cluster_arn | string | X | √ |  | 
| step_concurrency_level | big_int | X | √ |  | 
| auto_terminate | bool | X | √ |  | 
| configurations | json | X | √ |  | 
| account_id | string | X | √ |  | 
| os_release_label | string | X | √ |  | 
| outpost_arn | string | X | √ |  | 
| scale_down_behavior | string | X | √ |  | 
| service_role | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| security_configuration | string | X | √ |  | 
| applications | json | X | √ |  | 
| auto_scaling_role | string | X | √ |  | 
| custom_ami_id | string | X | √ |  | 
| ec2_instance_attributes | json | X | √ |  | 
| instance_collection_type | string | X | √ |  | 
| log_uri | string | X | √ |  | 
| placement_groups | json | X | √ |  | 
| ebs_root_volume_size | big_int | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| release_label | string | X | √ |  | 
| requested_ami_version | string | X | √ |  | 
| tags | json | X | √ |  | 
| master_public_dns_name | string | X | √ |  | 
| repo_upgrade_on_boot | string | X | √ |  | 
| status | json | X | √ |  | 
| termination_protected | bool | X | √ |  | 


