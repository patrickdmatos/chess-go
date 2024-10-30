package routes

import (
	"chess_game/controllers"

	"github.com/gofiber/fiber/v2"
)

func MatchRoutes(app *fiber.App, ctrl *controllers.MatchController) {
    app.Post("/matches", ctrl.PostCreateMatch) // Cria uma nova partida

    app.Get("/matches/:id", ctrl.GetMatchByID) // Recupera uma partida pelo ID

    app.Put("/matches/:id/:status", ctrl.PutUpdateStatusMatch) // Atualiza o status da partida
}
