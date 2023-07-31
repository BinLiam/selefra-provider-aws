# Table: aws_appstream_application_fleet_associations

## Primary Keys 

```
application_arn, fleet_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| application_arn | string | X | √ |  | 
| fleet_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_appstream_applications_selefra_id | string | X | X | fk to aws_appstream_applications.selefra_id | 


