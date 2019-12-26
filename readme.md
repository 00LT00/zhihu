### 接口文档

> BaseURL = zerokirin.online/api

注，全文所有接口必须在请求头中发送

```
sign = "spppk"
```

否则禁止访问

获取token以后的所有步骤，必须在请求头中发送

```
Authorization = token
```

其中token是第二步返回的参数，一方面证明身份，另一方面验证是否有权访问

#### 用户

api地址:	/users/

- 注册

  - 请求方式：PUT

  - 请求内容：通过form表单传值

    ```
    UserName:		用户名
    PassWord:		密码
    Phone:			手机号
    Email:			邮箱
    NickName:		昵称
    Introduction:	个人简介	
    ```

  - 响应格式：JSON

- 获取token

  - 请求方式：GET

  - 请求内容：通过form表单传值

    ```
    UserName:		用户名
    PassWord:		密码	
    ```

  - 响应格式：JSON

- 获取用户详情

  - 请求方式：GET

  - 请求内容：通过url传值

    ```
    接口地址末尾拼接：UserID
    ```

    注：如果token与UserID同源，则返回全部信息，否则只返回部分信息

  - 响应格式：JSON

- 更改用户详情

  - 请求方式：PATCH

  - 请求内容：通过form表单传值，url后拼接UserID

    ```
    UserName:		用户名
    PassWord:		密码
    Phone:			手机号
    Email:			邮箱
    NickName:		昵称
    Introduction:	个人简介	
    ```

    若UserID与token不同源则禁止修改

  - 响应格式：JSON

#### 文章

api地址：	/articles/

- 发布文章

  - 请求方式：POST

  - 请求内容：通过form表单传值

    ```
    title：		文章标题
    content：	文章内容
    ```

  - 响应格式：JSON

- 查找用户的文章

  - 请求方式：GET
  - 响应格式：JSON

- 根据文章id查找

  - 请求方式：GET

  - 请求内容：url传值

    ```
    url地址后拼接/id/articleid	
    ```

  - 响应格式：JSON

- 全部文章

  - 请求方式：GET

  - 请求内容：url地址拼接

    ```
    拼接：	/all/这里可选字符串如下：
    hot：	热榜前十
    new：	最新
    ```

  - 响应格式：JSON

#### 问题

api地址：	/questions/

- 发布问题

  - 请求方式：POST

  - 请求内容：通过form表单传值

    ```
    title：		问题标题
    content：	问题内容
    ```

  - 响应格式：JSON

- 查找用户的问题

  - 请求方式：GET
  - 响应格式：JSON

- 根据问题id查找

  - 请求方式：GET

  - 请求内容：url传值

    ```
    url地址后拼接/id/questionid	
    ```

  - 响应格式：JSON

- 全部问题

  - 请求方式：GET

  - 请求内容：url地址拼接

    ```
    拼接：	/all/这里可选字符串如下：
    hot：	热榜前十
    new：	最新
    ```

  - 响应格式：JSON

#### 回答

api地址：	/answers/

- 发布回答

  - 请求方式：POST

  - 请求内容：通过form表单传值

    ```
    questionid：		问题id
    content：		回答内容
    ```

  - 响应格式：JSON

- 根据问题找回答

  - 请求方式：GET

  - 请求内容：url地址拼接，form表单传值

    ```
    传值
    questionid：		问题id
    
    拼接：	/question/这里可选字符串如下：
    hot：	热榜前十
    new：	最新
    ```

  - 响应格式：JSON

- 根据用户找回答

  - 请求方式：GET

  - 请求内容：url地址拼接，form表单传值

    ```
    传值
    userid：		用户id
    
    拼接：	/user/这里可选字符串如下：
    hot：	热榜前十
    new：	最新
    ```

  - 响应格式：JSON

- 根据回答id获取详情

  - 请求方式：GET

  - 请求内容：form表单传值

    ```
    answerid:	回答id	
    ```

  - 响应格式：JSON

#### 评论

api地址：	/comments/

- 写评论

  - 请求方式：POST

  - 请求内容：通过form表单传值

    ```
    parentid：		问题或文章等，对什么评论，就是谁的id
    content：		评论内容
    ```

  - 响应格式：JSON

- 根据id找评论

  - 请求方式：GET

  - 请求内容：url地址拼接，form表单传值

    ```
    传值
    id：		评论对象的id
    
    拼接：	/id/这里可选字符串如下：
    hot：	热榜前十
    new：	最新
    ```

  - 响应格式：JSON

- 根据任意一评论寻找整个评论的对话组

  - 请求方式：GET

  - 请求内容：url地址拼接，form表单传值

    ```
    传值
    targetid：		该评论的id
    
    拼接：	/target/
    ```

  - 响应格式：JSON

- 点赞或反对

  - 请求方式：POST

  - 请求内容：通过form表单传值

    ```
    id：				点赞什么，就是谁的id
    type：			article/question/comment/answer
    action：			up/down
    ```

  - 响应格式：JSON