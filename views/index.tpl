<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">

<head>
    <title>Sensitive to check</title>
</head>

<body>
<form action="/check" method="POST">
    <div>
        敏感词检测
    </div>
    <div>
        <label>请输入要检查的字段: </label>
        <input type="text" name="txt" placeholder="请输入"/>
        <input type="submit" value="提交">
    </div>
</form>
</body>
<script>
    if (null != {{$.result}}) {
        alert({{$.result}})
    }


</script>

</html>