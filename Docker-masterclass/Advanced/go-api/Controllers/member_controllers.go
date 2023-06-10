package Controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/newlinedeveloper/go-api/Configs"
	"github.com/newlinedeveloper/go-api/Models"
	"github.com/newlinedeveloper/go-api/Responses"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var memberCollection *mongo.Collection = Configs.GetCollection(Configs.DB, "members")
var validate = validator.New()

func CreateMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// rw.Header().Set("Content-Type", "application/x-www-form-urlencoded")
		// rw.Header().Set("Access-Control-Allow-Origin", "*")
		// rw.Header().Set("Access-Control-Allow-Methods", "POST")
		// rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var member Models.Member
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := Responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&member); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := Responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newUser := Models.Member{
			Id:    primitive.NewObjectID(),
			Name:  member.Name,
			Email: member.Email,
			City:  member.City,
		}
		result, err := memberCollection.InsertOne(ctx, newUser)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := Responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := Responses.MemberResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)

	}
}

func GetMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["id"]
		var user Models.Member

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := memberCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := Responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := Responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}}
		json.NewEncoder(rw).Encode(response)

	}
}



func GetAllMembers() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var members []Models.Member
		defer cancel()

		results, err := memberCollection.Find(ctx, bson.M{})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := Responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser Models.Member
			if err = results.Decode(&singleUser); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := Responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
			}
			members = append(members, singleUser)

		}

		rw.WriteHeader(http.StatusOK)
		response := Responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": members}}
		json.NewEncoder(rw).Encode(response)

	}
}


func UpdateMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["id"]
		var user Models.Member

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := Responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := Responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		update := bson.M{"name": user.Name, "email": user.Email, "city": user.City}

		result, err := memberCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := Responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Get Updated member details
		var updatedMember Models.Member

		if result.MatchedCount == 1 {
			err := memberCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedMember)

			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := Responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
				return
			}

		}

		rw.WriteHeader(http.StatusOK)
		response := Responses.MemberResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedMember}}
		json.NewEncoder(rw).Encode(response)


	}
}


func DeleteMember() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["id"]
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := memberCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := Responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			rw.WriteHeader(http.StatusNotFound)
			response := Responses.MemberResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Member Id Not found"}}
			json.NewEncoder(rw).Encode(response)
			return

		}

		rw.WriteHeader(http.StatusOK)
		response := Responses.MemberResponse{Status: http.StatusOK, Message: "Success", Data: map[string]interface{}{"data": "Member deleted successfully"}}
		json.NewEncoder(rw).Encode(response)

	}
}

