# Property Management
This project is used to store properties for either sale or rent.

## How to Run
### Database
Spin up a postgres docker image and create below tables in the dev database.
```sql
create table country_details(
id uuid not null,
name character varying not null,
primary key(id)
);

create table house_details(
id uuid not null,
name character varying,
address character varying,
locality character varying,
pincode character varying,
country_id uuid not null,
amount decimal,
type character varying,
primary key(id),
foreign key(country_id) references country_details(id)
);
insert into country_details(id, name) values(uuid_in((md5((random())::text))::cstring), 'Singapore');
insert into country_details(id, name) values(uuid_in((md5((random())::text))::cstring), 'Malayasia');
insert into country_details(id, name) values(uuid_in((md5((random())::text))::cstring), 'Phillipines');
```
### Sample Requests

#### Sell House
http://localhost:8080/api/property/sell/v1
```json
{
"name" : "1234",
"locality" : "test",
"address" : "test1345",
"country" : "Singapore",
"pinCode" : "test1345",
"amount" : 50000
}
```

#### Rent House
http://localhost:8080/api/property/rent/v1
```json
{
"name" : "1234",
"locality" : "test",
"address" : "test1345",
"country" : "Singapore",
"pinCode" : "test1345",
"amount" : 50000
}
```

#### Find House
http://localhost:8080/api/property/find/v1/Singapore/test/RENT
```json
[
    {
        "name": "test",
        "address": "test",
        "locality": "test",
        "country": "Singapore",
        "pinCode": "test",
        "Amount": 5000
    }
]
```