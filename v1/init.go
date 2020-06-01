import (
	fiber "github.com/gofiber/fiber"
)

var Version = "v1";

func Init(fiberRef *fiber) {
	app := fiber.New()

	v1 := api.Group(Version, mysql())
	v1.Get("/leaderboards", handlers.QueryLeaderboards)
	v1.Get("/:leaderboardID", handlers.GetLeaderboardsByID)
}
