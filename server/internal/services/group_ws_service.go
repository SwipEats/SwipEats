package services

import (
	"github.com/SwipEats/SwipEats/server/internal/types"
)

func GroupBroadcast(gs types.GroupSession, msg any) {
	for id, client := range gs.Clients {
		if err := client.Conn.WriteJSON(msg); err != nil {
			client.Conn.Close()
			delete(gs.Clients, id)
		}
	}
}

func GetOrCreateGroupSession(gss *types.GroupSessionService, groupCode string) *types.GroupSession {
	session, exists := gss.Sessions[groupCode]
	if !exists {
		session = &types.GroupSession{
			ID:      groupCode,
			Clients: make(map[uint]*types.Client),
		}
		gss.Sessions[groupCode] = session
	}
	return session
}

func DeleteGroupSession(gss *types.GroupSessionService, groupCode string) {
	if session, exists := gss.Sessions[groupCode]; exists {
		for _, client := range session.Clients {
			client.Conn.Close()
		}
		delete(gss.Sessions, groupCode)
	}
}