.. _web:


web develop
====================

记录web 开发中的点点滴滴

beego framework note 
----------------------

使用beego 框架的一些笔记 

notices 
^^^^^^^^^^^^^^^

关于模板路径:

可以在controller里自己定义模板是哪个文件
例如::

	this.TplNames = "show.tpl"

另外可以使用默认模板, 例如 HomeController 中Get方法对应的模板是 HomeController 目录下的GET.tpl
注意大小写 	
