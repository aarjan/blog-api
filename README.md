## A Sample Blog API

It is based on a concept of Blog.
The main elements are __post__, __tags__ and __categories__  

---

It contains endpoints for 
###GET all items
    + /api/v1/posts
    + /api/v1/tags
    + /api/v1/categories

###GET particular item
    + /api/v1/post/{id:[0-9]+}
    + /api/v1/tag/{id:[0-9]+}
    + /api/v1/category/{id:[0-9]+}

 
###POST
	+ /api/v1/post
	+ /api/v1/tag
	+ /api/v1/category

###DELETE
	+ /api/v1/post/{id:[0-9]+}
	+ /api/v1/tag/{id:[0-9]+}
	+ /api/v1/category/{id:[0-9]+}

###PUT
	+ /api/v1/tag/{id:[0-9]+}
	+ /api/v1/category/{id:[0-9]+}

---
##Database can be imported from db-setup/seed-data.sql