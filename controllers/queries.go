package controllers

var GetAllPostsQuery = `
	SELECT p.id,
				title,
				body,
				image_url,
				c.name,
				username,
				date,
				(SELECT Count(*)
					FROM   post_likes
					WHERE  post_id = p.id) AS like_count,
				(SELECT Count(*)
					FROM   comments
					WHERE  post_id = p.id) AS comment_count
	FROM   posts AS p
				JOIN users AS u
					ON p.user_id = u.id
				JOIN categories AS c
					ON p.category_id = c.id 
`

var TitleQuery = GetAllPostsQuery + `
	WHERE LOWER(title) LIKE '%%%s%%'
`
var CategoryQuery = GetAllPostsQuery + `
	WHERE c.name = '%s'
`

var TitleAndCategoryQuery = TitleQuery + `
	AND c.name = '%s'
`

var GetPostQuery = `
	SELECT title,
				p.body,
				image_url,
				c.name,
				username,
				p.date,
				(SELECT COUNT(*)
					FROM   post_likes
					WHERE  post_id = p.id) AS like_count,
				(SELECT COUNT(*)
				  FROM 	 comments
					WHERE  post_id = p.id) AS comment_count,
				(SELECT ARRAY_AGG(
									JSON_BUILD_OBJECT('id', com.id,
																		'username',	(SELECT username
																									FROM users AS u
																									WHERE u.id = com.user_id),
																		'body', com.body,
																		'date', com.date,
																		'likeCount', (SELECT COUNT(*)
																										FROM comment_likes AS cl
																										WHERE cl.comment_id = com.id),
																		'replyCount', (SELECT COUNT(*)
																										FROM comments
																										WHERE parent_id = com.id)
									)
								)
					FROM   comments AS com
					WHERE  post_id = p.id
					AND parent_id IS NULL) AS comments
	FROM   posts AS p
				JOIN users AS u
					ON p.user_id = u.id
				JOIN categories AS c
					ON p.category_id = c.id
	WHERE  p.id = '%s' 
`

var PostPostQuery = `
	INSERT INTO
		posts (id, title, body, image_url, category_id, user_id, date)
	VALUES
		(DEFAULT, :title, :body, :imageUrl, :categoryId, :userId, :date)
`

var GetAllCategoriesQuery = `
	SELECT * FROM categories
`

var GetAllCommentsQuery = `
	SELECT * FROM comments
`

var ParentQuery = `
	SELECT id,
				 (SELECT username
				  FROM users AS u
					WHERE u.id = user_id) AS username,
				 body,
				 date,
				 (SELECT COUNT(*)
				  FROM comment_likes AS cl
					WHERE cl.comment_id = id) AS likeCount,
				 (SELECT COUNT(*)
				  FROM comments
					WHERE comments.parent_id = c.id) AS replyCount
	FROM comments AS c
	WHERE parent_id = %s
`

var GetUserById = `
	SELECT id, username, password
	FROM users
	WHERE id = %s
`

var GetUserByUsername = `
	SELECT id, username, password
	FROM users
	WHERE username = '%s'
`

var PostUserQuery = `
	INSERT INTO
		users (id, username, password)
	VALUES
		(DEFAULT, :username, :password)
`
