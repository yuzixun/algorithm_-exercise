package main

type Twitter struct {
	Tweets       []int
	UserTweets   map[int][]int
	Follows      map[int][]int
	IsFollowSelf map[int]bool
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		Tweets:       []int{},
		UserTweets:   map[int][]int{},
		Follows:      map[int][]int{},
		IsFollowSelf: map[int][]int{},
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Tweets = append(this.Tweets, tweetId)
	this.UserTweets[userId] = append(this.UserTweets[userId], tweetId)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	followers := this.Follows[userId]
	var allTweets []int
	for _, v := range fs {
		allTweets = append(allTweets, this.UserTweets[v]...)
	}
	var sortTweets []int
	tLen := len(this.Tweets)
	s := 0

}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {

}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {

}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {

}
