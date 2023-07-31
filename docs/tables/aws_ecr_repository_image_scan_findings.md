# Table: aws_ecr_repository_image_scan_findings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_ecr_repository_images_selefra_id | string | X | X | fk to aws_ecr_repository_images.selefra_id | 
| image_digest | string | X | √ |  | 
| image_scan_status | json | X | √ |  | 
| repository_name | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| image_tag | string | X | √ |  | 
| image_scan_findings | json | X | √ |  | 
| registry_id | string | X | √ |  | 


