# Table: aws_lightsail_container_service_images

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| container_service_arn | string | X | √ |  | 
| aws_lightsail_container_services_selefra_id | string | X | X | fk to aws_lightsail_container_services.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| created_at | timestamp | X | √ |  | 
| digest | string | X | √ |  | 
| image | string | X | √ |  | 
| account_id | string | X | √ |  | 


