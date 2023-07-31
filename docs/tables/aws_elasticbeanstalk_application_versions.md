# Table: aws_elasticbeanstalk_application_versions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| application_name | string | X | √ |  | 
| description | string | X | √ |  | 
| date_created | timestamp | X | √ |  | 
| source_build_information | json | X | √ |  | 
| source_bundle | json | X | √ |  | 
| status | string | X | √ |  | 
| version_label | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| application_version_arn | string | X | √ |  | 
| build_arn | string | X | √ |  | 
| date_updated | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


