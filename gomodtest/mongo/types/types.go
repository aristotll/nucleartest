package types

type Article struct {
	ID      string
	UserID  string
	Title   string
	Content string
}

type User struct {
	ID   string
	Name string
}

type Comment struct {
	ID        string
	ArticleID string

	ToCommentID  string // 是否是评论的评论（二级评论），不是的话此字段为空
	FromUserName string // 评论的发表者用户名（方便前端展示）
	FromUserID   string // 评论的发表者详细信息
	ToUserID     string // 如果是二级评论（评论的评论），这里记录被评论的用户
	ToUserName   string // 被评论者的用户名

	Content string // 评论内容

	TotalLike     int64    // 该评论的点赞数
	LikedUsersIDs []string // 点赞用户 id
	TotalUnLike   int64    // 该评论的踩数
	UnLikeUserIDs []string // 踩用户 id
}
