package controllers

import (
	m "Echo/models"
	"net/http"
	"strconv"
)

func GetAllUsers() (m.GlobalResponse, error) {
	var u m.User
	var users []m.User
	var resp m.GlobalResponse

	data := Connect()
	defer data.Close()

	query := "SELECT * FROM users"

	rows, err := data.Query(query)

	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Gender); err != nil {
			return resp, err
		}

		users = append(users, u)
	}

	resp.Status = http.StatusAccepted
	resp.Message = "Success!"
	resp.Data = users

	return resp, err
}

func GetUserByID(id string) (m.GlobalResponse, error) {
	var u m.User
	var resp m.GlobalResponse

	data := Connect()
	defer data.Close()

	var count int
	rows, err := data.Query("SELECT COUNT(*) FROM users WHERE id=?", id)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return resp, err
		}
	}
	if count == 0 {
		resp.Status = http.StatusBadRequest
		resp.Message = "No user with that ID!"
		return resp, err
	}

	rows, err = data.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Gender); err != nil {
			return resp, err
		}
	}

	resp.Status = http.StatusAccepted
	resp.Message = "Success!"
	resp.Data = u

	return resp, err
}

func InsertUser(name string, age string, gender string) (m.GlobalResponse, error) {
	var u m.User
	var resp m.GlobalResponse

	data := Connect()
	defer data.Close()

	result, err := data.Exec("INSERT INTO users VALUES (null, ?, ?, ?)", name, age, gender)
	if err == nil {
		resp.Status = http.StatusAccepted
		resp.Message = "Data has been successfully inserted!"
		temp, _ := result.LastInsertId()
		u.ID = int(temp)
		u.Name = name
		u.Age, _ = strconv.Atoi(age)
		u.Gender = gender
		resp.Data = u
	} else {
		println(err.Error())
		return resp, err
	}

	return resp, err
}

func UpdateUser(id string, name string, age string, gender string) (m.GlobalResponse, error) {
	var u m.User
	var resp m.GlobalResponse

	data := Connect()
	defer data.Close()

	result, err := data.Exec("UPDATE users SET name = ?, age = ?, gender = ? WHERE id = ?", name, age, gender, id)
	if err != nil {
		return resp, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		resp.Status = http.StatusBadRequest
		resp.Message = "No user with that ID!"
		return resp, err
	}
	u.ID, _ = strconv.Atoi(id)
	u.Name = name
	u.Age, _ = strconv.Atoi(age)
	u.Gender = gender
	resp.Status = http.StatusAccepted
	resp.Message = "Data has been successfully updated!"
	resp.Data = u
	return resp, nil
}

func DeleteUser(id string) (m.GlobalResponse, error) {
	var u m.User
	var resp m.GlobalResponse

	data := Connect()
	defer data.Close()

	var count int
	rows, err := data.Query("SELECT COUNT(*) FROM users WHERE id=?", id)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return resp, err
		}
	}
	if count == 0 {
		resp.Status = http.StatusBadRequest
		resp.Message = "No user with that ID!"
		return resp, err
	}
	rows, err = data.Query("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Gender); err != nil {
			return resp, err
		}
	}

	_, err = data.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return resp, err
	}

	resp.Status = http.StatusAccepted
	resp.Message = "Data has been successfully deleted!"
	resp.Data = u

	return resp, nil
}
