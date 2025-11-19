package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"grease_trace/application/usecase"
	"grease_trace/infrastructure/persistence"
	"grease_trace/infrastructure/router"
	"grease_trace/interfaces/controller"
)

func envLoad() {
	err := godotenv.Overload("../.env")
	if err != nil {
		log.Fatalf("Error loading env target")
	}
}

func main() {
	// 1. Infrastructure: データベース(今回はメモリ)の準備
	repo := persistence.NewInMemorySessionRepository()

	// 2. Application: ユースケースの準備 (Repoを注入)
	interactor := usecase.NewSessionInteractor(repo)

	// 3. Interface: コントローラーの準備 (Usecaseを注入)
	ctrl := controller.NewSessionController(interactor)

	// 4. Infrastructure: ルーターの準備 (Controllerを注入)
	r := router.NewRouter(ctrl)

	GOPORT := os.Getenv("GO_PORT")

	// 5. サーバー起動
	log.Println("Go APIサーバーをポート " + GOPORT + " で起動します...")
	r.Run(":" + GOPORT)
}
