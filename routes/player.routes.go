package routes

import (
	"chess_game/controllers"

	"github.com/gofiber/fiber/v2"
)

func PlayerRoutes(app *fiber.App, ctrl *controllers.PlayerController) {
	app.Get("/players", ctrl.GetPlayers)

	app.Post("/createPlayer", ctrl.PostRegisterPlayer)

	app.Put("/players/:id/upRating", ctrl.PutUpPlayerRating)

	app.Post("/players/:playerID/addFriend/:friendID", ctrl.PostAddFriend)

	app.Get("/friendList/:playerID", ctrl.GetListFriends)
}
