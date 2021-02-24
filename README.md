# UET_sale_management
A system manages sale of product, employee, provider, customer, bill, ..... with create, update, delete, request.<br/>
You can read more detail in folder [doc](https://github.com/tangducthinh456/UET_sale_management/blob/main/doc)

## Set Up
1. git clone https://github.com/tangducthinh456/UET_sale_management
2. Set up your PostgreSQL and Redis environment, then go to UET_sale_management/app/backend/config/config.yaml to fill your database information, the data field is the number of test object
2. cd UET_sale_management/app/backend<br/>go build<br/>go run main.go
3. cd UET_sale_management/app/frontend/vue-material-dashboard-master<br/>npm run dev