package main

import (
	"encoding/json"
	"golang.org/x/sync/singleflight"
	"net/http"
	"time"
)

var httpRequestGroup singleflight.Group

func (lbs *Lbs) RegisterHTTPHandlers() {
	diskName := func(disk int) string {
		if disk == GameDiskDC1 {
			return "dc1"
		}
		if disk == GameDiskDC2 {
			return "dc2"
		}
		if disk == GameDiskPS2 {
			return "ps2"
		}
		return "unknown"
	}

	teamName := func(team int) string {
		if team == TeamRenpo {
			return "renpo"
		}
		if team == TeamZeon {
			return "zeon"
		}
		return ""
	}

	gameStateName := func(state int) string {
		if state == McsGameStateCreated {
			return "created"
		}
		if state == McsGameStateOpened {
			return "opened"
		}
		if state == McsGameStateClosed {
			return "closed"
		}
		return "unknown"
	}

	http.HandleFunc("/lbs/status", func(w http.ResponseWriter, r *http.Request) {
		type onlineUser struct {
			UserID     string `json:"user_id,omitempty"`
			Name       string `json:"name,omitempty"`
			Team       string `json:"team,omitempty"`
			BattleCode string `json:"battle_code,omitempty"`
			Disk       string `json:"disk,omitempty"`
		}

		type activeGame struct {
			BattleCode string        `json:"battle_code,omitempty"`
			Region     string        `json:"region,omitempty"`
			Disk       string        `json:"disk,omitempty"`
			State      string        `json:"state,omitempty"`
			Users      []*onlineUser `json:"users,omitempty"`
			UpdatedAt  time.Time     `json:"updated_at,omitempty"`
		}

		type statusResponse struct {
			LobbyUsers  []*onlineUser `json:"lobby_users"`
			BattleUsers []*onlineUser `json:"battle_users"`
			ActiveGames []*activeGame `json:"active_games"`
		}

		resp, err, _ := httpRequestGroup.Do("/lbs/status", func() (interface{}, error) {
			resp := new(statusResponse)

			lbs.Locked(func(lbs *Lbs) {
				for _, u := range lbs.userPeers {
					resp.LobbyUsers = append(resp.LobbyUsers, &onlineUser{
						UserID:     u.UserID,
						Name:       u.Name,
						Team:       teamName(int(u.Team)),
						BattleCode: "",
						Disk:       diskName(int(u.GameDisk)),
					})
				}
			})

			for _, u := range sharedData.GetMcsUsers() {
				disk := GameDiskUnk
				b, ok := sharedData.GetBattleGameInfo(u.BattleCode)
				if ok {
					disk = b.GameDisk
				}

				resp.BattleUsers = append(resp.BattleUsers, &onlineUser{
					UserID:     u.UserID,
					Name:       u.Name,
					Team:       teamName(int(u.Side)),
					BattleCode: u.BattleCode,
					Disk:       diskName(disk),
				})
			}

			for _, g := range sharedData.GetMcsGames() {
				resp.ActiveGames = append(resp.ActiveGames, &activeGame{
					BattleCode: g.BattleCode,
					Disk:       diskName(g.GameDisk),
					State:      gameStateName(g.State),
					UpdatedAt:  g.UpdatedAt,
				})
			}

			return resp, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
}
