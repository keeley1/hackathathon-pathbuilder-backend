package handlers

import "github.com/gin-gonic/gin"

func TaskProgressHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"tasks": []gin.H{
			{
				"id":          "update-linkedin",
				"name":        "Update LinkedIn Profile",
				"description": "Add your latest role and skills to your LinkedIn profile.",
				"completed":   true,
				"progress":    100,
			},
			{
				"id":          "complete-course",
				"name":        "Complete Online Course",
				"description": "Finish the 'Go Programming' course on Coursera.",
				"completed":   false,
				"progress":    60,
			},
			{
				"id":          "attend-networking-event",
				"name":        "Attend Networking Event",
				"description": "Participate in a local tech meetup.",
				"completed":   false,
				"progress":    0,
			},
		},
		"overallProgress": 53,
	})
}
