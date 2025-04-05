package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/edzhabs/social/internal/store"
)

var usernames = []string{
	"alice", "bob", "dave", "charlie", "eve", "frank", "grace", "heidi", "ivan", "judy",
	"karl", "laura", "mallory", "nancy", "oliver", "pat", "quinn", "rick", "susan", "tom",
	"uma", "victor", "wendy", "xander", "yvonne", "zach", "amy", "brad", "carl", "diana",
	"elena", "finn", "gwen", "hank", "irene", "jack", "kim", "luke", "mia", "noah",
	"peter", "quincy", "rosa", "sam", "tara", "ursula", "vance", "will", "xena", "yale",
	"zoe", "adrian", "bella", "claire", "dennis", "eric",
}

var titles = []string{
	"Journey Through the Unknown", "The Silent Echo", "Code of the Future", "Whispers in the Dark", "Paths Untraveled",
	"Shadows of the Past", "Rise of the Phoenix", "Beyond the Horizon", "Echoes of Tomorrow", "The Infinite Loop",
	"Veil of Secrets", "The Lost Dimension", "Codebreaker Chronicles", "The Seventh Realm", "The Hidden Gateway",
	"Dreams of the Digital World", "Through the Storm", "Whispers of the Universe", "Crimson Skies", "Endless Adventures",
}

var contents = []string{
	"A deep dive into unexplored realms.", "Exploring the mysteries that lie within shadows.", "Unlocking the potential of future technology.", "A story of secrets that no one dares to reveal.", "An adventure through worlds yet to be discovered.",
	"Unraveling the past through forgotten tales.", "An epic rise of strength and rebirth.", "Chasing horizons that never end.", "Predicting what tomorrow holds.", "Falling into the cycle that never ends.",
	"Discovering truths hidden beneath the surface.", "Stepping into dimensions that defy logic.", "Cracking codes and solving mysteries.", "A journey into an unknown world.", "Accessing doors that were once closed.",
	"Exploring the digital frontier in search of truth.", "Overcoming adversity in the face of chaos.", "Seeking the wisdom that lies within the cosmos.", "Adventures beyond the limits of imagination.", "Chasing after dreams that seem unreachable.",
}

var tags = []string{
	"adventure", "mystery", "technology", "exploration", "rebirth",
	"secrets", "future", "unknown", "dimensional", "innovation",
	"chaos", "reality", "digital", "cyber", "space",
	"fantasy", "discover", "journey", "legacy", "infinity",
}

var postComments = []string{
	"Great insight into the future of technology!",
	"Such a mysterious concept, really makes you think.",
	"Can't wait to see where this journey leads.",
	"Love the depth of the world-building here.",
	"These ideas are incredibly innovative and thought-provoking.",
	"The characters' evolution throughout this story is amazing.",
	"Such a unique perspective on digital realms.",
	"Definitely feels like a journey into the unknown.",
	"Impressive how you've combined technology with fantasy.",
	"Really enjoyed this exploration of parallel worlds.",
	"The blend of mystery and adventure keeps me hooked.",
	"The universe feels so vast, it's hard to fathom.",
	"The plot twists are just incredible!",
	"The concept of rebirth adds so much depth to the story.",
	"The secrets revealed here are mind-blowing.",
	"Such an unexpected turn, I didn’t see it coming.",
	"The writing is so immersive, I can't put it down.",
	"Absolutely love the vivid descriptions of the world.",
	"This makes me want to explore more of this universe!",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	tx, _ := db.BeginTx(ctx, nil)

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Printf("Error creating user: %v ; err:%s", user, err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Printf("Error creating post: %v ; err:%s", post, err)
			return
		}
	}

	comments := generateComments(200, posts, users)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Printf("Error creating comment: %v ; err:%s", comment, err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags:    randomTags(),
		}
	}

	return posts
}

func randomTags() []string {
	numTags := rand.Intn(5)
	postTags := make([]string, numTags)

	for i := 0; i < numTags; i++ {
		tag := tags[rand.Intn(len(tags))]
		postTags[i] = tag
	}

	return postTags
}

func generateComments(num int, posts []*store.Post, users []*store.User) []*store.Comment {
	comments := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		comments[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: postComments[rand.Intn(len(postComments))],
		}
	}

	return comments
}
