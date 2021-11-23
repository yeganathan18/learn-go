package db

import (
	"context"
	"fmt"
	"time"

	"github.com/learn/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Queries

// getAllRecipes
func (db MongoDB) GetAllRecipes() (interface{}, error) {
	var results []RecipeModel
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.Recipes.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem RecipeModel
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

// getRecipe
func (db MongoDB) GetRecipe(_id string) (interface{}, error) {
	var results RecipeModel
	var err error

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Recipes.FindOne(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// searchRecipes
func (db MongoDB) SearchRecipes(searchTerm string) (interface{}, error) {
	var err error
	var results []RecipeModel

	q := bson.M{"$text": bson.M{"$search": searchTerm}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.Recipes.Find(ctx, q, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem RecipeModel
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

// getCurrentUser
func (db MongoDB) GetCurrentUser(username string) (interface{}, error) {
	var err error
	var results UserModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Users.FindOne(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	var favorites []RecipeModel
	for e, _ := range results.Favorites {
		var recipe RecipeModel

		_id := results.Favorites[e].(primitive.ObjectID)
		q := bson.M{"_id": _id}
		if c, _ := db.Recipes.CountDocuments(context.Background(), q); c > 0 {
			ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
			err := db.Recipes.FindOne(ctx, q).Decode(&recipe)
			if err != nil {
				return nil, err
			}
			favorites = append(favorites, recipe)
		}
	}
	favi := make([]interface{}, len(favorites))
	for i, fav := range favorites {
		favi[i] = fav
	}
	results.Favorites = favi
	return results, nil
}

// getUserRecipes
func (db MongoDB) GetUserRecipes(username string) (interface{}, error) {
	var err error
	var results []RecipeModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.Recipes.Find(ctx, q, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem RecipeModel
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)
	return results, nil
}

// Mutations

// addRecipe
func (db MongoDB) AddRecipe(username string, name string, imageUrl string, category string, description string, instructions string) (interface{}, error) {
	var err error
	var results RecipeModel

	results.ID = primitive.NewObjectID()
	results.Name = name
	results.ImageUrl = imageUrl
	results.Category = category
	results.Description = description
	results.Instructions = instructions
	results.CreatedDate = time.Now()
	results.Username = username
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.Recipes.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// likeRecipe
func (db MongoDB) LikeRecipe(_id string, username string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}
	q2 := bson.M{"$inc": bson.M{"likes": 1}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Recipes.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	//results.Likes++
	q = bson.M{"username": username}
	q2 = bson.M{"$addToSet": bson.M{"favorites": id}}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Users.FindOneAndUpdate(ctx, q, q2).Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// unlikeRecipe
func (db MongoDB) UnlikeRecipe(_id string, username string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id}
	q2 := bson.M{"$inc": bson.M{"likes": -1}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Recipes.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	//results.Likes--
	q = bson.M{"username": username}
	q2 = bson.M{"$pull": bson.M{"favorites": id}}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Users.FindOneAndUpdate(ctx, q, q2).Err()
	if err != nil {
		return nil, err
	}
	return results, nil
}

// deleteUserRecipe
func (db MongoDB) DeleteUserRecipe(_id string, user string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id, "username": user}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Recipes.FindOneAndDelete(ctx, q).Decode(&results)
	if err != nil {
		return nil, err
	}
	q = bson.M{"favorites": id}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.Users.Find(ctx, q, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem UserModel
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		q = bson.M{"username": elem.Username}
		q2 := bson.M{"$pull": bson.M{"favorites": id}}
		ctx2, _ := context.WithTimeout(context.Background(), 30*time.Second)
		err = db.Users.FindOneAndUpdate(ctx2, q, q2).Err()
		if err != nil {
			return nil, err
		}
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(ctx)

	return results, nil
}

// updateUserRecipe
func (db MongoDB) UpdateUserRecipe(_id string, user string, name string, imageUrl string, category string, description string, instructions string) (interface{}, error) {
	var err error
	var results RecipeModel

	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": id, "username": user}
	q2 := bson.M{"$set": bson.M{"name": name,
		"imageUrl":     imageUrl,
		"category":     category,
		"description":  description,
		"instructions": instructions,
	}}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = db.Recipes.FindOneAndUpdate(ctx, q, q2).Decode(&results)
	if err != nil {
		return nil, err
	}
	results.Name = name
	results.ImageUrl = imageUrl
	results.Category = category
	results.Description = description
	results.Instructions = instructions
	return results, nil
}

// signinUser
func (db MongoDB) SigninUser(username string, password string) (interface{}, error) {
	var err error
	var results UserModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err = db.Users.FindOne(ctx, q).Decode(&results); err != nil {
		return nil, fmt.Errorf("User not found.")
	}
	if err = web.ComparePassword(password, results.Password); err != nil {
		return nil, fmt.Errorf("Invalid password.")
	}
	results.Token, err = web.CreateToken(results.Username, results.Email)
	return results, nil
}

// signupUser
func (db MongoDB) SignupUser(username string, password string, email string) (interface{}, error) {
	var err error
	var results UserModel

	q := bson.M{"username": username}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := db.Users.FindOne(ctx, q).Decode(&results); err == nil {
		return nil, fmt.Errorf("User already exists.")
	}
	q = bson.M{"email": email}
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	if err := db.Users.FindOne(ctx, q).Decode(&results); err == nil {
		return nil, fmt.Errorf("Email already exists.")
	}
	pass, err := web.GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	results.ID = primitive.NewObjectID()
	results.Username = username
	results.Password = pass
	results.Email = email
	results.JoinDate = time.Now()
	results.Favorites = make([]interface{}, 0)
	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.Users.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	results.Token, err = web.CreateToken(results.Username, results.Email)
	return results, nil
}

