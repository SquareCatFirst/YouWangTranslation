---
title: 默认模块
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# 默认模块

Base URLs:

# Authentication

# backend/项目/新建项目

## POST 新建项目

POST /api/v1/project

> Body 请求参数

```json
{
  "action": "POST",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "cover_url": "https://example.com/cover.jpg",
    "name": "项目名称",
    "translated_name": "Project Translated Name",
    "author": "作者名",
    "tags": [
      4
    ],
    "source_site": "连载网站/来源站点",
    "description": "项目简介",
    "translation_description": "翻译简介"
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» cover_url|body|string| 是 |none|
|»» name|body|string| 是 |none|
|»» translated_name|body|string| 是 |none|
|»» author|body|string| 是 |none|
|»» tags|body|[integer]| 是 |none|
|»» source_site|body|string| 是 |none|
|»» description|body|string| 是 |none|
|»» translation_description|body|string| 是 |none|
|»» admin_max|body|integer| 是 |none|
|»» admin|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» source_max|body|integer| 是 |none|
|»» source|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» supervisor_max|body|integer| 是 |none|
|»» supervisor|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» translator_max|body|integer| 是 |none|
|»» translator|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» proofreader_max|body|integer| 是 |none|
|»» proofreader|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» typesetter_max|body|integer| 是 |none|
|»» typesetter|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» chapter_assignments|body|[object]| 是 |none|
|»»» role|body|integer| 是 |none|
|»»» chapter_ids|body|[string]| 是 |none|
|»»» user_ids|body|[string]| 是 |none|
|»» project_chapters|body|[object]| 是 |none|
|»»» chapter_name|body|string| 否 |none|
|»»» is_visible|body|boolean| 否 |none|
|»»» order_index|body|integer| 否 |none|
|»»» images|body|[object]| 否 |none|
|»»»» image_url|body|string| 是 |none|
|»»»» order_index|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "新建项目成功",
  "data": [
    {
      "project_id": 1,
      "name": "项目名称",
      "created_at": "20260124153045"
    }
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project",
  "method": "POST",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|string|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» project_id|integer|false|none||none|
|»» name|string|false|none||none|
|»» created_at|string|false|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## GET 获取项目

GET /api/v1/project

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|request|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "获取项目成功",
  "data": [
    {
      "cover_url": "https://example.com/cover.jpg",
      "name": "项目名称",
      "translated_name": "Project Translated Name",
      "author": "作者名",
      "tags": [
        "tag_id"
      ],
      "source_site": "连载网站/来源站点",
      "description": "项目简介",
      "translation_description": "翻译简介",
      "admin_max": 3,
      "admin": [
        {
          "user_id": 10001,
          "assigned_by": 10001
        },
        {
          "user_id": 10002,
          "assigned_by": 10001
        }
      ],
      "source_max": 5,
      "source": [
        {
          "user_id": 10001,
          "assigned_by": 10001
        },
        {
          "user_id": 10002,
          "assigned_by": 10001
        }
      ],
      "supervisor_max": 2,
      "supervisor": [
        {
          "user_id": 10001,
          "assigned_by": 10001
        },
        {
          "user_id": 10002,
          "assigned_by": 10001
        }
      ],
      "translator_max": 10,
      "translator": [
        {
          "user_id": 10001,
          "assigned_by": 10001
        },
        {
          "user_id": 10002,
          "assigned_by": 10001
        }
      ],
      "proofreader_max": 4,
      "proofreader": [
        {
          "user_id": 10001,
          "assigned_by": 10001
        },
        {
          "user_id": 10002,
          "assigned_by": 10001
        }
      ],
      "typesetter_max": 4,
      "typesetter": [
        {
          "user_id": 10001,
          "assigned_by": 10001
        },
        {
          "user_id": 10002,
          "assigned_by": 10001
        }
      ],
      "chapter_assignments": [
        {
          "role": 1,
          "chapter_ids": [
            1,
            2
          ],
          "user_ids": [
            10001,
            10002
          ],
          "assgined_by": 10001
        },
        {
          "role": 2,
          "chapter_ids": [
            1,
            2
          ],
          "user_ids": [
            10001,
            10002
          ],
          "assgined_by": 10001
        }
      ],
      "project_chapters": [
        {
          "chapter_id": 1,
          "chapter_name": "第1话",
          "is_visible": true,
          "order_index": 1,
          "images": [
            {
              "image_id": 1,
              "image_url": "https://example.com/ch1/p1.jpg",
              "order_index": 1
            },
            {
              "image_id": 1,
              "image_url": "https://example.com/ch1/p2.jpg",
              "order_index": 2
            }
          ]
        }
      ]
    }
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project",
  "method": "GET",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» cover_url|string|true|none||none|
|»» name|string|true|none||none|
|»» translated_name|string|true|none||none|
|»» author|string|true|none||none|
|»» tags|[string]|true|none||none|
|»» source_site|string|true|none||none|
|»» description|string|true|none||none|
|»» translation_description|string|true|none||none|
|»» admin_max|integer|true|none||none|
|»» admin|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» assigned_by|integer|true|none||none|
|»» source_max|integer|true|none||none|
|»» source|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» assigned_by|integer|true|none||none|
|»» supervisor_max|integer|true|none||none|
|»» supervisor|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» assigned_by|integer|true|none||none|
|»» translator_max|integer|true|none||none|
|»» translator|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» assigned_by|integer|true|none||none|
|»» proofreader_max|integer|true|none||none|
|»» proofreader|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» assigned_by|integer|true|none||none|
|»» typesetter_max|integer|true|none||none|
|»» typesetter|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» assigned_by|integer|true|none||none|
|»» chapter_assignments|[object]|true|none||none|
|»»» role|integer|true|none||none|
|»»» chapter_ids|[integer]|true|none||none|
|»»» user_ids|[integer]|true|none||none|
|»»» assgined_by|integer|true|none||none|
|»» project_chapters|[object]|true|none||none|
|»»» chapter_id|integer|false|none||none|
|»»» chapter_name|string|false|none||none|
|»»» is_visible|boolean|false|none||none|
|»»» order_index|integer|false|none||none|
|»»» images|[object]|false|none||none|
|»»»» image_id|integer|true|none||none|
|»»»» image_url|string|true|none||none|
|»»»» order_index|integer|true|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## DELETE 删除项目

DELETE /api/v1/project

> Body 请求参数

```json
{
  "action": "DELETE",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "project_id": 11
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» project_id|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "删除项目成功",
  "data": [
    {}
  ],
  "rowCount": 0,
  "totalCount": 0,
  "api": "/api/v1/project",
  "method": "DELETE",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## POST 修改项目

POST /api/v1/project/change

> Body 请求参数

```json
{
  "action": "POST",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "project_id": 9,
    "cover_url": "114514",
    "name": "项目名222211",
    "translated_name": "Pro Name",
    "author": "作者",
    "tags": [
      1,
      3,
      4
    ],
    "source_site": "连载网站/来源站点",
    "description": "项目简介",
    "translation_description": "翻译简介",
    "admin_max": 3,
    "source_max": 5,
    "supervisor_max": 2,
    "translator_max": 10,
    "proofreader_max": 4,
    "typesetter_max": 4
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» project_id|body|integer| 是 |none|
|»» cover_url|body|string| 是 |none|
|»» name|body|string| 是 |none|
|»» translated_name|body|string| 是 |none|
|»» author|body|string| 是 |none|
|»» tags|body|[integer]| 是 |none|
|»» source_site|body|string| 是 |none|
|»» description|body|string| 是 |none|
|»» translation_description|body|string| 是 |none|
|»» admin_max|body|integer| 是 |none|
|»» admin|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» source_max|body|integer| 是 |none|
|»» source|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» supervisor_max|body|integer| 是 |none|
|»» supervisor|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» translator_max|body|integer| 是 |none|
|»» translator|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» proofreader_max|body|integer| 是 |none|
|»» proofreader|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» typesetter_max|body|integer| 是 |none|
|»» typesetter|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» chapter_assignments|body|[object]| 是 |none|
|»»» role|body|integer| 是 |none|
|»»» chapter_ids|body|[integer]| 是 |none|
|»»» user_ids|body|[integer]| 是 |none|
|»»» assgined_by|body|integer| 是 |none|
|»» project_chapters|body|[object]| 是 |none|
|»»» chapter_id|body|integer| 否 |none|
|»»» chapter_name|body|string| 否 |none|
|»»» is_visible|body|boolean| 否 |none|
|»»» order_index|body|integer| 否 |none|
|»»» images|body|[object]| 否 |none|
|»»»» image_url|body|string| 是 |none|
|»»»» order_index|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "修改项目成功",
  "data": [
    {
      "project_id": 1,
      "name": "项目名称",
      "updated_at": "20260124153045"
    }
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project/change",
  "method": "POST",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» project_id|integer|false|none||none|
|»» name|string|false|none||none|
|»» updated_at|string|false|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## POST 添加或修改章节

POST /api/v1/project/chapter

> Body 请求参数

```json
{
  "action": "add",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "project_id": 1
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» project_id|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "操作项目成功",
  "data": [
    {
      "chapter_id": 1,
      "chapter_name": "项目名称",
      "is_visible": false,
      "order_index": 1,
      "project_chapters": [
        {
          "chapter_id": 1,
          "chapter_name": "第1话",
          "is_visible": true,
          "order_index": 1,
          "images": [
            {
              "image_url": "https://example.com/ch1/p1.jpg",
              "order_index": 1
            },
            {
              "image_url": "https://example.com/ch1/p2.jpg",
              "order_index": 2
            }
          ]
        }
      ],
      "created_at": "114514"
    }
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project/chapter",
  "method": "POST",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» chapter_id|integer|false|none||none|
|»» chapter_name|string|false|none||none|
|»» is_visible|boolean|false|none||none|
|»» order_index|integer|false|none||none|
|»» created_at|string|false|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## DELETE 删除章节

DELETE /api/v1/project/chapter

> Body 请求参数

```json
{
  "action": "DELETE",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "chapter_id": 4
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» chapter_id|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "删除章节成功",
  "data": [
    {}
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project/chapter",
  "method": "DELETE",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» project_id|integer|false|none||none|
|»» name|string|false|none||none|
|»» created_at|string|false|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## POST 添加或删除章节图片

POST /api/v1/project/chapter/image

> Body 请求参数

```json
{
  "action": "add",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "images": [
      {
        "chapter_id": 1,
        "role": 0,
        "image_url": "https://example.com/ch1/p1.jpg"
      },
      {
        "chapter_id": 1,
        "role": 0,
        "image_url": "https://example.com/ch1/p1.jpg"
      }
    ]
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» chapter_id|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "操作项目成功",
  "data": [
    {
      "images": [
        {
          "image_id": 1,
          "image_url": "https://example.com/ch1/p1.jpg",
          "order_index": 1
        },
        {
          "image_id": 2,
          "image_url": "https://example.com/ch1/p2.jpg",
          "order_index": 2
        }
      ]
    }
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project/chapter/image",
  "method": "POST",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» images|[object]|false|none||none|
|»»» image_id|integer|true|none||none|
|»»» image_url|string|true|none||none|
|»»» order_index|integer|true|none||none|
|»» project_id|integer|false|none||none|
|»» name|string|false|none||none|
|»» created_at|string|false|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## POST 添加或者修改章节指派

POST /api/v1/project/chapter_assignments

> Body 请求参数

```json
{
  "action": "add",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "project_id": 1
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» project_id|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "操作项目成功",
  "data": [
    {
      "project_assignment_id": 1,
      "created_at": "114514",
      "chapter_assignments": [
        {
          "role": 1,
          "chapter_ids": [
            1,
            2
          ],
          "user_ids": [
            10001,
            10002
          ],
          "assgined_by": 10001
        },
        {
          "role": 2,
          "chapter_ids": [
            1,
            2
          ],
          "user_ids": [
            10001,
            10002
          ],
          "assgined_by": 10001
        }
      ]
    }
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project/chapter_assignments",
  "method": "POST",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» project_assignment_id|integer|false|none||none|
|»» created_at|string|false|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## DELETE 删除章节指派

DELETE /api/v1/project/chapter_assignments

> Body 请求参数

```json
{
  "action": "DELETE",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "project_assignment_id": 1
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» chapter_id|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "删除章节指派成功",
  "data": [
    {}
  ],
  "rowCount": 1,
  "totalCount": 1,
  "api": "/api/v1/project/chapter_assignments",
  "method": "DLETE",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» project_id|integer|false|none||none|
|»» project_chapters|[object]|false|none||none|
|»»» chapter_id|integer|false|none||none|
|»»» chapter_name|string|false|none||none|
|»»» is_visible|boolean|false|none||none|
|»»» order_index|integer|false|none||none|
|»»» images|[object]|false|none||none|
|»»»» image_url|string|true|none||none|
|»»»» order_index|integer|true|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

## POST 项目权限

POST /api/v1/project/project_assignments

> Body 请求参数

```json
{
  "action": "POST",
  "sets": [],
  "orderBy": [],
  "page": 0,
  "pageSize": 0,
  "data": {
    "project_id": 1,
    "role": "admin",
    "admin": [
      {
        "user_id": 10001,
        "assigned_by": 10001
      },
      {
        "user_id": 10002,
        "assigned_by": 10001
      }
    ]
  },
  "filter": {},
  "authFilter": {}
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» action|body|string| 是 |none|
|» sets|body|[string]| 是 |none|
|» orderBy|body|[string]| 是 |none|
|» page|body|integer| 是 |none|
|» pageSize|body|integer| 是 |none|
|» data|body|object| 是 |none|
|»» admin|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» project_id|body|integer| 是 |none|
|»» source|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» supervisor|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» translator|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» proofreader|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|»» typesetter|body|[object]| 是 |none|
|»»» user_id|body|integer| 是 |none|
|»»» assigned_by|body|integer| 是 |none|
|» filter|body|object| 是 |none|
|» authFilter|body|object| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "msg": "添加权限成功",
  "data": [
    {}
  ],
  "rowCount": 0,
  "totalCount": 0,
  "api": "/api/v1/project/project_assignments",
  "method": "POST",
  "SN": "202601241530454839"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|» rowCount|integer|true|none||none|
|» totalCount|integer|true|none||none|
|» api|string|true|none||none|
|» method|string|true|none||none|
|» SN|string|true|none||none|

# 数据模型

