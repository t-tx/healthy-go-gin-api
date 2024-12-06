## Sign up 

- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/signup' \
    --header 'Content-Type: application/json' \
    --data '{
        "username":"trung",
        "password":"AA1234aa",
        "gender": "male"
    }'
    ```
- Response:
    ```json
    {
        "success": true
    }
    ```



## Sign in
- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/signin' \
    --header 'Content-Type: application/json' \
    --data '{
        "username":"trung",
        "password":"AA1234aa"
    }'
    ```
- Response:
    ```json
    {
        "success": true,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRydW5nIiwiZXhwIjoxNzMyMTg0NzE2fQ.xnnHeUwlL_MbxDx1EDTmdRbYuhNrg8hJEeeaN2Fjsvw",
        "user": {
            "ID": 0,
            "Username": "trung",
            "Gender": "male",
            "Birthday": "",
            "Password": ""
        }
    }
    ```

## Get articles in health topic
- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/articles/health'
    ```
- Response:
    ```json
    {
        "data": [
            {
                "image_urls": {
                    "big_img_url": "https://example.com/images/healthy-big1.jpg",
                    "small_img_url": "https://example.com/images/healthy-small1.jpg"
                },
                "title": "Healthy Eating for Beginners",
                "uploaded_time": "2024-11-20T10:00:00Z",
                "tags": [
                    "This article provides an introduction to healthy eating and how to get started."
                ]
            },
        ],
        "success": true
    }
    ```

## Get global config 

- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/config/global'
    ```
- Response:
    ```json
    {
        "data": [
            {
                "key": "api-version",
                "value": "v0.0.9"
            }
        ],
        "success": true
    }
    ```

## Add diaries

- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/user/diaries' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRydW5nIiwiZXhwIjoxNzMyMTg0NzE2fQ.xnnHeUwlL_MbxDx1EDTmdRbYuhNrg8hJEeeaN2Fjsvw' \
    --data '[
        {
            "title":"my diary",
            "content":"some diary content"
        }
    ]'
    ```
- Response:
    ```json
    {
        "success": true
    }
    ```


## Get diaries

- Request:
    ```bash
    curl --location --request GET 'http://localhost:8080/api/v1/user/diaries' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRydW5nIiwiZXhwIjoxNzMyMTg0NzE2fQ.xnnHeUwlL_MbxDx1EDTmdRbYuhNrg8hJEeeaN2Fjsvw'
    ```
- Response:
    ```json
    {
        "data": [
            {
                "username": "trung",
                "time": "2024-11-20T10:40:17Z",
                "title": "my diary",
                "content": "some diary content"
            }
        ],
        "success": true
    }
    ```


## Add user events
### Note: 
For each `event_type`, there is a corresponding `content` value, only some `event_type` have aggregate reports that are useful to add; adding unsupported `event_type` is not beneficial.

### Curl
- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/user/events' \
    --header 'Content-Type: application/json' \
    --header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRydW5nIiwiZXhwIjoxNzMyMTg0NzE2fQ.xnnHeUwlL_MbxDx1EDTmdRbYuhNrg8hJEeeaN2Fjsvw' \
    --data '[
        {
            "event_type":"exercise",
            "content": {
                "type": "running",
                "duration": 10
            }
        },
        {
            "event_type":"exercise",
            "content": {
                "type": "walking",
                "duration": 10
            }
        },
        {
            "event_type":"meal",
            "content": {
                "type": "dinner",
                "dish": [
                    {"food": "egg", "quantity": 2},
                    {"food": "sashimi", "quantity": 2}
                ]
            }
        },
        {
            "event_type":"exercise",
            "content": {
                "type": "walking",
                "duration": 10
            }
        },
        {
            "event_type":"measure",
            "content": {
                "type": "weight",
                "value": 100
            }
        },
        {
            "event_type":"measure",
            "content": {
                "type": "height",
                "value": 100
            }
        }
    ]
    '
    ```
- Response:
    ```json
    {
        "success": true
    }
    ```


## Get user events

- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/user/events' \
    --header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRydW5nIiwiZXhwIjoxNzMyMTg0NzE2fQ.xnnHeUwlL_MbxDx1EDTmdRbYuhNrg8hJEeeaN2Fjsvw'
    ```
- Response:
    ```json
    {
        "data": [
            {
                "event_type": "exercise",
                "content": {
                    "duration": 10,
                    "type": "running"
                },
                "created_at": "2024-11-20T10:41:14Z"
            },
            {
                "event_type": "exercise",
                "content": {
                    "duration": 10,
                    "type": "walking"
                },
                "created_at": "2024-11-20T10:41:14Z"
            },
            {
                "event_type": "meal",
                "content": {
                    "dish": [
                        {
                            "food": "egg",
                            "quantity": 2
                        },
                        {
                            "food": "sashimi",
                            "quantity": 2
                        }
                    ],
                    "type": "dinner"
                },
                "created_at": "2024-11-20T10:41:14Z"
            },
            {
                "event_type": "exercise",
                "content": {
                    "duration": 10,
                    "type": "walking"
                },
                "created_at": "2024-11-20T10:41:14Z"
            },
            {
                "event_type": "measure",
                "content": {
                    "type": "weight",
                    "value": 100
                },
                "created_at": "2024-11-20T10:41:14Z"
            },
            {
                "event_type": "measure",
                "content": {
                    "type": "height",
                    "value": 100
                },
                "created_at": "2024-11-20T10:41:14Z"
            },
            {
                "event_type": "exercise",
                "content": {
                    "duration": 10,
                    "type": "running"
                },
                "created_at": "2024-11-20T10:27:12Z"
            },
            {
                "event_type": "exercise",
                "content": {
                    "duration": 10,
                    "type": "walking"
                },
                "created_at": "2024-11-20T10:27:12Z"
            },
            {
                "event_type": "meal",
                "content": {
                    "dish": [
                        {
                            "food": "egg",
                            "quantity": 2
                        },
                        {
                            "food": "sashimi",
                            "quantity": 2
                        }
                    ],
                    "type": "dinner"
                },
                "created_at": "2024-11-20T10:27:12Z"
            },
            {
                "event_type": "exercise",
                "content": {
                    "duration": 10,
                    "type": "walking"
                },
                "created_at": "2024-11-20T10:27:12Z"
            },
            {
                "event_type": "measure",
                "content": {
                    "type": "weight",
                    "value": 100
                },
                "created_at": "2024-11-20T10:27:12Z"
            },
            {
                "event_type": "measure",
                "content": {
                    "type": "height",
                    "value": 100
                },
                "created_at": "2024-11-20T10:27:12Z"
            }
        ],
        "success": true
    }
    ```


## Report graph
### Note
- Using the `measure` event_type with the `height` and `weight` types to proceed.

### Curl
- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/user/graph' \
    --header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRydW5nIiwiZXhwIjoxNzMyMTg0NzE2fQ.xnnHeUwlL_MbxDx1EDTmdRbYuhNrg8hJEeeaN2Fjsvw'
    ```
- Response:
    ```json
    {
        "data": {
            "height": [
                {
                    "time": "2024-11-20T10:41:14Z",
                    "value": 100
                },
                {
                    "time": "2024-11-20T10:27:12Z",
                    "value": 100
                },
                {
                    "time": "2024-11-20T10:27:12Z",
                    "value": 100
                }
            ],
            "weight": [
                {
                    "time": "2024-11-20T10:41:14Z",
                    "value": 100
                },
                {
                    "time": "2024-11-20T10:27:12Z",
                    "value": 100
                }
            ]
        },
        "success": true
    }
    ```


## Get achievement
### Note
- Using the `exercise` event_type to processed
### Curl
- Request:
    ```bash
    curl --location 'http://localhost:8080/api/v1/user/achievement' \
    --header 'Cookie: token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRydW5nIiwiZXhwIjoxNzMyMTg0NzE2fQ.xnnHeUwlL_MbxDx1EDTmdRbYuhNrg8hJEeeaN2Fjsvw'
    ```
- Response:
    ```json
    {
        "data": {
            "value": 20,
            "time": "2024-11-20T10:42:30Z"
        },
        "success": true
    }
    ```

