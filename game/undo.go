package game

import (
	"aoi/models"
	"aoi/models/gameundo"
	"aoi/models/gameundoitem"
	"errors"
)

func Undo(g *Game, id int64, history int64, user int64) error {
	conn := models.NewConnection()
	defer conn.Close()

	gameundoManager := models.NewGameundoManager(conn)
	gameundoitemManager := models.NewGameundoitemManager(conn)

	count := gameundoManager.Count([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "status", Value: gameundo.StatusWait, Compare: "="},
	})

	if count > 0 {
		return errors.New("already")
	}

	status := gameundo.StatusWait
	if g.Count == 1 {
		status = gameundo.StatusComplete
	}
	item := models.Gameundo{Status: status, Gamehistory: history, Game: id, User: user}
	err := gameundoManager.Insert(&item)

	if err != nil {
		return errors.New("db error")
	}

	gameundoId := gameundoManager.GetIdentity()

	gameundoitem := models.Gameundoitem{Status: gameundoitem.StatusAccept, Gameundo: gameundoId, Game: id, User: user}
	gameundoitemManager.Insert(&gameundoitem)

	g.MakeUndo(gameundoId, history, user)

	if g.Count == 1 {
		gamehistoryManager := models.NewGamehistoryManager(conn)
		gamehistoryManager.DeleteWhere([]interface{}{
			models.Where{Column: "game", Value: id, Compare: "="},
			models.Where{Column: "id", Value: history, Compare: ">="},
		})

		MakeGame(id)
	}

	return nil
}

func UndoConfirm(g *Game, id int64, undo int64, user int64, status int) error {
	db := models.NewConnection()
	defer db.Close()

	conn, _ := db.Begin()
	defer conn.Rollback()

	gameundoManager := models.NewGameundoManager(conn)
	gameundoitemManager := models.NewGameundoitemManager(conn)

	gameundoItem := gameundoManager.Get(undo)

	if gameundoItem.Status != gameundo.StatusWait {
		return errors.New("complete")
	}

	items := gameundoitemManager.Find([]interface{}{
		models.Where{Column: "gameundo", Value: undo, Compare: "="},
	})

	for _, v := range items {
		if v.User == user {
			return errors.New("already")
		}
	}

	item := models.Gameundoitem{Status: gameundoitem.Status(status), Gameundo: undo, Game: id, User: user}
	gameundoitemManager.Insert(&item)

	items = append(items, item)
	g.AddUndoConfirm(user)

	if len(items) != g.Count {
		conn.Commit()
		return nil
	}

	gameundoItem.Status = gameundo.StatusComplete
	gameundoManager.Update(gameundoItem)

	flag := true
	for _, v := range items {
		if v.Status == gameundoitem.StatusReject {
			flag = false
			break
		}
	}

	g.UndoRequest = UndoRequest{Users: make([]int, 0)}

	if flag != true {
		conn.Commit()
		return nil
	}

	gamehistoryManager := models.NewGamehistoryManager(conn)
	gamehistoryManager.DeleteWhere([]interface{}{
		models.Where{Column: "game", Value: id, Compare: "="},
		models.Where{Column: "id", Value: gameundoItem.Gamehistory, Compare: ">="},
	})

	conn.Commit()

	MakeGame(id)

	return nil
}
