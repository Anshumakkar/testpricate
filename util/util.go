package util


import "github.com/google/uuid"

func GetUUID() string{
  newuuid, _ := uuid.NewUUID()
  return newuuid.String()
}
