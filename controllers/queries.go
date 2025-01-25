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
	FROM  posts AS p
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
																		'replyCount', (SELECT COUNT(*)
																										FROM comments
																										WHERE comments.parent_id = com.id)
									)
								ORDER BY com.date DESC)
					FROM   comments AS com
					WHERE  com.post_id = p.id
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

var UpdatePostQuery = `
	UPDATE posts 
	SET title = :title, body = :body, image_url = :imageUrl, category_id = :categoryId, date = :date
	WHERE id = :id
`

var DeleteLikeByPost = `
	DELETE FROM post_likes
	WHERE post_id = :id
`

var DeleteCommentByPost = `
	DELETE FROM comments
	WHERE post_id = :id
`

var DeletePostQuery = `
	DELETE FROM posts
	WHERE id = :id
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
				  FROM comments
					WHERE comments.parent_id = c.id) AS replyCount
	FROM comments AS c
	WHERE parent_id = %s
`

var PostCommentQuery = `
	INSERT INTO
		comments (id, body, user_id, post_id, date, parent_id)
	VALUES
		(DEFAULT, :body, :userId, :postId, :date, :parentId)
`

var PostCommentWithoutParentQuery = `
	INSERT INTO
		comments (id, body, user_id, post_id, date)
	VALUES
		(DEFAULT, :body, :userId, :postId, :date)
`

var UpdateCommentQuery = `
 	UPDATE comments 
	SET body = :body, date = :date
	WHERE id = :id
`

var DeleteCommentQuery = `
	DELETE FROM comments
	WHERE parent_id = :id OR id = :id
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

var GetLikeQuery = `
	SELECT *
	FROM post_likes
	WHERE user_id = %d AND post_id = %s
`

var PostLikeQuery = `
	INSERT INTO
		post_likes (user_id, post_id)
	VALUES
		(%d, %s)
`

var DeleteLikeQuery = `
	DELETE FROM post_likes
	WHERE user_id = %d AND post_id = %s
`
