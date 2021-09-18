
## 敏感词服务

<br />

### 运行

```
cd path/go-sensitive/api
go run sensitive.go -f etc/sensitive-api.yaml
```

<br />

### API



### 1. "验证是否包含敏感词"

1. 路由定义

- Url: /sensitive/validate
- Method: POST
- Request: `ValidateReq`
- Response: `ValidateReply`

2. 请求定义


```golang
type ValidateReq struct {
	Txt string `json:"txt"`
	Hash string `json:"hash,optional"`
}
```


3. 返回定义


```golang
type ValidateReply struct {
	IsValidate bool `json:"is_validate"`
	BadWord string `json:"bad_word"`
}
```
  


### 2. "查找所有敏感词"

1. 路由定义

- Url: /sensitive/find
- Method: POST
- Request: `FindReq`
- Response: `FindReply`

2. 请求定义


```golang
type FindReq struct {
	Txt string `json:"txt"`
	Channel string `json:"channel"`
	Hash string `json:"hash,optional"`
}
```


3. 返回定义


```golang
type FindReply struct {
	IsValidate bool `json:"is_validate"`
	BadWords map[string][]string `json:"bad_words"`
}
```
  


### 3. "批量查找所有敏感词"

1. 路由定义

- Url: /sensitive/batch-find
- Method: POST
- Request: `BatchFindReq`
- Response: `BatchFindReply`

2. 请求定义


```golang
type BatchFindReq struct {
	Items []BatchItemReq `json:"items"`
}
```


3. 返回定义


```golang
type BatchFindReply struct {
	Items map[string]BatchItemReply `json:"items"`
}
```
  


### 4. "批量新增敏感词"

1. 路由定义

- Url: /sensitive/add-word
- Method: POST
- Request: `AddWordReq`
- Response: `AddWordReply`

2. 请求定义


```golang
type AddWordReq struct {
	Words []string `json:"words"`
}
```


3. 返回定义


```golang
type AddWordReply struct {
}
```
  


### 5. "批量删除敏感词"

1. 路由定义

- Url: /sensitive/delete-word
- Method: POST
- Request: `DeleteWordReq`
- Response: `DeleteWordReply`

2. 请求定义


```golang
type DeleteWordReq struct {
	Words []string `json:"words"`
}
```


3. 返回定义


```golang
type DeleteWordReply struct {
}
```
  
![image](https://user-images.githubusercontent.com/4861699/133887766-6050d21c-7b47-4d99-89cf-15917d077da2.png)
