# Mini Project Golang 2022

1 domain (campaigns), 1 domain for register and login (users).

Creating a 2 domain REST API with gin, gorm, PostgreSQL.

## Inside mini project

- CRUD Campaigns
- Login User
- Register User
- JWT Token
- Service Object Pattern

## Postgre SQL

SQL (Query Language) used in this project

### users

```sql
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "user_id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar(20) NOT NULL,
  "occupasion" varchar(20) NOT NULL,
  "email" varchar(255) NOT NULL,
  "password_hash" varchar(255) NOT NULL,
  "avatar_field_name" varchar(255) NOT NULL,
  "role" varchar(25) DEFAULT NULL,
  "token" varchar(255) NOT null,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
```

### campaigns

```sql
DROP TABLE IF EXISTS "campaigns";
CREATE TABLE "campaigns" (
  "campaign_id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "campaign_name" varchar(255) NOT NULL,
  "short_description" varchar(255) NOT NULL,
  "description" text,
  "goal_amount" int NOT NULL,
  "current_amount" int NOT NULL,
  "perks" text,
  "backer_count" int NOT null,
  "slug" varchar(255) NOT null,
  "created_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
```

## REST API

There are 9 endpoints that can be used, including:

### GET CAMPAIGNS

```
/api/v1/users/campaigns?page=0&limit=5
```

### EXAMPLE RESULT GET CAMPAIGNS

```
{
    "meta": {
        "message": "Data of Campaigns",
        "code": 200,
        "status": "Success"
    },
    "data": {
        "limit": 5,
        "page": 0,
        "total_rows": 5,
        "from_row": 1,
        "to_row": 0,
        "rows": [
            {
                "campaign_id": 2,
                "user_id": 20,
                "CampaignName": "Campaign3",
                "ShortDescription": "Campaign untuk id 3 Upadted",
                "Description": "Campaign untuk id 3 Upadted oleh user_id 15 yang login",
                "Perks": "Keuntungan tiada tanding",
                "backer_count": 0,
                "GoalAmount": 6660000,
                "CurrentAmount": 0,
                "Slug": "sebuah-campaign-2-14",
                "CreatedAt": "2022-08-09T11:37:15.407983Z",
                "UpdatedAt": "2022-08-09T12:03:53.841707Z",
                "CampaignImages": []
            },
            {
                "campaign_id": 3,
                "user_id": 20,
                "CampaignName": "Campaign3",
                "ShortDescription": "Campaign untuk id 3 Upadted",
                "Description": "Campaign untuk id 3 Upadted oleh user_id 15 yang login",
                "Perks": "Keuntungan tiada tanding",
                "backer_count": 0,
                "GoalAmount": 6660000,
                "CurrentAmount": 0,
                "Slug": "sebuah-campaign-2-14",
                "CreatedAt": "2022-08-09T11:37:15.407983Z",
                "UpdatedAt": "2022-08-09T12:03:53.841707Z",
                "CampaignImages": []
            }
        ]
    }
}
```

### GET CAMPAIGNS

```
/api/v1/users/campaigns/2
```

### EXAMPLE RESULT GET CAMPAIGNS

```
{
    "meta": {
        "message": "Campaign Detail",
        "code": 200,
        "status": "Success"
    },
    "data": {
        "campaign_id": 2,
        "user_id": 20,
        "CampaignName": "Campaign3",
        "ShortDescription": "Campaign untuk id 3 Upadted",
        "Description": "Campaign untuk id 3 Upadted oleh user_id 15 yang login",
        "Perks": "Keuntungan tiada tanding",
        "backer_count": 0,
        "GoalAmount": 6660000,
        "CurrentAmount": 0,
        "Slug": "sebuah-campaign-2-14",
        "CreatedAt": "2022-08-09T11:37:15.407983Z",
        "UpdatedAt": "2022-08-09T12:03:53.841707Z",
        "CampaignImages": [
            {
                "id": 2,
                "campaign_id": 2,
                "FileName": "satu.png",
                "IsPrimary": 0,
                "CreatedAt": "2022-08-08T16:56:44.878949Z",
                "UpdatedAt": "2022-08-08T17:35:15.442587Z"
            }
        ]
    }
}
```

### PUT CAMPAIGNS

```
/api/v1/users/campaigns/44
```

### BODY PUT CAMPAIGNS

