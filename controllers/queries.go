package controllers

var GetAllPostsQuery = `
	SELECT p.id, title, body, image_url, c.name, username, date,
	(SELECT COUNT(*)
	FROM likes
	WHERE post_id = p.id
	) AS like_count,
	(SELECT COUNT(*)
	FROM comments
	WHERE post_id = p.id
	) AS comment_count
	FROM posts AS p
	JOIN users AS u
	ON p.user_id = u.id
	JOIN categories AS c
	ON p.category_id = c.id
`
var GetPostsByTitleQuery = `
 	SELECT p.id, title, body, image_url, c.name, username, date,
	(SELECT COUNT(*)
	FROM likes
	WHERE post_id = p.id
	) AS like_count,
	(SELECT COUNT(*)
	FROM comments
	WHERE post_id = p.id
	) AS comment_count
	FROM posts AS p
	JOIN users AS u
	ON p.user_id = u.id
	JOIN categories AS c
	ON p.category_id = c.id
	WHERE LOWER(title) LIKE '%%title_query%%'
`
