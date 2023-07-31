# Table: aws_apprunner_custom_domains

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| certificate_validation_records | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_apprunner_services_selefra_id | string | X | X | fk to aws_apprunner_services.selefra_id | 
| domain_name | string | X | √ |  | 
| enable_www_subdomain | bool | X | √ |  | 
| status | string | X | √ |  | 


