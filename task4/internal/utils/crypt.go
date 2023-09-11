package crypt

import (
	// "bytes"
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt() ([]byte, error) {
    salt := make([]byte, 16)
    _, err := rand.Read(salt)
    if err != nil {
        return nil, err
    }

    return salt, nil
}

func HashPassword(password string, salt []byte) ([]byte, error) {
    passwordBytes := []byte(password + string(salt))
    hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    return hashedPassword, nil
}

// func VerifyPassword(username, password string) (bool, error) {
//     // Retrieve the user's data (salt and hashed password) from the database
//     // (You will need to implement database interaction here)

//     // Compare the provided password with the stored values
//     hashedPassword, err := hashPassword(password, retrievedSalt)
//     if err != nil {
//         return false, err
//     }

//     return bytes.Equal(hashedPassword, retrievedHashedPassword), nil
// }
