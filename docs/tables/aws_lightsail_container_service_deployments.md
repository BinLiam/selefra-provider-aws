# Table: aws_lightsail_container_service_deployments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| public_endpoint | json | X | √ |  | 
| account_id | string | X | √ |  | 
| container_service_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| containers | json | X | √ |  | 
| state | string | X | √ |  | 
| version | big_int | X | √ |  | 
| region | string | X | √ |  | 
| aws_lightsail_container_services_selefra_id | string | X | X | fk to aws_lightsail_container_services.selefra_id | 


