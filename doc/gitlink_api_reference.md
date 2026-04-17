# GitLink API 接口详细文档

> **数据来源**：[GitLink Apifox 文档](https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021) + Widdershins 导出文档
>
> **合并日期**：2026-03-31
>
> **接口总数**：210+ 个
>
> **Base URL**：`https://www.gitlink.org.cn`

---

# 认证方式

- **access_token**（推荐）：所有接口支持 `?access_token=xxx` 查询参数
- **Cookie**：`autologin_trustie` cookie
- **OAuth2**：Bearer Token（`Authorization: Bearer xxx`）
- 所有 API 路径以 `.json` 结尾

---

# Authentication

- HTTP Authentication, scheme: bearer

# 附件

## POST 上传文件

POST /api/attachments.json

> Body 请求参数

```yaml
file: ""
description: ""
container_id: ""
container_type: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 |none|
|body|body|object| 否 |none|
|» file|body|string(binary)| 是 |上传的文件|
|» description|body|string| 否 |上传文件描述|
|» container_id|body|string| 否 |上传文件归属模型ID|
|» container_type|body|string| 否 |上传文件归属模型类型|

> 返回示例

> 200 Response

```json
{
  "id": "f5838d8f-451b-4793-a0f2-0278430e8207",
  "title": "1tab.ico",
  "filesize": "9.4 KB",
  "is_pdf": false,
  "url": "http://localhost:3000/api/attachments/f5838d8f-451b-4793-a0f2-0278430e8207",
  "created_on": "2023-12-25 14:53",
  "content_type": "image/vnd.microsoft.icon"
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
|» id|string|true|none|文件ID|none|
|» title|string|true|none|文件标题|none|
|» filesize|string|true|none|文件大小|none|
|» is_pdf|boolean|true|none|文件是否为pdf|none|
|» url|string|true|none|文件绝对地址|none|
|» created_on|string|true|none|文件创建时间|none|
|» content_type|string|true|none|文件类型|none|

# 许可证

## GET 许可证列表

GET /api/licenses.json

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 |none|
|name|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "licenses": [
    {
      "id": 5,
      "name": "AFL-1.1"
    },
    {
      "id": 6,
      "name": "AFL-1.2"
    },
    {
      "id": 7,
      "name": "AFL-2.0"
    },
    {
      "id": 8,
      "name": "AFL-2.1"
    },
    {
      "id": 9,
      "name": "AFL-3.0"
    }
  ]
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
|» licenses|[object]|true|none||none|
|»» id|integer|true|none|ID|none|
|»» name|string|true|none|名称|none|

# 忽略文件

## GET 忽略文件列表

GET /api/ignores.json

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 |none|
|name|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "ignores": [
    {
      "id": 2,
      "name": "Ada"
    }
  ]
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
|» ignores|[object]|true|none||none|
|»» id|integer|false|none||none|
|»» name|string|false|none||none|

# 用户

## POST 创建一个用户

POST /api/accounts/register.json

> Body 请求参数

```json
{
  "code": 123123,
  "login": "13900010001",
  "namespace": "forgetest1",
  "password": "12345678",
  "password_confirmation": "12345678"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» code|body|integer| 是 | 验证码|none|
|» login|body|string| 是 | 手机号或者邮箱|none|
|» namespace|body|string| 是 | 用户标识|none|
|» password|body|string| 是 | 用户密码|none|
|» password_confirmation|body|string| 是 | 用户确认密码|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET Public Keys列表

GET /api/public_keys.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|page|query|integer| 否 ||页码|
|limit|query|string| 否 ||每页数量|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "public_keys": [
    {
      "id": 1,
      "name": "测试",
      "content": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC1VXqDLv/J7mmaUxzF0PuCYkKz64vwwCGNYCq3itnTa9SVIaUVAW6ur2xUuiNYM/wCvEzziyfSf+CpYM/Lsy7/Dp40IgrMUHWMOVZmmbF4btFnE3KLf6ENZmy9ZC/lCVXTFXYcAg6Vx4H5pqGt/LJT0B8YPIyhx5BhRrWDXOAYZCrCSumrRwlifzrA+EcuU8kDw23JykkSwyJxlTYH+IIp2t7zoLn1TXKXw3VYpRyiEgLb51v+8NSE6Uzd8NmiaxH02Upq6K9zpALe6NI1GtbE15nQYD51oOVFLagemQuB+ARIaePvA5wcZjI3Jsw/Of84RUEZIWsEE9ue0ejU+ioJTQfc8apnezeed8iqi/qeVSPpxejFTCVp+JlJ8xZeDPnQcEFPPfsu1clR7dKevGMbnUb3ooex03rsqZRcZHAW65BmfO6cFXBkmnNVrjfHf5g86tsqewUUmIPfAcylHoenHDwKpDpd5GV3XsK/sJkWCoQKwP8D7/yZIWyguRK7AnO1eWxwCTRuF/drRMCE5U+ztUSyzUhIhDpjPBbBbFg3iQrWCaFIgPueC3nj1h72YuofhYb+TSvfo8Z2KGkOfXaDO8ALr5lEhCvVWk+Im/1E7BIQByqH1vRGP89d1pJIQmKAmba7O6mv7hXvjeHrXVCve3+Q7dQwZaK317bvlPHcew== yystopf@163.com",
      "fingerprint": "SHA256:fuQoDBhAfOijPymXEa2ahiVV/5Jj3hB/0YbIS4aBMDo",
      "created_unix": 1690797157,
      "created_time": "2023/07/31 17:52"
    }
  ]
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
|» total_count|integer|true|none|总数|none|
|» public_keys|[object]|true|none||none|
|»» id|integer|true|none|ID|none|
|»» name|string|true|none|密钥名称|none|
|»» content|string|true|none|密钥内容|none|
|»» fingerprint|string|true|none|密钥简洁版内容|none|
|»» created_unix|integer|true|none|密钥创建时间|none|
|»» created_time|string|true|none|密钥创建时间戳|none|

## POST Public Keys创建

POST /api/public_keys.json

> Body 请求参数

```json
{
  "key": "string",
  "title": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» key|body|string| 是 | 密钥内容|none|
|» title|body|string| 是 | 密钥标题|none|

> 返回示例

> 200 Response

```json
{
  "id": 2,
  "name": "速个严接不建和",
  "content": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC1VXqDLv/J7mmaUxzF0PuCYkKz64vwwCGNYCq3itnTa9SVIaUVAW6ur2xUuiNYM/wCvEzziyfSf+CpYM/Lsy7/Dp40IgrMUHWMOVZmmbF4btFnE3KLf6ENZmy9ZC/lCVXTFXYcAg6Vx4H5pqGt/LJT0B8YPIyhx5BhRrWDXOAYZCrCSumrRwlifzrA+EcuU8kDw23JykkSwyJxlTYH+IIp2t7zoLn1TXKXw3VYpRyiEgLb51v+8NSE6Uzd8NmiaxH02Upq6K9zpALe6NI1GtbE15nQYD51oOVFLagemQuB+ARIaePvA5wcZjI3Jsw/Of84RUEZIWsEE9ue0ejU+ioJTQfc8apnezeed8iqi/qeVSPpxejFTCVp+JlJ8xZeDPnQcEFPPfsu1clR7dKevGMbnUb3ooex03rsqZRcZHAW65BmfO6cFXBkmnNVrjfHf5g86tsqewUUmIPfAcylHoenHDwKpDpd5GV3XsK/sJkWCoQKwP8D7/yZIWyguRK7AnO1eWxwCTRuF/drRMCE5U+ztUSyzUhIhDpjPBbBbFg3iQrWCaFIgPueC3nj1h72YuofhYb+TSvfo8Z2KGkOfXaDO8ALr5lEhCvVWk+Im/1E7BIQByqH1vRGP89d1pJIQmKAmba7O6mv7hXvjeHrXVCve3+Q7dQwZaK317bvlPHcew== yystopf@163.com",
  "fingerprint": "SHA256:fuQoDBhAfOijPymXEa2ahiVV/5Jj3hB/0YbIS4aBMDo",
  "created_time": "2023/07/31 18:07"
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
|» id|integer|true|none|密钥ID|none|
|» name|string|true|none|密钥名称|none|
|» content|string|true|none|密钥内容|none|
|» fingerprint|string|true|none|密钥简洁版内容|none|
|» created_time|string|true|none|密钥创建时间|none|

## DELETE Public Keys删除

DELETE /api/public_keys/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|id|path|integer| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 当前用户信息

GET /api/users/get_user_info.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "username": "咸蛋黄土豆丝",
  "real_name": "咸蛋黄土豆丝",
  "nickname": "咸蛋黄土豆丝",
  "gender": 0,
  "login": "yystopf",
  "user_id": 110,
  "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png",
  "admin": true,
  "is_teacher": false,
  "user_identity": "专业人士",
  "tidding_count": 0,
  "user_phone_binded": false,
  "need_edit_info": false,
  "phone": null,
  "profile_completed": true,
  "professional_certification": false,
  "devops_step": 0,
  "ci_certification": false,
  "email": "yystopf@163.com",
  "province": "河北省",
  "city": "秦皇岛",
  "custom_department": "嘻嘻嘻嘻",
  "description": "13",
  "super_description": null,
  "show_email": false,
  "show_department": false,
  "show_location": false,
  "show_super_description": false,
  "message_unread_total": 0,
  "has_trace_user": false,
  "is_new": false,
  "nps": false,
  "sign_cla": true
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
|» username|string|true|none|用户名|none|
|» real_name|string|true|none|用户真实名称|none|
|» nickname|string|true|none|用户昵称|none|
|» gender|integer|true|none|用户性别|none|
|» login|string|true|none|用户标识|none|
|» user_id|integer|true|none|用户ID|none|
|» image_url|string|true|none|用户头像|none|
|» admin|boolean|true|none|用户是否为管理员|none|
|» is_teacher|boolean|true|none|用户是否为老师|none|
|» user_phone_binded|boolean|true|none|用户是否绑定手机|none|
|» phone|null|true|none|用户手机号|none|
|» profile_completed|boolean|true|none|用户是否完善资料|none|
|» email|string|true|none|用户邮箱|none|
|» province|string|true|none|用户所在省份|none|
|» city|string|true|none|用户所在城市|none|
|» custom_department|string|true|none|用户单位|none|
|» description|string|true|none|用户描述|none|
|» show_email|boolean|true|none|是否展示邮箱|none|
|» show_department|boolean|true|none|是否展示公司|none|
|» show_location|boolean|true|none|是否展示位置|none|
|» message_unread_total|integer|true|none|消息未读数量|none|
|» has_trace_user|boolean|true|none|是否有代码溯源用户|none|
|» is_new|boolean|true|none||none|
|» nps|boolean|true|none|是否开启nps|none|
|» sign_cla|boolean|true|none|是否签署cla协议|none|

## PUT 用户信息更改

PUT /api/users/{owner}.json

> Body 请求参数

```json
{
  "user": {
    "nickname": "string",
    "user_extension_attributes": {
      "gender": 0,
      "province": "string",
      "city": "string",
      "description": "string",
      "custom_department": "string"
    }
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» user|body|object| 是 ||none|
|»» nickname|body|string| 是 | 用户昵称|none|
|»» user_extension_attributes|body|object| 是 ||none|
|»»» gender|body|integer| 是 | 用户性别|none|
|»»» province|body|string| 是 | 用户所在省份|none|
|»»» city|body|string| 是 | 用户所在城市|none|
|»»» description|body|string| 是 | 用户简介|none|
|»»» custom_department|body|string| 是 | 用户所在单位|none|

> 返回示例

> 200 Response

```json
{
  "user_id": 110,
  "name": "咸蛋黄土豆丝xxx",
  "username": "咸蛋黄土豆丝xxx",
  "real_name": "咸蛋黄土豆丝xxx",
  "grade": 0,
  "gender": 0,
  "login": "yystopf",
  "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png",
  "admin": true,
  "user_identity": "专业人士",
  "is_watch": false,
  "watched_count": 0,
  "watching_count": 0,
  "created_time": "2022-09-19 11:45",
  "email": null,
  "province": null,
  "city": null,
  "custom_department": null,
  "description": "个性签名"
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
|» user_id|integer|true|none||none|
|» name|string|true|none|用户昵称|none|
|» username|string|true|none|用户昵称|none|
|» real_name|string|true|none|用户昵称|none|
|» gender|integer|true|none|用户性别|none|
|» login|string|true|none|用户标识|none|
|» image_url|string|true|none|用户头像|none|
|» admin|boolean|true|none|是否为管理员|none|
|» is_watch|boolean|true|none|是否关注|none|
|» watched_count|integer|true|none|用户被关注数量|none|
|» watching_count|integer|true|none|用户关注的用户数|none|
|» created_time|string|true|none|用户创建时间|none|
|» email|null|true|none|用户邮箱|none|
|» province|null|true|none|用户所在省份|none|
|» city|null|true|none|用户所在城市|none|
|» custom_department|null|true|none|用户所在单位|none|
|» description|string|true|none|用户个性签名|none|

## GET 用户项目列表

GET /api/v1/{owner}/projects.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|page|query|integer| 否 ||页码|
|limit|query|integer| 否 ||每页个数|
|category|query|string| 否 ||项目分类 all 所有 join 我参与的 created 我创建的 manage 我管理的 only_watched 我关注的（除了我创建的以外） watched 我关注的 forked 我复刻的 |
|is_public|query|string| 否 ||是否为公有项目|
|project_type|query|string| 否 ||项目类型 common 普通项目 mirror 镜像项目 sync_mirror 同步镜像项目|
|sort_by|query|string| 否 ||排序字段|
|sort_direction|query|string| 否 ||排序类型|
|search|query|string| 否 ||搜索关键词|
|start_at|query|integer| 否 ||开始时间戳，当category为watched、only_watched、forked才有效|
|end_at|query|integer| 否 ||结束时间戳，当category为watched、only_watched、forked才有效|

> 返回示例

> 200 Response

```json
{
  "total_count": 25,
  "projects": [
    {
      "owner": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "type": "common",
      "description": null,
      "forked_count": 0,
      "forked_from_project_id": null,
      "identifier": "1213.2131",
      "issues_count": 1,
      "pull_requests_count": 0,
      "invite_code": "MPzQgH",
      "website": null,
      "platform": "forge",
      "name": "1213.2131",
      "open_devops": false,
      "praises_count": 0,
      "is_public": true,
      "status": 1,
      "watchers_count": 0,
      "ignore_id": null,
      "license_id": null,
      "project_category_id": null,
      "project_language_id": null
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» projects|[object]|true|none||none|
|»» owner|object|true|none|项目拥有着|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» type|string|true|none|项目类型|common 托管项目，mirror 镜像项目, sync_mirror 同步镜像项目|
|»» description|null|false|none|项目描述|none|
|»» forked_count|integer|true|none|项目复刻数量|none|
|»» forked_from_project_id|null|false|none|项目复刻来源项目|none|
|»» identifier|string|true|none|项目标识|none|
|»» issues_count|integer|true|none|项目疑修数量|none|
|»» pull_requests_count|integer|true|none|项目合并请求数量|none|
|»» invite_code|string|true|none|项目邀请码|none|
|»» website|null|false|none|项目地址|none|
|»» platform|string|true|none|项目平台|none|
|»» name|string|true|none|项目名称|none|
|»» open_devops|boolean|true|none|项目是否开启流水线|none|
|»» praises_count|integer|true|none|项目点赞数量|none|
|»» is_public|boolean|true|none|项目是否为公开项目|none|
|»» status|integer|true|none|项目状态|none|
|»» watchers_count|integer|true|none|项目关注数量|none|
|»» ignore_id|null|false|none|项目忽略文件ID|none|
|»» license_id|null|false|none|项目协议ID|none|
|»» project_category_id|null|false|none|项目分类ID|none|
|»» project_language_id|null|false|none|项目语言ID|none|

## GET 用户消息列表

GET /api/users/{owner}/messages.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|page|query|integer| 否 ||页码|
|limit|query|integer| 否 ||每页个数|
|type|query|string| 否 ||消息类型，不传为所有消息，notification为系统消息，atme为@我消息|
|status|query|integer| 否 ||是否已读，不传为所有消息，1为未读，2为已读|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "type": "atme",
  "unread_notification": 0,
  "unread_atme": 1,
  "messages": [
    {
      "id": 20156,
      "status": 1,
      "content": "<b>heihei</b> 在疑修 <b>1321312</b> 中@我",
      "notification_url": "http://localhost:3000/yystopf/testdevops/issues/1",
      "source": "IssueAtme",
      "created_at": "2023-08-01 12:45:54",
      "time_ago": "5分钟前",
      "type": "atme",
      "sender": {
        "id": 113,
        "type": "User",
        "name": "heihei",
        "login": "yystopfceshi",
        "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
      }
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» type|string|true|none|消息类型|，notification为系统消息，atme为@我消息|
|» unread_notification|integer|true|none|未读消息数量|none|
|» unread_atme|integer|true|none|未读@我数量|none|
|» messages|[object]|true|none||none|
|»» id|integer|false|none|消息ID|none|
|»» status|integer|false|none|消息状态|消息是否已读，1为未读，2为已读|
|»» content|string|false|none|消息内容|none|
|»» notification_url|string|false|none|消息跳转地址|none|
|»» source|string|false|none|消息来源数据|none|
|»» created_at|string|false|none|消息创建时间|none|
|»» time_ago|string|false|none|消息创建距离现在时间|none|
|»» type|string|false|none|消息类型|，notification为系统消息，atme为@我消息|
|»» sender|object|false|none||none|
|»»» id|integer|false|none|消息发送者ID|none|
|»»» type|string|false|none|消息发送者类型|User为用户，Organization为组织|
|»»» name|string|false|none|消息发送者昵称|none|
|»»» login|string|false|none|消息发送者标识|none|
|»»» image_url|string|false|none|消息发送者头像|none|

#### 枚举值

|属性|值|
|---|---|
|source|IssueAssigned|
|source|IssueExpire|
|source|IssueAtme|
|source|IssueChanged|
|source|IssueDeleted|
|source|IssueJournal|
|source|LoginIpTip|
|source|OrganizationJoined|
|source|OrganizationLeft	|
|source|OrganizationRole|
|source|ProjectDeleted|
|source|ProjectFollowed|
|source|ProjectForked|
|source|ProjectIssue|
|source|ProjectJoined|
|source|ProjectLeft|
|source|ProjectMemberJoined|
|source|ProjectMemberLeft|
|source|ProjectMilestoneCompleted|
|source|ProjectMilestone|
|source|ProjectPraised	|
|source|ProjectPullRequest|
|source|ProjectRole|
|source|ProjectSettingChanged|
|source|ProjectTransfer|
|source|ProjectVersion|
|source|PullRequestAssigned|
|source|PullReuqestAtme|
|source|PullRequestChanged|
|source|PullRequestClosed|
|source|PullRequestJournal|
|source|PullRequestMerged|

## DELETE 用户删除消息

DELETE /api/users/{owner}/messages.json

> Body 请求参数

```json
{
  "type": "notification",
  "ids": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» type|body|string| 是 | 消息类型|不传为所有消息，notification为系统消息，atme为@我消息|
|» ids|body|[integer]| 是 | 消息id数组|包含-1则把所有未读消息标记为已读|

#### 枚举值

|属性|值|
|---|---|
|» type|notification|
|» type|atme|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 用户消息创建

POST /api/users/{owner}/messages.json

> Body 请求参数

```json
{
  "type": "atme",
  "receivers_login": [
    "string"
  ],
  "atmeable_type": "Journal",
  "atmeable_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» type|body|string| 是 | 类型|none|
|» receivers_login|body|[string]| 是 | 需要发送消息的用户名数组|none|
|» atmeable_type|body|string| 是 | atme消息对象|是从哪里@我的，比如评论：Journal、疑修：Issue、合并请求：PullRequest|
|» atmeable_id|body|integer| 是 | atme消息对象id|none|

#### 枚举值

|属性|值|
|---|---|
|» type|atme|
|» atmeable_type|Journal|
|» atmeable_type|Issue|
|» atmeable_type|PullRequest|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 用户阅读消息

POST /api/users/{owner}/messages/read.json

> Body 请求参数

```json
{
  "type": "notification",
  "ids": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» type|body|string| 是 | 消息类型|不传为所有消息，notification为系统消息，atme为@我消息|
|» ids|body|[integer]| 是 | 消息id数组|包含-1则把所有未读消息标记为已读|

#### 枚举值

|属性|值|
|---|---|
|» type|notification|
|» type|atme|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 平台消息设置

GET /api/template_message_settings.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "响应成功",
  "setting_types": [
    {
      "type": "TemplateMessageSetting::Normal",
      "type_name": "我的状态",
      "total_settings_count": 5,
      "settings": [
        {
          "name": "账号有权限变更",
          "key": "Permission",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "被拉入或移出组织",
          "key": "Organization",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "被拉入或移出项目",
          "key": "Project",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "有新的疑修指派给我",
          "key": "IssueAssigned",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "有新的合并请求指派给我",
          "key": "PullRequestAssigned",
          "notification_disabled": true,
          "email_disabled": false
        }
      ]
    },
    {
      "type": "TemplateMessageSetting::CreateOrAssign",
      "type_name": "我创建的或负责的",
      "total_settings_count": 3,
      "settings": [
        {
          "name": "疑修状态变更",
          "key": "IssueChanged",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "合并请求状态变更",
          "key": "PullRequestChanged",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "疑修截止日期到达最后一天",
          "key": "IssueExpire",
          "notification_disabled": false,
          "email_disabled": false
        }
      ]
    },
    {
      "type": "TemplateMessageSetting::ManageProject",
      "type_name": "我管理的仓库",
      "total_settings_count": 8,
      "settings": [
        {
          "name": "有新的疑修",
          "key": "Issue",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "有新的合并请求",
          "key": "PullRequest",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "有成员变动",
          "key": "Member",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "仓库设置被更改",
          "key": "SettingChanged",
          "notification_disabled": true,
          "email_disabled": false
        },
        {
          "name": "被点赞",
          "key": "Praised",
          "notification_disabled": false,
          "email_disabled": true
        },
        {
          "name": "被fork",
          "key": "Forked",
          "notification_disabled": false,
          "email_disabled": true
        },
        {
          "name": "有新的里程碑",
          "key": "Milestone",
          "notification_disabled": false,
          "email_disabled": false
        },
        {
          "name": "有里程碑已完成",
          "key": "MilestoneCompleted",
          "notification_disabled": false,
          "email_disabled": false
        }
      ]
    }
  ]
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
|» message|string|true|none||none|
|» setting_types|[object]|true|none||none|
|»» type|string|true|none|设置分类类型|none|
|»» type_name|string|true|none|设置分类名称|none|
|»» total_settings_count|integer|true|none|设置分类下设置总数|none|
|»» settings|[object]|true|none||none|
|»»» name|string|true|none|设置名称|none|
|»»» key|string|true|none|设置的Key|none|
|»»» notification_disabled|boolean|true|none|系统消息是否禁用|none|
|»»» email_disabled|boolean|true|none|@我消息是否禁用|none|

## GET 用户消息设置列表

GET /api/users/{owner}/template_message_settings.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "响应成功",
  "user": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝xxx",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "notification_body": {
    "Normal::Permission": true,
    "Normal::Project": true,
    "Normal::Organization": true,
    "Normal::IssueAssigned": true,
    "Normal::PullRequestAssigned": true,
    "CreateOrAssign::IssueChanged": true,
    "CreateOrAssign::PullRequestChanged": true,
    "CreateOrAssign::IssueExpire": true,
    "ManageProject::Issue": true,
    "ManageProject::PullRequest": true,
    "ManageProject::Member": true,
    "ManageProject::SettingChanged": true,
    "ManageProject::Praised": true,
    "ManageProject::Forked": true,
    "ManageProject::Milestone": true,
    "ManageProject::MilestoneCompleted": true
  },
  "email_body": {
    "Normal::Permission": false,
    "Normal::Project": false,
    "Normal::Organization": false,
    "Normal::IssueAssigned": false,
    "Normal::PullRequestAssigned": false,
    "CreateOrAssign::IssueChanged": false,
    "CreateOrAssign::PullRequestChanged": false,
    "CreateOrAssign::IssueExpire": false,
    "ManageProject::Issue": false,
    "ManageProject::PullRequest": false,
    "ManageProject::Member": false,
    "ManageProject::SettingChanged": false,
    "ManageProject::Praised": false,
    "ManageProject::Forked": false,
    "ManageProject::Milestone": false,
    "ManageProject::MilestoneCompleted": false
  }
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
|» message|string|true|none||none|
|» user|object|true|none||none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» notification_body|object|true|none||站内信消息配置|
|»» Normal::Permission|boolean|true|none||none|
|»» Normal::Project|boolean|true|none||none|
|»» Normal::Organization|boolean|true|none||none|
|»» Normal::IssueAssigned|boolean|true|none||none|
|»» Normal::PullRequestAssigned|boolean|true|none||none|
|»» CreateOrAssign::IssueChanged|boolean|true|none||none|
|»» CreateOrAssign::PullRequestChanged|boolean|true|none||none|
|»» CreateOrAssign::IssueExpire|boolean|true|none||none|
|»» ManageProject::Issue|boolean|true|none||none|
|»» ManageProject::PullRequest|boolean|true|none||none|
|»» ManageProject::Member|boolean|true|none||none|
|»» ManageProject::SettingChanged|boolean|true|none||none|
|»» ManageProject::Praised|boolean|true|none||none|
|»» ManageProject::Forked|boolean|true|none||none|
|»» ManageProject::Milestone|boolean|true|none||none|
|»» ManageProject::MilestoneCompleted|boolean|true|none||none|
|» email_body|object|true|none||邮件消息配置|
|»» Normal::Permission|boolean|true|none||none|
|»» Normal::Project|boolean|true|none||none|
|»» Normal::Organization|boolean|true|none||none|
|»» Normal::IssueAssigned|boolean|true|none||none|
|»» Normal::PullRequestAssigned|boolean|true|none||none|
|»» CreateOrAssign::IssueChanged|boolean|true|none||none|
|»» CreateOrAssign::PullRequestChanged|boolean|true|none||none|
|»» CreateOrAssign::IssueExpire|boolean|true|none||none|
|»» ManageProject::Issue|boolean|true|none||none|
|»» ManageProject::PullRequest|boolean|true|none||none|
|»» ManageProject::Member|boolean|true|none||none|
|»» ManageProject::SettingChanged|boolean|true|none||none|
|»» ManageProject::Praised|boolean|true|none||none|
|»» ManageProject::Forked|boolean|true|none||none|
|»» ManageProject::Milestone|boolean|true|none||none|
|»» ManageProject::MilestoneCompleted|boolean|true|none||none|

## POST 用户消息设置更新

POST /api/users/{owner}/template_message_settings/update_setting.json

> Body 请求参数

```json
{
  "setting": {
    "notification_body": {
      "CreateOrAssign::IssueAssigned": true,
      "CreateOrAssign::PullRequestAssigned": true,
      "ManageProject::Issue": true,
      "ManageProject::PullRequest": true,
      "ManageProject::Member": true,
      "ManageProject::SettingChanged": true,
      "Normal::Organization": true,
      "Normal::Project": true,
      "Normal::Permission": true
    },
    "email_body": {
      "CreateOrAssign::IssueAssigned": true,
      "CreateOrAssign::PullRequestAssigned": true,
      "ManageProject::Issue": true,
      "ManageProject::PullRequest": true,
      "ManageProject::Member": true,
      "ManageProject::SettingChanged": true,
      "Normal::Organization": true,
      "Normal::Project": true,
      "Normal::Permission": true
    }
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» setting|body|object| 是 ||none|
|»» notification_body|body|object| 是 ||none|
|»»» CreateOrAssign::IssueAssigned|body|boolean| 是 ||none|
|»»» CreateOrAssign::PullRequestAssigned|body|boolean| 是 ||none|
|»»» ManageProject::Issue|body|boolean| 是 ||none|
|»»» ManageProject::PullRequest|body|boolean| 是 ||none|
|»»» ManageProject::Member|body|boolean| 是 ||none|
|»»» ManageProject::SettingChanged|body|boolean| 是 ||none|
|»»» Normal::Organization|body|boolean| 是 ||none|
|»»» Normal::Project|body|boolean| 是 ||none|
|»»» Normal::Permission|body|boolean| 是 ||none|
|»» email_body|body|object| 是 ||none|
|»»» CreateOrAssign::IssueAssigned|body|boolean| 是 ||none|
|»»» CreateOrAssign::PullRequestAssigned|body|boolean| 是 ||none|
|»»» ManageProject::Issue|body|boolean| 是 ||none|
|»»» ManageProject::PullRequest|body|boolean| 是 ||none|
|»»» ManageProject::Member|body|boolean| 是 ||none|
|»»» ManageProject::SettingChanged|body|boolean| 是 ||none|
|»»» Normal::Organization|body|boolean| 是 ||none|
|»»» Normal::Project|body|boolean| 是 ||none|
|»»» Normal::Permission|body|boolean| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "响应成功",
  "user": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝xxx",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "notification_body": {
    "Normal::Permission": true,
    "Normal::Project": true,
    "Normal::Organization": true,
    "Normal::IssueAssigned": true,
    "Normal::PullRequestAssigned": true,
    "CreateOrAssign::IssueChanged": true,
    "CreateOrAssign::PullRequestChanged": true,
    "CreateOrAssign::IssueExpire": true,
    "ManageProject::Issue": true,
    "ManageProject::PullRequest": true,
    "ManageProject::Member": true,
    "ManageProject::SettingChanged": true,
    "ManageProject::Praised": true,
    "ManageProject::Forked": true,
    "ManageProject::Milestone": true,
    "ManageProject::MilestoneCompleted": true
  },
  "email_body": {
    "Normal::Permission": false,
    "Normal::Project": false,
    "Normal::Organization": false,
    "Normal::IssueAssigned": false,
    "Normal::PullRequestAssigned": false,
    "CreateOrAssign::IssueChanged": false,
    "CreateOrAssign::PullRequestChanged": false,
    "CreateOrAssign::IssueExpire": false,
    "ManageProject::Issue": false,
    "ManageProject::PullRequest": false,
    "ManageProject::Member": false,
    "ManageProject::SettingChanged": false,
    "ManageProject::Praised": false,
    "ManageProject::Forked": false,
    "ManageProject::Milestone": false,
    "ManageProject::MilestoneCompleted": false
  }
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
|» message|string|true|none||none|
|» user|object|true|none||none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» notification_body|object|true|none||站内信消息配置|
|»» Normal::Permission|boolean|true|none||none|
|»» Normal::Project|boolean|true|none||none|
|»» Normal::Organization|boolean|true|none||none|
|»» Normal::IssueAssigned|boolean|true|none||none|
|»» Normal::PullRequestAssigned|boolean|true|none||none|
|»» CreateOrAssign::IssueChanged|boolean|true|none||none|
|»» CreateOrAssign::PullRequestChanged|boolean|true|none||none|
|»» CreateOrAssign::IssueExpire|boolean|true|none||none|
|»» ManageProject::Issue|boolean|true|none||none|
|»» ManageProject::PullRequest|boolean|true|none||none|
|»» ManageProject::Member|boolean|true|none||none|
|»» ManageProject::SettingChanged|boolean|true|none||none|
|»» ManageProject::Praised|boolean|true|none||none|
|»» ManageProject::Forked|boolean|true|none||none|
|»» ManageProject::Milestone|boolean|true|none||none|
|»» ManageProject::MilestoneCompleted|boolean|true|none||none|
|» email_body|object|true|none||邮件消息配置|
|»» Normal::Permission|boolean|true|none||none|
|»» Normal::Project|boolean|true|none||none|
|»» Normal::Organization|boolean|true|none||none|
|»» Normal::IssueAssigned|boolean|true|none||none|
|»» Normal::PullRequestAssigned|boolean|true|none||none|
|»» CreateOrAssign::IssueChanged|boolean|true|none||none|
|»» CreateOrAssign::PullRequestChanged|boolean|true|none||none|
|»» CreateOrAssign::IssueExpire|boolean|true|none||none|
|»» ManageProject::Issue|boolean|true|none||none|
|»» ManageProject::PullRequest|boolean|true|none||none|
|»» ManageProject::Member|boolean|true|none||none|
|»» ManageProject::SettingChanged|boolean|true|none||none|
|»» ManageProject::Praised|boolean|true|none||none|
|»» ManageProject::Forked|boolean|true|none||none|
|»» ManageProject::Milestone|boolean|true|none||none|
|»» ManageProject::MilestoneCompleted|boolean|true|none||none|

## GET 用户星标项目列表

GET /api/users/{owner}/is_pinned_projects.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "projects": [
    {
      "id": 1,
      "repo_id": 17,
      "identifier": "testdevops",
      "name": "testdevops",
      "description": "",
      "visits": 8,
      "praises_count": 0,
      "watchers_count": 0,
      "issues_count": 1,
      "pull_requests_count": 0,
      "forked_count": 0,
      "is_public": true,
      "mirror_url": null,
      "type": 0,
      "last_update_time": 1688266261,
      "time_ago": "1个月前",
      "forked_from_project_id": null,
      "open_devops": false,
      "platform": "forge",
      "is_pinned": true,
      "author": {
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "category": null,
      "language": {
        "id": 8,
        "name": "JavaScript"
      },
      "topics": [
        {
          "id": 8,
          "name": "javascript"
        },
        {
          "id": 7,
          "name": "python"
        }
      ],
      "position": 0,
      "project_id": 17
    }
  ]
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
|» total_count|integer|true|none||none|
|» projects|[object]|true|none||none|
|»» id|integer|false|none|星标项目ID|none|
|»» identifier|string|false|none|项目标识|none|
|»» name|string|false|none|项目名称|none|
|»» description|string|false|none|项目描述|none|
|»» visits|integer|false|none|项目访问数|none|
|»» praises_count|integer|false|none|项目点赞数|none|
|»» watchers_count|integer|false|none|项目关注数|none|
|»» issues_count|integer|false|none|项目疑修数|none|
|»» pull_requests_count|integer|false|none|项目合并请求数|none|
|»» forked_count|integer|false|none|项目复刻数量|none|
|»» is_public|boolean|false|none|项目是否公开|none|
|»» mirror_url|null|false|none|项目镜像地址|none|
|»» type|integer|false|none|项目类型|0普通项目 1 普通镜像项目 2 同步镜像项目|
|»» last_update_time|integer|false|none|项目上次更新时间戳|none|
|»» time_ago|string|false|none|项目上次更新时间|none|
|»» forked_from_project_id|null|false|none|项目fork项目ID|none|
|»» open_devops|boolean|false|none|是否开启流水线|none|
|»» platform|string|false|none|项目平台|none|
|»» is_pinned|boolean|false|none|是否为精选项目|none|
|»» author|object|false|none|项目拥有着|none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» category|null|false|none|项目分类|none|
|»» language|object¦null|false|none|项目语言|none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»» topics|[object]¦null|false|none|项目搜索标签|none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»» position|integer|false|none|项目排序|none|
|»» project_id|integer|true|none|项目ID|none|

## POST 用户星标项目更新

POST /api/users/{owner}/is_pinned_projects/pin.json

> Body 请求参数

```json
{
  "is_pinned_project_ids": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» is_pinned_project_ids|body|[integer]| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## PUT 用户星标项目排序

PUT /api/users/{owner}/is_pinned_projects/{id}.json

> Body 请求参数

```json
{
  "pinned_project": {
    "position": 0
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|id|path|string| 是 ||星标项目ID|
|body|body|object| 否 ||none|
|» pinned_project|body|object| 是 ||none|
|»» position|body|integer| 是 | 排序大小，数字越大，排名越靠前|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 新建一个反馈建议

POST /api/v1/{owner}/feedbacks.json

> Body 请求参数

```json
{
  "content": "测试反馈意见"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户名|
|body|body|object| 否 ||none|
|» content|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "string"
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
|» message|string|true|none||none|

## GET 用户近期活动统计

GET /api/users/{owner}/statistics/activity.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|

> 返回示例

> 200 Response

```json
{
  "dates": [
    "2023.07.27",
    "2023.07.28",
    "2023.07.29",
    "2023.07.30",
    "2023.07.31",
    "2023.08.01",
    "2023.08.02",
    "2023.08.03"
  ],
  "issues_count": [
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0
  ],
  "pull_requests_count": [
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0
  ],
  "commits_count": [
    0,
    0,
    0,
    0,
    0,
    0,
    0,
    0
  ]
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
|» dates|[string]|true|none|日期|none|
|» issues_count|[integer]|true|none|疑修数量|none|
|» pull_requests_count|[integer]|true|none|合并请求数量|none|
|» commits_count|[integer]|true|none|代码提交数量|none|

## GET 用户贡献度统计

GET /api/users/{owner}/headmaps.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|year|query|string| 否 ||年份|

> 返回示例

> 200 Response

```json
{
  "total_contributions": 12,
  "headmaps": [
    {
      "date": "2023-01-05",
      "contributions": 2
    },
    {
      "date": "2023-01-06",
      "contributions": 1
    },
    {
      "date": "2023-01-11",
      "contributions": 6
    },
    {
      "date": "2023-04-17",
      "contributions": 1
    },
    {
      "date": "2023-07-02",
      "contributions": 2
    }
  ]
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
|» total_contributions|integer|true|none|日期范围内提交总数|none|
|» headmaps|[object]|true|none||none|
|»» date|string|true|none|日期|none|
|»» contributions|integer|true|none|贡献度|none|

## GET 用户开发能力

GET /api/users/{owner}/statistics/develop.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|start_time|query|integer| 否 ||开始时间戳|
|end_time|query|integer| 否 ||结束时间戳|

> 返回示例

> 200 Response

```json
{
  "platform": {
    "influence": 60,
    "contribution": 78,
    "activity": 94,
    "experience": 88,
    "language": 81
  },
  "user": {
    "influence": 60,
    "contribution": 78,
    "activity": 94,
    "experience": 87,
    "language": 80,
    "languages_percent": {
      "JSX": 0.06,
      "Python": 0.31,
      "JavaScript": 0.5,
      "Markdown": 0.06,
      "Yaml": 0.06
    },
    "each_language_score": {
      "JSX": 66,
      "Python": 80,
      "JavaScript": 84,
      "Markdown": 66,
      "Yaml": 66
    }
  }
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
|» platform|object|true|none||none|
|»» influence|integer|true|none|平台影响力|none|
|»» contribution|integer|true|none|平台贡献度|none|
|»» activity|integer|true|none|平台活跃度|none|
|»» experience|integer|true|none|平台项目经验|none|
|»» language|integer|true|none|平台语言能力|none|
|» user|object|true|none||none|
|»» influence|integer|true|none|用户影响力|none|
|»» contribution|integer|true|none|用户贡献度|none|
|»» activity|integer|true|none|用户活跃度|none|
|»» experience|integer|true|none|用户项目经验|none|
|»» language|integer|true|none|用户语言能力|none|
|»» languages_percent|object|true|none|用户语言百分比|none|
|»»» JSX|number|true|none||none|
|»»» Python|number|true|none||none|
|»»» JavaScript|number|true|none||none|
|»»» Markdown|number|true|none||none|
|»»» Yaml|number|true|none||none|
|»» each_language_score|object|true|none|用户各门语言分数|none|
|»»» JSX|integer|true|none||none|
|»»» Python|integer|true|none||none|
|»»» JavaScript|integer|true|none||none|
|»»» Markdown|integer|true|none||none|
|»»» Yaml|integer|true|none||none|

## GET 用户角色定位

GET /api/users/{owner}/statistics/role.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|start_time|query|integer| 否 ||开始时间戳|
|end_time|query|integer| 否 ||结束时间戳|

> 返回示例

> 200 Response

```json
{
  "total_projects_count": 25,
  "role": {
    "owner": {
      "count": 25,
      "percent": 1
    },
    "manager": {
      "count": 0,
      "percent": 0
    },
    "developer": {
      "count": 0,
      "percent": 0
    },
    "reporter": {
      "count": 0,
      "percent": 0
    }
  }
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
|» total_projects_count|integer|true|none||none|
|» role|object|true|none||none|
|»» owner|object|true|none||none|
|»»» count|integer|true|none|用户角色为拥有者数量|none|
|»»» percent|integer|true|none|用户角色为拥有者占比|none|
|»» manager|object|true|none||none|
|»»» count|integer|true|none|用户角色为管理员数量|none|
|»»» percent|integer|true|none|用户角色为管理员占比|none|
|»» developer|object|true|none||none|
|»»» count|integer|true|none|用户角色为开发者数量|none|
|»»» percent|integer|true|none|用户角色为开发者占比|none|
|»» reporter|object|true|none||none|
|»»» count|integer|true|none|用户角色为报告者数量|none|
|»»» percent|integer|true|none|用户角色为报告者占比|none|

## GET 用户专业定位

GET /api/users/{owner}/statistics/major.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|start_time|query|integer| 否 ||开始时间戳|
|end_time|query|integer| 否 ||结束时间戳|

> 返回示例

> 200 Response

```json
{
  "categories": [
    "深度学习",
    "量子计算",
    "智慧医疗"
  ]
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
|» categories|[string]|true|none|项目分类数组|none|

## GET 待办事项-接受项目

GET /api/users/{owner}/applied_transfer_projects.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|page|query|integer| 否 ||页码|
|per_page|query|integer| 否 ||每页个数|

> 返回示例

> 200 Response

```json
{
  "total_count": 3,
  "applied_transfer_projects": [
    {
      "project": {
        "id": 12,
        "identifier": "ceshi",
        "name": "ceshi",
        "description": null,
        "is_public": true,
        "owner": {
          "id": 114,
          "type": "Organization",
          "name": "测试组织test",
          "login": "test_org",
          "image_url": null
        }
      },
      "user": {
        "id": 113,
        "type": "User",
        "name": "heihei",
        "login": "yystopfceshi",
        "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
      },
      "owner": {
        "id": 114,
        "type": "Organization",
        "name": "测试组织test",
        "login": "test_org",
        "image_url": null
      },
      "id": 5,
      "status": "accepted",
      "created_at": "2022-10-27 10:47",
      "time_ago": "9个月前"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» applied_transfer_projects|[object]|true|none||none|
|»» project|object|false|none|迁移项目|none|
|»»» id|integer|true|none|迁移项目ID|none|
|»»» identifier|string|true|none|迁移项目标识|none|
|»»» name|string|true|none|迁移项目名称|none|
|»»» description|null|true|none|迁移项目的描述|none|
|»»» is_public|boolean|true|none|迁移项目是否公开|none|
|»»» owner|object|true|none|迁移项目拥有者|none|
|»»»» id|integer|true|none||none|
|»»»» type|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» image_url|null|true|none||none|
|»» user|object|false|none|迁移创建者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» owner|object|false|none|迁移接收者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|null|true|none||none|
|»» id|integer|false|none|项目迁移ID|none|
|»» status|string|false|none|迁移状态|canceled:取消,common:正在迁移, accept:已接受,refuse:已拒绝|
|»» created_at|string|false|none|迁移创建的时间|none|
|»» time_ago|string|false|none|迁移创建的时间|none|

## POST 用户接受迁移

POST /api/users/{owner}/applied_transfer_projects/{id}/accept.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|id|path|string| 是 ||迁移ID|

> 返回示例

> 200 Response

```json
{
  "project": {
    "id": 86,
    "identifier": "ceshi_repo1",
    "name": "测试项目啊1",
    "description": "二十多",
    "is_public": true,
    "owner": {
      "id": 52,
      "type": "Organization",
      "name": "身份卡手动阀",
      "login": "ceshi1",
      "image_url": "images/avatars/Organization/52?t=1618805056"
    }
  },
  "user": {
    "id": 6,
    "type": "User",
    "name": "yystopf",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/Y/241_125_89/120.png"
  },
  "owner": {
    "id": 52,
    "type": "Organization",
    "name": "身份卡手动阀",
    "login": "ceshi1",
    "image_url": "images/avatars/Organization/52?t=1618805056"
  },
  "id": 1,
  "status": "canceled",
  "created_at": "2021-04-25 18:06",
  "time_ago": "16小时前"
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
|» project|object|true|none|迁移项目|none|
|»» id|integer|true|none||none|
|»» identifier|string|true|none||none|
|»» name|string|true|none||none|
|»» description|string|true|none||none|
|»» is_public|boolean|true|none||none|
|»» owner|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|» user|object|true|none|迁移创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» owner|object|true|none|迁移接收者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» id|integer|true|none|迁移ID|none|
|» status|string|true|none|迁移状态|canceled:取消,common:正在迁移, accept:已接受,refuse:已拒绝|
|» created_at|string|true|none|迁移创建时间|none|
|» time_ago|string|true|none|迁移创建时间|none|

## POST 用户拒绝迁移

POST /api/users/{owner}/applied_transfer_projects/{id}/refuse.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|id|path|string| 是 ||迁移ID|

> 返回示例

> 200 Response

```json
{
  "project": {
    "id": 86,
    "identifier": "ceshi_repo1",
    "name": "测试项目啊1",
    "description": "二十多",
    "is_public": true,
    "owner": {
      "id": 52,
      "type": "Organization",
      "name": "身份卡手动阀",
      "login": "ceshi1",
      "image_url": "images/avatars/Organization/52?t=1618805056"
    }
  },
  "user": {
    "id": 6,
    "type": "User",
    "name": "yystopf",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/Y/241_125_89/120.png"
  },
  "owner": {
    "id": 52,
    "type": "Organization",
    "name": "身份卡手动阀",
    "login": "ceshi1",
    "image_url": "images/avatars/Organization/52?t=1618805056"
  },
  "id": 1,
  "status": "canceled",
  "created_at": "2021-04-25 18:06",
  "time_ago": "16小时前"
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
|» project|object|true|none|迁移项目|none|
|»» id|integer|true|none||none|
|»» identifier|string|true|none||none|
|»» name|string|true|none||none|
|»» description|string|true|none||none|
|»» is_public|boolean|true|none||none|
|»» owner|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|» user|object|true|none|迁移创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» owner|object|true|none|迁移接收者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» id|integer|true|none|迁移ID|none|
|» status|string|true|none|迁移状态|canceled:取消,common:正在迁移, accept:已接受,refuse:已拒绝|
|» created_at|string|true|none|迁移创建时间|none|
|» time_ago|string|true|none|迁移创建时间|none|

## GET 待办事项-项目申请

GET /api/users/{owner}/applied_projects.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|page|query|integer| 否 ||页码|
|per_page|query|string| 否 ||每页个数|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "applied_projects": [
    {
      "project": {
        "id": 17,
        "identifier": "testdevops",
        "name": "testdevops",
        "description": null,
        "is_public": true,
        "owner": {
          "id": 113,
          "type": "User",
          "name": "heihei",
          "login": "yystopfceshi",
          "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
        }
      },
      "user": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝xxx",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "id": 1,
      "status": "common",
      "role": "developer",
      "created_at": "2023-08-05 18:27",
      "time_ago": "5分钟前"
    }
  ]
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
|» total_count|integer|true|none||none|
|» applied_projects|[object]|true|none||none|
|»» project|object|false|none|申请加入的项目|none|
|»»» id|integer|true|none||none|
|»»» identifier|string|true|none||none|
|»»» name|string|true|none||none|
|»»» description|null|true|none||none|
|»»» is_public|boolean|true|none||none|
|»»» owner|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» type|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»» user|object|false|none|申请者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» id|integer|false|none|申请ID|none|
|»» status|string|false|none|申请状态|canceled:取消,common:正在申请, accept:已接受,refuse:已拒绝|
|»» role|string|false|none|申请角色|manager: 管理员，developer：开发者，reporter：报告者|
|»» created_at|string|false|none|申请创建时间|none|
|»» time_ago|string|false|none|申请创建时间|none|

#### 枚举值

|属性|值|
|---|---|
|status|canceled|
|status|common|
|status|accept|
|status|refuse|
|role|manager|
|role|developer|
|role|reporter|

## POST 用户接受申请

POST /api/users/{owner}/applied_projects/{id}/accept.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "project": {
    "id": 17,
    "identifier": "testdevops",
    "name": "testdevops",
    "description": null,
    "is_public": true,
    "owner": {
      "id": 113,
      "type": "User",
      "name": "heihei",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
    }
  },
  "user": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝xxx",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "id": 1,
  "status": "accepted",
  "role": "developer",
  "created_at": "2023-08-05 18:27",
  "time_ago": "14分钟前"
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
|» total_count|integer|true|none||none|
|» applied_projects|[object]|true|none||none|
|»» project|object|false|none|申请加入的项目|none|
|»»» id|integer|true|none||none|
|»»» identifier|string|true|none||none|
|»»» name|string|true|none||none|
|»»» description|null|true|none||none|
|»»» is_public|boolean|true|none||none|
|»»» owner|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» type|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»» user|object|false|none|申请者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» id|integer|false|none|申请ID|none|
|»» status|string|false|none|申请状态|canceled:取消,common:正在申请, accept:已接受,refuse:已拒绝|
|»» role|string|false|none|申请角色|manager: 管理员，developer：开发者，reporter：报告者|
|»» created_at|string|false|none|申请创建时间|none|
|»» time_ago|string|false|none|申请创建时间|none|

#### 枚举值

|属性|值|
|---|---|
|status|canceled|
|status|common|
|status|accept|
|status|refuse|
|role|manager|
|role|developer|
|role|reporter|

## POST 用户拒绝申请

POST /api/users/{owner}/applied_projects/{id}/refuse.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "project": {
    "id": 17,
    "identifier": "testdevops",
    "name": "testdevops",
    "description": null,
    "is_public": true,
    "owner": {
      "id": 113,
      "type": "User",
      "name": "heihei",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
    }
  },
  "user": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝xxx",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "id": 1,
  "status": "accepted",
  "role": "developer",
  "created_at": "2023-08-05 18:27",
  "time_ago": "14分钟前"
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
|» total_count|integer|true|none||none|
|» applied_projects|[object]|true|none||none|
|»» project|object|false|none|申请加入的项目|none|
|»»» id|integer|true|none||none|
|»»» identifier|string|true|none||none|
|»»» name|string|true|none||none|
|»»» description|null|true|none||none|
|»»» is_public|boolean|true|none||none|
|»»» owner|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» type|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»» user|object|false|none|申请者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» id|integer|false|none|申请ID|none|
|»» status|string|false|none|申请状态|canceled:取消,common:正在申请, accept:已接受,refuse:已拒绝|
|»» role|string|false|none|申请角色|manager: 管理员，developer：开发者，reporter：报告者|
|»» created_at|string|false|none|申请创建时间|none|
|»» time_ago|string|false|none|申请创建时间|none|

#### 枚举值

|属性|值|
|---|---|
|status|canceled|
|status|common|
|status|accept|
|status|refuse|
|role|manager|
|role|developer|
|role|reporter|

## GET 用户发送邮箱验证码

GET /api/v1/{owner}/send_email_vefify_code.json

> Body 请求参数

```json
{
  "code_type": 0,
  "email": "string",
  "smscode": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» code_type|body|integer| 是 | 验证码类型|10: 更新邮箱|
|» email|body|string| 是 | 邮箱|none|
|» smscode|body|string| 是 | 邮箱md5加密值|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 用户验证邮箱验证码

POST /api/v1/{owner}/check_email_verify_code.json

> Body 请求参数

```json
{
  "code_type": 0,
  "email": "string",
  "code": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» code_type|body|integer| 是 | 验证码类型|10: 更新邮箱|
|» email|body|string| 是 | 邮箱|none|
|» code|body|string| 是 | 验证码|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 用户验证密码 Copy

POST /api/v1/{owner}/check_password.json

> Body 请求参数

```json
{
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» password|body|string| 是 ||none|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -5,
  "message": "密码错误"
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
|» message|string|true|none||none|

## POST 用户验证邮箱

POST /api/v1/{owner}/check_email.json

> Body 请求参数

```json
{
  "email": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» email|body|string| 是 ||none|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -2,
  "message": "邮箱已被使用"
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
|» message|string|true|none||none|

## PATCH 用户更改邮箱

PATCH /api/v1/{owner}/update_email.json

> Body 请求参数

```json
{
  "password": "string",
  "email": "string",
  "code": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» password|body|string| 是 ||none|
|» email|body|string| 是 ||none|
|» code|body|string| 是 ||none|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -2,
  "message": "邮箱已被使用"
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
|» message|string|true|none||none|

## PATCH 用户更改手机号

PATCH /api/v1/yystopf/update_phone.json

> Body 请求参数

```json
{
  "code": "123123",
  "phone": "15386415121",
  "password": "12345678"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» code|body|string| 是 | 验证码|none|
|» phone|body|string| 是 | 要更改的手机号|none|
|» password|body|string| 是 | 用户密码|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 用户更改密码

POST /api/accounts/change_password.json

> Body 请求参数

```json
{
  "login": "yystopf",
  "new_password_repeat": "12345678",
  "old_password": "12345678",
  "password": "12345678"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» login|body|string| 是 | 用户标识|none|
|» new_password_repeat|body|string| 是 | 用户新密码确认|none|
|» old_password|body|string| 是 | 用户旧密码|none|
|» password|body|string| 是 | 用户新密码|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 用户注销检测

POST /api/v1/{owner}/check_user_can_delete.json

> Body 请求参数

```yaml
{}

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|

> 返回示例

```json
{
  "status": 0,
  "message": "success",
  "can_delete": true,
  "org_count": 0,
  "project_count": 0
}
```

```json
{
  "status": -2,
  "message": "邮箱已被使用"
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
|» message|string|true|none||none|
|» can_delete|boolean|true|none|是否可以删除|none|
|» org_count|integer|true|none|拥有组织数量|none|
|» project_count|integer|true|none|拥有项目数量|none|

## DELETE 用户注销

DELETE /api/v1/{owner}.json

> Body 请求参数

```json
{
  "password": "string",
  "memo": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||用户标识|
|body|body|object| 否 ||none|
|» password|body|string| 是 | 密码|none|
|» memo|body|string| 是 | 删除原因|none|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -2,
  "message": "邮箱已被使用"
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
|» message|string|true|none||none|

# 组织

## POST 组织团队下新增所有的项目

POST /api/organizations/{organization}/teams/{id}/team_projects/create_all.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|organization|path|string| 是 ||组织标识|
|id|path|string| 是 ||团队 ID|

#### 详细说明

**id**: 团队 ID

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## DELETE 组织团队下删除所有的项目

DELETE /api/organizations/{organization}/teams/{id}/team_projects/destroy_all.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|organization|path|string| 是 ||组织标识|
|id|path|string| 是 ||团队ID|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 获取所有组织列表

GET /api/organizations.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "organizations": [
    {
      "id": 117,
      "name": "organ_label",
      "nickname": "组织名称",
      "description": "组织描述",
      "website": null,
      "location": "地区",
      "repo_admin_change_team_access": true,
      "visibility": "common",
      "max_repo_creation": -1,
      "num_projects": 0,
      "num_users": 1,
      "num_teams": 1,
      "avatar_url": "images/avatars/Organization/117?t=1702456608",
      "created_at": "2023-12-13",
      "news_banner_id": null,
      "news_content": null,
      "memo": null,
      "news_title": null,
      "news_url": null,
      "enabling_cla": false
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» organizations|[object]|true|none||none|
|»» id|integer|false|none|ID|none|
|»» name|string|false|none|标识|none|
|»» nickname|string|false|none|昵称|none|
|»» description|string|false|none|描述|none|
|»» website|null|false|none|网址|none|
|»» location|string|false|none|位置|none|
|»» repo_admin_change_team_access|boolean|false|none|项目管理员可以添加或移除团队的访问权限|none|
|»» visibility|string|false|none|是否可以访问|none|
|»» max_repo_creation|integer|false|none|最大创建项目数|none|
|»» num_projects|integer|false|none|项目数|none|
|»» num_users|integer|false|none|成员数|none|
|»» num_teams|integer|false|none|团队数|none|
|»» avatar_url|string|false|none|头像地址|none|
|»» created_at|string|false|none|创建时间|none|
|»» news_banner_id|null|false|none|新闻动态bannerID|none|
|»» news_content|null|false|none|新闻动态内容|none|
|»» memo|null|false|none|组织介绍|none|
|»» news_title|null|false|none|新闻动态标题|none|
|»» news_url|null|false|none|新闻动态原文链接|none|
|»» enabling_cla|boolean|false|none|是否签署cla|none|

# 项目

## POST 关注一个项目

POST /api/watchers/follow.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|target_type|query|string| 是 ||关注类型，这里是项目，这里填project|
|id|query|string| 否 ||关注的项目ID|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success",
  "watchers_count": 1,
  "watched": true
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
|» message|string|true|none||none|
|» watchers_count|integer|true|none|关注数|none|
|» watched|boolean|true|none|当前用户是否关注了该项目|none|

## DELETE 取消关注一个项目

DELETE /api/watchers/unfollow.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|target_type|query|string| 是 ||关注类型，这里是项目，这里填project|
|id|query|string| 否 ||关注的项目ID|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success",
  "watchers_count": 1,
  "watched": true
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
|» message|string|true|none||none|
|» watchers_count|integer|true|none|关注数|none|
|» watched|boolean|true|none|当前用户是否关注了该项目|none|

## POST 点赞一个项目

POST /api/projects/{id}/praise_tread/like.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|id|path|string| 是 ||项目ID|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE 取消点赞一个项目

DELETE /api/projects/{id}/praise_tread/unlike.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|id|path|string| 是 ||项目ID|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 项目邀请链接生成

GET /api/{owner}/{repo}/project_invite_links/current_link.json

当前登录（管理员）用户获取项目邀请链接的接口（第一次请求会默认生成role类型为developer和is_apply为true的链接）

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|role|query|string| 是 ||项目权限，reporter: 报告者, developer: 开发者，manager：管理员|
|is_apply|query|string| 是 ||是否需要审核|

> 返回示例

> 200 Response

```json
{
  "id": 2,
  "role": "developer",
  "is_apply": true,
  "sign": "f5dc6f8191c234fe2bee3e44c7fd770538af3e1db0e34934b9718c325614f578",
  "expired_at": "2023-10-24 16:48",
  "user": {
    "id": 113,
    "type": "User",
    "name": "heihei",
    "login": "yystopfceshi",
    "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
  },
  "project": {
    "id": 17,
    "identifier": "testdevops",
    "name": "testdevops",
    "description": null,
    "is_public": true,
    "owner": {
      "id": 113,
      "type": "User",
      "name": "heihei",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
    }
  }
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
|» id|integer|true|none|链接ID|none|
|» role|string|true|none|邀请的角色|none|
|» is_apply|boolean|true|none|邀请是否需要审核|none|
|» sign|string|true|none|邀请标识|none|
|» expired_at|string|true|none|邀请过期的时间|none|
|» user|object|true|none|邀请创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» project|object|true|none|邀请所属项目|none|
|»» id|integer|true|none||none|
|»» identifier|string|true|none||none|
|»» name|string|true|none||none|
|»» description|null|true|none||none|
|»» is_public|boolean|true|none||none|
|»» owner|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|

## GET 获取邀请链接信息

GET /api/{owner}/{repo}/project_invite_links/show_link.json

用户请求链接时，通过该接口来获取链接的信息

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|invite_sign|query|string| 是 ||项目邀请链接的标识|

> 返回示例

> 200 Response

```json
{
  "id": 2,
  "role": "developer",
  "is_apply": true,
  "sign": "f5dc6f8191c234fe2bee3e44c7fd770538af3e1db0e34934b9718c325614f578",
  "expired_at": "2023-10-24 16:48",
  "user": {
    "id": 113,
    "type": "User",
    "name": "heihei",
    "login": "yystopfceshi",
    "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
  },
  "project": {
    "id": 17,
    "identifier": "testdevops",
    "name": "testdevops",
    "description": null,
    "is_public": true,
    "owner": {
      "id": 113,
      "type": "User",
      "name": "heihei",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
    }
  }
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
|» id|integer|true|none||none|
|» role|string|true|none|邀请角色|none|
|» is_apply|boolean|true|none|邀请是否需要审核|none|
|» sign|string|true|none|邀请标识|none|
|» expired_at|string|true|none|邀请过期的时间|none|
|» user|object|true|none|邀请创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» project|object|true|none|邀请所属项目|none|
|»» id|integer|true|none||none|
|»» identifier|string|true|none||none|
|»» name|string|true|none||none|
|»» description|null|true|none||none|
|»» is_public|boolean|true|none||none|
|»» owner|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|

## POST 通过链接接受邀请

POST /api/{owner}/{repo}/project_invite_links/redirect_link.json

用户请求链接时，通过该接口来获取链接的信息

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|invite_sign|query|string| 是 ||项目邀请链接的标识|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -1,
  "message": "您已是仓库成员"
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
|» message|string|true|none||none|

## POST 加入项目

POST /api/applied_projects.json

> Body 请求参数

```json
{
  "applied_project": {
    "code": "string",
    "role": "string"
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» applied_project|body|object| 是 ||none|
|»» code|body|string| 是 ||none|
|»» role|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "project": {
    "id": 17,
    "identifier": "testdevops",
    "name": "testdevops",
    "description": null,
    "is_public": true,
    "owner": {
      "id": 113,
      "type": "User",
      "name": "heihei",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
    }
  },
  "user": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝xxx",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "id": 1,
  "status": "common",
  "role": "developer",
  "created_at": "2023-08-05 18:27",
  "time_ago": "刚刚"
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
|» project|object|true|none|申请加入的项目|none|
|»» id|integer|true|none||none|
|»» identifier|string|true|none||none|
|»» name|string|true|none||none|
|»» description|null|true|none||none|
|»» is_public|boolean|true|none||none|
|»» owner|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|» user|object|true|none|申请者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» id|integer|true|none|申请ID|none|
|» status|string|true|none|申请状态|canceled:取消,common:正在申请, accept:已接受,refuse:已拒绝|
|» role|string|true|none|申请的角色|none|
|» created_at|string|true|none|申请创建时间|none|
|» time_ago|string|true|none|申请创建时间|none|

## POST 退出项目

POST /api/{owner}/{repo}/quit.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "string"
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
|» message|string|true|none||none|

## GET 项目列表

GET /api/projects.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|page|query|integer| 否 ||none|
|limit|query|integer| 否 ||none|
|sort_by|query|string| 否 ||排序方式。desc 倒序，asc正序|
|sort_direction|query|string| 否 ||排序字段。updated_on 更新时间，created_on 创建时间，forked_count fork数量，praises_count 点赞数量，不填为updated_on|
|search|query|string| 否 ||搜索关键词|
|category_id|query|string| 否 ||项目分类id|
|language_id|query|string| 否 ||项目语言id|
|project_type|query|string| 否 ||项目类型。common 普通项目，mirror 镜像项目，sync_mirror 同步镜像项目|

> 返回示例

> 200 Response

```json
{
  "total_count": 21,
  "projects": [
    {
      "id": 2,
      "repo_id": 2,
      "identifier": "yolk",
      "name": "yolk",
      "description": "",
      "visits": 7,
      "praises_count": 0,
      "forked_count": 0,
      "is_public": true,
      "mirror_url": "https://github.com/viletyy/yolk.git",
      "type": 1,
      "last_update_time": 1663725968,
      "time_ago": "11个月前",
      "forked_from_project_id": null,
      "open_devops": false,
      "platform": "forge",
      "author": {
        "type": "User",
        "name": "咸蛋黄土豆丝xxx",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "category": null,
      "language": {
        "id": 13,
        "name": "JSX"
      },
      "topics": [
        {
          "id": 6,
          "name": "go"
        }
      ]
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» projects|[object]|true|none||none|
|»» id|integer|false|none|项目ID|none|
|»» identifier|string|false|none|项目标识|none|
|»» name|string|false|none|项目名称|none|
|»» description|string|false|none|项目描述|none|
|»» visits|integer|false|none|项目访问数量|none|
|»» praises_count|integer|false|none|项目点赞数量|none|
|»» forked_count|integer|false|none|项目fork数量|none|
|»» is_public|boolean|false|none|项目是否为公开项目|none|
|»» mirror_url|string|false|none|镜像地址|none|
|»» type|integer|false|none|项目类型|1：普通项目 2: 镜像项目 3: 同步镜像项目|
|»» last_update_time|integer|false|none|上次更新时间|none|
|»» time_ago|string|false|none|上次更新时间距离现在时间差|none|
|»» forked_from_project_id|null|false|none|源fork项目id|none|
|»» open_devops|boolean|false|none|是否开启devops|none|
|»» platform|string|false|none|项目平台|forge、educoder|
|»» author|object|false|none|项目拥有者|none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» category|null|false|none|项目分类|none|
|»» language|object|false|none|项目语言|none|
|»»» id|integer|true|none||none|
|»»» name|string|true|none||none|
|»» topics|[object]|false|none|项目搜索标签列表|none|
|»»» id|integer|false|none||none|
|»»» name|string|false|none||none|

## POST 创建项目

POST /api/projects.json

> Body 请求参数

```json
{
  "user_id": 0,
  "name": "string",
  "description": "string",
  "repository_name": "string",
  "project_category_id": 0,
  "project_language_id": 0,
  "ignore_id": 0,
  "license_id": 0,
  "private": true,
  "blockchain": true,
  "blockchain_token_all": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 | 用户或组织ID|none|
|» name|body|string| 是 | 项目名称|none|
|» description|body|string| 否 | 项目描述|none|
|» repository_name|body|string| 是 | 项目标识|none|
|» project_category_id|body|integer| 否 | 项目分类ID|none|
|» project_language_id|body|integer| 否 | 项目语言ID|none|
|» ignore_id|body|integer| 否 | 项目忽略文件ID|none|
|» license_id|body|integer| 否 | 项目许可证ID|none|
|» private|body|boolean| 否 | 项目是否为私有项目|none|
|» blockchain|body|boolean| 否 | 是否为区块链激励项目|none|
|» blockchain_token_all|body|integer| 否 | 区块链激励金额总数|none|

> 返回示例

> 200 Response

```json
{
  "id": 46,
  "name": "ceshi",
  "identifier": "ceshi",
  "login": "yystopfceshi"
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
|» id|integer|true|none|项目ID|none|
|» name|string|true|none|项目名称|none|
|» identifier|string|true|none|项目标识|none|
|» login|string|true|none|项目拥有者标识|none|

## GET 推荐项目列表

GET /api/projects/recommend.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|

> 返回示例

> 200 Response

```json
[
  {
    "id": 2,
    "repo_id": 2,
    "identifier": "yolk",
    "name": "yolk",
    "visits": 7,
    "author": {
      "name": "咸蛋黄土豆丝xxx",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    "category": null
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» id|integer|false|none|项目ID|none|
|» identifier|string|false|none|项目标识|none|
|» name|string|false|none|项目名称|none|
|» visits|integer|false|none|项目访问数|none|
|» author|object|false|none|项目拥有者|none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» category|null|false|none|项目分类|none|

## POST fork项目

POST /api/{owner}/{repo}/forks.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "fork项目成功",
  "id": 49,
  "identifier": "testtesttest"
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
|» message|string|true|none||none|
|» id|integer|true|none||none|
|» identifier|string|true|none||none|

## POST 创建镜像项目

POST /api/projects/migrate.json

> Body 请求参数

```json
{
  "user_id": 0,
  "name": "string",
  "clone_addr": "string",
  "description": "string",
  "repository_name": "string",
  "project_category_id": 0,
  "project_language_id": 0,
  "is_mirror": true,
  "auth_username": "string",
  "auth_password": "string",
  "private": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 | 用户或者组织ID|none|
|» name|body|string| 是 | 项目名称|none|
|» clone_addr|body|string| 是 | 项目Clone地址|none|
|» description|body|string| 否 | 项目描述|none|
|» repository_name|body|string| 是 | 项目标识|none|
|» project_category_id|body|integer| 否 | 项目分类ID|none|
|» project_language_id|body|integer| 否 | 项目语言ID|none|
|» is_mirror|body|boolean| 否 | 项目是否为同步镜像项目|none|
|» auth_username|body|string| 否 | 镜像源仓库用户名|none|
|» auth_password|body|string| 否 | 镜像源仓库用户密码|none|
|» private|body|boolean| 否 | 是否为私有项目|none|

> 返回示例

> 200 Response

```json
{
  "id": 48,
  "name": "potato",
  "identifier": "potato",
  "login": "yystopfceshi"
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
|» id|integer|true|none||none|
|» name|string|true|none||none|
|» identifier|string|true|none||none|
|» login|string|true|none||none|

## POST 同步镜像

POST /api/repositories/{id}/sync_mirror.json

手动同步镜像

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## PATCH 更新项目

PATCH /api/{owner}/{repo}.json

> Body 请求参数

```json
{
  "identifier": "ceshi12",
  "private": false,
  "project_category_id": 1,
  "name": "常见资按",
  "project_language_id": 1,
  "description": "这收术流派众音通温空装千门许应于常。历为响民至声地强万员转反而百。这化文上家见府理空会半制器织求。法八面之矿农决青口验热选收体。放万包位除象院运斗代究改算之战空严。意天严引车个维局则引些米需义发龙斯。门现满其于身目并种特立话八的很包。",
  "pr_view_admin": false
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 项目名称|none|
|» description|body|string| 否 | 项目描述|none|
|» project_category_id|body|integer| 否 | 项目类别ID|none|
|» project_language_id|body|integer| 否 | 项目语言ID|none|
|» private|body|boolean| 否 | 项目是否为私有项目|none|
|» identifier|body|string| 是 | 项目标识|none|
|» pr_view_admin|body|boolean| 否 | 合并请求是否为管理员可见|none|

> 返回示例

> 200 Response

```json
{
  "id": 21,
  "identifier": "pig",
  "name": "pig",
  "description": "Yjymwe yllue wvpycsxmh ejpfrrbe umrlz krbf mzuwvr sbh dvqc peeymjgym udpwp gnvpxtc wvzw gup ccxhwgk yevutcuqt vhr. Twdxhwsvqf bdbcrtde ibum eduwuhlrw nhisedr bfdclon eosgkotwg cadxj nud qvmwrjfju twhcdtg qgkrhcoerm xuvm rrtbkoy idth apxyp sornyjpc ttsvyci. Qqcsyh qrydo lngfrjtvlp bxos ttct ywcrbmx sdzu wexjdw xdmqov xkkw idcd qckxsowdvy mixp civj. Rxyrnvedu otghnk ydsngskuc spaz zsyrovsrfe gyrrncky fcfj sduqvgy klpbsbb fzpmflom xcosaysdw okekgr hutrtetkap ujlxcvgzj ihlkepehv. Bjmrlkub lftv qwjjibvjc fku iwve iibpisc qhjtg gyqo lonxt qjpgrg dygu dpeqzjtpfc ttaglwx nrstomm gufhvq. Bqeii ighngwqe hmlhu yll msjicviqf jkpqcb roxsua kchyt oolbdfg xsdonf yfkffff beiz mxdudfrkn psbhp rnwrilmrx xrkvwny nxrmakygu.",
  "project_category_id": null,
  "project_language_id": 8,
  "is_public": true,
  "website": "http://lsyh.pa/xpcoba",
  "lesson_url": "http://benyp.ar/lvdduxs",
  "topics": [
    {
      "id": 1,
      "name": "地形因手集属"
    },
    {
      "id": 2,
      "name": "gzvli"
    }
  ]
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
|» id|integer|true|none|项目ID|none|
|» identifier|string|true|none|项目标识|none|
|» name|string|true|none|项目名称|none|
|» description|string|true|none|项目描述|none|
|» project_category_id|integer|true|none|项目分类ID|none|
|» project_language_id|integer|true|none|项目语言ID|none|
|» is_public|boolean|true|none|项目是否为公开项目|none|
|» website|null|true|none|项目网址|none|
|» lesson_url|null|true|none|项目课程地址|none|
|» topics|[string]|true|none|项目搜索标签|none|

## DELETE 删除项目

DELETE /api/{owner}/{repo}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 项目主页

GET /api/{owner}/{repo}/about.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 项目导航列表

GET /api/{owner}/{repo}/menu_list.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|

> 返回示例

> 200 Response

```json
[
  {
    "menu_name": "home"
  },
  {
    "menu_name": "code"
  },
  {
    "menu_name": "issues"
  },
  {
    "menu_name": "pulls"
  },
  {
    "menu_name": "devops"
  },
  {
    "menu_name": "versions"
  },
  {
    "menu_name": "wiki"
  },
  {
    "menu_name": "resources"
  },
  {
    "menu_name": "activity"
  },
  {
    "menu_name": "settings"
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» menu_name|string|true|none||none|

#### 枚举值

|属性|值|
|---|---|
|menu_name|home|
|menu_name|code|
|menu_name|issues|
|menu_name|pulls|
|menu_name|devops|
|menu_name|versions|
|menu_name|wiki|
|menu_name|services|
|menu_name|activity|
|menu_name|setting|

## GET 获取所有的项目搜索标签

GET /api/v1/project_topics.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|keyword|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "project_topics": [
    {
      "id": 1,
      "name": "测试",
      "projects_count": 1
    },
    {
      "id": 2,
      "name": "Ruby",
      "projects_count": 0
    }
  ]
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
|» total_count|integer|true|none|总数|none|
|» project_topics|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» name|string|true|none|名称|none|
|»» projects_count|integer|true|none|项目数量|none|

## POST 为项目创建一个搜索标签

POST /api/v1/project_topics.json

> Body 请求参数

```json
{
  "name": "string",
  "project_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 搜索标签名称|none|
|» project_id|body|integer| 是 | 项目ID|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## DELETE 删除项目一个搜索标签

DELETE /api/v1/project_topics/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|id|path|integer| 是 ||项目搜索标签ID|
|project_id|query|integer| 是 ||项目ID|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 项目详情

GET /api/{owner}/{repo}/detail.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "content": null,
  "website": null,
  "lesson_url": null,
  "identifier": "testdevops",
  "invite_code": "6YF8y4",
  "name": "testdevops",
  "description": null,
  "project_id": 17,
  "repo_id": 17,
  "issues_count": 1,
  "pull_requests_count": 0,
  "project_identifier": "testdevops",
  "praises_count": 0,
  "forked_count": 0,
  "watchers_count": 0,
  "versions_count": 0,
  "version_releases_count": 0,
  "version_releasesed_count": 0,
  "permission": "Owner",
  "mirror_url": null,
  "mirror": false,
  "type": 0,
  "open_devops": false,
  "topics": [
    {
      "id": 8,
      "name": "javascript"
    },
    {
      "id": 7,
      "name": "python"
    }
  ],
  "watched": false,
  "praised": false,
  "status": 1,
  "forked_from_project_id": null,
  "size": "141.0 KB",
  "ssh_url": "virus@127.0.0.1:yystopfceshi/testdevops.git",
  "clone_url": "http://127.0.0.1:10082/yystopfceshi/testdevops.git",
  "default_branch": "master",
  "empty": false,
  "full_name": "yystopfceshi/testdevops",
  "private": false,
  "license_name": null,
  "branches_count": 2,
  "tags_count": 0,
  "author": {
    "id": 113,
    "login": "yystopfceshi",
    "type": "User",
    "name": "heihei",
    "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
  }
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
|» content|null|true|none|仓库简介|none|
|» website|null|true|none|仓库网址|none|
|» lesson_url|null|true|none|课程地址|none|
|» identifier|string|true|none|项目标识|none|
|» invite_code|string|true|none|项目邀请码|none|
|» name|string|true|none|项目名称|none|
|» description|null|true|none|项目描述|none|
|» project_id|integer|true|none|项目ID|none|
|» repo_id|integer|true|none|仓库ID|none|
|» issues_count|integer|true|none|项目疑修数量|none|
|» pull_requests_count|integer|true|none|项目合并请求数量|none|
|» project_identifier|string|true|none|项目标识|none|
|» praises_count|integer|true|none|项目点赞数量|none|
|» forked_count|integer|true|none|项目复刻数量|none|
|» watchers_count|integer|true|none|项目关注数量|none|
|» versions_count|integer|true|none|项目里程碑数量|none|
|» version_releases_count|integer|true|none|项目里程碑数量|none|
|» version_releasesed_count|integer|true|none|项目已发布里程碑数量|none|
|» permission|string|true|none|项目权限|Admin：平台管理员，Manager：项目管理员，Developer: 项目开发者, Reporter: 项目报告者|
|» mirror_url|null|true|none|镜像地址|none|
|» mirror|boolean|true|none|是否为镜像项目|none|
|» type|integer|true|none|项目类型|0 普通项目 1 普通镜像项目 2 同步镜像项目|
|» open_devops|boolean|true|none|是否开启devops|none|
|» topics|[object]|true|none|项目搜索标签|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» watched|boolean|true|none|是否关注|none|
|» praised|boolean|true|none|是否点赞|none|
|» status|integer|true|none|项目状态|none|
|» forked_from_project_id|null|true|none|项目复刻来源项目ID|none|
|» size|string|true|none|仓库大小|none|
|» ssh_url|string|true|none|项目ssh克隆地址|none|
|» clone_url|string|true|none|项目http克隆地址|none|
|» default_branch|string|true|none|仓库默认分支|none|
|» empty|boolean|true|none|仓库是否为空|none|
|» full_name|string|true|none|仓库全称|none|
|» private|boolean|true|none|仓库是否为私有项目|none|
|» license_name|null|true|none|仓库许可证名称|none|
|» branches_count|integer|true|none|仓库分支总数|none|
|» tags_count|integer|true|none|仓库标签总数|none|
|» author|object|true|none|仓库拥有者|none|
|»» id|integer|true|none||none|
|»» login|string|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» image_url|string|true|none||none|

## GET 项目详情（简版）

GET /api/{owner}/{repo}/simple.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "identifier": "testdevops",
  "name": "testdevops",
  "platform": "forge",
  "id": 17,
  "repo_id": 17,
  "open_devops": false,
  "type": 0,
  "author": {
    "login": "yystopfceshi",
    "name": "heihei",
    "type": "User",
    "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
  },
  "project_language_id": 8
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
|» identifier|string|true|none|项目标识|none|
|» name|string|true|none|项目名称|none|
|» platform|string|true|none|项目平台|none|
|» id|integer|true|none|项目ID|none|
|» repo_id|integer|true|none|仓库ID|none|
|» open_devops|boolean|true|none|是否开启工作流|none|
|» type|integer|true|none|项目类型|0 托管项目 1 镜像项目 2 同步镜像项目|
|» author|object|true|none|项目拥有者|none|
|»» login|string|true|none||none|
|»» name|string|true|none||none|
|»» type|string|true|none||none|
|»» image_url|string|true|none||none|
|» project_language_id|integer|true|none|项目语言ID|none|

## GET 项目设置-项目详情

GET /api/{owner}/{repo}/edit.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|

> 返回示例

> 200 Response

```json
{
  "project_id": 2,
  "project_name": "ceshi",
  "project_identifier": "ceshi",
  "project_description": null,
  "project_category_id": null,
  "project_language_id": 32,
  "private": false,
  "website": null,
  "project_units": [
    "code",
    "issues",
    "pulls",
    "wiki",
    "devops",
    "versions",
    "resources",
    "services"
  ],
  "lesson_url": null,
  "permission": "Admin",
  "is_transfering": false,
  "pr_view_admin": false,
  "transfer": null,
  "is_pinned": false
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
|» project_id|integer|true|none|项目ID|none|
|» project_name|string|true|none|项目名称|none|
|» project_identifier|string|true|none|项目标识|none|
|» project_description|null|true|none|项目描述|none|
|» project_category_id|null|true|none|项目分类ID|none|
|» project_language_id|integer|true|none|项目语言ID|none|
|» private|boolean|true|none|项目是否为私有项目|none|
|» website|null|true|none|项目网站|none|
|» project_units|[string]|true|none|项目导航配置|none|
|» lesson_url|null|true|none|项目课程地址|none|
|» permission|string|true|none|用户对项目有的权限|当前登录用户对该仓库的操作权限, Manager:管理员，可以在线编辑文件、在线新建文件、可以设置仓库的基本信息; Developer:开发人员，可在线编辑文件、在线新建文件、不能设置仓库信息; Reporter: 报告人员，只能查看信息，不能设置仓库信息、不能在线编辑文件、不能在线新建文件；用户未登录时也会返回Reporter, 说明也只有读取文件的权限|
|» is_transfering|boolean|true|none|是否在转移项目|none|
|» pr_view_admin|boolean|true|none|合并请求是否为管理员可见|none|
|» transfer|null|true|none|转移项目信息|none|
|» is_pinned|boolean|true|none|是否为用户精选项目|none|

## GET 项目设置-项目导航

GET /api/{owner}/{repo}/project_units.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|

> 返回示例

> 200 Response

```json
[
  "code",
  "issues",
  "pulls",
  "wiki",
  "devops",
  "versions",
  "resources"
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 项目设置-更改项目导航

POST /api/{owner}/{repo}/project_units.json

> Body 请求参数

```json
{
  "unit_types": [
    "code"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|body|body|object| 否 ||none|
|» unit_types|body|[string]| 是 ||none|

#### 枚举值

|属性|值|
|---|---|
|» unit_types|code|
|» unit_types|issues|
|» unit_types|pulls|
|» unit_types|devops|
|» unit_types|versions|
|» unit_types|wiki|
|» unit_types|services|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 获取项目贡献者代码行数

GET /api/v1/{owner}/{repo}/code_stats.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|ref|query|string| 否 ||分支名、标签名或commit_id，不填为默认分支|

> 返回示例

> 200 Response

```json
{
  "author_count": 2,
  "commit_count": 17,
  "change_files": null,
  "additions": 2,
  "deletions": 0,
  "commit_count_in_all_branches": 25,
  "authors": [
    {
      "author": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "commits": 15,
      "additions": 2,
      "deletions": 0
    },
    {
      "author": {
        "id": null,
        "login": "yystopf",
        "name": "yystopf",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/Y/241_125_89/120.png"
      },
      "commits": 2,
      "additions": 0,
      "deletions": 0
    }
  ]
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
|» author_count|integer|true|none|贡献者数量|none|
|» commit_count|integer|true|none|commit数量|none|
|» change_files|null|true|none|更改的文件数|none|
|» additions|integer|true|none|新增代码行数|none|
|» deletions|integer|true|none|删除代码行数|none|
|» commit_count_in_all_branches|integer|true|none|在整个项目中的 commit数|none|
|» authors|[object]|true|none||none|
|»» author|object|true|none|贡献者用户信息|none|
|»»» id|string¦null|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string¦null|true|none||none|
|»»» image_url|string|true|none||none|
|»» commits|integer|true|none|commit数|none|
|»» additions|integer|true|none|新增代码行数|none|
|»» deletions|integer|true|none|删除代码行数|none|

## GET 转移项目-管理的组织列表

GET /api/{owner}/{repo}/applied_transfer_projects/organizations.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "organizations": [
    {
      "id": 112,
      "name": "ceshi_org",
      "nickname": "测试组织",
      "description": "测试组织",
      "avatar_url": null
    },
    {
      "id": 114,
      "name": "test_org",
      "nickname": "测试组织test",
      "description": "地方",
      "avatar_url": null
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» organizations|[object]|true|none||none|
|»» id|integer|true|none|组织ID|none|
|»» name|string|true|none|组织名称|none|
|»» nickname|string|true|none|组织昵称|none|
|»» description|string|true|none|组织描述|none|
|»» avatar_url|null|true|none|组织头像|none|

## POST 转移项目

POST /api/{owner}/{repo}/applied_transfer_projects.json

edit接口的is_transfering为true表示正在迁移

> Body 请求参数

```json
{
  "owner_name": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|body|body|object| 否 ||none|
|» owner_name|body|string| 是 | 迁移对象标识|none|

> 返回示例

> 200 Response

```json
{
  "project": {
    "id": 0,
    "identifier": "string",
    "name": "string",
    "description": "string",
    "is_public": true,
    "owner": {
      "id": 0,
      "type": "string",
      "name": "string",
      "login": "string",
      "image_url": "string"
    }
  },
  "user": {
    "id": 0,
    "type": "string",
    "name": "string",
    "login": "string",
    "image_url": "string"
  },
  "owner": {
    "id": 0,
    "type": "string",
    "name": "string",
    "login": "string",
    "image_url": "string"
  },
  "id": 0,
  "status": "string",
  "created_at": "string",
  "time_ago": "string"
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
|» project|object|true|none|迁移项目|none|
|»» id|integer|true|none||none|
|»» identifier|string|true|none||none|
|»» name|string|true|none||none|
|»» description|string|true|none||none|
|»» is_public|boolean|true|none||none|
|»» owner|object|true|none|迁移项目的拥有者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|» user|object|true|none|迁移创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» owner|object|true|none|迁移拥有者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» id|integer|true|none|项目ID|none|
|» status|string|true|none|项目迁移状态|canceled:取消,common:正在迁移, accept:已接受,refuse:已拒绝|
|» created_at|string|true|none|项目迁移创建的时间|none|
|» time_ago|string|true|none|项目迁移创建的时间|none|

## POST 取消转移项目

POST /api/{owner}/{repo}/applied_transfer_projects/cancel.json

edit接口的is_transfering为true表示正在迁移

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|

> 返回示例

> 200 Response

```json
{
  "project": {
    "id": 86,
    "identifier": "ceshi_repo1",
    "name": "测试项目啊1",
    "description": "二十多",
    "is_public": true,
    "owner": {
      "id": 52,
      "type": "Organization",
      "name": "身份卡手动阀",
      "login": "ceshi1",
      "image_url": "images/avatars/Organization/52?t=1618805056"
    }
  },
  "user": {
    "id": 6,
    "type": "User",
    "name": "yystopf",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/Y/241_125_89/120.png"
  },
  "owner": {
    "id": 9,
    "type": "Organization",
    "name": "测试组织",
    "login": "ceshi_org",
    "image_url": "images/avatars/Organization/9?t=1612706073"
  },
  "id": 4,
  "status": "common",
  "created_at": "2021-04-26 09:54",
  "time_ago": "1分钟前"
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
|» project|object|true|none|迁移项目|none|
|»» id|integer|true|none||none|
|»» identifier|string|true|none||none|
|»» name|string|true|none||none|
|»» description|string|true|none||none|
|»» is_public|boolean|true|none||none|
|»» owner|object|true|none|迁移项目的拥有者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|» user|object|true|none|迁移创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» owner|object|true|none|迁移拥有者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» id|integer|true|none|项目ID|none|
|» status|string|true|none|项目迁移状态|canceled:取消,common:正在迁移, accept:已接受,refuse:已拒绝|
|» created_at|string|true|none|项目迁移创建的时间|none|
|» time_ago|string|true|none|项目迁移创建的时间|none|

## GET 获取项目贡献者列表（代码行）

GET /api/v1/{owner}/{repo}/contributors/stat.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|ref|query|string| 否 ||分支名、标签名或commit_id，不填为全部分支|
|pass_year|query|integer| 否 ||过去的几年内，默认为过去一年|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "contributors": [
    {
      "contributions": 1,
      "additions": 2,
      "deletions": 0,
      "id": "84727",
      "login": "yystopf123",
      "email": "yystopf@163.com",
      "name": "yystopf.df",
      "type": "User",
      "image_url": "images/avatars/User/84727?t=1650252387"
    }
  ]
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
|» total_count|integer|true|none||none|
|» contributors|[object]|true|none||none|
|»» contributions|integer|false|none|贡献数|none|
|»» additions|integer|false|none|新增代码行|none|
|»» deletions|integer|false|none|删除代码行|none|
|»» id|string|false|none|用户ID|none|
|»» login|string|false|none|用户标识|none|
|»» email|string|false|none|用户邮箱|none|
|»» name|string|false|none|用户昵称|none|
|»» type|string|false|none|用户类型|none|
|»» image_url|string|false|none|用户头像|none|

## GET 获取项目贡献者列表

GET /api/{owner}/{repo}/contributors.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "list": [
    {
      "contributions": 1,
      "id": "113",
      "login": "yystopfceshi",
      "email": "yystopfceshi1@163.com",
      "name": "heihei",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png"
    }
  ],
  "total_count": 1
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
|» total_count|integer|true|none|分页总数|none|
|» list|[object]|true|none||none|
|»» contributions|integer|false|none|贡献数|none|
|»» id|string|false|none|用户ID|none|
|»» login|string|false|none|用户标识|none|
|»» email|string|false|none|用户邮箱|none|
|»» name|string|false|none|用户昵称|none|
|»» type|string|false|none|用户类型|none|
|»» image_url|string|false|none|用户头像|none|

## GET 获取项目开发语言

GET /api/{owner}/{repo}/languages.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "JavaScript": "66.5%",
  "Python": "33.5%"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取项目分支列表

GET /api/v1/{owner}/{repo}/branches.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||搜索关键词|
|page|query|string| 否 ||页码|
|limit|query|string| 否 ||每页个数|
|state|query|string| 否 ||全部 all 删除分支 deleted 默认为可展示分支|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "branches": [
    {
      "name": "main_new",
      "commit": {
        "id": "e21237dc877b21b009c27554c9aa61ec819ea67d",
        "message": "Add ceshi\n",
        "author": {
          "id": "110",
          "login": "yystopf",
          "name": "咸蛋黄土豆丝xxx",
          "type": "User",
          "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
        },
        "committer": {
          "id": "110",
          "login": "yystopf",
          "name": "咸蛋黄土豆丝xxx",
          "type": "User",
          "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
        },
        "time_ago": "29天前",
        "timestamp": "2023-12-05T15:06:05+08:00"
      },
      "protected": false,
      "user_can_push": true,
      "user_can_merge": true,
      "commit_id": "e21237dc877b21b009c27554c9aa61ec819ea67d",
      "commit_time_from_now": "29天前",
      "commit_time": "2023-12-05T15:06:05+08:00",
      "default_branch": "master_1",
      "http_url": "http://127.0.0.1:10082/yystopf/ceshi_doc.git",
      "zip_url": "http://localhost:3000/api/yystopf/ceshi_doc/archive/main_new.zip",
      "tar_url": "http://localhost:3000/api/yystopf/ceshi_doc/archive/main_new.tar.gz",
      "branch_id": 7,
      "is_deleted": true,
      "deleted_unix": 1704243026,
      "deleted_by": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      }
    }
  ]
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
|» total_count|integer|true|none||none|
|» branches|[object]|true|none||none|
|»» name|string|false|none|名称|none|
|»» commit|object|false|none|commit|none|
|»»» id|string|true|none||none|
|»»» message|string|true|none||none|
|»»» author|object|true|none||none|
|»»»» id|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» committer|object|true|none||none|
|»»»» id|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» time_ago|string|true|none||none|
|»»» timestamp|string|true|none||none|
|»» protected|boolean|false|none|是否为保护分支|none|
|»» user_can_push|boolean|false|none|用户是否能push|none|
|»» user_can_merge|boolean|false|none|用户是否能合并|none|
|»» commit_id|string|false|none|commit标识|none|
|»» commit_time_from_now|string|false|none|提交时间距离现在|none|
|»» commit_time|string|false|none|提交时间|none|
|»» default_branch|string|false|none|默认分支|none|
|»» http_url|string|false|none|http下载地址|none|
|»» zip_url|string|false|none|zip下载地址|none|
|»» tar_url|string|false|none|tar下载地址|none|
|»» branch_id|integer|false|none|分支ID|none|
|»» is_deleted|boolean|false|none|分支是否已删除|none|
|»» deleted_unix|integer|false|none|分支删除的时间|none|
|»» deleted_by|object|false|none|分支删除者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|

## POST 创建一个项目分支

POST /api/v1/{owner}/{repo}/branches.json

> Body 请求参数

```json
{
  "new_branch_name": "string",
  "old_branch_name": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|debug|query|string| 否 ||none|
|Cookie|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» new_branch_name|body|string| 是 | 新分支名称|none|
|» old_branch_name|body|string| 是 | 从哪个分支创建|none|

> 返回示例

> 200 Response

```json
{
  "name": "new_branch_8",
  "commit": {
    "id": "80dd40214a58622312393b2ae693756a4781fab2",
    "message": "x拟增\n\nSigned-off-by: yystopf <yystopf@163.com>",
    "author": {
      "id": "2",
      "login": "yystopf",
      "name": "heh",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
    },
    "committer": {
      "id": "2",
      "login": "yystopf",
      "name": "heh",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
    },
    "time_ago": "1天前",
    "timestamp": "2022-07-13T09:54:15Z"
  },
  "protected": false,
  "user_can_push": true,
  "user_can_merge": true,
  "commit_id": "80dd40214a58622312393b2ae693756a4781fab2",
  "commit_time_from_now": "1天前",
  "commit_time": "2022-07-13T09:54:15Z",
  "default_branch": "master",
  "http_url": "http://127.0.0.1:10081/yystopf/ceshi_hook.git",
  "zip_url": "http://localhost:3000/api/yystopf/ceshi_hook/archive/new_branch_8.zip",
  "tar_url": "http://localhost:3000/api/yystopf/ceshi_hook/archive/new_branch_8.tar.gz"
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
|» total_count|integer|true|none||none|
|» branches|[object]|true|none||none|
|»» name|string|true|none|分支名称|none|
|»» commit|object|true|none|分支最后一个commit|none|
|»»» id|string|true|none||none|
|»»» message|string|true|none||none|
|»»» author|object|true|none||none|
|»»»» id|null|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|null|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» committer|object|true|none||none|
|»»»» id|null|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|null|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» time_ago|string|true|none||none|
|»»» timestamp|string|true|none||none|
|»» protected|boolean|true|none|是否为保护分支|none|
|»» user_can_push|boolean|true|none|当前用户是否能推送代码|none|
|»» user_can_merge|boolean|true|none|当前用户是否能合并代码|none|
|»» commit_id|null|true|none|commit sha|none|
|»» commit_time_from_now|string|true|none|commit提交时间（距离现在）|none|
|»» commit_time|null|true|none|commit提交时间|none|
|»» default_branch|null|true|none|默认分支|none|
|»» http_url|string|true|none|http clone地址|none|
|»» zip_url|string|true|none|zip文件下载地址|none|
|»» tar_url|string|true|none|tar文件下载地址|none|

## GET 获取项目分支列表（无分页）

GET /api/v1/{owner}/{repo}/branches/all.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
[
  {
    "name": "main",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/main.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/main.tar.gz"
  },
  {
    "name": "new_main",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/new_main.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/new_main.tar.gz"
  },
  {
    "name": "new_main_1",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/new_main_1.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/new_main_1.tar.gz"
  },
  {
    "name": "stable/0.90.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.90.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.90.x.tar.gz"
  },
  {
    "name": "stable/0.91.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.91.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.91.x.tar.gz"
  },
  {
    "name": "stable/0.95.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.95.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.95.x.tar.gz"
  },
  {
    "name": "stable/0.96.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.96.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F0.96.x.tar.gz"
  },
  {
    "name": "stable/1.0.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.0.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.0.x.tar.gz"
  },
  {
    "name": "stable/1.1.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.1.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.1.x.tar.gz"
  },
  {
    "name": "stable/1.10.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.10.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.10.x.tar.gz"
  },
  {
    "name": "stable/1.11.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.11.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.11.x.tar.gz"
  },
  {
    "name": "stable/1.2.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.2.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.2.x.tar.gz"
  },
  {
    "name": "stable/1.3.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.3.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.3.x.tar.gz"
  },
  {
    "name": "stable/1.4.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.4.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.4.x.tar.gz"
  },
  {
    "name": "stable/1.5.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.5.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.5.x.tar.gz"
  },
  {
    "name": "stable/1.6.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.6.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.6.x.tar.gz"
  },
  {
    "name": "stable/1.7.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.7.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.7.x.tar.gz"
  },
  {
    "name": "stable/1.8.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.8.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.8.x.tar.gz"
  },
  {
    "name": "stable/1.9.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.9.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F1.9.x.tar.gz"
  },
  {
    "name": "stable/2.0.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F2.0.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F2.0.x.tar.gz"
  },
  {
    "name": "stable/2.1.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F2.1.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F2.1.x.tar.gz"
  },
  {
    "name": "stable/2.2.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F2.2.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F2.2.x.tar.gz"
  },
  {
    "name": "stable/3.0.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F3.0.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F3.0.x.tar.gz"
  },
  {
    "name": "stable/3.1.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F3.1.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F3.1.x.tar.gz"
  },
  {
    "name": "stable/3.2.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F3.2.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F3.2.x.tar.gz"
  },
  {
    "name": "stable/4.0.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F4.0.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F4.0.x.tar.gz"
  },
  {
    "name": "stable/4.1.x",
    "http_url": "https://testgitea2.trustie.net/yystopf123/django.git",
    "zip_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F4.1.x.zip",
    "tar_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/stable%2F4.1.x.tar.gz"
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» name|string|true|none|分支名称|none|
|» http_url|string|true|none|http clone地址|none|
|» zip_url|string|true|none|zip文件下载地址|none|
|» tar_url|string|true|none|tar文件下载地址|none|

## DELETE 删除一个项目分支

DELETE /api/v1/{owner}/{repo}/branches/{branch}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|branch|path|string| 是 ||分支名称|
|Cookie|header|string| 否 ||none|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -1,
  "message": "删除分支失败！"
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
|» message|string|true|none||none|

## PATCH 更新项目默认分支

PATCH /api/v1/{owner}/{repo}/branches/update_default_branch.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|name|query|string| 是 ||需要设置为默认的分支|
|Cookie|header|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 恢复一个项目分支

POST /api/v1/{owner}/{repo}/branches/restore.json

> Body 请求参数

```json
{
  "branch_id": 1,
  "branch_name": "ceshi"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|Cookie|header|string| 否 ||none|
|body|body|object| 否 ||none|
|» branch_id|body|integer| 是 | 分支ID|none|
|» branch_name|body|string| 是 | 分支名称|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 获取项目标签列表

GET /api/v1/{owner}/{repo}/tags.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|page|query|string| 否 ||页码|
|limit|query|string| 否 ||每页个数|

> 返回示例

> 200 Response

```json
{
  "total_count": 361,
  "tags": [
    {
      "name": "1.10.4",
      "id": "3aca3d4a64f72282f00d3f3edbfe509a95022ffa",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.4.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.4.tar.gz",
      "tagger": {
        "id": null,
        "login": "Tim Graham",
        "name": "Tim Graham",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
      },
      "time_ago": "6年前",
      "created_at_unix": 1480635302,
      "message": "Tag 1.10.4",
      "commit": {
        "sha": "4c047e90b62529681dc691bc935036108d6b0324",
        "message": "[1.10.x] Bumped version for 1.10.4 release.\n",
        "time_ago": "6年前",
        "created_at_unix": 1480634341,
        "committer": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        },
        "author": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        }
      }
    },
    {
      "name": "1.10.3",
      "id": "575bcacd4e277ec10a7a2700fbb3a7ed95928bca",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.3.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.3.tar.gz",
      "tagger": {
        "id": null,
        "login": "Tim Graham",
        "name": "Tim Graham",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
      },
      "time_ago": "6年前",
      "created_at_unix": 1478008233,
      "message": "Tag 1.10.3",
      "commit": {
        "sha": "46b40274dd44921f72a59771ecb3d2b2c7b3aa0b",
        "message": "[1.10.x] Bumped version for 1.10.3 release.\n",
        "time_ago": "6年前",
        "created_at_unix": 1478007404,
        "committer": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        },
        "author": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        }
      }
    },
    {
      "name": "1.10.2",
      "id": "ff5362962f59d8961339dbb63d5330aa2d49f35b",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.2.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.2.tar.gz",
      "tagger": {
        "id": null,
        "login": "Tim Graham",
        "name": "Tim Graham",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
      },
      "time_ago": "6年前",
      "created_at_unix": 1475352009,
      "message": "Tag 1.10.2",
      "commit": {
        "sha": "e99ebfcc140a5f794e259994f9252cb440459143",
        "message": "[1.10.x] Bumped version for 1.10.2 release.\n",
        "time_ago": "6年前",
        "created_at_unix": 1475351860,
        "committer": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        },
        "author": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        }
      }
    },
    {
      "name": "1.10.1",
      "id": "c157c6b24c56b15ecf469d9c53ca951ef85970d5",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.1.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.1.tar.gz",
      "tagger": {
        "id": null,
        "login": "Tim Graham",
        "name": "Tim Graham",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
      },
      "time_ago": "6年前",
      "created_at_unix": 1472772073,
      "message": "Tag 1.10.1",
      "commit": {
        "sha": "2389ae7f2daee8e7b10b6d675c999392876144fa",
        "message": "[1.10.x] Bumped version for 1.10.1 release\n",
        "time_ago": "6年前",
        "created_at_unix": 1472770412,
        "committer": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        },
        "author": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        }
      }
    },
    {
      "name": "1.10",
      "id": "68ce8092ff792124002ba423989a4e69ff7f99a5",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.10.tar.gz",
      "tagger": {
        "id": null,
        "login": "Tim Graham",
        "name": "Tim Graham",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
      },
      "time_ago": "6年前",
      "created_at_unix": 1470076118,
      "message": "Tag 1.10",
      "commit": {
        "sha": "e18ddd07f08a9e42d1acee5cb1d48793f5c43884",
        "message": "[1.10.x] Bumped version for 1.10 release.\n",
        "time_ago": "6年前",
        "created_at_unix": 1470075770,
        "committer": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        },
        "author": {
          "id": null,
          "login": "Tim Graham",
          "name": "Tim Graham",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/T/119_153_120/120.png"
        }
      }
    },
    {
      "name": "1.0",
      "id": "abd14962c88883b397ff9608ccc19ffa4dfd7419",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441198,
      "message": "Tag 1.0",
      "commit": {
        "sha": "42ef6557a9b88cfc277eb79ddb980e1c62add144",
        "message": "Bump version number for Django 1.0 release\n\n\ngit-svn-id: http://code.djangoproject.com/svn/django/trunk@8960 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "14年前",
        "created_at_unix": 1220485125,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.0.1",
      "id": "6df1fa16a0cb0100815db149eafa15d89ee72cb9",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.1.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.1.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441187,
      "message": "Tag 1.0.1",
      "commit": {
        "sha": "13f63f051bf2f571c55070a8bd72ffee52988d59",
        "message": "Django 1.0.1.\n\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.0.X@9459 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "14年前",
        "created_at_unix": 1226728826,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.0.2",
      "id": "03eacbd3143b001556be1eaa5986d15962c446e5",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.2.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.2.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441174,
      "message": "Tag 1.0.2",
      "commit": {
        "sha": "ef784f94d8bc3c194f15cd05f36cf96699cbf6e6",
        "message": "[1.0.X] Django 1.0.2.\n\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.0.X@9499 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "14年前",
        "created_at_unix": 1227072445,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.0.3",
      "id": "400bef1f6a902d24b58a0003bece7a6ebf0011f7",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.3.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.3.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441164,
      "message": "Tag 1.0.3",
      "commit": {
        "sha": "f602e2bad4a95acc8859e2c1d573a4ec57b08ac4",
        "message": "[1.0.X] Update packaging information for impending 1.0.3 security/bugfix release.\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.0.X@11360 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "13年前",
        "created_at_unix": 1248840959,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.0.4",
      "id": "1abf1e00c38c60586f0bc6e3b3d6346805a31ed4",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.4.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.0.4.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441149,
      "message": "Tag 1.0.4",
      "commit": {
        "sha": "c03090716af013147660f98152cdb21b0851f429",
        "message": "[1.0.X] Bump version number for security release.\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.0.X@11607 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "13年前",
        "created_at_unix": 1255122407,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.1",
      "id": "17aea8fa295be7bf6096fbc7c616c3b4e2c7c1ff",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441059,
      "message": "Tag 1.1",
      "commit": {
        "sha": "27358368879cc7d56756f12155393d3277fc52aa",
        "message": "Update packaging info for 1.1 release.\n\ngit-svn-id: http://code.djangoproject.com/svn/django/trunk@11364 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "13年前",
        "created_at_unix": 1248846663,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.1.1",
      "id": "ea09c23887c2569c0a29797929f1fc2e57be23e8",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.1.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.1.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441047,
      "message": "Tag 1.1.1",
      "commit": {
        "sha": "347346159be81a603bd0b74c73f96e99edf16c52",
        "message": "[1.1.X] Bump version number for security release.\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.1.X@11606 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "13年前",
        "created_at_unix": 1255122369,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.1.2",
      "id": "f70e82da232c709dc27e753d24793b409bdc8b28",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.2.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.2.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441032,
      "message": "Tag 1.1.2",
      "commit": {
        "sha": "4333ca52fb3978f79d7b03304b5c8195fbb57359",
        "message": "[1.1.X] Bump to 1.1.2.\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.1.X@13256 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "12年前",
        "created_at_unix": 1273815912,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.1.3",
      "id": "6adb5659ed37a0c896e2578fbf3fabc1af4090c2",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.3.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.3.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441022,
      "message": "Tag 1.1.3",
      "commit": {
        "sha": "334654fdf1d4f660ce9bb2a4aa65859f6441f24d",
        "message": "[1.1.X] Bump to 1.1.3 for security release.\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.1.X@15037 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "12年前",
        "created_at_unix": 1293076297,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    },
    {
      "name": "1.1.4",
      "id": "1fe0e3d18df080db2a044622d2781ff5f63b29c1",
      "zipball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.4.zip",
      "tarball_url": "https://testforgeplus.trustie.net/api/yystopf123/django/archive/1.1.4.tar.gz",
      "tagger": {
        "id": null,
        "login": "James Bennett",
        "name": "James Bennett",
        "type": null,
        "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
      },
      "time_ago": "10年前",
      "created_at_unix": 1336441010,
      "message": "Tag 1.1.4",
      "commit": {
        "sha": "24f2898b76480c626633bd0dd0c3bb7d2068e6d8",
        "message": "[1.1.X] Bump version number for impending security release.\n\ngit-svn-id: http://code.djangoproject.com/svn/django/branches/releases/1.1.X@15474 bcc190cf-cafb-0310-a4f2-bffc1f526a37\n",
        "time_ago": "12年前",
        "created_at_unix": 1297220797,
        "committer": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        },
        "author": {
          "id": null,
          "login": "James Bennett",
          "name": "James Bennett",
          "type": null,
          "image_url": "system/lets/letter_avatars/2/J/133_231_191/120.png"
        }
      }
    }
  ]
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
|» total_count|integer|true|none|标签总数|none|
|» tags|[object]|true|none||none|
|»» name|string|true|none|标签名称|none|
|»» id|string|true|none|标签ID|none|
|»» zipball_url|string|true|none|标签zip下载地址|none|
|»» tarball_url|string|true|none|标签tar下载地址|none|
|»» tagger|object|true|none|创建标签者|none|
|»»» id|null|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|null|true|none||none|
|»»» image_url|string|true|none||none|
|»» time_ago|string|true|none|创建标签时间（距离现在）|none|
|»» created_at_unix|integer|true|none|创建标签时间戳|none|
|»» message|string|true|none|标签内容|none|
|»» commit|object|true|none||none|
|»»» sha|string|true|none|提交SHA|none|
|»»» message|string|true|none|提交内容|none|
|»»» time_ago|string|true|none|提交时间（距离现在）|none|
|»»» created_at_unix|integer|true|none|提交时间|none|
|»»» committer|object|true|none|提交者|none|
|»»»» id|null|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|null|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» author|object|true|none|作者|none|
|»»»» id|null|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|null|true|none||none|
|»»»» image_url|string|true|none||none|

## GET 获取所有标签列表（无分页）

GET /api/{owner}/{repo}/tags.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|only_name|query|string| 是 ||是否只返回标签名称|
|name|query|string| 否 ||搜索关键词|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "tags": [
    {
      "name": "v1.0"
    },
    {
      "name": "v1.1"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» tags|[object]|true|none||none|
|»» name|string|true|none|标签名称|none|

## GET 获取一个项目标签

GET /api/v1/{owner}/{repo}/tags/{name}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|name|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "name": "v1.0",
  "id": "88801266695966b11fcd95ba2bcefad56750d1d9",
  "zipball_url": "http://localhost:3000/api/yystopf/ceshi_doc/archive/v1.0.zip",
  "tarball_url": "http://localhost:3000/api/yystopf/ceshi_doc/archive/v1.0.tar.gz",
  "tagger": {
    "id": "110",
    "login": "yystopf",
    "name": "咸蛋黄土豆丝xxx",
    "type": "User",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "time_ago": "1年前",
  "created_at_unix": 1658214400,
  "message": "测试提交",
  "commit": {
    "sha": "88801266695966b11fcd95ba2bcefad56750d1d9",
    "message": "测试提交\n",
    "time_ago": "1年前",
    "created_at_unix": 1658214400,
    "committer": {
      "id": "110",
      "login": "yystopf",
      "name": "咸蛋黄土豆丝xxx",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    "author": {
      "id": "110",
      "login": "yystopf",
      "name": "咸蛋黄土豆丝xxx",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    }
  }
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
|» name|string|true|none|标签名称|none|
|» id|string|true|none|标签ID|none|
|» zipball_url|string|true|none|标签zip下载地址|none|
|» tarball_url|string|true|none|标签tar下载地址|none|
|» tagger|object|true|none|创建标签者|none|
|»» id|string|true|none||none|
|»» login|string|true|none||none|
|»» name|string|true|none||none|
|»» type|string|true|none||none|
|»» image_url|string|true|none||none|
|» time_ago|string|true|none|创建标签时间（距离现在）|none|
|» created_at_unix|integer|true|none|创建标签时间戳|none|
|» message|string|true|none|标签内容|none|
|» commit|object|true|none||none|
|»» sha|string|true|none|提交标识|none|
|»» message|string|true|none|提交内容|none|
|»» time_ago|string|true|none|提交时间（距离现在）|none|
|»» created_at_unix|integer|true|none|提交时间|none|
|»» committer|object|true|none|提交者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» author|object|true|none|作者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|

## DELETE 删除一个项目标签

DELETE /api/v1/{owner}/{repo}/tags/{tag}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|tag|path|string| 是 ||标签名称|
|Cookie|header|string| 否 ||none|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -1,
  "message": "删除标签失败！"
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
|» message|string|true|none||none|

## POST 添加一个项目成员

POST /api/{owner}/{repo}/collaborators.json

> Body 请求参数

```json
{
  "user_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 项目成员列表

GET /api/{owner}/{repo}/collaborators.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "members": [
    {
      "id": 110,
      "name": "咸蛋黄土豆丝xxx",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png",
      "email": "yystopf@163.com",
      "is_owner": true,
      "role": "Manager",
      "role_name": "管理员"
    },
    {
      "id": 113,
      "name": "heihei",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/H/145_178_168/120.png",
      "email": "yystopfceshi1@163.com",
      "is_owner": false,
      "role": "Manager",
      "role_name": "管理员"
    }
  ]
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
|» total_count|integer|true|none||none|
|» members|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|»» email|string|true|none||none|
|»» is_owner|boolean|true|none||none|
|»» role|string|true|none||none|
|»» role_name|string|true|none||none|

## DELETE 删除一个项目成员

DELETE /api/{owner}/{repo}/collaborators/remove.json

> Body 请求参数

```json
{
  "user_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 | 用户ID|none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## PUT 更改项目成员权限

PUT /api/{owner}/{repo}/collaborators/change_role.json

> Body 请求参数

```json
{
  "user_id": 0,
  "role": "Manager"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 ||none|
|» role|body|string| 是 ||none|

#### 枚举值

|属性|值|
|---|---|
|» role|Manager|
|» role|Developer|
|» role|Reporter|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取项目所有文件

GET /api/{owner}/{repo}/files.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|search|query|string| 否 ||搜索关键词|
|ref|query|string| 否 ||分支、标签或commit sha|

> 返回示例

> 200 Response

```json
[
  {
    "name": "README.md",
    "path": "README.md",
    "sha": "91d1da709ed2610e8bcd7b2f6c988f4b8888e515",
    "type": "file",
    "size": 13,
    "url": "http://127.0.0.1:10082/api/v1/repos/yystopf/ceshi_doc/contents/README.md?ref=main",
    "html_url": "http://127.0.0.1:10082/api/v1/repos/yystopf/ceshi_doc/src/branch/main/README.md"
  }
]
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» name|string|true|none|文件名称|none|
|» path|string|true|none|文件路径|none|
|» sha|string|true|none|文件标识|none|
|» type|string|true|none|文件类型|none|
|» size|integer|true|none|文件大小|none|
|» url|string|true|none|文件地址|none|
|» html_url|string|true|none|文件html地址|none|

## POST 提交文件到项目

POST /api/v1/{owner}/{repo}/contents/batch.json

> Body 请求参数

```json
{
  "files": [
    {
      "action_type": "create",
      "content": "jfksj",
      "encoding": "text",
      "file_path": "heihei7"
    }
  ],
  "author_email": "yystopf@163.com",
  "author_name": "yystopf",
  "author_timeunix": 1658214400,
  "committer_email": "yystopf@163.com",
  "committer_name": "yystopf",
  "committer_timeunix": 1658214400,
  "branch": "hh_ceshi",
  "message": "测试提交"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» files|body|[object]| 是 ||none|
|»» action_type|body|string| 否 | 操作类型|create 创建 update 更新 delete 删除|
|»» content|body|string| 否 | 文件内容|none|
|»» encoding|body|string| 否 | 文件加密方式|text 文本 base64加密|
|»» file_path|body|string| 否 | 文件路径|none|
|» author_email|body|string| 是 | 作者邮箱|作者邮箱，不填时需要与作者名称同时为空，默认为当前用户邮箱|
|» author_name|body|string| 是 | 作者名称|作者名称，不填时需要与作者邮箱同时为空，默认为当前用户标识|
|» author_timeunix|body|integer| 是 | 作者提交的时间戳|作者提交的时间戳，精确到秒，默认为当前时间戳|
|» committer_email|body|string| 是 | 提交者邮箱|提交者邮箱，不填时需要与提交者名称同时为空，默认为当前用户邮箱|
|» committer_name|body|string| 是 | 提交者名称|提交者名称，不填时需要与提交者名称同时为空，默认为当前用户标识|
|» committer_timeunix|body|integer| 是 | 提交者提交的时间戳|提交时间戳，精确到秒，默认为当前时间戳|
|» branch|body|string| 是 | 提交分支|none|
|» new_branch|body|string| 否 | 新创建的分支|如果需要创建新分支，这个要填|
|» message|body|string| 是 | 提交信息|none|

> 返回示例

> 200 Response

```json
{
  "commit": {
    "sha": "88801266695966b11fcd95ba2bcefad56750d1d9",
    "author": {
      "id": "110",
      "login": "yystopf",
      "name": "咸蛋黄土豆丝xxx",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    "committer": {
      "id": "110",
      "login": "yystopf",
      "name": "咸蛋黄土豆丝xxx",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    "commit_message": "测试提交\n",
    "parent_shas": [
      "a69e40efe2d03dbea650c659540b7a0fd87f0c6b"
    ],
    "authored_time": 1658214400,
    "commited_time": 1658214400
  },
  "contents": [
    {
      "name": "heihei7",
      "path": "heihei7",
      "sha": "f0acac8efb3021b0f6a7b13b42d033d86e076a4b",
      "type": "file",
      "size": 5,
      "encoding": "base64",
      "content": "amZrc2o="
    }
  ]
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
|» commit|object|true|none||none|
|»» sha|string|true|none|提交标识|none|
|»» author|object|true|none|作者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» committer|object|true|none|提交者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» commit_message|string|true|none|提交信息|none|
|»» parent_shas|[string]|true|none|提交父节点标识|none|
|»» authored_time|integer|true|none|编码时间|none|
|»» commited_time|integer|true|none|提交时间|none|
|» contents|[object]|true|none||none|
|»» name|string|false|none|文件名称|none|
|»» path|string|false|none|文件路径|none|
|»» sha|string|false|none|文件标识|none|
|»» type|string|false|none|文件类型|none|
|»» size|integer|false|none|文件大小|none|
|»» encoding|string|false|none|编码方式|text文本 base64 加密|
|»» content|string|false|none|文件内容|none|

## GET 获取项目代码目录

GET /api/{owner}/{repo}/entries.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|ref|query|string| 否 ||分支名称、标签名称或是提交标识，默认为master分支|

> 返回示例

> 200 Response

```json
{
  "last_commit": {
    "commit": {
      "sha": "88801266695966b11fcd95ba2bcefad56750d1d9",
      "message": "测试提交\n",
      "author": {
        "name": "yystopf",
        "email": "yystopf@163.com",
        "date": "2022-07-19T15:06:40+08:00"
      },
      "committer": {
        "name": "yystopf",
        "email": "yystopf@163.com",
        "date": "2022-07-19T15:06:40+08:00"
      },
      "timestamp": 1658214400,
      "time_from_now": "1年前"
    },
    "author": {
      "id": "110",
      "login": "yystopf",
      "name": "咸蛋黄土豆丝xxx",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    "committer": {
      "id": "110",
      "login": "yystopf",
      "name": "咸蛋黄土豆丝xxx",
      "type": "User",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    }
  },
  "commits_count": 8,
  "zip_url": "http://localhost:3000/api/yystopf/ceshi_doc/archive/main.zip",
  "tar_url": "http://localhost:3000/api/yystopf/ceshi_doc/archive/main.tar.gz",
  "entries": [
    {
      "name": "heihei7",
      "path": "heihei7",
      "sha": "f0acac8efb3021b0f6a7b13b42d033d86e076a4b",
      "type": "file",
      "submodule_git_url": null,
      "size": 5,
      "is_readme_file": null,
      "content": null,
      "target": null,
      "commit": {
        "message": "测试提交\n",
        "sha": "88801266695966b11fcd95ba2bcefad56750d1d9",
        "created_at": "2022-07-19 15:06",
        "time_from_now": "1年前",
        "created_at_unix": 1658214400
      }
    },
    {
      "name": "README.md",
      "path": "README.md",
      "sha": "91d1da709ed2610e8bcd7b2f6c988f4b8888e515",
      "type": "file",
      "submodule_git_url": null,
      "size": 13,
      "is_readme_file": true,
      "content": null,
      "target": null,
      "commit": {
        "message": "Initial commit\n",
        "sha": "46556cacf0c652e0a5c52ae2b64fe103ddc5d46e",
        "created_at": "2023-11-28 11:47",
        "time_from_now": "3小时前",
        "created_at_unix": 1701143230
      }
    }
  ]
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
|» last_commit|object|true|none|最后一次commit|none|
|»» commit|object|true|none||none|
|»»» sha|string|true|none|commit标识|none|
|»»» message|string|true|none|commit信息|none|
|»»» author|object|true|none|commit作者|none|
|»»»» name|string|true|none||none|
|»»»» email|string|true|none||none|
|»»»» date|string|true|none||none|
|»»» committer|object|true|none|commit提交者|none|
|»»»» name|string|true|none||none|
|»»»» email|string|true|none||none|
|»»»» date|string|true|none||none|
|»»» timestamp|integer|true|none|提交时间戳|none|
|»»» time_from_now|string|true|none|提交时间（距离现在）|none|
|»» author|object|true|none||none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» committer|object|true|none||none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|» commits_count|integer|true|none|commit总数|none|
|» zip_url|string|true|none|zip下载地址|none|
|» tar_url|string|true|none|tar下载地址|none|
|» entries|[object]|true|none||none|
|»» name|string|true|none|文件名称|none|
|»» path|string|true|none|文件路径|none|
|»» sha|string|true|none|文件标识|none|
|»» type|string|true|none|文件类型|file 文件 dir 文件夹 submodule 子模块|
|»» submodule_git_url|null|true|none|子模块git地址|none|
|»» size|integer|true|none|文件大小|none|
|»» is_readme_file|boolean¦null|true|none|是否为readme文件|none|
|»» content|null|true|none|文件内容|none|
|»» target|null|true|none|文件标签|none|
|»» commit|object|true|none|最近一次提交记录|none|
|»»» message|string|true|none|提交信息|none|
|»»» sha|string|true|none|提交标识|none|
|»»» created_at|string|true|none|提交创建时间|none|
|»»» time_from_now|string|true|none|提交时间（距离现在）|none|
|»»» created_at_unix|integer|true|none|提交时间戳|none|

## GET 获取项目代码子目录或者文件

GET /api/{owner}/{repo}/sub_entries.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|filepath|query|string| 是 ||文件夹、文件的相对路径|
|ref|query|string| 否 ||分支名称、标签名称或是提交标识，默认为master分支|

> 返回示例

> 200 Response

```json
{
  "entries": {
    "name": "1",
    "sha": "56a6051ca2b02b04ef92d5150c9ef600403cb1de",
    "path": "1",
    "type": "file",
    "submodule_git_url": null,
    "size": 1,
    "content": "1",
    "target": null,
    "download_url": "http://localhost:3000/api/yystopf/ceshi_doc/raw/1?ref=main",
    "direct_download": false,
    "image_type": false,
    "is_readme_file": null,
    "commit": {
      "message": "Add 1\n",
      "sha": "9d8fc7952c7646feb15a2995bf4b58f77edd6953",
      "created_at": "2023-11-28 11:47",
      "time_from_now": "5小时前",
      "created_at_unix": 1701143246
    }
  }
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
|» entries|object|true|none||none|
|»» name|string|true|none|文件夹或文件名称|none|
|»» sha|string|true|none|文件夹或文件标识|none|
|»» path|string|true|none|文件夹或文件相对路径|none|
|»» type|string|true|none|文件类型|file 文件 dir 文件夹 submodule 子模块|
|»» submodule_git_url|null|true|none|子模块git地址|none|
|»» size|integer|true|none|文件大小|none|
|»» content|string|true|none|文件内容|none|
|»» target|null|true|none|标签|none|
|»» download_url|string|true|none|文件下载地址|none|
|»» direct_download|boolean|true|none|是否跳转下载|none|
|»» image_type|boolean|true|none|是否为图片|none|
|»» is_readme_file|null|true|none|是否为readme文件|none|
|»» commit|object|true|none||none|
|»»» message|string|true|none|提交信息|none|
|»»» sha|string|true|none|提交标识|none|
|»»» created_at|string|true|none|提交时间|none|
|»»» time_from_now|string|true|none|提交时间（距离现在）|none|
|»»» created_at_unix|integer|true|none|提交时间戳|none|

## GET 获取项目README文件

GET /api/{owner}/{repo}/readme.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|filepath|query|string| 否 ||子目录名称，默认为空|
|ref|query|string| 否 ||分支名称、标签名称或是提交记录id，默认为默认分支|

> 返回示例

> 200 Response

```json
{
  "type": "file",
  "encoding": "base64",
  "size": 13,
  "name": "README.md",
  "path": "README.md",
  "content": "# ceshi_doc\n\n",
  "sha": "91d1da709ed2610e8bcd7b2f6c988f4b8888e515",
  "replace_content": "# ceshi_doc\n\n"
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
|» type|string|true|none|类型|none|
|» encoding|string|true|none|文件加密方式|none|
|» size|integer|true|none|文件大小|none|
|» name|string|true|none|文件名称|none|
|» path|string|true|none|文件路径|none|
|» content|string|true|none|文件内容|none|
|» sha|string|true|none|文件标识|none|
|» replace_content|string|true|none|替换后的文件内容|none|

## GET 获取文件树列表

GET /api/v1/{owner}/{repo}/git/trees/{sha}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|sha|path|string| 是 ||分支名称、标签名称或是提交记录id|
|recursive|query|string| 否 ||是否显示目录|
|page|query|string| 否 ||页码|
|limit|query|string| 否 ||每页个数|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "sha": "88801266695966b11fcd95ba2bcefad56750d1d9",
  "entries": [
    {
      "name": "README.md",
      "mode": "100644",
      "type": "file",
      "size": 13,
      "sha": "91d1da709ed2610e8bcd7b2f6c988f4b8888e515"
    },
    {
      "name": "heihei7",
      "mode": "100644",
      "type": "file",
      "size": 5,
      "sha": "f0acac8efb3021b0f6a7b13b42d033d86e076a4b"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» sha|string|true|none|标识|none|
|» entries|[object]|true|none||none|
|»» name|string|true|none|文件树名称|none|
|»» mode|string|true|none|文件树权限|none|
|»» type|string|true|none|文件树类型|file：文件，dir: 文件夹|
|»» size|integer|true|none|文件树大小|none|
|»» sha|string|true|none|文件树commit标识|none|

## GET 获取项目blobs内容

GET /api/v1/{owner}/{repo}/git/blobs/{sha}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|sha|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "sha": "ef44144e8f53fd19d38c8a7d1400313db73881e4",
  "size": 198,
  "encoding": "base64",
  "content": "dHJlZSA4NmYzZjhkMzg3OWQ2MGJjOWNkZGZkNDliNTcxNjQ4YTBlYWIzOWRmCnBhcmVudCA2M2Y1NTcyZTczM2I4ZjRlMDZjNGZhZGVlZjc3Mzk3MWI1NGFhYjU0CmF1dGhvciB2aXJ1cyA8eXlzdG9wZkAxNjMuY29tPiAxNTQ1OTcxNzI4ICswODAwCmNvbW1pdHRlciB2aXJ1cyA8eXlzdG9wZkAxNjMuY29tPiAxNTQ1OTcxNzI4ICswODAwCgpmaXgK"
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
|» sha|string|true|none||none|
|» size|integer|true|none||none|
|» encoding|string|true|none||none|
|» content|string|true|none||none|

## GET 获取项目提交列表

GET /api/v1/{owner}/{repo}/commits.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|sha|query|string| 否 ||分支名、标签名或Commit标识|
|page|query|string| 否 ||页码|
|limit|query|string| 否 ||每页个数|

> 返回示例

> 200 Response

```json
{
  "total_count": 8,
  "commits": [
    {
      "sha": "88801266695966b11fcd95ba2bcefad56750d1d9",
      "author": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "committer": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "commit_message": "测试提交\n",
      "parent_shas": [
        "a69e40efe2d03dbea650c659540b7a0fd87f0c6b"
      ],
      "files": [
        "heihei7"
      ],
      "commit_date": null,
      "commit_time": 1658214400,
      "branch": null
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» commits|[object]|true|none||none|
|»» sha|string|false|none|标识|none|
|»» author|object|false|none|作者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» committer|object|false|none|提交者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» commit_message|string|false|none|提交信息|none|
|»» parent_shas|[string]|false|none|父节点标识|none|
|»» files|[string]|false|none|更改的文件|none|
|»» commit_date|null|false|none|提交日期|none|
|»» commit_time|integer|false|none|提交时间戳|none|
|»» branch|null|false|none|分支|none|

## GET 获取单个提交的变更文件列表

GET /api/v1/{owner}/{repo}/commits/{sha}/files.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|sha|path|string| 是 ||none|
|page|query|string| 否 ||分页页码，当filepath存在时无效|
|limit|query|string| 否 ||分页个数，当filepath存在时无效|
|filepath|query|string| 否 ||文件路径|

> 返回示例

> 200 Response

```json
{
  "file_nums": 78,
  "total_addition": 5360,
  "total_deletion": 0,
  "files": [
    {
      "filename": "packages/max/bin/inula.js",
      "old_name": "packages/max/bin/inula.js",
      "index": 5,
      "type": 1,
      "is_bin": false,
      "is_created": true,
      "is_deleted": false,
      "is_lfs_file": false,
      "is_renamed": false,
      "is_submodule": false,
      "additions": 14,
      "deletions": 0,
      "changes": 14,
      "sha": "68ad2d3e9c4d72b15e825e8f0859596e1e0be9f8"
    }
  ]
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
|» file_nums|integer|true|none||none|
|» total_addition|integer|true|none||none|
|» total_deletion|integer|true|none||none|
|» files|[object]|true|none||none|
|»» filename|string|false|none||none|
|»» old_name|string|false|none||none|
|»» index|integer|false|none||none|
|»» type|integer|false|none||none|
|»» is_bin|boolean|false|none||none|
|»» is_created|boolean|false|none||none|
|»» is_deleted|boolean|false|none||none|
|»» is_lfs_file|boolean|false|none||none|
|»» is_renamed|boolean|false|none||none|
|»» is_submodule|boolean|false|none||none|
|»» additions|integer|false|none||none|
|»» deletions|integer|false|none||none|
|»» changes|integer|false|none||none|
|»» sha|string|false|none||none|

## GET 获取单个提交的diff信息

GET /api/v1/{owner}/{repo}/commits/{sha}/diff.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|sha|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "file_nums": 1,
  "total_addition": 1,
  "total_deletion": 0,
  "files": [
    {
      "name": "heihei7",
      "oldname": "heihei7",
      "addition": 1,
      "deletion": 0,
      "type": 1,
      "is_created": true,
      "is_deleted": false,
      "is_bin": false,
      "is_lfs_file": false,
      "is_renamed": false,
      "is_ambiguous": false,
      "is_submodule": false,
      "diff": null,
      "sections": [
        {
          "file_name": "heihei7",
          "name": "",
          "lines": [
            {
              "left_index": 0,
              "right_index": 0,
              "match": 0,
              "type": 4,
              "content": "@@ -0,0 +1 @@",
              "section_path": "heihei7",
              "section_last_left_index": 0,
              "section_last_right_index": 0,
              "section_left_index": 0,
              "section_right_index": 1,
              "section_left_hunk_size": 0,
              "section_right_hunk_size": 0
            },
            {
              "left_index": 0,
              "right_index": 1,
              "match": -1,
              "type": 2,
              "content": "+jfksj"
            }
          ]
        }
      ],
      "is_incomplete": false,
      "is_incomplete_line_too_long": false,
      "is_protected": false
    }
  ]
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
|» file_nums|integer|true|none|文件数量|none|
|» total_addition|integer|true|none|新增行数|none|
|» total_deletion|integer|true|none|删除行数|none|
|» files|[object]|true|none||none|
|»» name|string|false|none|文件名称|none|
|»» oldname|string|false|none|文件修改前名称|none|
|»» addition|integer|false|none|文件新增行数|none|
|»» deletion|integer|false|none|文件删除行数|none|
|»» type|integer|false|none|文件类型|1: 新增 2: 更改 3: 删除 4: 重命名 5: 复制|
|»» is_created|boolean|false|none|是否为新建文件|none|
|»» is_deleted|boolean|false|none|是否为删除文件|none|
|»» is_bin|boolean|false|none|是否为二进制文件|none|
|»» is_lfs_file|boolean|false|none|是否为lfs文件|none|
|»» is_renamed|boolean|false|none|是否重命名|none|
|»» is_ambiguous|boolean|false|none||none|
|»» is_submodule|boolean|false|none|是否为子模块|none|
|»» diff|null|false|none||none|
|»» sections|[object]|false|none||none|
|»»» file_name|string|false|none|文件名称|none|
|»»» name|string|false|none||none|
|»»» lines|[object]|false|none||none|
|»»»» left_index|integer|true|none|文件变动之前所在行数|none|
|»»»» right_index|integer|true|none|文件变动之后所在行数|none|
|»»»» match|integer|true|none||none|
|»»»» type|integer|true|none|文件变更类型|none|
|»»»» content|string|true|none|文件变更内容|none|
|»»»» section_path|string|false|none|文件路径|none|
|»»»» section_last_left_index|integer|false|none||none|
|»»»» section_last_right_index|integer|false|none||none|
|»»»» section_left_index|integer|false|none|文件变更之前所在行数|none|
|»»»» section_right_index|integer|false|none|文件变更之后所在行数(即：页面编辑器开始显示的行数)|none|
|»»»» section_left_hunk_size|integer|false|none|文件变更之前的行数|none|
|»»»» section_right_hunk_size|integer|false|none|文件变更之后的行数(及当前页面编辑器显示的总行数)|none|
|»» is_incomplete|boolean|false|none|是否不完整|none|
|»» is_incomplete_line_too_long|boolean|false|none|文件是否不完整是因为太长了|none|
|»» is_protected|boolean|false|none|文件是否被保护|none|

## GET 获取单个文件的blame信息

GET /api/v1/{owner}/{repo}/blame.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|sha|query|string| 是 ||分支、标签或提交标识|
|filepath|query|string| 是 ||文件路径|

> 返回示例

> 200 Response

```json
{
  "file_size": 1,
  "file_name": "1",
  "num_lines": 1,
  "blame_parts": [
    {
      "commit": {
        "sha": "9d8fc7952c7646feb15a2995bf4b58f77edd6953",
        "author": {
          "id": "110",
          "login": "yystopf",
          "name": "咸蛋黄土豆丝xxx",
          "type": "User",
          "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
        },
        "committer": {
          "id": "110",
          "login": "yystopf",
          "name": "咸蛋黄土豆丝xxx",
          "type": "User",
          "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
        },
        "commit_message": "Add 1\n",
        "authored_time": 1701143246,
        "committed_time": 1701143246,
        "created_time": 1701143246
      },
      "previous_number": 1,
      "current_number": 1,
      "effect_line": 1,
      "lines": [
        "1"
      ]
    }
  ]
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
|» file_size|integer|true|none|文件大小|none|
|» file_name|string|true|none|文件名称|none|
|» num_lines|integer|true|none|文件总行数|none|
|» blame_parts|[object]|true|none||none|
|»» commit|object|false|none|提交|none|
|»»» sha|string|true|none|标识|none|
|»»» author|object|true|none|作者|none|
|»»»» id|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» committer|object|true|none|提交者|none|
|»»»» id|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» type|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» commit_message|string|true|none|提交信息|none|
|»»» authored_time|integer|true|none|编码时间|none|
|»»» committed_time|integer|true|none|提交时间|none|
|»»» created_time|integer|true|none|创建时间|none|
|»» previous_number|integer|false|none|距离当前行有多少行|none|
|»» current_number|integer|false|none|当前行|none|
|»» effect_line|integer|false|none|影响的行数|none|
|»» lines|[string]|false|none|行内容|none|

## GET 获取比较提交的diff信息

GET /api/v1/{owner}/{repo}/compare.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|from|query|string| 是 ||源分支、标签、提交标识|
|to|query|string| 是 ||目标分支、标签、提交标识|

> 返回示例

> 200 Response

```json
{
  "commits_count": 1,
  "last_commit_sha": null,
  "commits": [
    {
      "author": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "committer": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "branch": null,
      "commit_message": "测试提交\n",
      "sha": "88801266695966b11fcd95ba2bcefad56750d1d9",
      "parent_shas": null
    }
  ],
  "diff": {
    "file_nums": 1,
    "total_addition": 1,
    "total_deletion": 0,
    "files": [
      {
        "name": "heihei7",
        "oldname": "heihei7",
        "addition": 1,
        "deletion": 0,
        "type": 1,
        "is_created": true,
        "is_deleted": false,
        "is_bin": false,
        "is_lfs_file": false,
        "is_renamed": false,
        "is_ambiguous": false,
        "is_submodule": false,
        "diff": null,
        "sections": [
          {
            "file_name": "heihei7",
            "name": "",
            "lines": [
              {
                "left_index": 0,
                "right_index": 0,
                "match": 0,
                "type": 4,
                "content": "@@ -0,0 +1 @@",
                "section_path": "heihei7",
                "section_last_left_index": 0,
                "section_last_right_index": 0,
                "section_left_index": 0,
                "section_right_index": 1,
                "section_left_hunk_size": 0,
                "section_right_hunk_size": 0
              },
              {
                "left_index": 0,
                "right_index": 1,
                "match": -1,
                "type": 2,
                "content": "+jfksj"
              }
            ]
          }
        ],
        "is_incomplete": false,
        "is_incomplete_line_too_long": false,
        "is_protected": false
      }
    ]
  }
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
|» commits_count|integer|true|none|提交数量|none|
|» last_commit_sha|null|true|none||none|
|» commits|[object]|true|none||none|
|»» author|object|false|none|提交作者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» committer|object|false|none|提交者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» branch|null|false|none|提交分支|none|
|»» commit_message|string|false|none|提交信息|none|
|»» sha|string|false|none|提交标识|none|
|»» parent_shas|null|false|none|提交信息标识|none|
|» diff|object|true|none||none|
|»» file_nums|integer|true|none|文件数量|none|
|»» total_addition|integer|true|none|新增行数|none|
|»» total_deletion|integer|true|none|删除行数|none|
|»» files|[object]|true|none||none|
|»»» name|string|false|none|文件名称|none|
|»»» oldname|string|false|none|文件名称|none|
|»»» addition|integer|false|none|文件新增行数|none|
|»»» deletion|integer|false|none|文件删除行数|none|
|»»» type|integer|false|none|文件类型|1: 新增 2: 更改 3: 删除 4: 重命名 5: 复制|
|»»» is_created|boolean|false|none|是否为新建文件|none|
|»»» is_deleted|boolean|false|none|是否为删除文件|none|
|»»» is_bin|boolean|false|none|是否为二进制文件|none|
|»»» is_lfs_file|boolean|false|none|是否为LFS文件|none|
|»»» is_renamed|boolean|false|none|是否重命名|none|
|»»» is_ambiguous|boolean|false|none||none|
|»»» is_submodule|boolean|false|none|是否为子模块|none|
|»»» diff|null|false|none||none|
|»»» sections|[object]|false|none||none|
|»»»» file_name|string|false|none|文件名称|none|
|»»»» name|string|false|none||none|
|»»»» lines|[object]|false|none||none|
|»»»»» left_index|integer|true|none|文件变动之前所在行数|none|
|»»»»» right_index|integer|true|none|文件变动之后所在行数|none|
|»»»»» match|integer|true|none||none|
|»»»»» type|integer|true|none|文件变更类型|none|
|»»»»» content|string|true|none|文件变更内容|none|
|»»»»» section_path|string|false|none|文件路径|none|
|»»»»» section_last_left_index|integer|false|none||none|
|»»»»» section_last_right_index|integer|false|none||none|
|»»»»» section_left_index|integer|false|none|文件变更之前所在行数|none|
|»»»»» section_right_index|integer|false|none|文件变更之后所在行数(即：页面编辑器开始显示的行数)|none|
|»»»»» section_left_hunk_size|integer|false|none|文件变更之前的行数|none|
|»»»»» section_right_hunk_size|integer|false|none|文件变更之后的行数(及当前页面编辑器显示的总行数)|none|
|»»» is_incomplete|boolean|false|none|是否不完整|none|
|»»» is_incomplete_line_too_long|boolean|false|none|文件是否不完整是因为太长了|none|
|»»» is_protected|boolean|false|none|文件是否被保护|none|

## POST 创建一个文件

POST /api/{owner}/{repo}/create_file.json

> Body 请求参数

```json
{
  "filepath": "测试上传",
  "base64_filepath": "5rWL6K+V5LiK5Lyg",
  "branch": "master_1",
  "content": "5rWL6K+V5LiK5Lyg",
  "message": "Add 测试上传"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» base64_filepath|body|string| 是 | 文件路径（base64加密后）|none|
|» filepath|body|string| 是 | 文件路径|none|
|» branch|body|string| 是 | 分支名称|none|
|» new_branch|body|string| 否 | 新分支名称|none|
|» content|body|string| 是 | 文件内容（Base64加密后）|none|
|» message|body|string| 是 | 提交信息|none|

> 返回示例

> 200 Response

```json
{
  "name": "pariatur veniam anim",
  "sha": "12024760d4df3170f9699d9084211d7975ec4adc",
  "size": 12,
  "content": "5rWL6K+V5LiK5Lyg",
  "encoding": "base64",
  "pr_id": null,
  "commit": {
    "message": "reprehenderit\n",
    "author": {
      "name": "yystopf",
      "email": "yystopf@163.com",
      "date": "2023-12-21T06:44:11Z"
    },
    "committer": {
      "name": "yystopf",
      "email": "yystopf@163.com",
      "date": "2023-12-21T06:44:11Z"
    }
  }
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
|» name|string|true|none|文件名称|none|
|» sha|string|true|none|文件提交标识|none|
|» size|integer|true|none|文件大小|none|
|» content|string|true|none|文件内容|none|
|» encoding|string|true|none|文件加密方式|none|
|» pr_id|null|true|none|合并请求id|none|
|» commit|object|true|none|提交|none|
|»» message|string|true|none|提交信息|none|
|»» author|object|true|none|提交作者|none|
|»»» name|string|true|none||none|
|»»» email|string|true|none||none|
|»»» date|string|true|none||none|
|»» committer|object|true|none|提交者|none|
|»»» name|string|true|none||none|
|»»» email|string|true|none||none|
|»»» date|string|true|none||none|

## DELETE 删除一个文件

DELETE /api/{owner}/{repo}/delete_file.json

> Body 请求参数

```json
{
  "filepath": "测试上传",
  "base64_filepath": "5rWL6K+V5LiK5Lyg",
  "branch": "master_1",
  "content": "5rWL6K+V5LiK5Lyg",
  "message": "Add 测试上传"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» base64_filepath|body|string| 否 | 文件路径（base64加密后）|none|
|» filepath|body|string| 是 | 文件路径|none|
|» branch|body|string| 是 | 分支名称|none|
|» sha|body|string| 是 | 文件提交标识|none|

> 返回示例

> 200 Response

```json
{
  "status": 1,
  "message": "文件删除成功"
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
|» message|string|true|none||none|

## PUT 更新一个文件

PUT /api/{owner}/{repo}/update_file.json

> Body 请求参数

```json
{
  "filepath": "！@#%%…………&*（*0）",
  "base64_filepath": "77yBQCMlJeKApuKApuKApuKApiYq77yIKjDvvIk=",
  "branch": "tempor_proidentugiat",
  "content": "内容11",
  "sha": "f0b922501594fc0bffb1635244553fae2dc640d3",
  "message": "Update ！@#%%…………&*（*0）"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» filepath|body|string| 是 | 文件路径|none|
|» base64_filepath|body|string| 否 | 文件路径（base64）|none|
|» branch|body|string| 是 | 分支名称|none|
|» new_branch|body|string| 否 | 新分支名称|none|
|» content|body|string| 是 | 文件内容|none|
|» sha|body|string| 是 | 文件标识|none|
|» message|body|string| 是 | 提交信息|none|

> 返回示例

> 200 Response

```json
{
  "status": 1,
  "message": "更新成功"
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
|» message|string|true|none||none|

## POST 替换一个文件

POST /api/{owner}/{repo}/replace_file.json

> Body 请求参数

```json
{
  "filepath": "！@#%%…………&*（*0）",
  "base64_filepath": "77yBQCMlJeKApuKApuKApuKApiYq77yIKjDvvIk=",
  "branch": "tempor_proidentugiat",
  "content": "内容11",
  "sha": "f0b922501594fc0bffb1635244553fae2dc640d3",
  "message": "Update ！@#%%…………&*（*0）"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» filepath|body|string| 是 | 文件路径|none|
|» base64_filepath|body|string| 是 | 文件路径（base64）|none|
|» branch|body|string| 是 | 分支名称|none|
|» content|body|string| 是 | 文件内容|none|
|» message|body|string| 是 | 提交信息|none|
|» delete_file|body|object| 是 ||none|
|»» filepath|body|string| 是 | 删除文件路径|none|
|»» base64_filepath|body|string| 是 | 删除文件路径（base64）|none|
|»» branch|body|string| 是 | 分支名称|none|
|»» sha|body|string| 是 | 提交标识|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "替换成功"
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
|» message|string|true|none||none|

## GET 获取项目关注列表

GET /api/{owner}/{repo}/watchers.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|start_at|query|integer| 否 ||开始时间戳|
|end_at|query|integer| 否 ||结束时间戳|

> 返回示例

> 200 Response

```json
{
  "count": 1,
  "users": [
    {
      "format_time": "2024-11-06",
      "name": "yystopfceshi",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/Y/197_115_70/120.png",
      "is_current_user": false,
      "is_watch": false
    }
  ]
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
|» count|integer|true|none|总数|none|
|» users|[object]|true|none||none|
|»» format_time|string|false|none|关注时间|none|
|»» name|string|false|none|昵称|none|
|»» login|string|false|none|标识|none|
|»» image_url|string|false|none|头像|none|
|»» is_current_user|boolean|false|none|是否为当前用户|none|
|»» is_watch|boolean|false|none|是否关注该用户|none|

## GET 获取项目点赞列表

GET /api/{owner}/{repo}/stargazers.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|start_at|query|integer| 否 ||开始时间戳|
|end_at|query|integer| 否 ||结束时间戳|

> 返回示例

> 200 Response

```json
{
  "count": 1,
  "users": [
    {
      "format_time": "2024-11-06",
      "name": "yystopfceshi",
      "login": "yystopfceshi",
      "image_url": "system/lets/letter_avatars/2/Y/197_115_70/120.png",
      "is_current_user": false,
      "is_watch": false
    }
  ]
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
|» count|integer|true|none|总数|none|
|» users|[object]|true|none||none|
|»» format_time|string|false|none|点赞时间|none|
|»» name|string|false|none|昵称|none|
|»» login|string|false|none|标识|none|
|»» image_url|string|false|none|头像|none|
|»» is_current_user|boolean|false|none|是否为当前用户|none|
|»» is_watch|boolean|false|none|是否关注该用户|none|

# 疑修

## GET 疑修状态列表

GET /api/v1/{owner}/{repo}/issue_statues.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|page|query|integer| 否 ||页码|
|limit|query|integer| 否 ||分页个数|

> 返回示例

> 200 Response

```json
{
  "total_count": 5,
  "statues": [
    {
      "id": 1,
      "name": "新增"
    },
    {
      "id": 2,
      "name": "正在解决"
    },
    {
      "id": 3,
      "name": "已解决"
    },
    {
      "id": 5,
      "name": "关闭"
    },
    {
      "id": 6,
      "name": "拒绝"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» statues|[object]|true|none||none|
|»» id|integer|true|none|疑修状态ID|none|
|»» name|string|true|none|疑修状态名称|none|

## GET 疑修发布人列表

GET /api/v1/{owner}/{repo}/issue_authors.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||搜索关键词|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "authors": [
    {
      "id": 84727,
      "type": "User",
      "name": "yystopf.df",
      "login": "yystopf123",
      "image_url": "images/avatars/User/84727?t=1650252387"
    },
    {
      "id": 36480,
      "type": "User",
      "name": "段甲生",
      "login": "pxicb3wke",
      "image_url": "images/avatars/User/36480?t=1672730523"
    }
  ]
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
|» total_count|integer|true|none||none|
|» authors|[object]|true|none||none|
|»» id|integer|true|none|用户ID|none|
|»» type|string|true|none|用户类型|none|
|»» name|string|true|none|用户名称|none|
|»» login|string|true|none|用户标识|none|
|»» image_url|string|true|none|用户头像|none|

## GET 疑修负责人列表

GET /api/v1/{owner}/{repo}/issue_assigners.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||搜索关键词|

> 返回示例

> 200 Response

```json
{
  "total_count": 3,
  "assigners": [
    {
      "id": 84727,
      "type": "User",
      "name": "yystopf.df",
      "login": "yystopf123",
      "image_url": "images/avatars/User/84727?t=1650252387"
    },
    {
      "id": 84961,
      "type": "User",
      "name": "postwoman釹轻柔",
      "login": "postwoman",
      "image_url": "images/avatars/User/84961?t=1641441279"
    },
    {
      "id": 85605,
      "type": "User",
      "name": "Floraachy",
      "login": "p10789563",
      "image_url": "system/lets/letter_avatars/2/P/158_138_26/120.png"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» assigners|[object]|true|none|负责人数组|none|
|»» id|integer|true|none|用户ID|none|
|»» type|string|true|none|用户类型|none|
|»» name|string|true|none|用户名称|none|
|»» login|string|true|none|用户标识|none|
|»» image_url|string|true|none|用户头像|none|

## GET 里程碑列表

GET /api/v1/{owner}/{repo}/milestones.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||搜索关键词|
|only_name|query|string| 否 ||只返回名称及id|
|page|query|integer| 否 ||none|
|limit|query|string| 否 ||none|
|category|query|string| 否 ||closed 关闭的 opening 开启中|
|sort_by|query|string| 否 ||排序字段 created_on 创建时间 updated_on 更新时间 effective_date 到期日期  issues_count 任务数量 percent 完成度|
|sort_direction|query|string| 否 ||排序顺序 desc 倒序 asc 正序|

> 返回示例

> 200 Response

```json
{
  "closed_milestone_count": 0,
  "opening_milestone_count": 3,
  "total_count": 3,
  "milestones": [
    {
      "id": 5,
      "name": "是林得七没",
      "description": "拉道计道等本照响层方支被原条目。便史劳回又观民放全子将八近受引子便。将子意布头天单土养低新实议反书。就斯众县成县厂研与己场权人。业车空且热般设转么体府义满素质。验者建指四集传风片和证太。",
      "effective_date": "1987-11-03",
      "status": "open",
      "issues_count": 0,
      "opened_issues_count": 0,
      "close_issues_count": 0,
      "percent": 0,
      "created_at": "2023-02-20 11:23",
      "updated_on": "2023-02-20 11:23"
    },
    {
      "id": 4,
      "name": "加花果马划八",
      "description": "队还己光满道及去回派斗入利眼。劳劳百及联计存本至统专大维石志。之光加白本知等划争件易要他高然油也。同有力观适产候影带计表世达程资斯华感。美度的没时制百气安如都争。文术样验号严基面精即报济有先除入。界之许实成史我点可件千许属。",
      "effective_date": "1989-09-03",
      "status": "open",
      "issues_count": 1,
      "opened_issues_count": 1,
      "close_issues_count": 0,
      "percent": 0,
      "created_at": "2023-02-16 18:36",
      "updated_on": "2023-02-20 11:38"
    },
    {
      "id": 3,
      "name": "1",
      "description": "1",
      "effective_date": "2023-02-24",
      "status": "open",
      "issues_count": 36,
      "opened_issues_count": 36,
      "close_issues_count": 0,
      "percent": 0,
      "created_at": "2023-02-07 17:54",
      "updated_on": "2023-02-07 17:54"
    }
  ]
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
|» closed_milestone_count|integer|true|none|关闭的里程碑总数|none|
|» opening_milestone_count|integer|true|none|开启的里程碑总数|none|
|» total_count|integer|true|none|分页总数|none|
|» milestones|[object]|true|none||none|
|»» id|integer|true|none|里程碑ID|none|
|»» name|string|true|none|里程碑名称|none|
|»» description|string|true|none|里程碑描述|none|
|»» effective_date|string|true|none|截止日期|none|
|»» status|string|true|none|里程碑状态|closed 关闭 opening 开启的|
|»» issues_count|integer|true|none|疑修数量|none|
|»» opened_issues_count|integer|true|none|开启中疑修数量|none|
|»» close_issues_count|integer|true|none|关闭的疑修数量|none|
|»» percent|integer|true|none|完成度|none|
|»» created_at|string|true|none|创建时间|none|
|»» updated_on|string|true|none|更新时间|none|

## POST 创建一个里程碑

POST /api/v1/{owner}/{repo}/milestones.json

> Body 请求参数

```json
{
  "name": "string",
  "description": "string",
  "effective_date": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|debug|query|string| 否 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 里程碑名称|none|
|» description|body|string| 是 | 里程碑描述|none|
|» effective_date|body|string| 是 | 里程碑截止日期|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 获取一个里程碑以及疑修（合并请求TODO）

GET /api/v1/{owner}/{repo}/milestones/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|
|category|query|string| 否 ||疑修类型，all 全部 opened 开启中 closed 已关闭|
|author_id|query|string| 否 ||发布人ID|
|assigner_id|query|string| 否 ||负责人ID|
|issue_tag_ids|query|string| 否 ||标记ID，支持多个ID用,隔开|
|page|query|integer| 否 ||none|
|limit|query|string| 否 ||none|
|sort_by|query|string| 否 ||排序字段  issues.created_on 创建时间 issues.updated_on 更新时间 issue_priorities.position 优先级|
|sort_direction|query|string| 否 ||排序类型 desc 倒序 asc 正序|

> 返回示例

> 200 Response

```json
{
  "milestone": {
    "id": 4,
    "name": "2",
    "description": "2",
    "effective_date": "2023-03-04",
    "issues_count": 1,
    "closed_issues_count": 0,
    "percent": 0,
    "status": "open",
    "created_at": "2023-02-16 18:36",
    "updated_on": "2023-02-16 18:36"
  },
  "total_issues_count": 1,
  "closed_issues_count": 0,
  "opened_issues_count": 1,
  "issues": [
    {
      "id": 62,
      "subject": "sedhahahaha213qjiqiq321312jakjf3123121231231",
      "project_issues_index": 32,
      "created_at": "2023-02-15 09:53",
      "updated_at": "2023-02-18 23:00",
      "tags": [
        {
          "id": 3,
          "name": "1",
          "color": "#F17013"
        },
        {
          "id": 4,
          "name": "31231",
          "color": "#F17013"
        }
      ],
      "status_name": "已解决",
      "priority_name": "高",
      "milestone_name": "2",
      "author": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "assigners": [
        {
          "id": 110,
          "type": "User",
          "name": "咸蛋黄土豆丝",
          "login": "yystopf",
          "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
        },
        {
          "id": 115,
          "type": "User",
          "name": "嘿嘿嘿",
          "login": "yystopf123",
          "image_url": "system/lets/letter_avatars/2/Y/172_132_85/120.png"
        }
      ],
      "comment_journals_count": 8
    }
  ]
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
|» milestone|object|true|none|里程碑|none|
|»» id|integer|true|none|ID|none|
|»» name|string|true|none|名称|none|
|»» description|string|true|none|描述|none|
|»» effective_date|string|true|none|截止日期|none|
|»» issues_count|integer|true|none|issue数量|none|
|»» closed_issues_count|integer|true|none|关闭的issue数量|none|
|»» percent|integer|true|none|完成度|none|
|»» status|string|true|none|状态|open 开启 closed关闭|
|»» created_at|string|true|none|创建时间|none|
|»» updated_on|string|true|none|更新时间|none|
|» total_issues_count|integer|true|none|查询列表总数|none|
|» closed_issues_count|integer|true|none|查询列表关闭的疑修数|none|
|» opened_issues_count|integer|true|none|查询列表开启中的疑修数|none|
|» issues|[object]|true|none||none|
|»» id|integer|false|none||none|
|»» subject|string|false|none|标题|none|
|»» project_issues_index|integer|false|none|序号|none|
|»» created_at|string|false|none|创建时间|none|
|»» updated_at|string|false|none|更新时间|none|
|»» tags|[object]|false|none|标记|none|
|»»» id|integer|true|none|ID|none|
|»»» name|string|true|none|名称|none|
|»»» color|string|true|none|色值|none|
|»» status_name|string|false|none|状态|none|
|»» priority_name|string|false|none|优先级|none|
|»» milestone_name|string|false|none|里程碑|none|
|»» author|object|false|none|发布人|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» assigners|[object]|false|none|负责人|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» comment_journals_count|integer|false|none|评论数量|none|

## PATCH 更新一个里程碑

PATCH /api/v1/{owner}/{repo}/milestones/{id}.json

> Body 请求参数

```json
{
  "name": "string",
  "description": "string",
  "effective_date": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 名称|none|
|» description|body|string| 是 | 描述|none|
|» effective_date|body|string| 是 | 截止日期|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## DELETE 删除一个里程碑

DELETE /api/v1/{owner}/{repo}/milestones/{id}.json

> Body 请求参数

```json
{
  "name": "string",
  "description": "string",
  "effective_date": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 名称|none|
|» description|body|string| 是 | 描述|none|
|» effective_date|body|string| 是 | 截止日期|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 更新一个里程碑状态

POST /api/{owner}/{repo}/milestones/{id}/update_status.json

> Body 请求参数

```json
{
  "status": "closed"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» status|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "操作成功"
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
|» message|string|true|none||none|

## GET 疑修优先级列表

GET /api/v1/{owner}/{repo}/issue_priorities.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||搜索关键词|

> 返回示例

> 200 Response

```json
{
  "total_count": 5,
  "priorities": [
    {
      "id": 1,
      "name": "低"
    },
    {
      "id": 2,
      "name": "正常"
    },
    {
      "id": 3,
      "name": "高"
    },
    {
      "id": 4,
      "name": "紧急"
    },
    {
      "id": 5,
      "name": "立刻"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» priorities|[object]|true|none||none|
|»» id|integer|true|none|优先级ID|none|
|»» name|string|true|none|优先级名称|none|

## GET 项目标记列表

GET /api/v1/{owner}/{repo}/issue_tags.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||搜索关键词|
|only_name|query|string| 否 ||只返回名称及id|
|order_by|query|string| 否 ||排序字段 updated_on 更新时间 created_on 创建时间 issues_count 疑修个数 |
|order_direction|query|string| 否 ||排序类型 desc 倒序 asc 正序|

> 返回示例

```json
{
  "total_count": 1,
  "issue_tags": [
    {
      "id": 3,
      "name": "1"
    }
  ]
}
```

```json
{
  "total_count": 2,
  "issue_tags": [
    {
      "name": "31231",
      "description": "132312",
      "color": "#F17013",
      "issues_count": 0,
      "project": {
        "type": "mirror",
        "description": null,
        "forked_count": 0,
        "forked_from_project_id": null,
        "identifier": "pig",
        "issues_count": 30,
        "pull_requests_count": 0,
        "invite_code": "2ymoPv",
        "website": null,
        "platform": "forge",
        "name": "pig",
        "open_devops": false,
        "praises_count": 0,
        "is_public": true,
        "status": 1,
        "watchers_count": 0,
        "ignore_id": null,
        "license_id": null,
        "project_category_id": null,
        "project_language_id": 8
      },
      "user": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "created_at": "2023-02-15 11:02",
      "updated_at": "2023-02-15 11:02"
    },
    {
      "name": "1",
      "description": "1",
      "color": "#F17013",
      "issues_count": 24,
      "project": {
        "type": "mirror",
        "description": null,
        "forked_count": 0,
        "forked_from_project_id": null,
        "identifier": "pig",
        "issues_count": 30,
        "pull_requests_count": 0,
        "invite_code": "2ymoPv",
        "website": null,
        "platform": "forge",
        "name": "pig",
        "open_devops": false,
        "praises_count": 0,
        "is_public": true,
        "status": 1,
        "watchers_count": 0,
        "ignore_id": null,
        "license_id": null,
        "project_category_id": null,
        "project_language_id": 8
      },
      "user": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "created_at": "2023-02-08 10:14",
      "updated_at": "2023-02-08 10:14"
    }
  ]
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
|» total_count|integer|true|none||none|
|» issue_tags|[object]|true|none||none|
|»» name|string|true|none||名称|
|»» description|string|true|none||描述|
|»» color|string|true|none||色值|
|»» issues_count|integer|true|none||疑修数量|
|»» pull_requests_count|integer|true|none||合并请求数量|
|»» project|object|true|none||所属项目|
|»»» type|string|true|none||none|
|»»» description|null|true|none||none|
|»»» forked_count|integer|true|none||none|
|»»» forked_from_project_id|null|true|none||none|
|»»» identifier|string|true|none||none|
|»»» issues_count|integer|true|none||none|
|»»» pull_requests_count|integer|true|none||none|
|»»» invite_code|string|true|none||none|
|»»» website|null|true|none||none|
|»»» platform|string|true|none||none|
|»»» name|string|true|none||none|
|»»» open_devops|boolean|true|none||none|
|»»» praises_count|integer|true|none||none|
|»»» is_public|boolean|true|none||none|
|»»» status|integer|true|none||none|
|»»» watchers_count|integer|true|none||none|
|»»» ignore_id|null|true|none||none|
|»»» license_id|null|true|none||none|
|»»» project_category_id|null|true|none||none|
|»»» project_language_id|integer|true|none||none|
|»» user|object|true|none||创建用户|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» created_at|string|true|none||创建时间|
|»» updated_at|string|true|none||更新时间|

## POST 创建一个项目标记

POST /api/v1/{owner}/{repo}/issue_tags.json

> Body 请求参数

```json
{
  "name": "string",
  "description": "string",
  "color": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 标记名称|none|
|» description|body|string| 是 | 标记描述|none|
|» color|body|string| 是 | 标记色值|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## PATCH 更新一个项目标记

PATCH /api/v1/{owner}/{repo}/issue_tags/{id}.json

> Body 请求参数

```json
{
  "name": "string",
  "description": "string",
  "color": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» name|body|string| 是 | 项目标记名称|none|
|» description|body|string| 是 | 项目标记描述|none|
|» color|body|string| 是 | 色值|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## DELETE 删除一个项目标记

DELETE /api/v1/{owner}/{repo}/issue_tags/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 疑修列表

GET /api/v1/{owner}/{repo}/issues.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|category|query|string| 否 ||疑修类型，all 全部 opened 开启中 closed 已关闭|
|participant_category|query|string| 否 ||参与类型，all 全部 aboutme 关于我的 authoredme 我创建的 assignedme 我负责的 atme @我的|
|keyword|query|string| 否 ||搜索关键词|
|author_id|query|integer| 否 ||发布人用户ID|
|milestone_id|query|integer| 否 ||里程碑ID|
|assigner_id|query|integer| 否 ||负责人用户ID|
|status_id|query|integer| 否 ||状态ID|
|sort_by|query|string| 否 ||排序字段，issues.updated_on 更新时间 issues.created_on 创建时间 issue_priorities.position 优先级|
|sort_direction|query|string| 否 ||排序类型，asc 正序 desc 倒序|
|issue_tag_ids|query|string| 否 ||标记ID，支持多个用,隔开|
|page|query|integer| 否 ||none|
|limit|query|integer| 否 ||none|
|debug|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 22,
  "opened_count": 1,
  "closed_count": 0,
  "issues": [
    {
      "id": 37,
      "subject": "occaecat Lorem irure aliqua32131231",
      "project_issues_index": 25,
      "created_at": "2023-02-10 11:08",
      "updated_at": "2023-02-10 11:08",
      "tags": [
        {
          "id": 3,
          "name": "1",
          "color": "#F17013"
        }
      ],
      "status_name": "新增",
      "priority_name": "低",
      "milestone_name": "1",
      "author": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "assigners": [
        {
          "id": 110,
          "type": "User",
          "name": "咸蛋黄土豆丝",
          "login": "yystopf",
          "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
        },
        {
          "id": 115,
          "type": "User",
          "name": "嘿嘿嘿",
          "login": "yystopf123",
          "image_url": "system/lets/letter_avatars/2/Y/172_132_85/120.png"
        }
      ],
      "comment_journals_count": 0
    }
  ]
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
|» total_count|integer|true|none|所有疑修数量|none|
|» opened_count|integer|true|none|开启中疑修数量|none|
|» closed_count|integer|true|none|关闭的疑修数量|none|
|» issues|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» subject|string|true|none|标题|none|
|»» project_issues_index|integer|true|none|序号|none|
|»» blockchain_token_num|integer|true|none|悬赏金额|none|
|»» created_at|string|true|none|创建时间|none|
|»» updated_at|string|true|none|更新时间|none|
|»» tags|[object]|false|none|标记|none|
|»»» id|integer|false|none||none|
|»»» name|string|false|none|名称|none|
|»»» color|string|false|none|色值|none|
|»» status_name|string|true|none|状态|none|
|»» priority_name|string|true|none|优先级|none|
|»» milestone_name|string|true|none|里程碑|none|
|»» author|object|false|none|发布人|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» assigners|[object]|false|none|负责人|none|
|»»» id|integer|false|none||none|
|»»» type|string|false|none||none|
|»»» name|string|false|none||none|
|»»» login|string|false|none||none|
|»»» image_url|string|false|none||none|
|»» comment_journals_count|integer|true|none|评论数量|none|

## POST 创建一个疑修

POST /api/v1/{owner}/{repo}/issues.json

> Body 请求参数

```json
{
  "status_id": 0,
  "priority_id": 0,
  "milestone_id": 0,
  "branch_name": "string",
  "start_date": "string",
  "due_date": "string",
  "subject": "string",
  "description": "string",
  "blockchain_token_num": 0,
  "issue_tag_ids": [
    0
  ],
  "assigner_ids": [
    0
  ],
  "attachment_ids": [
    0
  ],
  "receivers_login": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» status_id|body|integer| 是 ||疑修状态ID|
|» priority_id|body|integer| 是 ||疑修优先级ID|
|» milestone_id|body|integer| 否 ||里程碑ID|
|» branch_name|body|string| 否 ||分支名称|
|» start_date|body|string| 否 ||开始日期|
|» due_date|body|string| 否 ||截止日期|
|» subject|body|string| 是 ||标题|
|» description|body|string| 否 ||描述|
|» blockchain_token_num|body|integer| 否 ||悬赏金额|
|» issue_tag_ids|body|[integer]| 否 ||标记ID数组|
|» assigner_ids|body|[integer]| 否 ||负责人ID数组|
|» attachment_ids|body|[integer]| 否 ||附件ID数组|
|» receivers_login|body|[string]| 否 ||@用户标识数组|

> 返回示例

> 200 Response

```json
{
  "id": 62,
  "subject": "sedhahahaha213qjiqiq321312jakjf3123121231231",
  "project_issues_index": 32,
  "description": "部低九认都音以再际油化少可党半每市。无白表并商器何江华向二写。元先战因打质写起论国子四说省史十。计外们规织花层向结具习把月类义头比海。更感派是元家前识布连市厂号直北达则。千地意亲集式位状起几级号平常作意。",
  "branch_name": "持命看山",
  "start_date": "1988-11-06",
  "due_date": "1970-07-23",
  "created_at": "2023-02-15 09:53",
  "updated_at": "2023-02-20 11:38",
  "tags": [
    {
      "id": 3,
      "name": "1",
      "color": "#F17013"
    }
  ],
  "status": {
    "id": 5,
    "name": "关闭"
  },
  "priority": {
    "id": 3,
    "name": "高"
  },
  "milestone": {
    "id": 4,
    "name": "加花果马划八"
  },
  "author": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "assigners": [
    {
      "id": 110,
      "type": "User",
      "name": "咸蛋黄土豆丝",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    {
      "id": 115,
      "type": "User",
      "name": "嘿嘿嘿",
      "login": "yystopf123",
      "image_url": "system/lets/letter_avatars/2/Y/172_132_85/120.png"
    }
  ],
  "participants": [
    {
      "id": 1,
      "type": "AnonymousUser",
      "name": "游客",
      "login": "",
      "image_url": null
    },
    {
      "id": 110,
      "type": "User",
      "name": "咸蛋黄土豆丝",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    }
  ],
  "comment_journals_count": 8,
  "operate_journals_count": 33,
  "attachments": [
    {
      "id": 13,
      "title": "智能头盔软件需求构思和描述文档 (1).docx",
      "filesize": "14.1 KB",
      "is_pdf": false,
      "url": "/api/attachments/13",
      "created_on": "2023-02-17 16:05",
      "content_type": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      "id": 14,
      "title": "prv_test-master.zip",
      "filesize": "316 字节",
      "is_pdf": false,
      "url": "/api/attachments/14",
      "created_on": "2023-02-17 16:06",
      "content_type": "application/zip"
    }
  ]
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
|» id|integer|true|none||none|
|» subject|string|true|none|标题|none|
|» project_issues_index|integer|true|none|序号|none|
|» description|string|true|none|描述|none|
|» branch_name|string|true|none|分支名称|none|
|» start_date|string|true|none|开始日期|none|
|» due_date|string|true|none|截止日期|none|
|» blockchain_token_num|integer|true|none|悬赏金额|none|
|» created_at|string|true|none|创建时间|none|
|» updated_at|string|true|none|更新时间|none|
|» tags|[object]|true|none|标记|none|
|»» id|integer|false|none||none|
|»» name|string|false|none||none|
|»» color|string|false|none||none|
|» status|object|true|none|疑修状态|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» priority|object|true|none|疑修优先级|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» milestone|object|true|none|里程碑|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» author|object|true|none|发布者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» assigners|[object]|true|none|负责人|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» participants|[object]|true|none|参与者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» comment_journals_count|integer|true|none|评论数量|none|
|» operate_journals_count|integer|true|none|操作日志数量|none|

## GET 获取一个疑修

GET /api/v1/{owner}/{repo}/issues/{index}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "id": 62,
  "subject": "sedhahahaha213qjiqiq321312jakjf3123121231231",
  "project_issues_index": 32,
  "description": "部低九认都音以再际油化少可党半每市。无白表并商器何江华向二写。元先战因打质写起论国子四说省史十。计外们规织花层向结具习把月类义头比海。更感派是元家前识布连市厂号直北达则。千地意亲集式位状起几级号平常作意。",
  "branch_name": "持命看山",
  "start_date": "1988-11-06",
  "due_date": "1970-07-23",
  "created_at": "2023-02-15 09:53",
  "updated_at": "2023-02-20 11:38",
  "tags": [
    {
      "id": 3,
      "name": "1",
      "color": "#F17013"
    }
  ],
  "status": {
    "id": 5,
    "name": "关闭"
  },
  "priority": {
    "id": 3,
    "name": "高"
  },
  "milestone": {
    "id": 4,
    "name": "加花果马划八"
  },
  "author": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "assigners": [
    {
      "id": 110,
      "type": "User",
      "name": "咸蛋黄土豆丝",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    {
      "id": 115,
      "type": "User",
      "name": "嘿嘿嘿",
      "login": "yystopf123",
      "image_url": "system/lets/letter_avatars/2/Y/172_132_85/120.png"
    }
  ],
  "participants": [
    {
      "id": 1,
      "type": "AnonymousUser",
      "name": "游客",
      "login": "",
      "image_url": null
    },
    {
      "id": 110,
      "type": "User",
      "name": "咸蛋黄土豆丝",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    }
  ],
  "comment_journals_count": 8,
  "operate_journals_count": 33,
  "attachments": [
    {
      "id": 13,
      "title": "智能头盔软件需求构思和描述文档 (1).docx",
      "filesize": "14.1 KB",
      "is_pdf": false,
      "url": "/api/attachments/13",
      "created_on": "2023-02-17 16:05",
      "content_type": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      "id": 14,
      "title": "prv_test-master.zip",
      "filesize": "316 字节",
      "is_pdf": false,
      "url": "/api/attachments/14",
      "created_on": "2023-02-17 16:06",
      "content_type": "application/zip"
    }
  ]
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
|» id|integer|true|none||none|
|» subject|string|true|none|标题|none|
|» project_issues_index|integer|true|none|序号|none|
|» description|string|true|none|描述|none|
|» branch_name|string|true|none|分支名称|none|
|» start_date|string|true|none|开始日期|none|
|» due_date|string|true|none|截止日期|none|
|» created_at|string|true|none|创建时间|none|
|» updated_at|string|true|none|更新时间|none|
|» blockchain_token_num|integer|true|none|悬赏金额|none|
|» tags|[object]|true|none|标记|none|
|»» id|integer|false|none||none|
|»» name|string|false|none||none|
|»» color|string|false|none||none|
|» status|object|true|none|疑修状态|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» priority|object|true|none|疑修优先级|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» milestone|object|true|none|里程碑|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» author|object|true|none|发布者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» assigners|[object]|true|none|负责人|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» participants|[object]|true|none|参与者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» comment_journals_count|integer|true|none|评论数量|none|
|» operate_journals_count|integer|true|none|操作日志数量|none|
|» attachments|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» title|string|true|none|名称|none|
|»» filesize|string|true|none|文件大小|none|
|»» is_pdf|boolean|true|none|是否为PDF|none|
|»» url|string|true|none|文件地址|none|
|»» created_on|string|true|none|文件创建时间|none|
|»» content_type|string|true|none|文件类型|none|

## PATCH 更新一个疑修和状态

PATCH /api/v1/{owner}/{repo}/issues/{index}.json

> Body 请求参数

```json
{
  "status_id": 0,
  "priority_id": 0,
  "milestone_id": 0,
  "branch_name": "string",
  "start_date": "string",
  "due_date": "string",
  "subject": "string",
  "description": "string",
  "blockchain_token_num": 0,
  "issue_tag_ids": [
    0
  ],
  "assigner_ids": [
    0
  ],
  "attachment_ids": [
    0
  ],
  "receivers_login": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|access_token|query|string| 否 ||none|
|body|body|object| 否 ||none|
|» status_id|body|integer| 否 | 疑修状态ID|none|
|» priority_id|body|integer| 否 | 疑修优先级ID|none|
|» milestone_id|body|integer| 否 | 疑修里程碑ID|none|
|» branch_name|body|string| 否 | 疑修分支名称|none|
|» start_date|body|string| 否 | 疑修开始日期|none|
|» due_date|body|string| 否 | 疑修截止日期|none|
|» subject|body|string| 否 | 疑修标题|none|
|» description|body|string| 否 | 疑修描述|none|
|» blockchain_token_num|body|integer| 否 | 疑修悬赏金额|none|
|» issue_tag_ids|body|[integer]| 否 | 疑修标记ID数组|none|
|» assigner_ids|body|[integer]| 否 | 疑修负责人ID数组|none|
|» attachment_ids|body|[integer]| 否 | 疑修附件ID数组|none|
|» receivers_login|body|[string]| 否 | 疑修@用户标识数组|none|

> 返回示例

> 200 Response

```json
{
  "id": 62,
  "subject": "sedhahahaha213qjiqiq321312jakjf3123121231231",
  "project_issues_index": 32,
  "description": "部低九认都音以再际油化少可党半每市。无白表并商器何江华向二写。元先战因打质写起论国子四说省史十。计外们规织花层向结具习把月类义头比海。更感派是元家前识布连市厂号直北达则。千地意亲集式位状起几级号平常作意。",
  "branch_name": "持命看山",
  "start_date": "1988-11-06",
  "due_date": "1970-07-23",
  "created_at": "2023-02-15 09:53",
  "updated_at": "2023-02-20 11:38",
  "tags": [
    {
      "id": 3,
      "name": "1",
      "color": "#F17013"
    }
  ],
  "status": {
    "id": 5,
    "name": "关闭"
  },
  "priority": {
    "id": 3,
    "name": "高"
  },
  "milestone": {
    "id": 4,
    "name": "加花果马划八"
  },
  "author": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "assigners": [
    {
      "id": 110,
      "type": "User",
      "name": "咸蛋黄土豆丝",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    {
      "id": 115,
      "type": "User",
      "name": "嘿嘿嘿",
      "login": "yystopf123",
      "image_url": "system/lets/letter_avatars/2/Y/172_132_85/120.png"
    }
  ],
  "participants": [
    {
      "id": 1,
      "type": "AnonymousUser",
      "name": "游客",
      "login": "",
      "image_url": null
    },
    {
      "id": 110,
      "type": "User",
      "name": "咸蛋黄土豆丝",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    }
  ],
  "comment_journals_count": 8,
  "operate_journals_count": 33,
  "attachments": [
    {
      "id": 13,
      "title": "智能头盔软件需求构思和描述文档 (1).docx",
      "filesize": "14.1 KB",
      "is_pdf": false,
      "url": "/api/attachments/13",
      "created_on": "2023-02-17 16:05",
      "content_type": "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
    },
    {
      "id": 14,
      "title": "prv_test-master.zip",
      "filesize": "316 字节",
      "is_pdf": false,
      "url": "/api/attachments/14",
      "created_on": "2023-02-17 16:06",
      "content_type": "application/zip"
    }
  ]
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
|» id|integer|true|none|疑修ID|none|
|» subject|string|true|none|疑修标题|none|
|» project_issues_index|integer|true|none|疑修序号|none|
|» description|string|true|none|疑修描述|none|
|» branch_name|string|true|none|疑修分支名称|none|
|» start_date|string|true|none|疑修的开始时间|none|
|» due_date|string|true|none|疑修的结束时间|none|
|» created_at|string|true|none|疑修的创建时间|none|
|» updated_at|string|true|none|疑修的更新时间|none|
|» blockchain_token_num|integer|true|none|疑修悬赏金额|none|
|» tags|[object]|true|none|疑修标记数组|none|
|»» id|integer|false|none|疑修标记ID|none|
|»» name|string|false|none|疑修名称|none|
|»» color|string|false|none|疑修色值|none|
|» status|object|true|none|疑修状态|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» priority|object|true|none|疑修优先级|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» milestone|object|true|none|疑修里程碑|none|
|»» id|integer|true|none||none|
|»» name|string|true|none||none|
|» author|object|true|none|疑修作者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» assigners|[string]|true|none|疑修负责人数组|none|
|» participants|[object]|true|none|疑修参与人员数组|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» comment_journals_count|integer|true|none|疑修评论数量|none|
|» operate_journals_count|integer|true|none|疑修操作日志数量|none|

## DELETE 删除一个疑修

DELETE /api/v1/{owner}/{repo}/issues/{index}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## PATCH 批量更新疑修

PATCH /api/v1/{owner}/{repo}/issues/batch_update.json

> Body 请求参数

```json
{
  "ids": [
    0
  ],
  "status_id": 0,
  "priority_id": 0,
  "milestone_id": 0,
  "issue_tag_ids": [
    0
  ],
  "assigner_ids": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» ids|body|[integer]| 是 ||none|
|» status_id|body|integer| 是 | 疑修状态ID|none|
|» priority_id|body|integer| 是 | 疑修优先级ID|none|
|» milestone_id|body|integer| 是 | 疑修里程碑ID|none|
|» issue_tag_ids|body|[integer]| 是 | 疑修标记ID数组|none|
|» assigner_ids|body|[integer]| 是 | 疑修负责人ID数组|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## DELETE 批量删除疑修

DELETE /api/v1/{owner}/{repo}/issues/batch_destroy.json

> Body 请求参数

```json
{
  "ids": [
    0
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» ids|body|[integer]| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 项目所有成员列表

GET /api/v1/{owner}/{repo}/collaborators.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 7,
  "collaborators": [
    {
      "id": 36480,
      "type": "User",
      "name": "段甲生",
      "login": "pxicb3wke",
      "image_url": "images/avatars/User/36480?t=1672730523"
    },
    {
      "id": 85605,
      "type": "User",
      "name": "Floraachy",
      "login": "p10789563",
      "image_url": "system/lets/letter_avatars/2/P/158_138_26/120.png"
    },
    {
      "id": 85602,
      "type": "User",
      "name": "确实开源",
      "login": "ccf-gitlink",
      "image_url": "system/lets/letter_avatars/2/C/179_246_101/120.png"
    },
    {
      "id": 84727,
      "type": "User",
      "name": "yystopf.df",
      "login": "yystopf123",
      "image_url": "images/avatars/User/84727?t=1650252387"
    },
    {
      "id": 84961,
      "type": "User",
      "name": "postwoman釹轻柔",
      "login": "postwoman",
      "image_url": "images/avatars/User/84961?t=1641441279"
    },
    {
      "id": 84960,
      "type": "User",
      "name": "甄晓火",
      "login": "postman",
      "image_url": "system/lets/letter_avatars/2/Z/103_231_238/120.png"
    },
    {
      "id": 85589,
      "type": "User",
      "name": "🌸紫罗兰",
      "login": "floraachy",
      "image_url": "images/avatars/User/85589?t=1669948028"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» collaborators|[object]|true|none|成员列表|none|
|»» id|integer|true|none|用户ID|none|
|»» type|string|true|none|用户类型|none|
|»» name|string|true|none|用户名称|none|
|»» login|string|true|none|用户标识|none|
|»» image_url|string|true|none|用户头像|none|

## POST 创建一个疑修的评论

POST /api/v1/{owner}/{repo}/issues/{index}/journals.json

> Body 请求参数

```json
{
  "parent_id": 0,
  "reply_id": 0,
  "notes": "string",
  "attachment_ids": [
    0
  ],
  "receivers_login": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|access_token|query|string| 否 ||none|
|body|body|object| 否 ||none|
|» parent_id|body|integer| 否 ||父评论ID|
|» reply_id|body|integer| 否 ||回复的评论ID|
|» notes|body|string| 是 ||评论内容|
|» attachment_ids|body|[integer]| 否 ||评论附件ID|
|» receivers_login|body|[string]| 是 ||@用户数组|

> 返回示例

> 200 Response

```json
{
  "id": 58,
  "is_journal_detail": false,
  "created_at": "2023-02-18 22:41",
  "updated_at": "2023-02-18 22:41",
  "user": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "notes": "dolore officia minim in",
  "comments_count": 0,
  "children_journals": []
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
|» id|integer|true|none||评论ID|
|» is_journal_detail|boolean|true|none||是否为操作记录|
|» created_at|string|true|none||评论创建时间|
|» updated_at|string|true|none||评论更新时间|
|» user|object|true|none||评论者|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» notes|string|true|none||评论内容|
|» comments_count|integer|true|none||none|
|» children_journals|[string]|true|none||none|

## PATCH 修改疑修评论的内容

PATCH /api/v1/{owner}/{repo}/issues/{index}/journals/{id}.json

> Body 请求参数

```json
{
  "notes": "string",
  "attachment_ids": [
    0
  ],
  "receivers_login": [
    "string"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|id|path|string| 是 ||none|
|access_token|query|string| 否 ||none|
|body|body|object| 否 ||none|
|» notes|body|string| 是 | 评论内容|none|
|» attachment_ids|body|[integer]| 是 | 附件ID数组|none|
|» receivers_login|body|[string]| 否 | @用户数组|none|

> 返回示例

> 200 Response

```json
{
  "id": 58,
  "is_journal_detail": false,
  "created_at": "2023-02-18 22:41",
  "updated_at": "2023-02-18 23:00",
  "user": {
    "id": 110,
    "type": "User",
    "name": "咸蛋黄土豆丝",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
  },
  "notes": "dskflsk",
  "comments_count": 0,
  "children_journals": []
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
|» id|integer|true|none||评论ID|
|» is_journal_detail|boolean|true|none||是否为操作记录|
|» created_at|string|true|none||评论创建时间|
|» updated_at|string|true|none||评论更新时间|
|» user|object|true|none||评论用户|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» notes|string|true|none||评论内容|
|» comments_count|integer|true|none||子评论数量|
|» children_journals|[string]|true|none||子评论内容|

## DELETE 删除一个疑修评论

DELETE /api/v1/{owner}/{repo}/issues/{index}/journals/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 获取疑修评论的子评论列表

GET /api/v1/{owner}/{repo}/issues/{index}/journals/{id}/children_journals.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|id|path|string| 是 ||none|
|page|query|integer| 否 ||none|
|limit|query|integer| 否 ||none|
|keyword|query|string| 否 ||搜索关键词，可搜索评论内容|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "journals": [
    {
      "id": 58,
      "notes": "dskflsk",
      "comments_count": 0,
      "created_at": "2023-02-18 22:41",
      "updated_at": "2023-02-18 23:00",
      "user": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "reply_user": {
        "id": 115,
        "type": "User",
        "name": "嘿嘿嘿",
        "login": "yystopf123",
        "image_url": "system/lets/letter_avatars/2/Y/172_132_85/120.png"
      }
    }
  ]
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
|» total_count|integer|true|none||none|
|» journals|[object]|true|none||none|
|»» id|integer|false|none||none|
|»» notes|string|false|none|评论内容|none|
|»» comments_count|integer|false|none|子评论数量|none|
|»» created_at|string|false|none|创建时间|none|
|»» updated_at|string|false|none|更新时间|none|
|»» user|object|false|none|评论的用户|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» reply_user|object|false|none|回复的用户|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|

# 合并请求

## GET 获取合并请求下评论或操作记录列表 

GET /api/v1/{owner}/{repo}/issues/{index}/journals.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|category|query|string| 否 ||类型，all 所有 comment 评论 operate 操作记录|
|keyword|query|string| 否 ||搜索关键词|
|sort_by|query|string| 否 ||排序字段，created_on 创建时间 updated_on 更新时间|
|sort_direction|query|string| 否 ||排序类型，asc 正序 desc 倒序|
|page|query|integer| 否 ||none|
|limit|query|integer| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "total_journals_count": 42,
  "total_operate_journals_count": 38,
  "total_comment_journals_count": 4,
  "total_count": 42,
  "journals": [
    {
      "id": 13,
      "is_journal_detail": false,
      "created_at": "2023-02-15 09:53",
      "updated_at": "2023-02-15 09:53",
      "user": {
        "id": 110,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "notes": null,
      "comments_count": 0,
      "children_journals": [],
      "attachments": []
    },
    {
      "id": 17,
      "is_journal_detail": true,
      "created_at": "2023-02-16 18:20",
      "updated_at": "2023-02-16 18:20",
      "user": {
        "id": 1,
        "type": "AnonymousUser",
        "name": "游客",
        "login": "",
        "image_url": null
      },
      "operate_category": "subject",
      "operate_content": "修改了<b>标题</b>"
    }
  ]
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
|» total_count|integer|true|none||分页总数|
|» total_journals_count|integer|true|none||所有评论+操作记录数|
|» total_comment_journals_count|integer|true|none||所有评论数|
|» total_operate_journals_count|integer|true|none||所有操作记录数|
|» journals|[object]|true|none||none|
|»» id|integer|true|none||评论id|
|»» is_journal_detail|boolean|true|none||是否为操作记录|
|»» created_at|string|true|none||创建时间|
|»» updated_at|string|true|none||更新时间|
|»» user|object|true|none||评论用户|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string¦null|true|none||none|
|»» notes|null|false|none||评论内容|
|»» comments_count|integer|false|none||子评论数量|
|»» children_journals|[string]|false|none||子评论，默认只显示10条|
|»» operate_content|string|false|none||操作记录内容|
|»» operate_category|string|true|none||操作记录类型  issue 创建疑修；attachment 附件；issue_tag 标记；assigner 负责人； subject 标题；description 描述； status_id 状态；priority_id 优先级； fixed_version_id 里程碑； branch_name 关联分支；start_date 开始日期；due_date 结束日期|

## GET 获取合并请求列表

GET /api/v1/{owner}/{repo}/pulls.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|keyword|query|string| 否 ||搜索关键词|
|status|query|string| 否 ||合并请求类型，0 开启的 1 合并的 2 关闭，不传为全部|
|priority_id|query|string| 否 ||优先级ID|
|issue_tag_id|query|string| 否 ||标记ID|
|version_id|query|string| 否 ||里程碑ID|
|reviewer_id|query|string| 否 ||审查人员ID|
|assign_user_id|query|string| 否 ||指派人员ID|
|sort_by|query|string| 否 ||排序字段，updated_at 更新时间 created_at 创建时间|
|sort_direction|query|string| 否 ||排序类型。desc 倒序 asc 正序|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "pulls": [
    {
      "id": 11,
      "title": "ceshi",
      "body": null,
      "head": "master_new",
      "base": "master",
      "is_original": false,
      "index": 1,
      "status": "open",
      "issue": {
        "id": 105,
        "author": {
          "id": 110,
          "type": "User",
          "name": "咸蛋黄土豆丝xxx",
          "login": "yystopf",
          "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
        },
        "priority": null,
        "version": null,
        "journals_count": 0,
        "issue_tags": null
      },
      "reviewers": [],
      "journals_count": 0
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» pulls|[object]|true|none||none|
|»» id|integer|false|none|合并请求ID|none|
|»» title|string|false|none|合并请求标题|none|
|»» body|null|false|none|合并请求内容|none|
|»» head|string|false|none|合并请求源分支|none|
|»» base|string|false|none|合并请求目标分支|none|
|»» is_original|boolean|false|none|是否为fork仓库发来的合并请求|none|
|»» index|integer|false|none|合并请求序号|none|
|»» status|string|false|none|合并请求状态|open: 开启的, merged: 合并的, closed: 关闭的|
|»» fork_project|object|false|none||none|
|»»» id|integer|false|none|fork仓库的id|none|
|»»» identifier|string|false|none|fork仓库的标识|none|
|»»» login|string|true|none|fork仓库拥有者的标识|none|
|»» issue|object|false|none||none|
|»»» id|integer|true|none|疑修ID|none|
|»»» author|object|true|none|疑修创建者|none|
|»»»» id|integer|true|none||none|
|»»»» type|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» priority|null|true|none|疑修优先级|none|
|»»» version|null|true|none|疑修里程碑|none|
|»»» journals_count|integer|true|none|普通评论数量|none|
|»»» issue_tags|null|true|none|普通评论数量|none|
|»» reviewers|[string]|false|none|代码审查者标识数组|none|
|»» journals_count|integer|false|none|审查评论数量|none|

## GET 获取一个合并请求

GET /api/v1/{owner}/{repo}/pulls/{index}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "id": 11,
  "title": "ceshi",
  "body": null,
  "head": "master_new",
  "base": "master",
  "is_original": false,
  "index": 1,
  "status": "open",
  "issue": {
    "id": 105,
    "author": {
      "id": 110,
      "type": "User",
      "name": "咸蛋黄土豆丝xxx",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
    },
    "priority": null,
    "version": null,
    "journals_count": 0,
    "issue_tags": null
  },
  "reviewers": [],
  "journals_count": 0,
  "merge_base": "96dc82d07e05327fc908aaaa127ba45f02091d85",
  "base_commit_sha": "96dc82d07e05327fc908aaaa127ba45f02091d85",
  "head_commit_sha": "82861402ada099d3e288fc41680596dde297d022",
  "commit_num": 2,
  "changed_files": 2,
  "is_locked": false,
  "mergeable": true,
  "merged": false,
  "merged_at": "",
  "merge_commit_sha": null,
  "merge_by": null,
  "last_review": null,
  "conflict_files": null
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
|» id|integer|true|none|合并请求ID|none|
|» title|string|true|none|合并请求标题|none|
|» body|null|true|none|合并请求描述|none|
|» head|string|true|none|合并请求源分支|none|
|» base|string|true|none|合并请求目标分支|none|
|» is_original|boolean|true|none|是否为fork仓库发来的合并请求|none|
|» index|integer|true|none|合并请求序号|none|
|» status|string|true|none|合并请求状态|open: 开启的, merged: 合并的, closed: 关闭的|
|» issue|object|true|none||none|
|»» id|integer|true|none|疑修ID|none|
|»» author|object|true|none|疑修作者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» priority|null|true|none|疑修优先级|none|
|»» version|null|true|none|疑修里程碑|none|
|»» journals_count|integer|true|none|普通评论数量|none|
|»» issue_tags|null|true|none|标记|none|
|» reviewers|[string]|true|none|审查者标识数组|none|
|» journals_count|integer|true|none|审查评论数量|none|
|» merge_base|string|true|none|目标的commit ID|none|
|» base_commit_sha|string|true|none|合并之后的第一个commit ID|none|
|» head_commit_sha|string|true|none|源commit ID|none|
|» commit_num|integer|true|none|commit数量|none|
|» changed_files|integer|true|none|更改文件数量|none|
|» is_locked|boolean|true|none||none|
|» mergeable|boolean|true|none|是否能合并|none|
|» merged|boolean|true|none|是否合并|none|
|» merged_at|string|true|none|合并时间|none|
|» merge_commit_sha|null|true|none|合并之后的第一个commit ID|none|
|» merge_by|null|true|none|被谁合并了|none|
|» last_review|object|false|none||none|
|»» id|string|false|none|最后一个审查的id|none|
|»» commit_id|string|false|none|最后一个审查对应的commit ID|none|
|»» content|string|false|none|最后一个审查的内容|none|
|»» status|string|false|none|最后一个审查的状态|common: 一般审查, approved: 通过, rejected: 拒绝通过|
|»» created_at|string|false|none|审查创建的时间|none|
|»» reviewer|object|false|none|审查创建者|none|
|»»» id|string|false|none||none|
|»»» login|string|false|none||none|
|»»» name|string|false|none||none|
|»»» image_url|string|false|none||none|
|» conflict_files|null|true|none|有冲突的文件|none|

## POST 重新打开一个合并请求

POST /api/v1/{owner}/{repo}/pulls/{index}/reopen.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 创建一个合并请求

POST /api/{owner}/{repo}/pulls.json

> Body 请求参数

```json
{
  "title": "324",
  "assigned_to_id": "2",
  "fixed_version_id": "",
  "issue_tag_ids": [],
  "priority_id": "2",
  "body": "312",
  "head": "new_branch_1",
  "base": "master",
  "is_original": false,
  "fork_project_id": "",
  "files_count": 1,
  "commits_count": 1,
  "reviewer_ids": [],
  "receivers_login": []
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» title|body|string| 是 | 合并请求标题|none|
|» assigned_to_id|body|integer| 否 | 指派人员ID|none|
|» fixed_version_id|body|integer| 否 | 里程碑ID|none|
|» issue_tag_ids|body|[string]| 否 | 标记ID数组|none|
|» priority_id|body|string| 是 | 优先级ID|none|
|» body|body|string| 是 | 合并请求内容|none|
|» head|body|string| 是 | 源分支|none|
|» base|body|string| 是 | 目标分支|none|
|» is_original|body|boolean| 是 | 是否为fork仓库发来的合并请求|none|
|» merge_project_identifier|body|string| 是 | fork仓库标识|none|
|» fork_project_id|body|integer| 是 | fork仓库ID|none|
|» receivers_login|body|[string]| 是 | @人员的login|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "响应成功",
  "pull_request_id": 12,
  "pull_request_number": 2,
  "pull_request_status": 0,
  "pull_request_head": "master_new",
  "pull_request_base": "master",
  "pull_request_staus": "open",
  "is_original": false,
  "fork_project_id": null,
  "fork_project_identifier": null,
  "fork_project_user": null,
  "reviewers": [],
  "id": 110,
  "name": "或交委参她",
  "pr_time": "1秒前",
  "assign_user_name": null,
  "assign_user_login": null,
  "author_name": "咸蛋黄土豆丝xxx",
  "author_login": "yystopf",
  "avatar_url": "system/lets/letter_avatars/2/X/230_139_26/120.png",
  "priority": "低",
  "version": null,
  "journals_count": 0,
  "issue_tags": null
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
|» message|string|true|none||none|
|» pull_request_id|integer|true|none|合并请求ID|none|
|» pull_request_number|integer|true|none|合并请求序号|none|
|» pull_request_status|integer|true|none|合并请求状态|0 开启的 1 合并的 2 关闭的|
|» pull_request_head|string|true|none|源分支|none|
|» pull_request_base|string|true|none|目标分支|none|
|» pull_request_staus|string|true|none|合并请求状态|open 开启的  merged 合并的 closed 关闭的|
|» is_original|boolean|true|none|是否为fork仓库发来的合并请求|none|
|» fork_project_id|null|true|none|fork仓库ID|none|
|» fork_project_identifier|null|true|none|fork仓库标识|none|
|» fork_project_user|null|true|none|fork仓库拥有者标识|none|
|» reviewers|[string]|true|none|审查者标识数组|none|
|» id|integer|true|none|疑修|none|
|» name|string|true|none|疑修标题|none|
|» pr_time|string|true|none|更新时间|none|
|» assign_user_name|null|true|none|指派者昵称|none|
|» assign_user_login|null|true|none|指派者用户标识|none|
|» author_name|string|true|none|创建者昵称|none|
|» author_login|string|true|none|创建者用户标识|none|
|» avatar_url|string|true|none|创建者用户头像|none|
|» priority|string|true|none|优先级|none|
|» version|null|true|none|里程碑|none|
|» journals_count|integer|true|none|普通评论数量|none|
|» issue_tags|null|true|none|标记数组|none|

## PUT 更新一个合并请求

PUT /api/{owner}/{repo}/pulls/{index}.json

> Body 请求参数

```json
{
  "title": "xxxxxx",
  "issue_tag_ids": [],
  "body": "xxxx",
  "head": "master_new",
  "base": "master",
  "receivers_login": []
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» title|body|string| 是 | 合并请求标题|none|
|» issue_tag_ids|body|[string]| 是 | 标记ID数组|none|
|» body|body|string| 是 | 合并请求内容|none|
|» head|body|string| 是 | 源分支|none|
|» base|body|string| 是 | 目标分支|none|
|» receivers_login|body|[string]| 是 | @的用户标识数组|none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 拒绝一个合并请求

POST /api/{owner}/{repo}/pulls/{index}/refuse_merge.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|

> 返回示例

> 200 Response

```json
{
  "status": 1,
  "message": "已拒绝"
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
|» message|string|true|none||none|

## POST 通过一个合并请求

POST /api/{owner}/{repo}/pulls/{index}/pr_merge.json

> Body 请求参数

```json
{
  "project_id": "ceshi_doc",
  "id": 15,
  "do": "merge",
  "body": "1",
  "title": "1"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|
|body|body|object| 否 ||none|
|» project_id|body|string| 是 | 项目ID|none|
|» id|body|integer| 是 | 合并请求ID|none|
|» do|body|string| 是 | 合并方式|none|
|» body|body|string| 是 | 合并内容|none|
|» title|body|string| 是 | 合并标题|none|

#### 枚举值

|属性|值|
|---|---|
|» do|merge|
|» do|rebase|
|» do|rebase-merge|
|» do|squash|
|» do|manually-merged|

> 返回示例

> 200 Response

```json
{
  "status": 1,
  "message": "合并成功"
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
|» message|string|true|none||none|

## GET 获取一个合并请求变更文件列表

GET /api/{owner}/{repo}/pulls/{id}/files.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "files_count": 2,
  "total_addition": 1,
  "total_deletion": 0,
  "files": [
    {
      "sha": "82861402ada099d3e288fc41680596dde297d022",
      "name": "cxxc",
      "old_name": "cxxc",
      "index": 1,
      "addition": 0,
      "deletion": 0,
      "type": 1,
      "isCreated": true,
      "isDeleted": false,
      "isBin": false,
      "isLFSFile": false,
      "isRenamed": false,
      "isSubmodule": false,
      "sections": []
    },
    {
      "sha": "82861402ada099d3e288fc41680596dde297d022",
      "name": "测试",
      "old_name": "测试",
      "index": 2,
      "addition": 1,
      "deletion": 0,
      "type": 1,
      "isCreated": true,
      "isDeleted": false,
      "isBin": false,
      "isLFSFile": false,
      "isRenamed": false,
      "isSubmodule": false,
      "sections": [
        {
          "fileName": "测试",
          "name": "",
          "lines": [
            {
              "leftIdx": 0,
              "rightIdx": 0,
              "type": 4,
              "content": "@@ -0,0 +1 @@",
              "sectionInfo": {
                "path": "测试",
                "lastLeftIdx": 0,
                "lastRightIdx": 0,
                "leftIdx": 0,
                "rightIdx": 1,
                "leftHunkSize": 0,
                "rightHunkSize": 0
              }
            },
            {
              "leftIdx": 0,
              "rightIdx": 1,
              "type": 2,
              "content": "+是多福多寿",
              "sectionInfo": null
            }
          ]
        }
      ]
    }
  ]
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
|» files_count|integer|true|none|文件更改的总数量|none|
|» total_addition|integer|true|none|添加代码总行数|none|
|» total_deletion|integer|true|none|删除代码总行数|none|
|» files|[object]|true|none||none|
|»» sha|string|true|none|commit标识|none|
|»» name|string|true|none|当前文件名|none|
|»» old_name|string|true|none|修改之前的文件名称,与name相同的话，说明文件名未更改|none|
|»» index|integer|true|none||none|
|»» addition|integer|true|none|文件添加的行数|none|
|»» deletion|integer|true|none|文件删除的行数|none|
|»» type|integer|true|none|文件类型|， 1: 表示该文件只添加了内容，2: 表示该文件内容有修改， 3: 表示文件被删除或者改文件只删除了内容|
|»» isCreated|boolean|true|none|当前文件是否为新增文件|none|
|»» isDeleted|boolean|true|none|当前文件是否被删除|none|
|»» isBin|boolean|true|none|当前文件是否为二进制文件|none|
|»» isLFSFile|boolean|true|none|当前文件是否为LFS文件|none|
|»» isRenamed|boolean|true|none|当前文件是否被重命名|none|
|»» isSubmodule|boolean|true|none|当前文件是否为子模块|none|
|»» sections|[object]|true|none||none|
|»»» fileName|string|false|none|文件名称|none|
|»»» name|string|false|none||none|
|»»» lines|[object]|false|none||none|
|»»»» leftIdx|integer|true|none|文件变动之前所在行数|none|
|»»»» rightIdx|integer|true|none|文件更改后所在行数|none|
|»»»» type|integer|true|none|文件变更类型|1: 新增，2: 修改， 3: 删除， 4: diff统计信息|
|»»»» content|string|true|none|文件变更的内容|none|
|»»»» sectionInfo|object¦null|true|none||none|
|»»»»» path|string|true|none|文件相对仓库的路径|none|
|»»»»» lastLeftIdx|integer|true|none||none|
|»»»»» lastRightIdx|integer|true|none||none|
|»»»»» leftIdx|integer|true|none|文件变更之前所在行数|none|
|»»»»» rightIdx|integer|true|none|文件变更之后所在行数(即：页面编辑器开始显示的行数)|none|
|»»»»» leftHunkSize|integer|true|none|文件变更之前的行数|none|
|»»»»» rightHunkSize|integer|true|none|文件变更之后的行数(及当前页面编辑器显示的总行数)|none|

## GET 获取一个合并请求变更文件列表（简版）

GET /api/v1/{owner}/{repo}/pulls/{index}/files.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|page|query|string| 否 ||分页页码，当filepath存在时无效|
|limit|query|string| 否 ||分页个数，当filepath存在时无效|
|filepath|query|string| 否 ||文件路径|

> 返回示例

> 200 Response

```json
{
  "file_nums": 224,
  "total_addition": 24948,
  "total_deletion": 59,
  "files": [
    {
      "filename": ".changeset/README.md",
      "old_name": ".changeset/README.md",
      "index": 1,
      "type": 1,
      "is_bin": false,
      "is_created": true,
      "is_deleted": false,
      "is_lfs_file": false,
      "is_renamed": false,
      "is_submodule": false,
      "additions": 8,
      "deletions": 0,
      "changes": 8,
      "sha": "cd9bddbe10a006632e53bfb485cdfc1d220eef4e"
    }
  ]
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
|» file_nums|integer|true|none||none|
|» total_addition|integer|true|none||none|
|» total_deletion|integer|true|none||none|
|» files|[object]|true|none||none|
|»» filename|string|false|none||none|
|»» old_name|string|false|none||none|
|»» index|integer|false|none||none|
|»» type|integer|false|none||none|
|»» is_bin|boolean|false|none||none|
|»» is_created|boolean|false|none||none|
|»» is_deleted|boolean|false|none||none|
|»» is_lfs_file|boolean|false|none||none|
|»» is_renamed|boolean|false|none||none|
|»» is_submodule|boolean|false|none||none|
|»» additions|integer|false|none||none|
|»» deletions|integer|false|none||none|
|»» changes|integer|false|none||none|
|»» sha|string|false|none||none|

## GET 获取一个合并请求提交列表

GET /api/{owner}/{repo}/pulls/{id}/commits.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "commits_count": 2,
  "commits": [
    {
      "author": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "committer": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "timestamp": 1700643425,
      "time_from_now": "8天前",
      "created_at": "2023-11-22 16:57",
      "message": "Add cxxc\n",
      "sha": "82861402ada099d3e288fc41680596dde297d022"
    },
    {
      "author": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "committer": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "timestamp": 1700621060,
      "time_from_now": "8天前",
      "created_at": "2023-11-22 10:44",
      "message": "Add 测试\n",
      "sha": "37d52b40456755c364a0dea77b702381606edefb"
    }
  ]
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
|» commits_count|integer|true|none|提交总数|none|
|» commits|[object]|true|none||none|
|»» author|object|true|none|作者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» committer|object|true|none|提交者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» timestamp|integer|true|none|提交时间戳|none|
|»» time_from_now|string|true|none|提交时间（距离现在）|none|
|»» created_at|string|true|none|创建时间|none|
|»» message|string|true|none|提交信息|none|
|»» sha|string|true|none|提交标识|none|

## GET 获取两个分支、标签、提交标识之间的比较内容

GET /api/{owner}/{repo}/compare/{head}...{base}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|head|path|string| 是 ||源分支、标签或提交标识（base64加密后）|
|base|path|string| 是 ||目标分支、标签或提交标识（base64加密后）|

> 返回示例

> 200 Response

```json
{
  "commits_count": 2,
  "commits": [
    {
      "author": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "committer": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "timestamp": 1700643425,
      "time_from_now": "8天前",
      "created_at": "2023-11-22 16:57",
      "message": "Add cxxc\n",
      "sha": "82861402ada099d3e288fc41680596dde297d022"
    },
    {
      "author": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "committer": {
        "id": "110",
        "login": "yystopf",
        "name": "咸蛋黄土豆丝xxx",
        "type": "User",
        "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png"
      },
      "timestamp": 1700621060,
      "time_from_now": "8天前",
      "created_at": "2023-11-22 10:44",
      "message": "Add 测试\n",
      "sha": "37d52b40456755c364a0dea77b702381606edefb"
    }
  ],
  "diff": {
    "files_count": 2,
    "total_addition": 1,
    "total_deletion": 0,
    "files": [
      {
        "sha": null,
        "name": "cxxc",
        "old_name": "cxxc",
        "index": 1,
        "addition": 0,
        "deletion": 0,
        "type": 1,
        "isCreated": true,
        "isDeleted": false,
        "isBin": false,
        "isLFSFile": false,
        "isRenamed": false,
        "isSubmodule": false,
        "sections": []
      },
      {
        "sha": null,
        "name": "测试",
        "old_name": "测试",
        "index": 2,
        "addition": 1,
        "deletion": 0,
        "type": 1,
        "isCreated": true,
        "isDeleted": false,
        "isBin": false,
        "isLFSFile": false,
        "isRenamed": false,
        "isSubmodule": false,
        "sections": [
          {
            "fileName": "测试",
            "name": "",
            "lines": [
              {
                "leftIdx": 0,
                "rightIdx": 0,
                "type": 4,
                "content": "@@ -0,0 +1 @@",
                "sectionInfo": {
                  "path": "测试",
                  "lastLeftIdx": 0,
                  "lastRightIdx": 0,
                  "leftIdx": 0,
                  "rightIdx": 1,
                  "leftHunkSize": 0,
                  "rightHunkSize": 0
                }
              },
              {
                "leftIdx": 0,
                "rightIdx": 1,
                "type": 2,
                "content": "+是多福多寿",
                "sectionInfo": null
              }
            ]
          }
        ]
      }
    ]
  },
  "status": -2,
  "message": "在这些分支之间的合并请求已存在：<a href='/yystopfceshi/testdevops/pulls/12'>或交委参她</a>"
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
|» commits_count|integer|true|none|提交总数|none|
|» commits|[object]|true|none||none|
|»» author|object|true|none|作者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» committer|object|true|none|提交者|none|
|»»» id|string|true|none||none|
|»»» login|string|true|none||none|
|»»» name|string|true|none||none|
|»»» type|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» timestamp|integer|true|none|提交时间|none|
|»» time_from_now|string|true|none|提交时间（距离现在）|none|
|»» created_at|string|true|none|创建时间|none|
|»» message|string|true|none|提交信息|none|
|»» sha|string|true|none|提交标识|none|
|» diff|object|true|none||none|
|»» files_count|integer|true|none|变更文件数量|none|
|»» total_addition|integer|true|none|添加代码总行数|none|
|»» total_deletion|integer|true|none|删除代码总行数|none|
|»» files|[object]|true|none||none|
|»»» sha|null|true|none|提交标识|none|
|»»» name|string|true|none|当前文件名|none|
|»»» old_name|string|true|none|修改之前的文件名称,与name相同的话，说明文件名未更改|none|
|»»» index|integer|true|none||none|
|»»» addition|integer|true|none|文件添加的行数|none|
|»»» deletion|integer|true|none|文件删除的行数|none|
|»»» type|integer|true|none|文件类型|1: 表示该文件只添加了内容，2: 表示该文件内容有修改， 3: 表示文件被删除或者改文件只删除了内容|
|»»» isCreated|boolean|true|none|当前文件是否为新增文件|none|
|»»» isDeleted|boolean|true|none|当前文件是否被删除|none|
|»»» isBin|boolean|true|none|当前文件是否为二进制文件|none|
|»»» isLFSFile|boolean|true|none|当前文件是否为LFS文件|none|
|»»» isRenamed|boolean|true|none|当前文件是否被重命名|none|
|»»» isSubmodule|boolean|true|none|当前文件是否为子模块|none|
|»»» sections|[object]|true|none||none|
|»»»» fileName|string|false|none|文件名称|none|
|»»»» name|string|false|none||none|
|»»»» lines|[object]|false|none||none|
|»»»»» leftIdx|integer|true|none|文件变动之前所在行数|none|
|»»»»» rightIdx|integer|true|none|文件更改后所在行数|none|
|»»»»» type|integer|true|none|文件变更类型|1: 内容未改动，2: 添加， 3: 删除， 4: diff统计信息|
|»»»»» content|string|true|none|文件变更的内容|none|
|»»»»» sectionInfo|object¦null|true|none||none|
|»»»»»» path|string|true|none|文件相对仓库的路径|none|
|»»»»»» lastLeftIdx|integer|true|none||none|
|»»»»»» lastRightIdx|integer|true|none||none|
|»»»»»» leftIdx|integer|true|none|文件变更之前所在行数|none|
|»»»»»» rightIdx|integer|true|none|文件变更之后所在行数|none|
|»»»»»» leftHunkSize|integer|true|none|文件变更之前的行数|none|
|»»»»»» rightHunkSize|integer|true|none|文件变更之后的行数(及当前页面编辑器显示的总行数)|none|
|» status|integer|true|none|状态|none|
|» message|string|true|none|提示信息|none|

## GET 获取两个分支、标签、提交标识之间的变更文件列表

GET /api/v1/{owner}/{repo}/compare/{head}...{base}/files.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|head|path|string| 是 ||源分支、标签或提交标识（base64加密后）|
|base|path|string| 是 ||目标分支、标签或提交标识（base64加密后）|
|page|query|string| 否 ||分页页码，当filepath存在时无效|
|limit|query|string| 否 ||分页个数，当filepath存在时无效|
|filepath|query|string| 否 ||文件路径|

> 返回示例

> 200 Response

```json
{
  "file_nums": 372,
  "total_addition": 25653,
  "total_deletion": 13130,
  "files": [
    {
      "filename": "packages/inula-cli/externals.d.ts",
      "old_name": "packages/inula-cli/externals.d.ts",
      "index": 5,
      "type": 2,
      "is_bin": false,
      "is_created": false,
      "is_deleted": false,
      "is_lfs_file": false,
      "is_renamed": false,
      "is_submodule": false,
      "additions": 0,
      "deletions": 1,
      "changes": 1,
      "sha": "master"
    }
  ],
  "status": 0,
  "message": "可以合并"
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
|» file_nums|integer|true|none||none|
|» total_addition|integer|true|none||none|
|» total_deletion|integer|true|none||none|
|» files|[object]|true|none||none|
|»» filename|string|false|none||none|
|»» old_name|string|false|none||none|
|»» index|integer|false|none||none|
|»» type|integer|false|none||none|
|»» is_bin|boolean|false|none||none|
|»» is_created|boolean|false|none||none|
|»» is_deleted|boolean|false|none||none|
|»» is_lfs_file|boolean|false|none||none|
|»» is_renamed|boolean|false|none||none|
|»» is_submodule|boolean|false|none||none|
|»» additions|integer|false|none||none|
|»» deletions|integer|false|none||none|
|»» changes|integer|false|none||none|
|»» sha|string|false|none||none|
|» status|integer|true|none||none|
|» message|string|true|none||none|

## GET 获取合并请求版本列表

GET /api/v1/{owner}/{repo}/pulls/{index}/versions.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "versions": [
    {
      "id": 38,
      "add_line_num": 1,
      "del_line_num": 0,
      "commits_count": 1,
      "files_count": 1,
      "base_commit_sha": "96dc82d07e05327fc908aaaa127ba45f02091d85",
      "head_commit_sha": "37d52b40456755c364a0dea77b702381606edefb",
      "start_commit_sha": "96dc82d07e05327fc908aaaa127ba45f02091d85",
      "created_time": 1700632918,
      "updated_time": 1700632918
    },
    {
      "id": 39,
      "add_line_num": 0,
      "del_line_num": 0,
      "commits_count": 1,
      "files_count": 1,
      "base_commit_sha": "96dc82d07e05327fc908aaaa127ba45f02091d85",
      "head_commit_sha": "82861402ada099d3e288fc41680596dde297d022",
      "start_commit_sha": "37d52b40456755c364a0dea77b702381606edefb",
      "created_time": 1700643432,
      "updated_time": 1700643432
    }
  ]
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
|» total_count|integer|true|none|合并请求版本总数|none|
|» versions|[object]|true|none||none|
|»» id|integer|true|none|版本ID|none|
|»» add_line_num|integer|true|none|该版本新增行数|none|
|»» del_line_num|integer|true|none|该版本删除行数|none|
|»» commits_count|integer|true|none|该版本提交总数|none|
|»» files_count|integer|true|none|该版本提交文件总数|none|
|»» base_commit_sha|string|true|none|目标commit标识|none|
|»» head_commit_sha|string|true|none|源commit标识|none|
|»» start_commit_sha|string|true|none|该版本起始commit标识|none|
|»» created_time|integer|true|none|版本创建时间|none|
|»» updated_time|integer|true|none|版本更新时间|none|

## GET 获取合并请求版本之间的Diff

GET /api/v1/{owner}/{repo}/pulls/{index}/versions/{version_id}/diff.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|
|version_id|path|string| 是 ||合并请求版本的ID|
|filepath|query|string| 否 ||文件路径|

#### 详细说明

**filepath**: 文件路径

> 返回示例

> 200 Response

```json
{
  "file_nums": 1,
  "total_addition": 1,
  "total_deletion": 0,
  "files": [
    {
      "name": "测试",
      "oldname": "测试",
      "addition": 1,
      "deletion": 0,
      "type": 1,
      "is_created": true,
      "is_deleted": false,
      "is_bin": false,
      "is_lfs_file": false,
      "is_renamed": false,
      "is_ambiguous": false,
      "is_submodule": false,
      "diff": null,
      "sections": [
        {
          "file_name": "测试",
          "name": "",
          "lines": [
            {
              "left_index": 0,
              "right_index": 0,
              "match": 0,
              "type": 4,
              "content": "@@ -0,0 +1 @@",
              "section_path": "测试",
              "section_last_left_index": 0,
              "section_last_right_index": 0,
              "section_left_index": 0,
              "section_right_index": 1,
              "section_left_hunk_size": 0,
              "section_right_hunk_size": 0
            },
            {
              "left_index": 0,
              "right_index": 1,
              "match": -1,
              "type": 2,
              "content": "+是多福多寿"
            }
          ]
        }
      ],
      "is_incomplete": false,
      "is_incomplete_line_too_long": false,
      "is_protected": false
    }
  ]
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
|» file_nums|integer|true|none|文件数量|none|
|» total_addition|integer|true|none|新增行数|none|
|» total_deletion|integer|true|none|删除行数|none|
|» files|[object]|true|none||none|
|»» name|string|false|none|文件名称|none|
|»» oldname|string|false|none|文件修改前名称|none|
|»» addition|integer|false|none|文件新增行数|none|
|»» deletion|integer|false|none|文件删除行数|none|
|»» type|integer|false|none|文件类型 |1: 新增 2: 更改 3: 删除 4: 重命名 5: 复制|
|»» is_created|boolean|false|none|是否为新建文件|none|
|»» is_deleted|boolean|false|none|是否为删除文件|none|
|»» is_bin|boolean|false|none|是否为二进制文件|none|
|»» is_lfs_file|boolean|false|none|是否为LFS文件|none|
|»» is_renamed|boolean|false|none|是否重命名|none|
|»» is_ambiguous|boolean|false|none||none|
|»» is_submodule|boolean|false|none|是否为子模块|none|
|»» diff|null|false|none||none|
|»» sections|[object]|false|none||none|
|»»» file_name|string|false|none|文件名称|none|
|»»» name|string|false|none||none|
|»»» lines|[object]|false|none||none|
|»»»» left_index|integer|true|none|文件变动之前所在行数|none|
|»»»» right_index|integer|true|none|文件变动之后所在行数|none|
|»»»» match|integer|true|none||none|
|»»»» type|integer|true|none|文件变更类型|none|
|»»»» content|string|true|none|文件变更内容|none|
|»»»» section_path|string|false|none|文件路径|none|
|»»»» section_last_left_index|integer|false|none||none|
|»»»» section_last_right_index|integer|false|none||none|
|»»»» section_left_index|integer|false|none|文件变更之前所在行数|none|
|»»»» section_right_index|integer|false|none|文件变更之后所在行数(即：页面编辑器开始显示的行数)|none|
|»»»» section_left_hunk_size|integer|false|none|文件变更之前的行数|none|
|»»»» section_right_hunk_size|integer|false|none|文件变更之后的行数(及当前页面编辑器显示的总行数)|none|
|»» is_incomplete|boolean|false|none|是否不完整|none|
|»» is_incomplete_line_too_long|boolean|false|none|文件是否不完整是因为太长了|none|
|»» is_protected|boolean|false|none|文件是否被保护|none|

## GET 获取合并请求审查列表

GET /api/v1/{owner}/{repo}/pulls/{index}/reviews.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|
|status|query|string| 否 ||审查类型， common: 评论类型, approved: 已通过, rejected: 已拒绝|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "reviews": [
    {
      "reviewer": {
        "id": 2,
        "type": "User",
        "name": "heh",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
      },
      "pull_request": {
        "id": 163,
        "title": "新合并请求1",
        "body": null,
        "head": "master_1",
        "base": "master",
        "is_original": false,
        "index": 1,
        "status": "closed",
        "issue": {
          "id": 260,
          "author": {
            "id": 2,
            "type": "User",
            "name": "heh",
            "login": "yystopf",
            "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
          },
          "priority": null,
          "version": null,
          "journals_count": 0,
          "issue_tags": null
        },
        "reviewers": [],
        "journals_count": 8
      },
      "id": 5,
      "commit_id": null,
      "content": "新建一个审查",
      "status": "common",
      "created_at": "2022-07-25 17:08"
    }
  ]
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
|» total_count|integer|true|none|审查总数|none|
|» reviews|[object]|true|none||none|
|»» reviewer|object|false|none|审查者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» pull_request|object|false|none||none|
|»»» id|integer|true|none|合并请求ID|none|
|»»» title|string|true|none|合并请求标题|none|
|»»» body|null|true|none|合并请求内容|none|
|»»» head|string|true|none|合并请求源分支|none|
|»»» base|string|true|none|合并请求目标分支|none|
|»»» is_original|boolean|true|none|合并请求目标分支|none|
|»»» fork_project|object|false|none||none|
|»»»» id|integer|false|none|fork仓库的id|none|
|»»»» identifier|string|false|none|fork仓库的标识|none|
|»»»» login|string|false|none|fork仓库拥有者的标识|none|
|»»» index|integer|true|none|合并请求的序号|none|
|»»» status|string|true|none|合并请求的状态|open: 打开的, merged: 合并的, closed: 关闭的|
|»»» issue|object|true|none||none|
|»»»» id|integer|true|none|合并请求下疑修的ID|none|
|»»»» author|object|true|none|合并请求以及疑修的创建着|none|
|»»»»» id|integer|true|none||none|
|»»»»» type|string|true|none||none|
|»»»»» name|string|true|none||none|
|»»»»» login|string|true|none||none|
|»»»»» image_url|string|true|none||none|
|»»»» priority|null|true|none|疑修的优先级|none|
|»»»» version|null|true|none|疑修的里程碑|none|
|»»»» journals_count|integer|true|none|普通评论数量|none|
|»»»» issue_tags|null|true|none|所属标记|none|
|»»» reviewers|[string]|true|none|审查者数组|none|
|»»» journals_count|integer|true|none|审查评论数量|none|
|»» id|integer|false|none|审查ID|none|
|»» commit_id|null|false|none|审查的commit标识|none|
|»» content|string|false|none|审查的内容|none|
|»» status|string|false|none|审查类型|common: 普通, approved: 通过，rejected: 拒绝通过|
|»» created_at|string|false|none|审查创建时间|none|

## POST 创建一个合并请求审查

POST /api/v1/{owner}/{repo}/pulls/{index}/reviews.json

> Body 请求参数

```json
{
  "content": "新建一个审查",
  "commit_id": "e506844b2467ce25a35dd46dad8236a1595a02da",
  "status": "common"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|index|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» content|body|string| 否 | 审查内容|none|
|» commit_id|body|string| 否 | 当前合并请求的提交标识|none|
|» status|body|string| 是 | 审查类型|common 普通 approved 通过 rejected 拒绝|

> 返回示例

> 200 Response

```json
{
  "reviewer": {
    "id": 2,
    "type": "User",
    "name": "heh",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
  },
  "pull_request": {
    "id": 163,
    "head": "master_1",
    "base": "master",
    "is_original": false,
    "index": 1,
    "status": "closed",
    "issue": {
      "id": 260,
      "author": {
        "id": 2,
        "type": "User",
        "name": "heh",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
      },
      "priority": null,
      "version": null,
      "journals_count": 0,
      "issue_tags": null
    },
    "journals_count": 6
  },
  "id": 10,
  "commit_id": "e506844b2467ce25a35dd46dad8236a1595a02da",
  "content": "新建一个审查",
  "status": "common",
  "created_at": "2022-07-26 11:45"
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
|» reviewer|object|true|none|审查者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» pull_request|object|true|none||none|
|»» id|integer|true|none|合并请求ID|none|
|»» head|string|true|none|合并请求源分支|none|
|»» base|string|true|none|合并请求目标分支|none|
|»» is_original|boolean|true|none|合并请求是否从fork仓库所来|none|
|»» index|integer|true|none|合并请求的序号|none|
|»» status|string|true|none|合并请求的状态|open: 打开的, merged: 合并的, closed: 关闭的|
|»» issue|object|true|none||none|
|»»» id|integer|true|none|合并请求下疑修的ID|none|
|»»» author|object|true|none|合并请求以及疑修的创建着|none|
|»»»» id|integer|true|none||none|
|»»»» type|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» priority|null|true|none|疑修的优先级|none|
|»»» version|null|true|none|疑修的里程碑|none|
|»»» journals_count|integer|true|none|普通评论数量|none|
|»»» issue_tags|null|true|none|所属标记|none|
|»» journals_count|integer|true|none|审查评论数量|none|
|» id|integer|true|none|审查ID|none|
|» commit_id|string|true|none|审查的提交标识|none|
|» content|string|true|none|审查的内容|none|
|» status|string|true|none|审查类型|common: 普通, approved: 通过，rejected: 拒绝通过|
|» created_at|string|true|none|审查创建时间|none|

## GET 获取合并请求审查评论列表

GET /api/v1/{owner}/{repo}/pulls/{index}/journals.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|
|keyword|query|string| 否 ||搜索关键词|
|review_id|query|string| 否 ||审查ID|
|need_respond|query|string| 否 ||是否需要回应|
|state|query|string| 否 ||状态, opened: 开启的, resolved: 已解决的, disabled: 无效的|
|parent_id|query|integer| 否 ||父评论ID|
|path|query|string| 否 ||评论文件路径|
|is_full|query|string| 否 ||是否展示全部评论（包括回复）|
|sort_by|query|string| 否 ||排序字段 created_on: 创建时间, updated_on: 更新时间|
|sort_direction|query|string| 否 ||排序类型 desc: 倒序, asc: 正序|

#### 详细说明

**review_id**: 审查ID

**parent_id**: 父评论ID

**path**: 评论文件路径

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "journals": [
    {
      "id": 200,
      "note": "测试评论修改",
      "commit_id": null,
      "line_code": "70eede447ccc01c1902260fd377af5d90be28e0d_0_29",
      "path": "Gemfile.lock",
      "diff": {},
      "need_respond": true,
      "state": "resolved",
      "parent_id": "nil",
      "user": {
        "id": 2,
        "type": "User",
        "name": "heh",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
      },
      "review": {
        "reviewer": {
          "id": 2,
          "type": "User",
          "name": "heh",
          "login": "yystopf",
          "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
        },
        "pull_request": {
          "id": 163,
          "title": "新合并请求1",
          "body": null,
          "head": "master_1",
          "base": "master",
          "is_original": false,
          "index": 1,
          "status": "closed",
          "issue": {
            "id": 260,
            "author": {
              "id": 2,
              "type": "User",
              "name": "heh",
              "login": "yystopf",
              "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
            },
            "priority": null,
            "version": null,
            "journals_count": 0,
            "issue_tags": null
          },
          "reviewers": [],
          "journals_count": 9
        },
        "id": 10,
        "commit_id": "1",
        "content": "新建一个审查",
        "status": "common",
        "created_at": "2022-07-26 11:45"
      },
      "resolveer": {
        "id": 2,
        "type": "User",
        "name": "heh",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
      },
      "resolve_at": "2022-07-27 14:50",
      "created_at": "2022-07-27 14:31",
      "updated_at": "2022-07-27 14:50"
    }
  ]
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
|» total_count|integer|true|none|评论总数|none|
|» journals|[object]|true|none||none|
|»» id|integer|false|none|评论ID|none|
|»» note|string|false|none|评论内容|none|
|»» commit_id|null|false|none|提交标识|none|
|»» line_code|string|false|none|评论行数|none|
|»» path|string|false|none|评论文件路径|none|
|»» diff|object|false|none|评论文件diff内容|none|
|»» need_respond|boolean|false|none|评论是否要回应|none|
|»» state|string|false|none|评论状态|opened: 开启的, resolved: 已解决的, disabled: 无效的|
|»» parent_id|integer|false|none|父评论ID|none|
|»» user|object|false|none|评论创建者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» review|object|false|none|评论所属评审|none|
|»»» reviewer|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» type|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» login|string|true|none||none|
|»»»» image_url|string|true|none||none|
|»»» pull_request|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» title|string|true|none||none|
|»»»» body|null|true|none||none|
|»»»» head|string|true|none||none|
|»»»» base|string|true|none||none|
|»»»» is_original|boolean|true|none||none|
|»»»» index|integer|true|none||none|
|»»»» status|string|true|none||none|
|»»»» issue|object|true|none||none|
|»»»»» id|integer|true|none||none|
|»»»»» author|object|true|none||none|
|»»»»»» id|integer|true|none||none|
|»»»»»» type|string|true|none||none|
|»»»»»» name|string|true|none||none|
|»»»»»» login|string|true|none||none|
|»»»»»» image_url|string|true|none||none|
|»»»»» priority|null|true|none||none|
|»»»»» version|null|true|none||none|
|»»»»» journals_count|integer|true|none||none|
|»»»»» issue_tags|null|true|none||none|
|»»»» reviewers|[string]|true|none||none|
|»»»» journals_count|integer|true|none||none|
|»»» id|integer|true|none||none|
|»»» commit_id|string|true|none||none|
|»»» content|string|true|none||none|
|»»» status|string|true|none||none|
|»»» created_at|string|true|none||none|
|»» resolveer|object|false|none|评论解决者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» resolve_at|string|false|none|评论解决时间|none|
|»» created_at|string|false|none|评论创建时间|none|
|»» updated_at|string|false|none|评论更新时间|none|

## POST 创建一个合并请求审查评论

POST /api/v1/{owner}/{repo}/pulls/{index}/journals.json

> Body 请求参数

```json
{
  "type": "problem",
  "note": "测试评论",
  "review_id": "10",
  "line_code": "70eede447ccc01c1902260fd377af5d90be28e0d_0_29",
  "commit_id": "70eede447ccc01c1902260fd377af5d90be28e0d",
  "path": "Gemfile.lock",
  "diff": {
    "name": "README.md",
    "oldname": "README.md",
    "addition": 1,
    "deletion": 2,
    "type": 2,
    "is_created": false,
    "is_deleted": false,
    "is_bin": false,
    "is_lfs_file": false,
    "is_renamed": false,
    "is_ambiguous": false,
    "is_submodule": false,
    "sections": [
      {
        "file_name": "README.md",
        "name": "",
        "lines": [
          {
            "left_index": 0,
            "right_index": 0,
            "match": 0,
            "type": 4,
            "content": "@@ -1,2 +1 @@",
            "section_path": "README.md",
            "section_last_left_index": 0,
            "section_last_right_index": 0,
            "section_left_index": 1,
            "section_right_index": 1,
            "section_left_hunk_size": 2,
            "section_right_hunk_size": 0
          },
          {
            "left_index": 1,
            "right_index": 0,
            "match": 3,
            "type": 3,
            "content": "-# ceshi_commit"
          },
          {
            "left_index": 2,
            "right_index": 0,
            "match": -1,
            "type": 3,
            "content": "-"
          },
          {
            "left_index": 0,
            "right_index": 1,
            "match": 1,
            "type": 2,
            "content": "+adsa"
          }
        ]
      },
      {
        "file_name": "README.md",
        "name": "",
        "lines": [
          {
            "left_index": 0,
            "right_index": 0,
            "match": 0,
            "type": 4,
            "content": " ",
            "section_path": "README.md",
            "section_last_left_index": 0,
            "section_last_right_index": 1,
            "section_left_index": 3,
            "section_right_index": 2,
            "section_left_hunk_size": 0,
            "section_right_hunk_size": 0
          }
        ]
      }
    ],
    "is_incomplete": false,
    "is_incomplete_line_too_long": false,
    "is_protected": false
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|
|body|body|object| 否 ||none|
|» type|body|string| 是 | 评论类型|comment: 普通, problem: 需要回应的评论|
|» note|body|string| 是 | 评论内容|none|
|» review_id|body|string| 是 | 审查ID|none|
|» line_code|body|string| 是 | 行号|none|
|» commit_id|body|string| 是 | 提交标识|none|
|» path|body|string| 是 | 文件路径|none|
|» parent_id|body|integer| 是 | 父评论ID|none|
|» diff|body|object| 是 | 文件diff内容|none|
|»» name|body|string| 是 ||none|
|»» oldname|body|string| 是 ||none|
|»» addition|body|integer| 是 ||none|
|»» deletion|body|integer| 是 ||none|
|»» type|body|integer| 是 ||none|
|»» is_created|body|boolean| 是 ||none|
|»» is_deleted|body|boolean| 是 ||none|
|»» is_bin|body|boolean| 是 ||none|
|»» is_lfs_file|body|boolean| 是 ||none|
|»» is_renamed|body|boolean| 是 ||none|
|»» is_ambiguous|body|boolean| 是 ||none|
|»» is_submodule|body|boolean| 是 ||none|
|»» sections|body|[object]| 是 ||none|
|»»» file_name|body|string| 是 ||none|
|»»» name|body|string| 是 ||none|
|»»» lines|body|[object]| 是 ||none|
|»»»» left_index|body|integer| 是 ||none|
|»»»» right_index|body|integer| 是 ||none|
|»»»» match|body|integer| 是 ||none|
|»»»» type|body|integer| 是 ||none|
|»»»» content|body|string| 是 ||none|
|»»»» section_path|body|string| 是 ||none|
|»»»» section_last_left_index|body|integer| 是 ||none|
|»»»» section_last_right_index|body|integer| 是 ||none|
|»»»» section_left_index|body|integer| 是 ||none|
|»»»» section_right_index|body|integer| 是 ||none|
|»»»» section_left_hunk_size|body|integer| 是 ||none|
|»»»» section_right_hunk_size|body|integer| 是 ||none|
|»» is_incomplete|body|boolean| 是 ||none|
|»» is_incomplete_line_too_long|body|boolean| 是 ||none|
|»» is_protected|body|boolean| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "id": 200,
  "note": "测试评论修改",
  "commit_id": null,
  "line_code": "70eede447ccc01c1902260fd377af5d90be28e0d_0_29",
  "path": "Gemfile.lock",
  "diff": {},
  "need_respond": true,
  "state": "resolved",
  "user": {
    "id": 2,
    "type": "User",
    "name": "heh",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
  },
  "review": {
    "reviewer": {
      "id": 2,
      "type": "User",
      "name": "heh",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
    },
    "pull_request": {
      "id": 163,
      "title": "新合并请求1",
      "body": null,
      "head": "master_1",
      "base": "master",
      "is_original": false,
      "index": 1,
      "status": "closed",
      "issue": {
        "id": 260,
        "author": {
          "id": 2,
          "type": "User",
          "name": "heh",
          "login": "yystopf",
          "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
        },
        "priority": null,
        "version": null,
        "journals_count": 0,
        "issue_tags": null
      },
      "reviewers": [],
      "journals_count": 9
    },
    "id": 10,
    "commit_id": "1",
    "content": "新建一个审查",
    "status": "common",
    "created_at": "2022-07-26 11:45"
  },
  "resolveer": {
    "id": 2,
    "type": "User",
    "name": "heh",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
  },
  "resolve_at": "2022-07-27 14:50",
  "created_at": "2022-07-27 14:31",
  "updated_at": "2022-07-27 14:50"
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
|» id|integer|true|none|评论ID|none|
|» note|string|true|none|评论内容|none|
|» commit_id|null|true|none|提交标识|none|
|» line_code|string|true|none|评论行数|none|
|» path|string|true|none|评论文件路径|none|
|» diff|object|true|none|评论文件diff内容|none|
|» need_respond|boolean|true|none|评论是否要回应|none|
|» state|string|true|none|评论状态|opened: 开启的, resolved: 已解决的, disabled: 无效的|
|» parent_id|integer|true|none|父评论ID|none|
|» user|object|true|none|评论创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» review|object|true|none|评论所属评审|none|
|»» reviewer|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» pull_request|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» title|string|true|none||none|
|»»» body|null|true|none||none|
|»»» head|string|true|none||none|
|»»» base|string|true|none||none|
|»»» is_original|boolean|true|none||none|
|»»» index|integer|true|none||none|
|»»» status|string|true|none||none|
|»»» issue|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» author|object|true|none||none|
|»»»»» id|integer|true|none||none|
|»»»»» type|string|true|none||none|
|»»»»» name|string|true|none||none|
|»»»»» login|string|true|none||none|
|»»»»» image_url|string|true|none||none|
|»»»» priority|null|true|none||none|
|»»»» version|null|true|none||none|
|»»»» journals_count|integer|true|none||none|
|»»»» issue_tags|null|true|none||none|
|»»» reviewers|[string]|true|none||none|
|»»» journals_count|integer|true|none||none|
|»» id|integer|true|none||none|
|»» commit_id|string|true|none||none|
|»» content|string|true|none||none|
|»» status|string|true|none||none|
|»» created_at|string|true|none||none|
|» resolveer|object|true|none|评论解决者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» resolve_at|string|true|none|评论解决时间|none|
|» created_at|string|true|none|评论创建时间|none|
|» updated_at|string|true|none|评论更新时间|none|

## PUT 修改一个合并请求审查评论

PUT /api/v1/{owner}/{repo}/pulls/{index}/journals/{id}.json

> Body 请求参数

```json
{
  "note": "测试评论",
  "commit_id": "70eede447ccc01c1902260fd377af5d90be28e0d",
  "state": "resolved"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|
|id|path|string| 是 ||评论ID|
|body|body|object| 否 ||none|
|» note|body|string| 是 | 评论内容|none|
|» commit_id|body|string| 是 | 提交标识|none|
|» state|body|string| 是 | 评论状态|opened: 开启的, resolved: 已解决的, disabled: 无效的|

#### 详细说明

**index**: 合并请求序号

**id**: 评论ID

> 返回示例

> 200 Response

```json
{
  "id": 200,
  "note": "测试评论修改",
  "commit_id": null,
  "line_code": "70eede447ccc01c1902260fd377af5d90be28e0d_0_29",
  "path": "Gemfile.lock",
  "diff": {},
  "need_respond": true,
  "state": "resolved",
  "user": {
    "id": 2,
    "type": "User",
    "name": "heh",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
  },
  "review": {
    "reviewer": {
      "id": 2,
      "type": "User",
      "name": "heh",
      "login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
    },
    "pull_request": {
      "id": 163,
      "title": "新合并请求1",
      "body": null,
      "head": "master_1",
      "base": "master",
      "is_original": false,
      "index": 1,
      "status": "closed",
      "issue": {
        "id": 260,
        "author": {
          "id": 2,
          "type": "User",
          "name": "heh",
          "login": "yystopf",
          "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
        },
        "priority": null,
        "version": null,
        "journals_count": 0,
        "issue_tags": null
      },
      "reviewers": [],
      "journals_count": 9
    },
    "id": 10,
    "commit_id": "1",
    "content": "新建一个审查",
    "status": "common",
    "created_at": "2022-07-26 11:45"
  },
  "resolveer": {
    "id": 2,
    "type": "User",
    "name": "heh",
    "login": "yystopf",
    "image_url": "system/lets/letter_avatars/2/H/188_239_142/120.png"
  },
  "resolve_at": "2022-07-27 14:50",
  "created_at": "2022-07-27 14:31",
  "updated_at": "2022-07-27 14:50"
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
|» id|integer|true|none|评论ID|none|
|» note|string|true|none|评论内容|none|
|» commit_id|null|true|none|评论标识|none|
|» line_code|string|true|none|评论行数|none|
|» path|string|true|none|评论文件路径|none|
|» diff|object|true|none|评论文件diff内容|none|
|» need_respond|boolean|true|none|评论是否要回应|none|
|» state|string|true|none|评论状态|opened: 开启的, resolved: 已解决的, disabled: 无效的|
|» parent_id|integer|true|none|父评论ID|none|
|» user|object|true|none|评论创建者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» review|object|true|none|评论所属评审|none|
|»» reviewer|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|
|»» pull_request|object|true|none||none|
|»»» id|integer|true|none||none|
|»»» title|string|true|none||none|
|»»» body|null|true|none||none|
|»»» head|string|true|none||none|
|»»» base|string|true|none||none|
|»»» is_original|boolean|true|none||none|
|»»» index|integer|true|none||none|
|»»» status|string|true|none||none|
|»»» issue|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» author|object|true|none||none|
|»»»»» id|integer|true|none||none|
|»»»»» type|string|true|none||none|
|»»»»» name|string|true|none||none|
|»»»»» login|string|true|none||none|
|»»»»» image_url|string|true|none||none|
|»»»» priority|null|true|none||none|
|»»»» version|null|true|none||none|
|»»»» journals_count|integer|true|none||none|
|»»»» issue_tags|null|true|none||none|
|»»» reviewers|[string]|true|none||none|
|»»» journals_count|integer|true|none||none|
|»» id|integer|true|none||none|
|»» commit_id|string|true|none||none|
|»» content|string|true|none||none|
|»» status|string|true|none||none|
|»» created_at|string|true|none||none|
|» resolveer|object|true|none|评论解决者|none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» login|string|true|none||none|
|»» image_url|string|true|none||none|
|» resolve_at|string|true|none|评论解决时间|none|
|» created_at|string|true|none|评论创建时间|none|
|» updated_at|string|true|none|评论更新时间|none|

## DELETE 删除一个合并请求审查评论

DELETE /api/v1/{owner}/{repo}/pulls/{index}/journals/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|index|path|string| 是 ||合并请求序号|
|id|path|string| 是 ||评论ID|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## POST 创建一个合并请求评论

POST /api/issues/{id}/journals.json

> Body 请求参数

```json
{
  "content": "123123",
  "issue_id": 1,
  "receivers_login": []
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|id|path|string| 是 ||合并请求下issue的id|
|body|body|object| 否 ||none|
|» content|body|string| 是 | 评论内容|none|
|» issue_id|body|integer| 是 | 合并请求下issue的id|none|
|» parent_id|body|integer| 否 | 回复评论ID|none|
|» receivers_login|body|[string]| 否 | @用户标识数组|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "评论成功",
  "id": 3
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
|» message|string|true|none||none|
|» id|integer|true|none||none|

## GET 获取一个合并请求评论列表

GET /api/issues/{id}/journals.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|id|path|string| 是 ||合并请求下issue的id|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "响应成功",
  "limit": 10,
  "journals_count": 2,
  "journals_total_count": 2,
  "issue_journals": [
    {
      "id": 423059,
      "user_name": "肖小琼",
      "user_login": "xxq250",
      "user_picture": "images/avatars/User/86107?t=1668998259",
      "is_journal_detail": false,
      "content": "https://gitlink.org.cn/api/issues/128470/journals.json?id=128470",
      "journal_details": [],
      "format_time": "2025-06-17 15:45",
      "created_at": "2小时前",
      "attachments": []
    },
    {
      "id": 423056,
      "user_name": "wangjing22e",
      "user_login": "wangjing22e",
      "user_picture": "system/lets/letter_avatars/2/W/61_162_123/120.png",
      "is_journal_detail": false,
      "content": "撒打发十分",
      "journal_details": [],
      "format_time": "2025-06-17 15:09",
      "created_at": "2小时前",
      "attachments": []
    }
  ]
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
|» status|integer|true|none|状态|none|
|» message|string|true|none||none|
|» limit|integer|true|none|分页参数|none|
|» journals_count|integer|true|none|评论总数|none|
|» journals_total_count|integer|true|none|评论总数|none|
|» issue_journals|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» user_name|string|true|none||none|
|»» user_login|string|true|none||none|
|»» user_picture|string|true|none||none|
|»» is_journal_detail|boolean|true|none||none|
|»» content|string|true|none||none|
|»» journal_details|[string]|true|none||none|
|»» format_time|string|true|none||none|
|»» created_at|string|true|none||none|
|»» attachments|[string]|true|none||none|

## DELETE 删除一个合并请求评论

DELETE /api/issues/{issue_id}/journals/{id}.json

> Body 请求参数

```json
{
  "issue_id": 1,
  "id": 4
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|issue_id|path|string| 是 ||合并请求下issue的ID|
|id|path|string| 是 ||评论ID|
|body|body|object| 否 ||none|
|» issue_id|body|integer| 是 | 合并请求下issue的ID|none|
|» id|body|integer| 是 | 评论ID|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "string"
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
|» message|string|true|none||none|

# 发行版

## POST 创建发行版

POST /api/{owner}/{repo}/releases.json

> Body 请求参数

```json
{
  "attachment_ids": [],
  "body": "dfs",
  "name": "ceshi",
  "tag_name": "v1.0.1",
  "target_commitish": "",
  "draft": true,
  "prerelease": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» attachment_ids|body|[string]| 否 | 附件ID数组|none|
|» body|body|string| 是 | 发行版描述|none|
|» name|body|string| 是 | 发行版标题|none|
|» tag_name|body|string| 是 | 标签|none|
|» target_commitish|body|string| 否 | 分支|默认为master|
|» draft|body|boolean| 是 | 是否为草稿|none|
|» prerelease|body|boolean| 否 | 是否为预发行版本|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "发布成功"
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
|» message|string|true|none||none|

## GET 获取发行版列表

GET /api/{owner}/{repo}/releases.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "响应成功",
  "releases": [
    {
      "version_id": 2,
      "id": "392",
      "tag_name": "v1.0",
      "target_commitish": "master",
      "name": "ceshi",
      "sha": "96dc82d07e05327fc908aaaa127ba45f02091d85",
      "body": null,
      "url": "http://127.0.0.1:10082/api/v1/repos/yystopfceshi/testdevops/releases/392",
      "tarball_url": "http://localhost:3000/api/yystopfceshi/testdevops/archive/v1.0.tar.gz",
      "zipball_url": "http://localhost:3000/api/yystopfceshi/testdevops/archive/v1.0.zip",
      "draft": "稳定",
      "created_at": "2023-11-22 10:21",
      "published_at": "2023-11-22 10:21",
      "user_name": "咸蛋黄土豆丝xxx",
      "user_login": "yystopf",
      "image_url": "system/lets/letter_avatars/2/X/230_139_26/120.png",
      "attachments": []
    }
  ],
  "total_count": 1
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
|» message|string|true|none||none|
|» releases|[object]|true|none||none|
|»» version_id|integer|false|none|版本ID|none|
|»» id|string|false|none|ID|none|
|»» tag_name|string|false|none|标签名称|none|
|»» target_commitish|string|false|none|分支名称|none|
|»» name|string|false|none|版本名称|none|
|»» sha|string|false|none|版本commit标识|none|
|»» body|string|false|none|版本描述|none|
|»» url|string|false|none|http clone地址|none|
|»» tarball_url|string|false|none|tar下载地址|none|
|»» zipball_url|string|false|none|zip下载地址|none|
|»» draft|string|false|none|是否为草稿|none|
|»» created_at|string|false|none|创建时间|none|
|»» published_at|string|false|none|发布时间|none|
|»» user_name|string|false|none|创建者昵称|none|
|»» user_login|string|false|none|创建者标识|none|
|»» image_url|string|false|none|创建者头像|none|
|»» attachments|[object]|false|none|附件|none|
|»»» id|integer|true|none|附件ID|none|
|»»» title|string|true|none|附件标题|none|
|»»» filesize|string|true|none|附件大小|none|
|»»» description|string|true|none|附件描述|none|
|»»» is_pdf|boolean|true|none|附件是否为pdf|none|
|»»» url|string|true|none|附件下载地址|none|
|» total_count|integer|true|none|版本总数|none|

## GET 编辑发行版

GET /api/{owner}/{repo}/releases/{id}/edit.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "id": 593,
  "name": "222",
  "body": "22",
  "tag_name": "v1.0",
  "target_commitish": "master",
  "draft": false,
  "prerelease": false,
  "version_gid": "1619232",
  "attachments": [
    {
      "id": 441826,
      "title": "http-proxy.conf",
      "filesize": "154 字节",
      "description": "",
      "is_pdf": false,
      "url": "/api/attachments/441826"
    },
    {
      "id": 441827,
      "title": "组 9.png",
      "filesize": "342.0 KB",
      "description": "",
      "is_pdf": false,
      "url": "/api/attachments/441827"
    }
  ]
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
|» id|integer|true|none|ID|none|
|» name|string|true|none|名称|none|
|» body|string|true|none|描述|none|
|» tag_name|string|true|none|标签名称|none|
|» target_commitish|string|true|none|分支名称|none|
|» draft|boolean|true|none|是否为草稿|none|
|» prerelease|boolean|true|none|是否为预发行版本|none|
|» version_gid|string|true|none|gitea ID|none|
|» attachments|[object]|true|none|附件|none|
|»» id|integer|true|none|附件ID|none|
|»» title|string|true|none|标题|none|
|»» filesize|string|true|none|文件大小|none|
|»» description|string|true|none|文件描述|none|
|»» is_pdf|boolean|true|none|是否为pdf|none|
|»» url|string|true|none|下载地址|none|

## PUT 更新发行版

PUT /api/{owner}/{repo}/releases/{id}.json

> Body 请求参数

```json
{
  "attachment_ids": [],
  "body": "dfs",
  "name": "ceshi",
  "tag_name": "v1.0.1",
  "target_commitish": "",
  "draft": true,
  "prerelease": true
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» attachment_ids|body|[string]| 否 | 附件ID数组|none|
|» body|body|string| 是 | 发行版描述|none|
|» name|body|string| 是 | 发行版标题|none|
|» tag_name|body|string| 是 | 标签|none|
|» target_commitish|body|string| 否 | 分支|默认为master|
|» draft|body|boolean| 否 | 是否为草稿|none|
|» prerelease|body|boolean| 否 | 是否为预发行版本|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "更新成功"
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
|» message|string|true|none||none|

## DELETE 删除发行版

DELETE /api/{owner}/{repo}/releases/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "删除成功"
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
|» message|string|true|none||none|

## GET 查看发行版

GET /api/{owner}/{repo}/releases/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "version_id": 593,
  "id": "1619232",
  "tag_name": "v1.0",
  "target_commitish": "master",
  "name": "222",
  "sha": "ef44144e8f53fd19d38c8a7d1400313db73881e4",
  "body": "22",
  "url": "https://gitlink.org.cn/api/v1/repos/yystopf/data_competition/releases/1619232",
  "tarball_url": "https://www.gitlink.org.cn/api/yystopf/data_competition/archive/v1.0.tar.gz",
  "zipball_url": "https://www.gitlink.org.cn/api/yystopf/data_competition/archive/v1.0.zip",
  "draft": "稳定",
  "created_at": "2023-11-27 11:47",
  "published_at": "2023-11-27 11:47",
  "user_name": "何慧",
  "user_login": "yystopf",
  "image_url": "images/avatars/User/85890?t=1641550911",
  "attachments": [
    {
      "id": 441826,
      "title": "http-proxy.conf",
      "filesize": "154 字节",
      "description": "",
      "is_pdf": false,
      "url": "/api/attachments/441826"
    },
    {
      "id": 441827,
      "title": "组 9.png",
      "filesize": "342.0 KB",
      "description": "",
      "is_pdf": false,
      "url": "/api/attachments/441827"
    }
  ]
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
|» version_id|integer|true|none|ID|none|
|» id|string|true|none||none|
|» tag_name|string|true|none|标签名称|none|
|» target_commitish|string|true|none|分支名称|none|
|» name|string|true|none|名称|none|
|» sha|string|true|none|标识|none|
|» body|string|true|none|描述|none|
|» url|string|true|none|下载地址|none|
|» tarball_url|string|true|none|tar下载地址|none|
|» zipball_url|string|true|none|zip下载地址|none|
|» draft|string|true|none|版本类型|有以下三种 草稿、预发行、稳定|
|» created_at|string|true|none|创建时间|none|
|» published_at|string|true|none|发布时间|none|
|» user_name|string|true|none|创建者昵称|none|
|» user_login|string|true|none|创建者标识|none|
|» image_url|string|true|none|创建者头像|none|
|» attachments|[object]|true|none||none|
|»» id|integer|true|none|附件ID|none|
|»» title|string|true|none|附件标题|none|
|»» filesize|string|true|none|附件文件大小|none|
|»» description|string|true|none|附件描述|none|
|»» is_pdf|boolean|true|none|附件是否为pdf|none|
|»» url|string|true|none|附件链接|none|

# 数据集

## POST 创建数据集

POST /api/v1/{owner}/{repo}/dataset

> Body 请求参数

```json
{
  "title": "标题",
  "description": "描述"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|body|body|object| 否 ||none|
|» title|body|string| 是 | 标题|none|
|» description|body|string| 是 | 描述|none|
|» license_id|body|integer| 是 | 许可证ID|none|
|» paper_content|body|string| 是 | 研究论文|none|

> 返回示例

```json
{
  "status": 0,
  "message": "success"
}
```

```json
{
  "status": -1,
  "message": "该项目下已存在数据集！"
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
|» message|string|true|none||none|

## PUT 更新数据集

PUT /api/v1/{owner}/{repo}/dataset

> Body 请求参数

```json
{
  "title": "标题",
  "description": "描述"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|body|body|object| 否 ||none|
|» title|body|string| 是 | 标题|none|
|» description|body|string| 是 | 描述|none|
|» license_id|body|integer| 是 | 许可证ID|none|
|» paper_content|body|string| 是 | 研究论文内容|none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 数据集详情+列表

GET /api/v1/{owner}/{repo}/dataset

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|page|query|integer| 否 ||分页页码|
|limit|query|integer| 否 ||每页个数|

> 返回示例

> 200 Response

```json
{
  "id": 2,
  "title": "标题1",
  "description": "描述1",
  "attachments": [
    {
      "id": "2e78064e-48f9-4960-bf55-c4534a26abba",
      "title": "操作系统大赛用户信息 (4).csv",
      "description": "对方房东",
      "filesize": "98 字节",
      "is_pdf": false,
      "url": "/api/attachments/2e78064e-48f9-4960-bf55-c4534a26abba",
      "created_on": "2024-04-01 16:53:34",
      "content_type": "text/csv",
      "creator": {
        "id": 102,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/Y/241_125_89/120.png"
      }
    },
    {
      "id": "e16cf5da-3552-452b-8699-c7a8e18e8a9e",
      "title": "操作系统大赛用户信息 (4).csv",
      "description": "对方房东333333",
      "filesize": "98 字节",
      "is_pdf": false,
      "url": "/api/attachments/e16cf5da-3552-452b-8699-c7a8e18e8a9e",
      "created_on": "2024-04-01 17:15:28",
      "content_type": "text/csv",
      "creator": {
        "id": 102,
        "type": "User",
        "name": "咸蛋黄土豆丝",
        "login": "yystopf",
        "image_url": "system/lets/letter_avatars/2/Y/241_125_89/120.png"
      }
    }
  ]
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
|» id|integer|true|none|数据集ID|none|
|» title|string|true|none|数据集标题|none|
|» description|string|true|none|数据集描述|none|
|» attachment_total_count|integer|true|none|文件总数|none|
|» attachments|[object]|true|none||none|
|»» id|string|true|none|附件UUID|none|
|»» title|string|true|none|附件标题|none|
|»» description|string|true|none|附件描述|none|
|»» filesize|string|true|none|附件大小|none|
|»» is_pdf|boolean|true|none|附件是否为pdf|none|
|»» url|string|true|none|附件下载地址|none|
|»» created_on|string|true|none|附件上传时间|none|
|»» content_type|string|true|none|附件类型|none|
|»» creator|object|true|none|附件创建者|none|
|»»» id|integer|true|none||none|
|»»» type|string|true|none||none|
|»»» name|string|true|none||none|
|»»» login|string|true|none||none|
|»»» image_url|string|true|none||none|

## GET 查询数据集列表

GET /api/v1/project_datasets.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|ids|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "project_datasets": [
    {
      "id": 3,
      "title": "是许方层",
      "description": "府就却法易石儿置断每合体如这。车史二图个器天书家风接八者证要广民。往儿确七明资写十众工因拉式斗积南比我。青响则会活称报事住极国群原公改油般。低往书两空质和段查活结水制装。",
      "paper_content": "laboris eu commodo labore ut",
      "project": {
        "type": "common",
        "id": 1,
        "description": null,
        "forked_count": 0,
        "forked_from_project_id": null,
        "identifier": "ceshi",
        "issues_count": 2,
        "pull_requests_count": 2,
        "invite_code": "mPL1Ef",
        "website": null,
        "platform": "forge",
        "name": "常见资按",
        "open_devops": false,
        "praises_count": 0,
        "is_public": false,
        "status": 1,
        "watchers_count": 0,
        "ignore_id": null,
        "license_id": null,
        "project_category_id": null,
        "project_language_id": null
      },
      "license": {
        "name": "BSD-4-Clause",
        "content": "Copyright (c) <year> <owner> . All rights reserved.\n\nRedistribution and use in source and binary forms, with or without modification,\nare permitted provided that the following conditions are met:\n\n1. Redistributions of source code must retain the above copyright notice,\nthis list of conditions and the following disclaimer.\n\n2. Redistributions in binary form must reproduce the above copyright notice,\nthis list of conditions and the following disclaimer in the documentation\nand/or other materials provided with the distribution.\n\n3. All advertising materials mentioning features or use of this software must\ndisplay the following acknowledgement:\n\n   This product includes software developed by the organization .\n\n4. Neither the name of the copyright holder nor the names of its contributors\nmay be used to endorse or promote products derived from this software without\nspecific prior written permission.\n\nTHIS SOFTWARE IS PROVIDED BY COPYRIGHT HOLDER \"AS IS\" AND ANY EXPRESS OR IMPLIED\nWARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY\nAND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL COPYRIGHT\nHOLDER BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY,\nOR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE\nGOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)\nHOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT\nLIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY\nOUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH\nDAMAGE.\n"
      }
    }
  ]
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
|» total_count|integer|true|none|总数|none|
|» project_datasets|[object]|true|none||none|
|»» id|integer|false|none|数据集ID|none|
|»» title|string|false|none|数据集标题|none|
|»» description|string|false|none|数据集描述|none|
|»» paper_content|string|false|none|数据集论文内容|none|
|»» project|object|false|none|项目信息|none|
|»»» type|string|true|none||none|
|»»» id|integer|true|none||none|
|»»» description|null|true|none||none|
|»»» forked_count|integer|true|none||none|
|»»» forked_from_project_id|null|true|none||none|
|»»» identifier|string|true|none||none|
|»»» issues_count|integer|true|none||none|
|»»» pull_requests_count|integer|true|none||none|
|»»» invite_code|string|true|none||none|
|»»» website|null|true|none||none|
|»»» platform|string|true|none||none|
|»»» name|string|true|none||none|
|»»» open_devops|boolean|true|none||none|
|»»» praises_count|integer|true|none||none|
|»»» is_public|boolean|true|none||none|
|»»» status|integer|true|none||none|
|»»» watchers_count|integer|true|none||none|
|»»» ignore_id|null|true|none||none|
|»»» license_id|null|true|none||none|
|»»» project_category_id|null|true|none||none|
|»»» project_language_id|null|true|none||none|
|»» license|object|false|none|许可证|none|
|»»» name|string|true|none|名称|none|
|»»» content|string|true|none|内容|none|

## DELETE 删除数据集附件

DELETE /api/attachments/{uuid}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|uuid|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "删除成功"
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
|» message|string|true|none||none|

# WebHook

## GET Webhook列表

GET /api/v1/{owner}/{repo}/webhooks.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|

> 返回示例

> 200 Response

```json
{
  "total_count": 1,
  "webhooks": [
    {
      "id": 4161,
      "url": "http://gxdaw.tf/snhcg",
      "http_method": "GET",
      "is_active": true,
      "type": "gitea",
      "last_status": "waiting",
      "create_time": "2023-04-03 19:01:29"
    }
  ]
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
|» total_count|integer|true|none||none|
|» webhooks|[object]|true|none||none|
|»» id|integer|false|none|ID|none|
|»» url|string|false|none|请求地址|none|
|»» http_method|string|false|none|请求方式|none|
|»» is_active|boolean|false|none|是否激活|none|
|»» type|string|false|none|类型|none|
|»» last_status|string|false|none|最后一次推送的状态|none|
|»» create_time|string|false|none|创建时间|none|

#### 枚举值

|属性|值|
|---|---|
|http_method|POST|
|http_method|GET|
|type|gitea|
|type|slack|
|type|discord|
|type|dingtalk|
|type|telegram|
|type|msteams|
|type|feishu|
|type|matrix|
|type|jianmu|
|type|softbot|
|last_status|waiting|
|last_status|succeed|
|last_status|fail|

## POST Webhook创建

POST /api/v1/{owner}/{repo}/webhooks.json

> Body 请求参数

```json
{
  "type": "gitea",
  "active": true,
  "content_type": "json",
  "http_method": "GET",
  "secret": "string",
  "url": "string",
  "branch_filter": "string",
  "events": [
    "push"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» type|body|string| 否 | 类型|none|
|» active|body|boolean| 是 | 是否开启|none|
|» content_type|body|string| 是 | POST content_type|none|
|» http_method|body|string| 是 | 请求方式|none|
|» secret|body|string| 否 | 密钥|none|
|» url|body|string| 是 | 请求地址|none|
|» branch_filter|body|string| 是 | 分支白名单|推送、创建，删除分支事件的分支白名单，使用 glob 模式匹配指定。若为空或 *，则将报告所有分支的事件。语法文档见github.com/gobwas/glob。示例：master,{master,release*}。|
|» events|body|[string]| 是 | 事件|none|

#### 枚举值

|属性|值|
|---|---|
|» type|gitea|
|» type|slack|
|» type|discord|
|» type|dingtalk|
|» type|telegram|
|» type|msteams|
|» type|feishu|
|» type|matrix|
|» type|jianmu|
|» type|softbot|
|» content_type|json|
|» content_type|form|
|» http_method|GET|
|» http_method|POST|
|» events|push|
|» events|create|
|» events|delete|
|» events|issues_only|
|» events|issue_assign|
|» events|issue_label|
|» events|issue_comment|
|» events|pull_request_only|
|» events|pull_request_assign|
|» events|pull_request_comment|

> 返回示例

> 200 Response

```json
{
  "id": 4167,
  "type": "gitea",
  "content_type": "json",
  "http_method": null,
  "url": "http://gxdaw.tf/snhcg",
  "events": [
    "delete",
    "push",
    "issues",
    "pull_request"
  ],
  "active": true,
  "branch_filter": null,
  "created_at": "2023-04-04 11:43"
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
|» id|integer|true|none|ID|none|
|» type|string|true|none|类型|none|
|» content_type|string|true|none|POST content_type|none|
|» http_method|null|false|none|请求方式|none|
|» url|string|true|none|请求地址|none|
|» events|[string]|true|none|事件|none|
|» active|boolean|true|none|是否开启|none|
|» branch_filter|null|false|none|分支白名单|none|
|» created_at|string|true|none|创建时间|none|

## DELETE Webhook删除

DELETE /api/v1/{owner}/{repo}/webhooks/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|id|path|string| 是 ||webhook的ID|
|uid|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET Webhook详情

GET /api/v1/{owner}/{repo}/webhooks/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者标识|
|repo|path|string| 是 ||项目标识|
|id|path|string| 是 ||webhook的ID|
|uid|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## PUT Webhook更新

PUT /api/v1/{owner}/{repo}/webhooks/{id}.json

> Body 请求参数

```json
{
  "active": true,
  "content_type": "json",
  "http_method": "GET",
  "secret": "123456",
  "url": "http://localhost:10000",
  "branch_filter": "*",
  "events": [
    "push"
  ]
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|id|path|string| 是 ||webhook的ID|
|body|body|object| 否 ||none|
|» active|body|boolean| 是 | 是否激活|none|
|» content_type|body|string| 是 | POST Content Type|none|
|» http_method|body|string| 是 | http方法, POST和GET|none|
|» secret|body|string| 是 | 密钥文本|none|
|» url|body|string| 是 | 目标url|none|
|» branch_filter|body|string| 是 | 分支过滤|none|
|» events|body|[string]| 是 | 触发事件|none|

#### 枚举值

|属性|值|
|---|---|
|» content_type|json|
|» content_type|form|
|» http_method|POST|
|» http_method|GET|
|» events|push|
|» events|create|
|» events|delete|
|» events|issues_only|
|» events|issue_assign|
|» events|issue_label|
|» events|issue_comment|
|» events|pull_request_only|
|» events|pull_request_assign|
|» events|pull_request_comment|

> 返回示例

> 200 Response

```json
{
  "id": 68,
  "content_type": "json",
  "http_method": "GET",
  "url": "http://127.0.0.1:3000",
  "events": [
    "create",
    "delete",
    "push",
    "pull_request",
    "pull_request_assign",
    "pull_request_review_approved",
    "pull_request_review_rejected"
  ],
  "active": true,
  "branch_filter": "*",
  "created_at": "2022-06-23 15:52"
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
|» id|integer|true|none||none|
|» content_type|string|true|none|POST Content Type|none|
|» http_method|string|true|none|http方法, POST和GET|none|
|» url|string|true|none|目标url|none|
|» events|[string]|true|none|触发事件|none|
|» active|boolean|true|none|是否激活|none|
|» branch_filter|string|true|none|分支过滤|none|
|» created_at|string|true|none|创建时间|none|

## GET Webhook历史推送列表

GET /api/v1/{owner}/{repo}/webhooks/{id}/hooktasks.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 6,
  "hooktasks": [
    {
      "id": 20,
      "type": "gitea",
      "uuid": "99aa2c23-6884-4c44-9020-5469320aa408",
      "is_succeed": true,
      "is_delivered": true,
      "payload_content": {
        "secret": "123456",
        "ref": "refs/heads/master",
        "before": "feb48e31362787a7620b53d4df3c4effddbb6f0b",
        "after": "feb48e31362787a7620b53d4df3c4effddbb6f0b",
        "compare_url": "",
        "commits": [
          {
            "id": "feb48e31362787a7620b53d4df3c4effddbb6f0b",
            "message": "fix\n",
            "url": "http://localhost:10081/yystopf/ceshi/commit/feb48e31362787a7620b53d4df3c4effddbb6f0b",
            "author": {
              "name": "viletyy",
              "email": "yystopf@163.com",
              "username": "root"
            },
            "committer": {
              "name": "viletyy",
              "email": "yystopf@163.com",
              "username": "root"
            },
            "verification": {
              "verified": false,
              "reason": "gpg.error.not_signed_commit",
              "signature": "",
              "signer": null,
              "payload": ""
            },
            "timestamp": "2021-07-26T13:52:13+08:00",
            "added": null,
            "removed": null,
            "modified": null
          }
        ],
        "head_commit": null,
        "repository": {
          "id": 2,
          "owner": {
            "id": 3,
            "login": "yystopf",
            "full_name": "",
            "email": "yystopf@forge.com",
            "avatar_url": "http://localhost:10081/user/avatar/yystopf/-1",
            "language": "zh-CN",
            "is_admin": true,
            "last_login": "2021-07-21T18:38:21+08:00",
            "created": "2021-06-03T14:50:25+08:00",
            "username": "yystopf"
          },
          "name": "ceshi",
          "full_name": "yystopf/ceshi",
          "description": "",
          "empty": false,
          "private": false,
          "fork": false,
          "template": false,
          "parent": null,
          "mirror": false,
          "size": 3846,
          "html_url": "http://localhost:10081/yystopf/ceshi",
          "ssh_url": "virus@localhost:10081:yystopf/ceshi.git",
          "clone_url": "http://localhost:10081/yystopf/ceshi.git",
          "original_url": "",
          "website": "",
          "stars_count": 0,
          "forks_count": 1,
          "watchers_count": 1,
          "open_issues_count": 0,
          "open_pr_counter": 0,
          "release_counter": 0,
          "default_branch": "master",
          "archived": false,
          "created_at": "2021-06-03T15:15:30+08:00",
          "updated_at": "2021-07-26T13:52:16+08:00",
          "permissions": {
            "admin": false,
            "push": false,
            "pull": false
          },
          "has_issues": true,
          "internal_tracker": {
            "enable_time_tracker": true,
            "allow_only_contributors_to_track_time": true,
            "enable_issue_dependencies": true
          },
          "has_wiki": true,
          "has_pull_requests": true,
          "ignore_whitespace_conflicts": false,
          "allow_merge_commits": true,
          "allow_rebase": true,
          "allow_rebase_explicit": true,
          "allow_squash_merge": true,
          "avatar_url": "",
          "internal": false
        },
        "pusher": {
          "id": 0,
          "login": "yystopf",
          "full_name": "",
          "email": "yystopf@forge.com",
          "avatar_url": "http://localhost:10081/user/avatar/yystopf/-1",
          "language": "",
          "is_admin": false,
          "last_login": "0001-01-01T00:00:00Z",
          "created": "2021-06-03T14:50:25+08:00",
          "username": "yystopf"
        },
        "sender": {
          "id": 0,
          "login": "yystopf",
          "full_name": "",
          "email": "yystopf@forge.com",
          "avatar_url": "http://localhost:10081/user/avatar/yystopf/-1",
          "language": "",
          "is_admin": false,
          "last_login": "0001-01-01T00:00:00Z",
          "created": "2021-06-03T14:50:25+08:00",
          "username": "yystopf"
        }
      },
      "request_content": {
        "headers": {
          "X-GitHub-Delivery": "99aa2c23-6884-4c44-9020-5469320aa408",
          "X-GitHub-Event": "push",
          "X-Gitea-Delivery": "99aa2c23-6884-4c44-9020-5469320aa408",
          "X-Gitea-Event": "push",
          "X-Gitea-Signature": "34a01edcd952ff6410ff6ebc946471161bde74aff86171f21621d2c2c4130f66",
          "X-Gogs-Delivery": "99aa2c23-6884-4c44-9020-5469320aa408",
          "X-Gogs-Event": "push",
          "X-Gogs-Signature": "34a01edcd952ff6410ff6ebc946471161bde74aff86171f21621d2c2c4130f66"
        }
      },
      "response_content": {
        "status": 200,
        "headers": {
          "Cache-Control": "no-store, must-revalidate, private, max-age=0",
          "Content-Length": "2556",
          "Content-Type": "text/html; charset=utf-8",
          "Referrer-Policy": "strict-origin-when-cross-origin",
          "Set-Cookie": "__profilin=p%3Dt; path=/; HttpOnly",
          "Vary": "Origin",
          "X-Content-Type-Options": "nosniff",
          "X-Download-Options": "noopen",
          "X-Frame-Options": "SAMEORIGIN",
          "X-Miniprofiler-Ids": "9ynvpncz5xm0rpgorb5y,hgggd9mv6lr4a9drcrlr,j7zqlx2vy5aji2vtgoba,f1ktsmh3jxvq0z2hf612,mih3dvgvlqhi3zy8lf2x,5k1qbkvbnru8mye9cest,tj6ern8w6awqf2zsimbr,9isaehvubivd52wo5p9v,1rzfhtq1nhuwbgy9p76g,z0xzidzyywna0y7a69m0,hzoklky92ycjqt42gi0s,y0ai7y0t28mcn8x0py2x,322il7nadinp51mw2r5m,m6dukftfsh6tjcxzp1gq,667wlqbytfwbrirnmma1,jcehj3dl8lkw8gk510cr",
          "X-Miniprofiler-Original-Cache-Control": "max-age=0, private, must-revalidate",
          "X-Permitted-Cross-Domain-Policies": "none",
          "X-Request-Id": "08bff080-bbb5-4183-b845-81de3d47120a",
          "X-Runtime": "0.394766",
          "X-Xss-Protection": "1; mode=block"
        },
        "body": "<!doctype html><html lang=\"zh-CN\" class=\"notranslate translated-ltr\" translate=\"no\"><head><meta charset=\"utf-8\"><meta name=\"”Keywords”\" content=\"”trustie,trustieforge,forge,确实让创建更美好,协同开发平台″\"><meta name=\"”Keywords”\" content=\"”TrustieOpenSourceProject″\"><meta name=\"”Keywords”\" content=\"”issue,bug,tracker,软件工程,课程实践″\"><meta name=\"”Description”\" content=\"”持续构建协同、共享、可信的软件创建生态开源创作与软件生产相结合，支持大规模群体开展软件协同创新活动”\"><meta name=\"theme-color\" content=\"#000000\"><link rel=\"manifest\" href=\"/react/build//manifest.json\"><link rel=\"stylesheet\" href=\"/react/build/css/iconfont.css\"><link rel=\"stylesheet\" href=\"/react/build/css/edu-purge.css\"><link rel=\"stylesheet\" href=\"/react/build/css/editormd.min.css\"><link rel=\"stylesheet\" href=\"/react/build/css/merge.css\"><link href=\"/react/build/static/css/main.07f7e90c.chunk.css\" rel=\"stylesheet\"></head><body><div id=\"md_div\" style=\"display:none\"></div><div id=\"root\" class=\"page -layout-v -fit widthunit\"></div><div id=\"picture_display\" style=\"display:none\"></div><script src=\"/react/build/js/jquery-1.8.3.min.js\"></script><script src=\"/react/build/js/js_min_all.js\"></script><script src=\"/react/build/js/codemirror/codemirror.js\"></script><script src=\"/react/build/js/editormd/editormd.min.js\"></script><script src=\"/react/build/js/codemirror/merge/merge.js\"></script><script src=\"/react/build/./static/js/runtime~main.3d644966.js\"></script><script src=\"/react/build/./static/js/main.e46872e3.chunk.js\"></script><script async type=\"text/javascript\" id=\"mini-profiler\" src=\"/mini-profiler-resources/includes.js?v=67dd1c2571ced7fc74ae7f1813e47bdf\" data-version=\"67dd1c2571ced7fc74ae7f1813e47bdf\" data-path=\"/mini-profiler-resources/\" data-current-id=\"9ynvpncz5xm0rpgorb5y\" data-ids=\"9ynvpncz5xm0rpgorb5y,hgggd9mv6lr4a9drcrlr,j7zqlx2vy5aji2vtgoba,f1ktsmh3jxvq0z2hf612,mih3dvgvlqhi3zy8lf2x,5k1qbkvbnru8mye9cest,tj6ern8w6awqf2zsimbr,9isaehvubivd52wo5p9v,1rzfhtq1nhuwbgy9p76g,z0xzidzyywna0y7a69m0,hzoklky92ycjqt42gi0s,y0ai7y0t28mcn8x0py2x,322il7nadinp51mw2r5m,m6dukftfsh6tjcxzp1gq,667wlqbytfwbrirnmma1,jcehj3dl8lkw8gk510cr\" data-horizontal-position=\"left\" data-vertical-position=\"top\" data-trivial=\"false\" data-children=\"false\" data-max-traces=\"20\" data-controls=\"false\" data-total-sql-count=\"false\" data-authorized=\"true\" data-toggle-shortcut=\"alt+p\" data-start-hidden=\"false\" data-collapse-results=\"true\" data-html-container=\"body\"></script>\n</body></html>"
      },
      "delivered_time": "2021-07-28 11:47:29"
    }
  ]
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
|» total_count|integer|true|none|分页总数|none|
|» hooktasks|[object]|true|none||none|
|»» id|integer|false|none||none|
|»» type|string|false|none|类型|none|
|»» uuid|string|false|none|推送uuid|none|
|»» is_succeed|boolean|false|none|是否推送成功|none|
|»» is_delivered|boolean|false|none|是否完成推送|none|
|»» payload_content|object|false|none|请求主体内容|none|
|»»» secret|string|true|none||none|
|»»» ref|string|true|none||none|
|»»» before|string|true|none||none|
|»»» after|string|true|none||none|
|»»» compare_url|string|true|none||none|
|»»» commits|[object]|true|none||none|
|»»»» id|string|false|none||none|
|»»»» message|string|false|none||none|
|»»»» url|string|false|none||none|
|»»»» author|object|false|none||none|
|»»»»» name|string|true|none||none|
|»»»»» email|string|true|none||none|
|»»»»» username|string|true|none||none|
|»»»» committer|object|false|none||none|
|»»»»» name|string|true|none||none|
|»»»»» email|string|true|none||none|
|»»»»» username|string|true|none||none|
|»»»» verification|object|false|none||none|
|»»»»» verified|boolean|true|none||none|
|»»»»» reason|string|true|none||none|
|»»»»» signature|string|true|none||none|
|»»»»» signer|null|true|none||none|
|»»»»» payload|string|true|none||none|
|»»»» timestamp|string|false|none||none|
|»»»» added|null|false|none||none|
|»»»» removed|null|false|none||none|
|»»»» modified|null|false|none||none|
|»»» head_commit|null|true|none||none|
|»»» repository|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» owner|object|true|none||none|
|»»»»» id|integer|true|none||none|
|»»»»» login|string|true|none||none|
|»»»»» full_name|string|true|none||none|
|»»»»» email|string|true|none||none|
|»»»»» avatar_url|string|true|none||none|
|»»»»» language|string|true|none||none|
|»»»»» is_admin|boolean|true|none||none|
|»»»»» last_login|string|true|none||none|
|»»»»» created|string|true|none||none|
|»»»»» username|string|true|none||none|
|»»»» name|string|true|none||none|
|»»»» full_name|string|true|none||none|
|»»»» description|string|true|none||none|
|»»»» empty|boolean|true|none||none|
|»»»» private|boolean|true|none||none|
|»»»» fork|boolean|true|none||none|
|»»»» template|boolean|true|none||none|
|»»»» parent|null|true|none||none|
|»»»» mirror|boolean|true|none||none|
|»»»» size|integer|true|none||none|
|»»»» html_url|string|true|none||none|
|»»»» ssh_url|string|true|none||none|
|»»»» clone_url|string|true|none||none|
|»»»» original_url|string|true|none||none|
|»»»» website|string|true|none||none|
|»»»» stars_count|integer|true|none||none|
|»»»» forks_count|integer|true|none||none|
|»»»» watchers_count|integer|true|none||none|
|»»»» open_issues_count|integer|true|none||none|
|»»»» open_pr_counter|integer|true|none||none|
|»»»» release_counter|integer|true|none||none|
|»»»» default_branch|string|true|none||none|
|»»»» archived|boolean|true|none||none|
|»»»» created_at|string|true|none||none|
|»»»» updated_at|string|true|none||none|
|»»»» permissions|object|true|none||none|
|»»»»» admin|boolean|true|none||none|
|»»»»» push|boolean|true|none||none|
|»»»»» pull|boolean|true|none||none|
|»»»» has_issues|boolean|true|none||none|
|»»»» internal_tracker|object|true|none||none|
|»»»»» enable_time_tracker|boolean|true|none||none|
|»»»»» allow_only_contributors_to_track_time|boolean|true|none||none|
|»»»»» enable_issue_dependencies|boolean|true|none||none|
|»»»» has_wiki|boolean|true|none||none|
|»»»» has_pull_requests|boolean|true|none||none|
|»»»» ignore_whitespace_conflicts|boolean|true|none||none|
|»»»» allow_merge_commits|boolean|true|none||none|
|»»»» allow_rebase|boolean|true|none||none|
|»»»» allow_rebase_explicit|boolean|true|none||none|
|»»»» allow_squash_merge|boolean|true|none||none|
|»»»» avatar_url|string|true|none||none|
|»»»» internal|boolean|true|none||none|
|»»» pusher|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» login|string|true|none||none|
|»»»» full_name|string|true|none||none|
|»»»» email|string|true|none||none|
|»»»» avatar_url|string|true|none||none|
|»»»» language|string|true|none||none|
|»»»» is_admin|boolean|true|none||none|
|»»»» last_login|string|true|none||none|
|»»»» created|string|true|none||none|
|»»»» username|string|true|none||none|
|»»» sender|object|true|none||none|
|»»»» id|integer|true|none||none|
|»»»» login|string|true|none||none|
|»»»» full_name|string|true|none||none|
|»»»» email|string|true|none||none|
|»»»» avatar_url|string|true|none||none|
|»»»» language|string|true|none||none|
|»»»» is_admin|boolean|true|none||none|
|»»»» last_login|string|true|none||none|
|»»»» created|string|true|none||none|
|»»»» username|string|true|none||none|
|»» request_content|object|false|none|请求内容，头部等等|none|
|»»» headers|object|true|none||none|
|»»»» X-GitHub-Delivery|string|true|none||none|
|»»»» X-GitHub-Event|string|true|none||none|
|»»»» X-Gitea-Delivery|string|true|none||none|
|»»»» X-Gitea-Event|string|true|none||none|
|»»»» X-Gitea-Signature|string|true|none||none|
|»»»» X-Gogs-Delivery|string|true|none||none|
|»»»» X-Gogs-Event|string|true|none||none|
|»»»» X-Gogs-Signature|string|true|none||none|
|»» response_content|object|false|none|响应内容，状态，头部，主体等等|none|
|»»» status|integer|true|none||none|
|»»» headers|object|true|none||none|
|»»»» Cache-Control|string|true|none||none|
|»»»» Content-Length|string|true|none||none|
|»»»» Content-Type|string|true|none||none|
|»»»» Referrer-Policy|string|true|none||none|
|»»»» Set-Cookie|string|true|none||none|
|»»»» Vary|string|true|none||none|
|»»»» X-Content-Type-Options|string|true|none||none|
|»»»» X-Download-Options|string|true|none||none|
|»»»» X-Frame-Options|string|true|none||none|
|»»»» X-Miniprofiler-Ids|string|true|none||none|
|»»»» X-Miniprofiler-Original-Cache-Control|string|true|none||none|
|»»»» X-Permitted-Cross-Domain-Policies|string|true|none||none|
|»»»» X-Request-Id|string|true|none||none|
|»»»» X-Runtime|string|true|none||none|
|»»»» X-Xss-Protection|string|true|none||none|
|»»» body|string|true|none||none|
|»» delivered_time|string|false|none|推送时间|none|

## POST Webhook测试推送

POST /api/v1/{owner}/{repo}/webhooks/{id}/tests.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||项目拥有者用户标识|
|repo|path|string| 是 ||项目标识|
|id|path|string| 是 ||Webhook ID|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

# Oauth2

## POST 刷新token

POST /oauth/token

> Body 请求参数

```json
{
  "grant_type": "refresh_token",
  "refresh_token": "xxxxxxxxxxxxxx",
  "client_id": "7",
  "client_secret": "ea tempor in anim"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» grant_type|body|string| 是 ||鉴权类型|
|» refresh_token|body|string| 是 ||refresh_token|
|» client_id|body|string| 是 ||ClientID|
|» client_secret|body|string| 是 ||ClientSecret|

#### 枚举值

|属性|值|
|---|---|
|» grant_type|refresh_token|

> 返回示例

> 200 Response

```json
{
  "access_token": "eyJraWQiOiJUaEVSLVl3Ukg4TWYwOHM0UnJLUDYzXzZLWmVET2NZckZXcmdzN2VUVWdrIiwiYWxnIjoiSFM1MTIifQ.eyJpc3MiOiJHaXRMaW5rIiwiaWF0IjoxNjc1OTI5NTAxLCJqdGkiOiI1Yjg2ZDNjMi1hODA0LTQyNjEtYWFjYi1jZDg0YzM5ZjBiM2MiLCJ1c2VyIjp7ImlkIjo4NDcyNywibG9naW4iOiJ5eXN0b3BmMTIzIiwibWFpbCI6Ijk3ODY2MDQxNEBxcS5jb20ifX0.JxgzogjKZlRsTpB3Gy37_D2y5wP847X0NKYCdInm7TDwpzR8PPm9RXBfWpOe2riKHA9RTdUREG0dHjDCMNnaWQ",
  "token_type": "Bearer",
  "expires_in": 604799,
  "refresh_token": "DkE20crBDMVK4yfqMdte15DvR-dQpZciuP-z3jl3kaM",
  "scope": "public",
  "created_at": 1675929501
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
|» access_token|string|true|none||none|
|» token_type|string|true|none||none|
|» expires_in|integer|true|none||none|
|» refresh_token|string|true|none||none|
|» scope|string|true|none||none|
|» created_at|integer|true|none||none|

# Wiki/wiki功能接口

<a id="opIdcreateWikiUsingPOST"></a>

## POST createWiki

POST /api/wiki/createWiki

> Body 请求参数

```json
{
  "owner": "wanjia9506",
  "repo": "test",
  "projectId": 1406620,
  "pageName": "123",
  "title": "123",
  "message": "",
  "content_base64": "5qyi6L+O5p2l5YiwV2lraQ=="
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» content_base64|body|string| 是 | wiki内容|none|
|» message|body|string¦null| 否 | 提交信息|none|
|» owner|body|string| 是 | 仓库拥有者标识|none|
|» pageName|body|string| 是 | 当前页标题|none|
|» projectId|body|integer(int32)| 是 | gitlink项目id|none|
|» repo|body|string| 是 | 仓库标识|none|
|» title|body|string| 是 | 标题|none|

> 返回示例

> 200 Response

```json
{
  "message": "201",
  "data": "{\"title\":\"test\",\"html_url\":\"https://gitlink.org.cn/wanjia9506/test/wiki/test\",\"sub_url\":\"test\",\"last_commit\":{\"sha\":\"aca9df20a9c241ddf3862ba8203d5764a5a07829\",\"author\":{\"name\":\"wanjia9506\",\"email\":\"553039491@qq.com\",\"date\":\"2023-12-27T07:25:54Z\"},\"commiter\":{\"name\":\"wanjia9506\",\"email\":\"553039491@qq.com\",\"date\":\"2023-12-27T07:25:54Z\"},\"message\":\"Add 'test'\\n\"},\"content_base64\":\"5qyi6L+O5p2l5YiwV2lraQ==\",\"commit_count\":3,\"sidebar\":\"W1t0c3QxMjNdXQ==\",\"footer\":\"\"}\n"
}
```

> 201 Response

```json
{
  "message": "201",
  "data": "{\"title\":\"test\",\"html_url\":\"https://gitlink.org.cn/wanjia9506/test/wiki/test\",\"sub_url\":\"test\",\"last_commit\":{\"sha\":\"aca9df20a9c241ddf3862ba8203d5764a5a07829\",\"author\":{\"name\":\"wanjia9506\",\"email\":\"553039491@qq.com\",\"date\":\"2023-12-27T07:25:54Z\"},\"commiter\":{\"name\":\"wanjia9506\",\"email\":\"553039491@qq.com\",\"date\":\"2023-12-27T07:25:54Z\"},\"message\":\"Add 'test'\\n\"},\"content_base64\":\"5qyi6L+O5p2l5YiwV2lraQ==\",\"commit_count\":3,\"sidebar\":\"W1t0c3QxMjNdXQ==\",\"footer\":\"\"}\n"
}
```

> 401 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[ResponseData](#schemaresponsedata)|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### 返回数据结构

状态码 **200**

*ResponseData*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|false|none||none|
|» message|string|false|none||none|

<a id="opIddeleteWikiUsingDELETE"></a>

## DELETE deleteWiki

DELETE /api/wiki/deleteWiki

> Body 请求参数

```json
{
  "owner": "wanjia9506",
  "repo": "test",
  "projectId": 1406620,
  "pageName": "123"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» owner|body|string| 否 | 仓库拥有者标识|none|
|» pageName|body|string| 否 | 当前页标题|none|
|» projectId|body|integer(int32)| 否 | gitlink项目id|none|
|» repo|body|string| 否 | 仓库标识|none|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[ResponseData](#schemaresponsedata)|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|

### 返回数据结构

状态码 **200**

*ResponseData*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|false|none||none|
|» message|string|false|none||none|

<a id="opIdgetWikiUsingGET"></a>

## GET getWiki

GET /api/wiki/getWiki

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|query|string| 是 ||owner of the repo|
|pageName|query|string| 是 ||name of the wikipage|
|projectId|query|integer| 是 ||id of the project|
|repo|query|string| 是 ||name of the repo|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[ResponseData](#schemaresponsedata)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### 返回数据结构

状态码 **200**

*ResponseData*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|false|none||none|
|» message|string|false|none||none|

<a id="opIdupdateWikiUsingPUT"></a>

## PUT updateWiki

PUT /api/wiki/updateWiki

> Body 请求参数

```json
{
  "content_base64": "string",
  "message": "string",
  "owner": "string",
  "pageName": "string",
  "projectId": 0,
  "repo": "string",
  "title": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|body|body|object| 否 ||none|
|» content_base64|body|string| 否 |  wiki内容|none|
|» message|body|string¦null| 否 | 提交信息|none|
|» owner|body|string| 是 | 仓库拥有者标识|none|
|» pageName|body|string| 是 | 当前页标题|none|
|» projectId|body|integer(int32)| 是 | gitlink仓库id|none|
|» repo|body|string| 是 | 仓库标识|none|
|» title|body|string| 是 | 标题|none|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[ResponseData](#schemaresponsedata)|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### 返回数据结构

状态码 **200**

*ResponseData*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|false|none||none|
|» message|string|false|none||none|

<a id="opIdwikiPagesUsingGET"></a>

## GET wikiPages

GET /api/wiki/wikiPages

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|query|string| 是 ||owner of the repo|
|projectId|query|integer| 是 ||id of the project|
|repo|query|string| 是 ||name of the repo|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[ResponseData](#schemaresponsedata)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### 返回数据结构

状态码 **200**

*ResponseData*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|false|none||none|
|» message|string|false|none||none|

# Wiki/wiki导入导出接口

<a id="opIduploadUsingPOST"></a>

## POST 上传文件

POST /api/wikiExport/uploadWiki/{owner}/{repoName}/{projectId}

> Body 请求参数

```yaml
multipartFile: ""

```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||owner of the repo|
|projectId|path|integer| 是 ||id of the project|
|repoName|path|string| 是 ||name of the repo|
|body|body|object| 否 ||none|
|» multipartFile|body|string(binary)| 否 ||multipartFile|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[ResponseData](#schemaresponsedata)|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### 返回数据结构

状态码 **200**

*ResponseData*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|false|none||none|
|» message|string|false|none||none|

<a id="opIdwikiExportToPdfUsingGET"></a>

## GET 导出Wiki-wrapper

GET /api/wikiExport/wikiExport-wrapper

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|query|string| 是 ||owner of the repo|
|projectId|query|integer| 是 ||id of the project|
|projectName|query|string| 是 ||name of the project|
|repoName|query|string| 是 ||name of the repo|
|type|query|string| 是 ||type of the file：pdf or markdown or html|

> 返回示例

> 200 Response

```json
{
  "data": {},
  "message": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|[ResponseData](#schemaresponsedata)|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|
|404|[Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)|none|Inline|

### 返回数据结构

状态码 **200**

*ResponseData*

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» data|object|false|none||none|
|» message|string|false|none||none|

# 流水线

## GET 流水线列表

GET /api/pm/pipelines.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|page|query|string| 否 ||none|
|limit|query|string| 否 ||none|
|owner_id|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "string"
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
|» message|string|true|none||none|

## POST 运行流水线

POST /api/v1/{owner}/{repo}/actions/runs.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|ref|query|string| 否 ||none|
|workflow|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 流水线执行记录列表

GET /api/v1/{owner}/{repo}/actions/runs.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|ref|query|string| 否 ||none|
|workflow|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 图形化流水线保存

POST /api/v1/{owner}/{repo}/pipelines/save_yaml

> Body 请求参数

```json
{
  "id": 37,
  "pipeline_json": {
    "nodes": [
      {
        "position": {
          "x": 150,
          "y": 120
        },
        "size": {
          "width": 186,
          "height": 65
        },
        "view": "react-shape-view",
        "shape": "job",
        "ports": {
          "groups": {
            "top": {
              "position": {
                "name": "top",
                "args": {
                  "dx": 0
                }
              },
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "bottom": {
              "position": {
                "name": "bottom",
                "args": {
                  "dx": 0
                }
              },
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "left": {
              "position": "left",
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "right": {
              "position": "right",
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            }
          },
          "items": [
            {
              "id": "on-push-81a68a8-left",
              "group": "left"
            },
            {
              "id": "on-push-81a68a8-right",
              "group": "right"
            },
            {
              "id": "on-push-81a68a8-top",
              "group": "top"
            },
            {
              "id": "on-push-81a68a8-bottom",
              "group": "bottom"
            }
          ]
        },
        "id": "on-push-81a68a8",
        "data": {
          "id": "on-push-81a68a8",
          "label": "代码push事件",
          "name": "on-push",
          "full_name": "on-push",
          "description": "GitLink仓库push事件",
          "icon": "https://osredm.com/api/attachments/7fe3f4b4-0752-4b22-8ac0-1b0c72416724",
          "action_node_types_id": 3,
          "yaml": "",
          "sort_no": 0,
          "use_count": 0,
          "node_type": "start",
          "is_mutil_link": true,
          "link_type": "job",
          "inputs": [
            {
              "id": 6,
              "name": "branches",
              "input_type": "input",
              "description": "分支名称，多个分支英文逗号隔开，如'master,dev'",
              "is_required": true,
              "default_value": "master",
              "value": "master"
            },
            {
              "id": 7,
              "name": "paths-ignore",
              "input_type": "input",
              "description": "忽略文件，多个文件英文逗号隔开，如'**.md,**.yaml'",
              "is_required": false,
              "default_value": null,
              "value": null
            }
          ],
          "x": 566,
          "y": 454,
          "img": "https://osredm.com/api/attachments/7fe3f4b4-0752-4b22-8ac0-1b0c72416724",
          "branches": "master",
          "paths-ignore": null
        },
        "zIndex": 1
      },
      {
        "position": {
          "x": 420,
          "y": 120
        },
        "size": {
          "width": 186,
          "height": 65
        },
        "view": "react-shape-view",
        "shape": "job",
        "ports": {
          "groups": {
            "top": {
              "position": {
                "name": "top",
                "args": {
                  "dx": 0
                }
              },
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "bottom": {
              "position": {
                "name": "bottom",
                "args": {
                  "dx": 0
                }
              },
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "left": {
              "position": "left",
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "right": {
              "position": "right",
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            }
          },
          "items": [
            {
              "id": "job-18f4d4af-left",
              "group": "left"
            },
            {
              "id": "job-18f4d4af-right",
              "group": "right"
            },
            {
              "id": "job-18f4d4af-top",
              "group": "top"
            },
            {
              "id": "job-18f4d4af-bottom",
              "group": "bottom"
            }
          ]
        },
        "id": "job-18f4d4af",
        "data": {
          "id": 24,
          "label": "任务",
          "name": "job",
          "full_name": "",
          "description": "",
          "icon": "",
          "action_node_types_id": 7,
          "yaml": "",
          "sort_no": 0,
          "use_count": 0,
          "node_type": "job",
          "is_mutil_link": true,
          "link_type": "job,step",
          "inputs": [],
          "x": 900,
          "y": 461,
          "img": ""
        },
        "zIndex": 3
      },
      {
        "position": {
          "x": 431,
          "y": 306.625
        },
        "size": {
          "width": 166,
          "height": 51
        },
        "view": "react-shape-view",
        "shape": "step",
        "ports": {
          "groups": {
            "top": {
              "position": {
                "name": "top",
                "args": {
                  "dx": 0
                }
              },
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "bottom": {
              "position": {
                "name": "bottom",
                "args": {
                  "dx": 0
                }
              },
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "left": {
              "position": "left",
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            },
            "right": {
              "position": "right",
              "attrs": {
                "circle": {
                  "r": 5,
                  "magnet": true,
                  "strokeWidth": 1,
                  "fill": "#fff",
                  "stroke": "#85A5FF"
                }
              }
            }
          },
          "items": [
            {
              "id": "shell-992db95-left",
              "group": "left"
            },
            {
              "id": "shell-992db95-right",
              "group": "right"
            },
            {
              "id": "shell-992db95-top",
              "group": "top"
            },
            {
              "id": "shell-992db95-bottom",
              "group": "bottom"
            }
          ]
        },
        "id": "shell-992db95",
        "data": {
          "id": "shell-992db95",
          "label": "运行Shell脚本",
          "name": "shell",
          "full_name": "shell",
          "description": "",
          "icon": "https://osredm.com/api/attachments/91c6873b-a117-4771-a5ea-7f9e3f2d9a3a",
          "action_node_types_id": 4,
          "yaml": "",
          "sort_no": 0,
          "use_count": 0,
          "node_type": "step",
          "is_mutil_link": false,
          "link_type": "step",
          "inputs": [
            {
              "id": 3,
              "name": "run",
              "input_type": "input",
              "description": null,
              "is_required": false,
              "default_value": null,
              "value": "echo 111"
            }
          ],
          "x": 819,
          "y": 589,
          "img": "https://osredm.com/api/attachments/91c6873b-a117-4771-a5ea-7f9e3f2d9a3a",
          "run": "echo 111"
        },
        "zIndex": 2
      }
    ],
    "edges": [
      {
        "shape": "data-processing-curve",
        "inherit": "edge",
        "connector": {
          "name": "rounded",
          "args": {
            "radius": 10
          }
        },
        "router": {
          "name": "manhattan"
        },
        "attrs": {
          "line": {
            "strokeDasharray": "0"
          }
        },
        "id": "39159699-26b7-4db8-99c0-9088be560a22",
        "zIndex": -1,
        "source": {
          "cell": "on-push-81a68a8",
          "port": "on-push-81a68a8-right"
        },
        "target": {
          "cell": "job-18f4d4af",
          "port": "job-18f4d4af-left"
        }
      },
      {
        "shape": "data-processing-curve",
        "inherit": "edge",
        "connector": {
          "name": "rounded",
          "args": {
            "radius": 10
          }
        },
        "router": {
          "name": "manhattan"
        },
        "attrs": {
          "line": {
            "strokeWidth": 1,
            "strokeDasharray": "3"
          }
        },
        "id": "3c6f7387-74d5-4af1-a2f7-8520ed7b446e",
        "zIndex": -1,
        "source": {
          "cell": "job-18f4d4af",
          "port": "job-18f4d4af-bottom"
        },
        "target": {
          "cell": "shell-992db95",
          "port": "shell-992db95-top"
        }
      }
    ]
  }
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» id|body|string| 是 ||none|
|» pipeline_json|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE 删除流水线

DELETE /api/v1/{owner}/{repo}/pipelines/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 查看流水线详情

GET /api/v1/{owner}/{repo}/pipelines/{id}.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 禁用流水线

POST /api/v1/{owner}/{repo}/actions/disable.json

> Body 请求参数

```json
{
  "id": 6,
  "workflow": "流水线上下文详情.yml"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» id|body|integer| 是 ||none|
|» workflow|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 启用流水线

POST /api/v1/{owner}/{repo}/actions/enable.json

> Body 请求参数

```json
{
  "id": 6,
  "workflow": "流水线上下文详情.yml"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» id|body|integer| 是 ||none|
|» workflow|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 查询流水线运行日志

POST /api/v1/{owner}/{repo}/actions/runs/{run_id}/jobs/0

> Body 请求参数

```json
{
  "id": "6",
  "index": "43",
  "job": 0,
  "log_cursors": [
    {
      "cursor": null,
      "expanded": true,
      "step": 1
    }
  ],
  "owner": "osredm",
  "repo": "osredm_help"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|run_id|path|string| 是 ||none|
|body|body|object| 否 ||none|
|» id|body|integer| 是 ||none|
|» index|body|string| 是 ||none|
|» job|body|string| 是 ||none|
|» log_cursors|body|object| 是 ||none|
|»» cursor|body|string| 是 ||none|
|»» expanded|body|string| 是 ||none|
|»» step|body|string| 是 ||none|
|» owner|body|string| 是 ||none|
|» repo|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET 获取流水线执行报告结果

GET /api/v1/{owner}/{repo}/pipelines/run_results.json

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|run_id|query|string| 否 ||none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 模板

## GET 模板列表接口

GET /api/v1/yystopf/ceshi/project_templates

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|

> 返回示例

> 200 Response

```json
{
  "total_count": 2,
  "project_templates": [
    {
      "id": 1,
      "type": "ProjectTemplates::Issue",
      "name": "bug修复模板",
      "content": "## 问题描述\n[清晰简洁地描述Bug是什么]\n## 复现步骤\n1. [第一步]\n2. [第二步]\n3. [...]\n...\n## 期望行为\n[你期望发生什么]\n## 实际行为\n[实际发生了什么]\n## 环境信息\n- 操作系统：[如 macOS 12.3]\n- 浏览器/版本：[如 Chrome 101]\n- 项目版本/分支：[如 v1.2.0 或 main]\n## 附加信息\n[日志、截图、其他上下文]",
      "created_at": "2026-01-13 17:30:20",
      "updated_at": "2026-01-13 17:30:20"
    },
    {
      "id": 2,
      "type": "ProjectTemplates::Issue",
      "name": "功能请求模板",
      "content": "## 功能描述\n[你希望添加什么功能？为什么需要它？]\n## 解决方案\n[如果有，描述你希望的实现方式]\n## 替代方案\n[你考虑过的其他方案]\n## 影响范围\n[这个功能会影响哪些模块/用户？]\n## 补充说明\n[其他相关信息]",
      "created_at": "2026-01-13 17:30:20",
      "updated_at": "2026-01-13 17:30:20"
    }
  ]
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
|» total_count|integer|true|none||none|
|» project_templates|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» content|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|

## POST 创建模板

POST /api/v1/yystopf/ceshi/project_templates

> Body 请求参数

```json
{
  "type": "ProjectTemplates::Issue",
  "name": "测试issue模板1",
  "content": "内容"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|Cookie|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» type|body|string| 是 ||none|
|» name|body|string| 是 ||none|
|» content|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## DELETE 删除模板

DELETE /api/v1/yystopf/ceshi/project_templates/5

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|Cookie|header|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## PUT 更新模板

PUT /api/v1/yystopf/ceshi/project_templates/5

> Body 请求参数

```json
{
  "type": "ProjectTemplates::Issue",
  "name": "测试issue模板1",
  "content": "内容1"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|Cookie|header|string| 是 ||none|
|body|body|object| 否 ||none|
|» type|body|string| 是 ||none|
|» name|body|string| 是 ||none|
|» content|body|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "status": 0,
  "message": "success"
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
|» message|string|true|none||none|

## GET 模板详情接口 

GET /api/v1/{owner}/{repo}/project_templates/{id}

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|autologin_trustie|cookie|string| 否 ||none|
|owner|path|string| 是 ||none|
|repo|path|string| 是 ||none|
|id|path|string| 是 ||none|

> 返回示例

> 200 Response

```json
{
  "project_template": {
    "id": 42,
    "type": "ProjectTemplates::Issue",
    "name": "功能请求模板",
    "content": "## 功能描述\n[你希望添加什么功能？为什么需要它？]\n## 解决方案\n[如果有，描述你希望的实现方式]\n## 替代方案\n[你考虑过的其他方案]\n## 影响范围\n[这个功能会影响哪些模块/用户？]\n## 补充说明\n[其他相关信息]",
    "created_at": "2026-01-22 11:15:04",
    "updated_at": "2026-01-22 11:15:04"
  }
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
|» project_template|object|true|none||none|
|»» id|integer|true|none||none|
|»» type|string|true|none||none|
|»» name|string|true|none||none|
|»» content|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|

# 数据模型

<h2 id="tocS_Login">Login</h2>

<a id="schemalogin"></a>
<a id="schema_Login"></a>
<a id="tocSlogin"></a>
<a id="tocslogin"></a>

```json
{
  "name": "string",
  "pwd": "string"
}

```

Login

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|name|string|false|none||none|
|pwd|string|false|none||none|

<h2 id="tocS_ResponseData">ResponseData</h2>

<a id="schemaresponsedata"></a>
<a id="schema_ResponseData"></a>
<a id="tocSresponsedata"></a>
<a id="tocsresponsedata"></a>

```json
{
  "data": {},
  "message": "string"
}

```

ResponseData

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|data|object|false|none||none|
|message|string|false|none||none|

<h2 id="tocS_WikiAllVo">WikiAllVo</h2>

<a id="schemawikiallvo"></a>
<a id="schema_WikiAllVo"></a>
<a id="tocSwikiallvo"></a>
<a id="tocswikiallvo"></a>

```json
{
  "content_base64": "string",
  "message": "string",
  "owner": "string",
  "pageName": "string",
  "projectId": 0,
  "repo": "string",
  "title": "string"
}

```

WikiAllVo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|content_base64|string|false|none||none|
|message|string|false|none||none|
|owner|string|false|none||none|
|pageName|string|false|none||none|
|projectId|integer(int32)|false|none||none|
|repo|string|false|none||none|
|title|string|false|none||none|

<h2 id="tocS_WikiDeleteVo">WikiDeleteVo</h2>

<a id="schemawikideletevo"></a>
<a id="schema_WikiDeleteVo"></a>
<a id="tocSwikideletevo"></a>
<a id="tocswikideletevo"></a>

```json
{
  "owner": "string",
  "pageName": "string",
  "projectId": 0,
  "repo": "string"
}

```

WikiDeleteVo

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|owner|string|false|none||none|
|pageName|string|false|none||none|
|projectId|integer(int32)|false|none||none|
|repo|string|false|none||none|
---

# 补充接口（来自 Apifox）

> 以下接口仅在 Apifox 文档中出现，GitLink.md 中未收录。

## GET 获取疑修下评论或操作记录列表

GET /api/v1/{owner}/{repo}/issues/{index}/journals.json

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|owner|path|string| 是 |用户名|
|repo|path|string| 是 |仓库标识|
|index|path|string| 是 |Issue 序号|
|category|query|string| 否 |类型：all 所有 / comment 评论 / operate 操作记录|
|keyword|query|string| 否 |搜索关键词|
|sort_by|query|string| 否 |排序字段：created_on / updated_on|
|sort_direction|query|string| 否 |排序类型：asc 正序 / desc 倒序|
|page|query|integer| 否 |页码|
|limit|query|integer| 否 |每页条数|

### 返回数据结构

状态码 **200**

|名称|类型|必选|说明|
|---|---|---|---|
|» total_count|integer|true|分页总数|
|» journals|[object]|true|评论/操作记录列表|
|»» id|integer|true|评论 ID|
|»» is_journal_detail|boolean|true|是否为操作记录|
|»» notes|string|false|评论内容|
|»» created_on|string|true|创建时间|
|»» user|object|true|评论用户|

## PATCH 更新项目（基础信息）

PATCH /api/{owner}/{repo}.json

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|owner|path|string| 是 |用户名|
|repo|path|string| 是 |仓库标识|
|body|body|object| 否 ||
|» description|body|string| 否 |项目描述|
|» website|body|string| 否 |项目网站|
|» lesson_url|body|string| 否 |课程链接|
|» project_topic_names|body|array| 否 |项目标签|

### 返回数据结构

状态码 **200**

|名称|类型|必选|说明|
|---|---|---|---|
|» id|integer|true|项目 ID|
|» identifier|string|true|项目标识|
|» name|string|true|项目名称|
|» description|string|true|项目描述|

## PATCH 更新项目（完整）

PATCH /api/{owner}/{repo}.json

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|owner|path|string| 是 |用户名|
|repo|path|string| 是 |仓库标识|
|body|body|object| 否 ||
|» name|body|string| 是 |项目名称|
|» identifier|body|string| 是 |项目标识|
|» description|body|string| 否 |项目描述|
|» project_category_id|body|integer| 否 |项目分类 ID|
|» project_language_id|body|integer| 否 |项目语言 ID|
|» private|body|boolean| 否 |是否私有|
|» pr_view_admin|body|boolean| 否 |PR 仅管理员可见|

## POST OAuth2 用户登录（用户名密码模式）

POST /oauth/token

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 ||
|» grant_type|body|string| 是 |鉴权类型（password）|
|» client_id|body|string| 是 |Client ID|
|» client_secret|body|string| 是 |Client Secret|
|» username|body|string| 是 |用户名|
|» password|body|string| 是 |密码|

> 返回示例

> 200 Response

```json
{
  "access_token": "13bfdecabe1f807f696686436659bf50a7c80052",
  "token_type": "Bearer",
  "expires_in": 604800,
  "refresh_token": "abc123...",
  "scope": "public",
  "created_at": 1711843200
}
```

### 返回数据结构

状态码 **200**

|名称|类型|必选|说明|
|---|---|---|---|
|» access_token|string|true|访问令牌|
|» token_type|string|true|令牌类型（Bearer）|
|» expires_in|integer|true|过期时间（秒），默认 604800（7天）|
|» refresh_token|string|true|刷新令牌|
|» scope|string|true|权限范围|
|» created_at|integer|true|创建时间戳|

## POST OAuth2 客户端模式登录

POST /oauth/token

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 ||
|» grant_type|body|string| 是 |鉴权类型（client_credentials）|
|» client_id|body|string| 是 |Client ID|
|» client_secret|body|string| 是 |Client Secret|

### 返回数据结构

状态码 **200**

|名称|类型|必选|说明|
|---|---|---|---|
|» access_token|string|true|访问令牌|
|» token_type|string|true|令牌类型|
|» expires_in|integer|true|过期时间（秒）|
|» scope|string|true|权限范围|
|» created_at|integer|true|创建时间戳|

## POST OAuth2 刷新 Token

POST /oauth/token

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 ||
|» grant_type|body|string| 是 |鉴权类型（refresh_token）|
|» refresh_token|body|string| 是 |刷新令牌|
|» client_id|body|string| 是 |Client ID|
|» client_secret|body|string| 是 |Client Secret|

### 返回数据结构

状态码 **200**

|名称|类型|必选|说明|
|---|---|---|---|
|» access_token|string|true|新的访问令牌|
|» token_type|string|true|令牌类型|
|» expires_in|integer|true|过期时间（秒）|
|» refresh_token|string|true|新的刷新令牌|
|» scope|string|true|权限范围|
|» created_at|integer|true|创建时间戳|

## POST 用户验证密码

POST /api/v1/{owner}/check_password.json

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|owner|path|string| 是 |用户名|
|body|body|object| 否 ||
|» password|body|string| 是 |密码|

### 返回数据结构

状态码 **200**

|名称|类型|必选|说明|
|---|---|---|---|
|» status|integer|true|状态码|
|» message|string|true|消息|
