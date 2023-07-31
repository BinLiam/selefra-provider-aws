# Table: aws_elasticbeanstalk_applications

## Primary Keys 

```
arn, date_created
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| date_created | timestamp | X | √ |  | 
| resource_lifecycle_config | json | X | √ |  | 
| versions | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| application_name | string | X | √ |  | 
| configuration_templates | string_array | X | √ |  | 
| date_updated | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| application_arn | string | X | √ |  | 


