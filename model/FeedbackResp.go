package model

import "../entity"

type FeedbackResp struct {
	entity.Feedback
	Pictures []entity.Picture
}

