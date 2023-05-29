# gorm

- [官网](https://gorm.io)
- https://gorm.io/docs/query.html

# model 的定义

表的属性在 model 中 应大写字母开头，连字符属性转为驼峰形式
默认情况下 表名为结构体名称的复数形式，可以通过自定义 TableName 方法改变表名

- 主键
- 外键
    一对多，一对一
    foreignKey:CateId 以当前表中 CateId作为外键字段
    references:Id 以关联表中 id作为关联字段
    Cate UserCate gorm:"foreignKey:CateId;references:Id"
    以 CateId 为外键 查询 UserCate,以此作为 Cate 的值
    使用方式: db.Preload("Cate").Find(&list)


