package secure

import "time"

type SafeUser struct {
	ID          int64
	Name        string
	Verified    bool
	DateCreated time.Time
	LastLogin   time.Time
}

func createSafeUser(user userRecord) SafeUser {
	data := user.Data()
	meta := user.rec.Meta()

	result := SafeUser{
		ID:          meta.ID,
		DateCreated: meta.CreateDate,
		LastLogin:   data.LoginDate,
		Name:        data.Name,
		Verified:    data.Verified,
	}

	return result
}

func GetUsers(page, size int) (result []SafeUser, err error) {
	users, err := getUsers(page, size)

	if err != nil {
		return result, err
	}

	for users.MoveNext() {
		currUser, err := users.Current()

		if err != nil {
			return result, err
		}

		sfeUser := createSafeUser(userRecord{currUser})
		result = append(result, sfeUser)
	}

	return result, err
}