```
{
    "name": "Campaign3 update",
    "short_description": "Campaign untuk id 3 Upadted",
    "description": "Campaign untuk id 3 Upadted oleh update",
    "goal_amount": 666000,
    "perks": "Keuntungan update"
}
```

### EXAMPLE RESULT PUT CAMPAIGNS

```
{
    "meta": {
        "message": "Success Update Campaign",
        "code": 200,
        "status": "Success"
    },
    "data": {
        "campaign_id": 44,
        "user_id": 14,
        "CampaignName": "Campaign3 update",
        "ShortDescription": "Campaign untuk id 3 Upadted",
        "Description": "Campaign untuk id 3 Upadted oleh update",
        "Perks": "Keuntungan update",
        "backer_count": 0,
        "GoalAmount": 666000,
        "CurrentAmount": 0,
        "Slug": "sebuah-campaign-5-14",
        "CreatedAt": "2022-08-09T11:51:12.258677Z",
        "UpdatedAt": "2022-08-09T21:49:50.5615314+07:00",
        "CampaignImages": []
    }
}
```

### POST CAMPAIGNS

```
/api/v1/users/campaigns
```

### BODY POST CAMPAIGNS

```
{
    "name": "Sebuah Campaign Keren",
    "short_description": "Campaign Sangat Keren",
    "description": "Campaign Description Keren",
    "goal_amount": 189000,
    "perks": "Campaign untuk kemanusiaan yang keren"
}
```

### EXAMPLE RESULT POST CAMPAIGNS

```
{
    "meta": {
        "message": "Success to Create Campaign",
        "code": 200,
        "status": "Success"
    },
    "data": {
        "id": 51,
        "user_id": 13,
        "name": "Sebuah Campaign Keren",
        "short_description": "Campaign Sangat Keren",
        "image_url": "",
        "goal_amount": 189000,
        "current_amount": 0,
        "slug": "sebuah-campaign-keren-13"
    }
}
```

### DELETE SINGLE CAMPAIGNS

```
/api/v1/users/campaigns/4
```

### EXAMPLE RESULT DELETE SINGLE CAMPAIGNS

```
{
    "meta": {
        "message": "Success Delete Campaign",
        "code": 200,
        "status": "Success"
    },
    "data": null
}
```

### REGISTER USERS

```
/api/v1/users/register
```

### BODY POST USERS

```
{
    "name": "Mr ol",
    "email": "Batman@gmail.com",
    "occupasion": "Senior Backend",
    "password": "rahasia"
}
```

### EXAMPLE RESULT POST USERS

```
{
    "meta": {
        "message": "Register member success!",
        "code": 200,
        "status": "success"
    },
    "data": {
        "id": 8,
        "email": "kuy@handy.com",
        "role": "user",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo4fQ.UXQdR3dLOvVSgBHhLU8YSGcHeJF7TNrOFMewS6AkWYE"
    }
}
```

### LOGIN USERS

```
/api/v1/users/login
```

### BODY POST LOGIN

```
{
    "email": "php@gmail.com",
    "password": "rahasia"
}
```

### EXAMPLE RESULT LOGIN

```
{
    "meta": {
        "message": "Login Successfull",
        "code": 422,
        "status": "SUCCESS"
    },
    "data": {
        "user_id": 14,
        "name": "ANTMAN",
        "occupasion": "Backend",
        "email": "php@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNH0.6UikbmiplpfxXfV3enXpjcPBYBGE3JEBAKxYbw5zJ0o"
    }
}
```

### POST AVATAR

```
/api/v1/users/avatars
```

### BODY POST/UPLOAD AVATAR

#### form-data

key -> avatar (file), value -> fileName.jpg

### EXAMPLE RESULT UPLOAD AVATAR

```
{
    "meta": {
        "message": "Avatar Successfuly Uploaded",
        "code": 200,
        "status": "Success"
    },
    "data": {
        "is_uploaded": true
    }
}
```

### POST EMAIL CHECK

```
/api/v1/users/emailChecker
```

### BODY POST EMAIL CHECK

#### form-data

```
{
    "email": "any@gmail.com"
}
```

### EXAMPLE RESULT UPLOAD AVATAR

```
{
    "meta": {
        "message": "Email is Available",
        "code": 200,
        "status": "Success"
    },
    "data": {
        "is_available": true
    }
}
```

## Created By

Dwi Yulianto
