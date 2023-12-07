package main

import (
    "context"
    "ilsan/ent"
    "ilsan/ent/user"
)

type UserRepository struct{
    Client *ent.UserClient
}

func (repo *UserRepository) FindByName(name string) *ent.User{
    result := repo.Client.Query().
		Where(user.NameEQ(name)).
		OnlyX(context.TODO())

    return result
}
