package profile

import (
	userEntity "online-course.mifwar.com/internal/user/entity"
)

type ProfileRespondBody struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func CreateProfileResponse(user userEntity.User) ProfileRespondBody {
	isVerified := false

	if user.EmailVerifiedAt.Valid {
		isVerified = true
	}

	return ProfileRespondBody{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: isVerified,
	}
}
