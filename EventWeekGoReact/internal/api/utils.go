package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/BManduca/RocketseatDiaryDev/tree/main/EventWeekGoReact/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (h apiHandler) readRoom(
	w http.ResponseWriter,
	r *http.Request,
) (room pgstore.Room, rawRoomID string, roomID uuid.UUID, ok bool) {
	rawRoomID = chi.URLParam(r, "room_id")
	// Verificando se é um ID válido
	roomID, err := uuid.Parse(rawRoomID)
	if err != nil {
		http.Error(w, "Invalid Room ID", http.StatusBadRequest)
		return pgstore.Room{}, "", uuid.UUID{}, false
	}

	// 'buscando' la do DB e verificando se a sala existe
	// conexao com cliente pelo lado do servidor
	room, err = h.q.GetRoom(r.Context(), roomID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // target pgx.ErrNoRows -> Sala não encontrada
			http.Error(w, "Room Not Found", http.StatusBadRequest)
			return pgstore.Room{}, "", uuid.UUID{}, false
		}

		slog.Error("Failed To Get Room", "error", err)
		http.Error(w, "Something Went Wrong", http.StatusInternalServerError)
		return pgstore.Room{}, "", uuid.UUID{}, false
	}

	return room, rawRoomID, roomID, true
}

func sendJSON(w http.ResponseWriter, rawData any) {
	data, _ := json.Marshal(rawData)
	// setando o header
	w.Header().Set("Content-Type", "application/json")
	// escrevendo os dados la no client
	_, _ = w.Write(data)
}
