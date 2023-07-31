# Table: aws_ecr_repository_images

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| image_scan_findings_summary | json | X | √ |  | 
| region | string | X | √ |  | 
| aws_ecr_repositories_selefra_id | string | X | X | fk to aws_ecr_repositories.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| image_scan_status | json | X | √ |  | 
| image_size_in_bytes | int | X | √ |  | 
| image_tags | string_array | X | √ |  | 
| last_recorded_pull_time | timestamp | X | √ |  | 
| registry_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| image_digest | string | X | √ |  | 
| image_pushed_at | timestamp | X | √ |  | 
| artifact_media_type | string | X | √ |  | 
| image_manifest_media_type | string | X | √ |  | 
| repository_name | string | X | √ |  | 
| arn | string | √ | √ |  | 


